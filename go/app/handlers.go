package app

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type Handler struct {
	db *sqlx.DB
}

func (h *Handler) getQueries(c *gin.Context) {
	c.JSON(200, entities)
}

func (h *Handler) postQuery(c *gin.Context) {
	var req Action
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if len(req.Queries) == 0 {
		c.JSON(400, gin.H{"error": "no queries provided"})
		return
	}

	// if there is only one query execute it
	if len(req.Queries) == 1 {
		resp, err := h.db.NamedQuery(req.Queries[0].Text, req.Queries[0].Params)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		m, err := rowsToSlice(resp.Rows)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, m)
		return
	}

	// if there are multiple queries execute them in a single transaction
	tx, err := h.db.Beginx()
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	for _, q := range req.Queries {
		if _, err := tx.NamedQuery(q.Text, q.Params); err != nil {
			_ = tx.Rollback()
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
	}

	if err := tx.Commit(); err != nil {
		_ = tx.Rollback()
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{})
}

type createBookInstanceReq struct {
	InventoryNumber string `json:"inventoryNumber"`
	BookISBN        string `json:"bookIsbn"`
	Shelf           string `json:"shelf"`
}

func (h *Handler) createBookInstance(c *gin.Context) {
	var req createBookInstanceReq
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	rows, err := h.db.NamedQuery(`
	SELECT EXISTS (
		SELECT * FROM book_instances
		WHERE book_isbn = :isbn
	) AND EXISTS(
		SELECT *
		FROM (
			 SELECT inventory_number
			 FROM book_instances
			 WHERE book_isbn = :isbn
		 ) AS ins
		WHERE NOT EXISTS(
			SELECT id
			FROM replacement_acts
			WHERE old_inventory_number = ins.inventory_number
		)
	) as present`,
		map[string]interface{}{
			"isbn": req.BookISBN,
		})
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	rows.Next()
	res := make(map[string]interface{})
	if err = rows.MapScan(res); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	if res["present"].(bool) {
		c.JSON(400, gin.H{"error": "book already exists"})
		return
	}

	_, err = h.db.NamedExec(` INSERT INTO book_instances (inventory_number, book_isbn, shelf) VALUES (:in, :isbn, :shelf)`,
		map[string]interface{}{
			"in":    req.InventoryNumber,
			"isbn":  req.BookISBN,
			"shelf": req.Shelf,
		})
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{})
}

func rowsToSlice(rows *sql.Rows) (res []map[string]interface{}, err error) {
	cols, _ := rows.Columns()
	for rows.Next() {
		columns := make([]interface{}, len(cols))
		columnPointers := make([]interface{}, len(cols))
		for i := range columns {
			columnPointers[i] = &columns[i]
		}

		// Scan the result into the column pointers...
		if err = rows.Scan(columnPointers...); err != nil {
			return nil, err
		}

		// Create our map, and retrieve the value for each column from the pointers slice,
		// storing it in the map with the name of the column as the key.
		m := make(map[string]interface{})
		for i, colName := range cols {
			val := columnPointers[i].(*interface{})
			m[colName] = *val
		}

		res = append(res, m)
	}

	return res, nil
}
