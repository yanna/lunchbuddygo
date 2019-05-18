package core

import (
	"errors"
	"math/rand"
	"time"
)

// People groups Person objects
type People struct {
	idToPerson map[int]Person
}

func NewPeople(people []Person) *People {
	idToPerson := make(map[int]Person)
	for _, person := range people {
		idToPerson[person.ID] = person
	}
	return &People{
		idToPerson: idToPerson,
	}
}

// GetPersonIDByAlias returns the id of the person
func (m *People) GetPersonIDByAlias(alias string) (int, error) {
	for _, person := range m.idToPerson {
		if person.Alias == alias {
			return person.ID, nil
		}
	}

	return -1, errors.New("Can't find person with alias " + alias)
}

// GetOptedIn returns all the people who are opted in at the moment
func (m *People) GetOptedIn() []Person {
	var optedIn []Person
	for _, person := range m.idToPerson {
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
	for _, personID := range personIDs {
		person := m.idToPerson[personID]
		result = append(result, person.Alias)
	}
	return result
}
