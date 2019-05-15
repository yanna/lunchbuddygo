package core

import "errors"

// Person - Information about a person
type Person struct {
	ID         int    `json:"id"`
	FullName   string `json:"fullname"`
	Alias      string `json:"alias"`
	Team       string `json:"team"`
	Discipline string `json:"discipline"`
	OptIn      string `json:"optin"`
}

// NewPerson constructs a Person
func NewPerson(personID int, fullName string, alias string, team string, discipline string, optIn string) *Person {

	return &Person{
		ID:         personID,
		FullName:   fullName,
		Alias:      alias,
		Team:       team,
		Discipline: discipline,
		OptIn:      optIn,
	}
}

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

// PeopleMatches contains the people and past matches
type PeopleMatches struct {
	People
	Matches
}

// People groups Person objects
type People struct {
	Data []Person `json:"people"`
}

// GetPersonIDByAlias returns the id of the person
func (m *People) GetPersonIDByAlias(alias string) (int, error) {
	for _, person := range m.Data {
		if person.Alias == alias {
			return person.ID, nil
		}
	}

	return -1, errors.New("Can't find person with alias " + alias)
}

// Matches stores the past matches for each person
type Matches struct {
	PersonIDToMatches map[int][]Match
}
