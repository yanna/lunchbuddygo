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

// NewPeople creates a new People object
func NewPeople(people []Person) *People {
	idToPerson := make(map[int]Person)
	for _, person := range people {
		idToPerson[person.ID] = person
	}
	return &People{
		idToPerson: idToPerson,
	}
}

// GetPerson returns a person based on the id
func (m *People) GetPerson(personID int) Person {
	return m.idToPerson[personID]
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

// GetActivePeople returns all the people who are actively available for matching
func (m *People) GetActivePeople() []Person {
	var activeUsers []Person
	for _, person := range m.idToPerson {
		if person.Active {
			activeUsers = append(activeUsers, person)
		}
	}
	return activeUsers
}

//SplitOptedInPeopleIntoTwoGroups separates the opted in people into two groups randomly.
//It is possible to have an odd person
func (m *People) SplitOptedInPeopleIntoTwoGroups() (group1 []Person, group2 []Person, oddPerson *Person) {
	activePeople := m.GetActivePeople()

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(
		len(activePeople),
		func(i, j int) {
			activePeople[i], activePeople[j] = activePeople[j], activePeople[i]
		})

	peopleLength := len(activePeople)
	midIndex := peopleLength / 2
	group1 = activePeople[:midIndex]
	group2 = activePeople[midIndex : midIndex*2]
	oddPerson = nil
	if peopleLength%2 != 0 {
		oddPerson = &activePeople[len(activePeople)-1]
	}
	return
}

// GetAliases returns the aliases representing the ids in the order of the ids
func (m *People) GetAliases(personIDs []int) []string {
	var result []string
	for _, personID := range personIDs {
		person := m.idToPerson[personID]
		result = append(result, person.Alias)
	}
	return result
}
