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
func (p *Person) GetScore(personScoreIsFor *Person, matchMode MatchMode) int {
	if matchMode == Similar {
		return p.GetScoreForMostSimilarMatch(personScoreIsFor)
	}

	return p.GetScoreForMostDifferentMatch(personScoreIsFor)

}

// GetScore returns the score of the current person with relation to the input person
// that will yield a higher score for more differences
func (p *Person) GetScoreForMostDifferentMatch(personScoreIsFor *Person) int {
	// TODO: support multiple teams

	score := 0
	if personScoreIsFor.Seniority != p.Seniority {
		score += 10
	}

	if personScoreIsFor.Team != p.Team {
		score += 8
	}

	if personScoreIsFor.Discipline != p.Discipline {
		score += 6
	}

	if personScoreIsFor.Gender != p.Gender {
		score += 4
	}

	return score
}

// GetScore returns the score of the current person with relation to the input person
// that will yield a higher score for more similarities
func (p *Person) GetScoreForMostSimilarMatch(personScoreIsFor *Person) int {
	// TODO: support multiple teams

	score := 0
	if personScoreIsFor.Team == p.Team {
		score += 16
	}

	if personScoreIsFor.Discipline == p.Discipline {
		score += 8
	}

	if personScoreIsFor.Gender == p.Gender {
		score += 4
	}

	if personScoreIsFor.Seniority == p.Seniority {
		score += 2
	}

	return score
}
