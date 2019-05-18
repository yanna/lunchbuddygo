package core

// Match - Information about a match
type Match struct {
	// In the form yyyymm e.g. 201901
	Date     string `json:"date"`
	Alias    string `json:"alias"`
	PersonID int    `json:"id"`
}

// NewMatch constructs a Match
func NewMatch(date string, alias string, personID int) *Match {
	return &Match{
		Date:     date,
		Alias:    alias,
		PersonID: personID,
	}
}
