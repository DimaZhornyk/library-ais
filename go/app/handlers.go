package app

import (
	"database/sql"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	db *sql.DB
}

func (h *Handler) getQueries(c *gin.Context) {
	c.JSON(200, queries)
}

func (h *Handler) postQuery(c *gin.Context) {
	type Req struct {
		Query string `json:"query"`
	}
	var req Req
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if _, err := h.db.Query(req.Query); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{})
}
