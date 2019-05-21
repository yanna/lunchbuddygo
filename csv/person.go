package csv

type person struct {
	ID          int     `json:"id"`
	FullName    string  `json:"fullname"`
	Alias       string  `json:"alias"`
	Team        string  `json:"team"`
	Discipline  string  `json:"discipline"`
	Seniority   string  `json:"seniority"`
	Gender      string  `json:"gender"`
	Active      bool    `json:"active"`
	PastMatches []match `json:"past_matches"`
}
