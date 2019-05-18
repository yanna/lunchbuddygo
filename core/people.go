package core

import (
	"errors"
	"math/rand"
	"time"
)

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

// GetOptedIn returns all the people who are opted in at the moment
func (m *People) GetOptedIn() []Person {
	var optedIn []Person
	for _, person := range m.Data {
		if person.OptIn {
			optedIn = append(optedIn, person)
		}
	}
	return optedIn
}

//SplitOptedInPeopleIntoTwoGroups separates the opted in people into two groups randomly.
//It is possible to have an odd person
func (m *People) SplitOptedInPeopleIntoTwoGroups() (group1 []Person, group2 []Person, oddPerson *Person) {
	optedInPeople := m.GetOptedIn()

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(
		len(optedInPeople),
		func(i, j int) {
			optedInPeople[i], optedInPeople[j] = optedInPeople[j], optedInPeople[i]
		})

	peopleLength := len(optedInPeople)
	midIndex := peopleLength / 2
	group1 = optedInPeople[:midIndex]
	group2 = optedInPeople[midIndex : midIndex*2]
	oddPerson = nil
	if peopleLength%2 != 0 {
		oddPerson = &optedInPeople[len(optedInPeople)-1]
	}
	return
}

// GetAliases returns the aliases representing the ids
func (m *People) GetAliases(personIDs []int) []string {
	var result []string
	for _, person := range m.Data {
		if isIntInSlice(person.ID, personIDs) {
			result = append(result, person.Alias)
		}
	}
	return result
}
