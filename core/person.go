package core

// Person - Information about a person
type Person struct {
	ID         int    `json:"id"`
	FullName   string `json:"fullname"`
	Alias      string `json:"alias"`
	Team       string `json:"team"`
	Discipline string `json:"discipline"`
	Seniority  string `json:"seniority"`
	Gender     string `json:"gender"`
	OptIn      bool   `json:"optin"`
}

// NewPerson constructs a Person
func NewPerson(personID int, fullName string, alias string, team string, discipline string, seniority string, gender string, optIn bool) *Person {

	return &Person{
		ID:         personID,
		FullName:   fullName,
		Alias:      alias,
		Team:       team,
		Discipline: discipline,
		Seniority:  seniority,
		Gender:     gender,
		OptIn:      optIn,
	}
}
