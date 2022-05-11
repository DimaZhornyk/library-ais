package app

const (
	String  = "string"
	Integer = "integer"
	Float   = "float"
	Date    = "date"
)

const (
	Reader    = "reader"
	Librarian = "librarian"
	Admin     = "admin"
)

type Action struct {
	Name    string  `json:"queryName"`
	Queries []Query `json:"queries"`
}

type Query struct {
	Text   string         `json:"text"`
	Params map[string]any `json:"params"`
}

type RoleEntities struct {
	Role     string   `json:"role"`
	Entities []Entity `json:"entities"`
}

type Entity struct {
	EntityName string   `json:"entityName"`
	BasicQuery string   `json:"basicQuery"`
	Actions    []Action `json:"actions"`
}

var entities = map[string][]Entity{
	Reader: {
		{"Books", "SELECT * FROM books", []Action{}},
		{"Knowledge areas", "SELECT * FROM knowledge_areas", []Action{
			{"Get books of area", []Query{{
				`SELECT book_instances.*
						FROM book_instances INNER JOIN book_areas ba ON book_instances.book_isbn = ba.book_isbn
						WHERE ba.book_isbn = :isbn AND NOT EXISTS(
							SELECT * FROM checkouts
							WHERE checkouts.book_inventory_number = book_instances.inventory_number
							AND checkouts.return_date IS NOT NULL
						) AND NOT EXISTS(
							SELECT * FROM replacement_acts
							WHERE replacement_acts.old_inventory_number = book_instances.inventory_number
						)`,
				map[string]any{
					"isbn": String,
				}},
			}},
		}},
		{"Book authors", "SELECT * FROM book_authors", []Action{
			{"Get books of author", []Query{{
				`SELECT books.*
						FROM books INNER JOIN book_authors ON books.isbn = book_authors.book_isbn
						WHERE book_authors.id = :id`,
				map[string]any{
					"id": Integer,
				}},
			}},
		}},
	},
	Librarian: {
		{"Books", "SELECT * FROM books",
			[]Action{
				{"Create book", []Query{{
					`INSERT INTO books (isbn, title, city, publishing_house, year, pages_quantity, price)
					VALUES (:isbn, :title, :city, :publishing_house, :year, :pages_quantity, :price)`,
					map[string]any{
						"isbn":             String,
						"title":            String,
						"city":             String,
						"publishing_house": String,
						"year":             Integer,
						"pages_quantity":   Integer,
						"price":            Float,
					},
				}}},
				{"Delete book", []Query{{
					"DELETE FROM books WHERE isbn = :isbn",
					map[string]any{
						"isbn": String,
					},
				}}},
			}},
		{"Book instances", "SELECT * FROM book_instances",
			[]Action{
				{"Delete book instance", []Query{{
					"DELETE FROM book_instances WHERE inventory_number = :inventory_number",
					map[string]any{
						"inventory_number": String,
					},
				}}},
				{"Check if book was replaced", []Query{{
					"SELECT EXISTS (SELECT * FROM replacement_acts WHERE old_inventory_number = :inventory_number)",
					map[string]any{
						"inventory_number": String,
					},
				}}},
				{"Create book instance", []Query{{
					`INSERT INTO book_instances (inventory_number, book_isbn, shelf) VALUES (:in, :isbn, :shelf)`,
					map[string]interface{}{
						"inventory_number": String,
						"isbn":             String,
						"shelf":            String,
					},
				}}},
			},
		},
		// MARK: BOOK AUTHORS ############################################################################################
		{"Book authors", "SELECT * FROM book_authors",
			[]Action{
				{"Create book author", []Query{{
					`INSERT INTO book_authors (book_isbn, surname)
					VALUES (:book_isbn, :surname)`,
					map[string]any{
						"book_isbn": String,
						"surname":   String,
					},
				}}},
				{"Delete book author", []Query{{
					"DELETE FROM book_authors WHERE id = :id",
					map[string]any{
						"id": Integer,
					},
				}}},
			},
		},
		// MARK: BOOK AREAS #############################################################################################
		{"Book areas", "SELECT * FROM book_areas",
			[]Action{
				{"Create book area", []Query{{
					`INSERT INTO book_areas (book_isbn, area_cipher)
					VALUES (:book_isbn, :area_cipher)`,
					map[string]any{
						"book_isbn":   String,
						"area_cipher": String,
					},
				}}},
				{"Delete book area", []Query{{
					"DELETE FROM book_areas WHERE book_isbn = :book_isbn AND area_cipher = :area_cipher",
					map[string]any{
						"book_isbn":   String,
						"area_cipher": String,
					},
				}}},
			},
		},
		// MARK: CHECKOUTS ##############################################################################################
		{"Checkouts", "SELECT * FROM checkouts",
			[]Action{
				{"Create checkout", []Query{{
					"SELECT f_raise('Reader has already taken that book')",
					map[string]any{},
				}, {
					`INSERT INTO checkouts (reader_card_number, book_inventory_number, checkout_date, expected_return_date)
					VALUES (:reader_card_number, :book_inventory_number, :checkout_date, :expected_return_date)`,
					map[string]any{
						"reader_card_number":    String,
						"book_inventory_number": String,
						"checkout_date":         Date,
						"expected_return_date":  Date,
					},
				}}},
				{"Complete checkout", []Query{{
					`UPDATE checkouts
					SET return_date = :return_date, repaid = :repaid
					WHERE checkout_number = :checkout_number`,
					map[string]any{
						"checkout_number": Integer,
						"return_date":     Date,
						"repaid":          Float,
					},
				}}},
				{"Delete checkout", []Query{{
					"DELETE FROM checkouts WHERE checkout_number = :checkout_number",
					map[string]any{
						"checkout_number": Integer,
					},
				}}},
				{"Get checkouts count for each reader", []Query{{
					`SELECT full_name, COUNT(*) AS cnt
					FROM (checkouts INNER JOIN readers r ON r.card_number = checkouts.reader_card_number)
					GROUP BY r.card_number`,
					map[string]any{},
				}}},
				{"Get checkouts quantity for books, that were taken more than once", []Query{{
					`SELECT title, COUNT(*) AS cnt
						FROM (checkouts INNER JOIN book_instances bi ON bi.inventory_number = checkouts.book_inventory_number
							INNER JOIN books b ON bi.book_isbn = b.isbn)
						GROUP BY isbn
						HAVING COUNT(*) > 1`,
					map[string]any{},
				}}},
			},
		},
		// MARK: KNOWLEDGE AREAS ##########################################################################################
		{"Knowledge areas", "SELECT * FROM knowledge_areas",
			[]Action{
				{"Create knowledge area", []Query{{
					`INSERT INTO knowledge_areas (cipher, title)
					VALUES (:cipher, :title)`,
					map[string]any{
						"cipher": String,
						"title":  String,
					},
				}}},
				{"Delete knowledge area", []Query{{
					"DELETE FROM knowledge_areas WHERE cipher = :cipher",
					map[string]any{
						"cipher": String,
					},
				}}},
				{"Find a knowledge area where all books are from replacement", []Query{{
					`SELECT cipher as area_cipher, title
						FROM (
							 book_areas INNER JOIN knowledge_areas ka ON book_areas.area_cipher = ka.cipher
						)
						WHERE NOT EXISTS(
							SELECT *
							FROM book_instances
							WHERE NOT EXISTS (
								SELECT * FROM replacement_acts
								WHERE replacement_acts.new_inventory_number = book_instances.inventory_number
							) AND book_instances.book_isbn = book_areas.book_isbn
						)`,
					map[string]any{},
				}}},
			},
		},
		// MARK: READERS ##################################################################################################
		{"Readers", "SELECT * FROM readers",
			[]Action{
				{"Create reader", []Query{{
					`INSERT INTO readers (card_number, full_name, home_address, seat, birth_date)
					VALUES (:card_number, :full_name, :home_address, :seat, :birth_date)`,
					map[string]any{
						"card_number":  String,
						"full_name":    String,
						"home_address": String,
						"seat":         String,
						"birth_date":   Date,
					},
				}}},
				{"Delete reader", []Query{{
					"DELETE FROM readers WHERE card_number = :card_number",
					map[string]any{
						"card_number": String,
					},
				}}},
				{"Get all books that some reader ever took", []Query{{
					`SELECT b.*
						FROM (
							checkouts INNER JOIN readers r ON r.card_number = checkouts.reader_card_number
							INNER JOIN book_instances bi ON checkouts.book_inventory_number = bi.inventory_number
							INNER JOIN books b ON bi.book_isbn = b.isbn
						)
						WHERE r.card_number = :card_number`,
					map[string]any{
						"card_number": String,
					},
				}}},
			},
		},
		// MARK: READERS PHONES ###########################################################################################
		{"Readers phones", "SELECT * FROM readers_phones",
			[]Action{
				{"Create reader phone", []Query{{
					`INSERT INTO readers_phones (reader_card, phone)
					VALUES (:reader_card, :phone)`,
					map[string]any{
						"reader_card": String,
						"phone":       String,
					},
				}}},
				{"Get all reader phones", []Query{{
					"SELECT * FROM readers_phones",
					map[string]any{},
				}}},
				{"Get reader phones", []Query{{
					"SELECT phone FROM readers_phones WHERE reader_card = :reader_card",
					map[string]any{
						"reader_card": String,
					},
				}}},
				{"Delete reader phone", []Query{{
					"DELETE FROM readers_phones WHERE reader_card = :reader_card",
					map[string]any{
						"reader_card": String,
					},
				}}},
			},
		},
		// MARK: REPLACEMENT ACTS #####################################################################################
		{"Replacement acts", "SELECT * FROM replacement_acts",
			[]Action{
				{"Create replacement act", []Query{{
					`INSERT INTO replacement_acts (replacement_date, old_inventory_number, new_inventory_number)
					VALUES (:replacement_date, :old_inventory_number, :new_inventory_number)`,
					map[string]any{
						"replacement_date":     Date,
						"old_inventory_number": String,
						"new_inventory_number": String,
					},
				}}},
				{"Get replacement act", []Query{{
					"SELECT * FROM replacement_acts WHERE id = :id",
					map[string]any{
						"id": Integer,
					},
				}}},
				{"Delete replacement act", []Query{{
					"DELETE FROM replacement_acts WHERE replacement_act_number = :replacement_act_number",
					map[string]any{
						"replacement_act_number": String,
					},
				}}},
			},
		},
	},
	// TODO: add debtors
	// - Отримувати відомості про читачів, що є боржниками бібліотеки, тобто не повернули вчасно примірники взятих книг.
	// - Отримувати відомості про вартість конкретної книги; це необхідно для того, щоб встановити можливість повернення вартості загубленої читачем книги або можливість заміни її іншою книгою.
	Admin: {
		{},
	},
}
