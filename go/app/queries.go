package app

const (
	String  = "string"
	Integer = "integer"
	Float   = "float"
	Date    = "date"
)

type Action struct {
	Name    string  `json:"queryName"`
	Queries []Query `json:"queries"`
}

type Query struct {
	Text   string         `json:"text"`
	Params map[string]any `json:"params"`
}

type Entity struct {
	EntityName string   `json:"entityName"`
	BasicQuery string   `json:"basicQuery"`
	Actions    []Action `json:"actions"`
}

var entities = []Entity{
	// MARK: BOOKS ###################################################################################################
	{"Books", "SELECT * FROM BOOKS",
		[]Action{
			{"Create book", []Query{{
				"INSERT INTO books (isbn, title, city, publishing_house, year, pages_quantity, price) " +
					"VALUES (:isbn, :title, :city, :publishing_house, :year, :pages_quantity, :price)",
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
		},
	},
	// MARK: BOOK INSTANCES #########################################################################################
	{"Book instances", "SELECT * FROM book_instances",
		[]Action{
			{"Create book instance", []Query{{
				"INSERT INTO book_instances (inventory_number, book_isbn, shelf) " +
					"VALUES (:inventory_number, :book_isbn, :shelf)",
				map[string]any{
					"inventory_number": String,
					"book_isbn":        String,
					"shelf":            String,
				},
			}}},
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
		},
	},
	// MARK: BOOK AUTHORS ############################################################################################
	{"Book authors", "SELECT * FROM book_authors",
		[]Action{
			{"Create book author", []Query{{
				"INSERT INTO book_authors (book_isbn, surname) " +
					"VALUES (:book_isbn, :surname)",
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
				"INSERT INTO book_areas (book_isbn, area_cipher) " +
					"VALUES (:book_isbn, :area_cipher)",
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
				"INSERT INTO checkouts (checkout_date, expected_return_date) " +
					"VALUES (:checkout_date, :expected_return_date)",
				map[string]any{
					"checkout_date":        Date,
					"expected_return_date": Date,
				},
			}}},
			{"Complete checkout", []Query{{
				"UPDATE checkouts " +
					"SET return_date = :return_date, repaid = :repaid " +
					"WHERE checkout_number = :checkout_number",
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
		},
	},
	// MARK: KNOWLEDGE AREAS ##########################################################################################
	{"Knowledge areas", "SELECT * FROM knowledge_areas",
		[]Action{
			{"Create knowledge area", []Query{{
				"INSERT INTO knowledge_areas (cipher, title) " +
					"VALUES (:cipher, :title)",
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
		},
	},
	// MARK: READERS ##################################################################################################
	{"Readers", "SELECT * FROM readers",
		[]Action{
			{"Create reader", []Query{{
				"INSERT INTO readers (card_number, full_name, home_address, seat, birth_date) " +
					"VALUES (:card_number, :full_name, :home_address, :seat, :birth_date)",
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
		},
	},
	// MARK: READERS PHONES ###########################################################################################
	{"Readers phones", "SELECT * FROM readers_phones",
		[]Action{
			{"Create reader phone", []Query{{
				"INSERT INTO readers_phones (reader_card, phone) " +
					"VALUES (:reader_card, :phone)",
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
				"INSERT INTO replacement_acts (replacement_date, old_inventory_number, new_inventory_number) " +
					"VALUES (:replacement_date, :old_inventory_number, :new_inventory_number)",
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
}
