package app

type QueryMessage struct {
	Name    string  `json:"queryName"`
	Queries []Query `json:"queries"`
}

type Query struct {
	Text   string         `json:"text"`
	Params map[string]any `json:"params"`
}

var queries = []QueryMessage{
	{"GetAllUsers", []Query{{
		"SELECT * FROM users",
		map[string]any{},
	}}},
}
