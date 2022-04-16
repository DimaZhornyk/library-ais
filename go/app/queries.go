package app

const (
	String  = "string"
	Integer = "integer"
	Float   = "float"
	Date    = "date"
)

type QueryMessage struct {
	Name    string  `json:"queryName"`
	Queries []Query `json:"queries"`
}

type Query struct {
	Text   string         `json:"text"`
	Params map[string]any `json:"params"`
}

var queries = []QueryMessage{
	// MARK: BOOKS ###################################################################################################
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
	{"Get all books", []Query{{
		"SELECT * FROM books",
		map[string]any{},
	}}},
	{"Delete book", []Query{{
		"DELETE FROM books WHERE isbn = :isbn",
		map[string]any{
			"isbn": String,
		},
	}}},

	// MARK: BOOK INSTANCES #########################################################################################
	{"Create book instance", []Query{{
		"INSERT INTO book_instances (inventory_number, book_isbn, shelf) " +
			"VALUES (:inventory_number, :book_isbn, :shelf)",
		map[string]any{
			"inventory_number": Integer,
			"book_isbn":        String,
			"shelf":            String,
		},
	}}},
	{"Get all book instances", []Query{{
		"SELECT * FROM book_instances",
		map[string]any{},
	}}},
	{"Delete book instance", []Query{{
		"DELETE FROM book_instances WHERE inventory_number = :inventory_number",
		map[string]any{
			"inventory_number": Integer,
		},
	}}},

	// MARK: BOOK AUTHORS ############################################################################################
	{"Create book author", []Query{{
		"INSERT INTO book_authors (book_isbn, surname) " +
			"VALUES (:book_isbn, :surname)",
		map[string]any{
			"book_isbn": String,
			"surname":   String,
		},
	}}},
	{"Get all book authors", []Query{{
		"SELECT * FROM book_authors",
		map[string]any{},
	}}},
	{"Delete book author", []Query{{
		"DELETE FROM book_authors WHERE id = :id",
		map[string]any{
			"id": Integer,
		},
	}}},

	// MARK: BOOK AREAS #############################################################################################
	{"Create book area", []Query{{
		"INSERT INTO book_areas (book_isbn, area_cipher) " +
			"VALUES (:book_isbn, :area_cipher)",
		map[string]any{
			"book_isbn":   String,
			"area_cipher": String,
		},
	}}},
	{"Get all book areas", []Query{{
		"SELECT * FROM book_areas",
		map[string]any{},
	}}},
	{"Delete book area", []Query{{
		"DELETE FROM book_areas WHERE book_isbn = :book_isbn AND area_cipher = :area_cipher",
		map[string]any{
			"book_isbn":   String,
			"area_cipher": String,
		},
	}}},

	// MARK: CHECKOUTS ##############################################################################################
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
	{"Get all checkouts", []Query{{
		"SELECT * FROM checkouts",
		map[string]any{},
	}}},
	{"Delete checkout", []Query{{
		"DELETE FROM checkouts WHERE checkout_number = :checkout_number",
		map[string]any{
			"checkout_number": Integer,
		},
	}}},

	// MARK: KNOWLEDGE AREAS ##########################################################################################
	{"Create knowledge area", []Query{{
		"INSERT INTO knowledge_areas (cipher, title) " +
			"VALUES (:cipher, :title)",
		map[string]any{
			"cipher": String,
			"title":  String,
		},
	}}},
	{"Get all knowledge areas", []Query{{
		"SELECT * FROM knowledge_areas",
		map[string]any{},
	}}},
	{"Delete knowledge area", []Query{{
		"DELETE FROM knowledge_areas WHERE cipher = :cipher",
		map[string]any{
			"cipher": String,
		},
	}}},

	// MARK: READERS ##################################################################################################
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
	{"Get all readers", []Query{{
		"SELECT * FROM readers",
		map[string]any{},
	}}},
	{"Delete reader", []Query{{
		"DELETE FROM readers WHERE card_number = :card_number",
		map[string]any{
			"card_number": String,
		},
	}}},

	// MARK: READERS_PHONES ###########################################################################################
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
}
