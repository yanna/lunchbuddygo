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
	Active     bool   `json:"active"`
}

// NewPerson constructs a Person
func NewPerson(personID int, fullName string, alias string, team string, discipline string, seniority string, gender string, active bool) *Person {

	return &Person{
		ID:         personID,
		FullName:   fullName,
		Alias:      alias,
		Team:       team,
		Discipline: discipline,
		Seniority:  seniority,
		Gender:     gender,
		Active:     active,
	}
}

// GetScore returns the score of the current person with relation to the input person
func (p *Person) GetScore(personScoreIsFor *Person) int {
	// Most important differences
	// 1. Seniority (we don't want Senior leadership ppl to be grouped together)
	// 2. Gender
	// 3. Team
	// 4. Discipline
	// Higher score is better
	score := 0

	if personScoreIsFor.Seniority != p.Seniority {
		score += 10
	}

	if personScoreIsFor.Gender != p.Gender {
		score += 8
	}

	if personScoreIsFor.Team != p.Team {
		score += 4
	}

	if personScoreIsFor.Discipline != p.Discipline {
		score += 4
	}

	return score
}
