package core

import "strings"

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
	LowPref    string `json:"lowpref"`
}

// NewPerson constructs a Person
func NewPerson(personID int, fullName string, alias string, team string, discipline string, seniority string, gender string, active bool, lowPref string) *Person {

	return &Person{
		ID:         personID,
		FullName:   fullName,
		Alias:      alias,
		Team:       team,
		Discipline: discipline,
		Seniority:  seniority,
		Gender:     gender,
		Active:     active,
		LowPref:    lowPref,
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
	score := 0
	if personScoreIsFor.Seniority != p.Seniority {
		score += 10
	}

	if !isInSameTeam(personScoreIsFor.Team, p.Team) {
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

	score := 0
	if (p.LowPref == personScoreIsFor.Alias) {
			return -1000;
	}
	if isInSameTeam(personScoreIsFor.Team, p.Team) {
		score += 12
	}
	
	if personScoreIsFor.Seniority != p.Seniority {
		score += 5
	}

	if personScoreIsFor.Discipline == p.Discipline {
		score += 4
	}

	if personScoreIsFor.Gender == p.Gender {
		score += 2
	}


	return score
}

func isInSameTeam(teams1 string, teams2 string) bool {
	// teams is a delimited by " "
	teams1Split := strings.Fields(teams1)
	teams2Split := strings.Fields(teams2)

	for _, team1 := range teams1Split {
		for _, team2 := range teams2Split {
			if team1 == team2 {
				return true
			}
		}
	}

	return false
}
