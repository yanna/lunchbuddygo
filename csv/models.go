package csv

type person struct {
	ID          int     `json:"id"`
	FullName    string  `json:"fullname"`
	Alias       string  `json:"alias"`
	Team        string  `json:"team"`
	Discipline  string  `json:"discipline"`
	OptIn       string  `json:"optin"`
	PastMatches []match `json:"past_matches"`
}

type match struct {
	Date  string `json:"date"`
	Alias string `json:"alias"`
}
