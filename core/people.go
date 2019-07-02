package core

import (
	"errors"
	"math/rand"
	"sort"
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
func (p *People) GetPerson(personID int) Person {
	return p.idToPerson[personID]
}

//GetPersonByAlias returns a person based on an alias
func (p *People) GetPersonByAlias(alias string) (Person, error) {
	personID, err := p.GetPersonIDByAlias(alias)
	if err != nil {
		return Person{}, err
	}

	return p.GetPerson(personID), nil
}

// GetPersonIDByAlias returns the id of the person
func (p *People) GetPersonIDByAlias(alias string) (int, error) {
	for _, person := range p.idToPerson {
		if person.Alias == alias {
			return person.ID, nil
		}
	}

	return -1, errors.New("Can't find person with alias " + alias)
}

// GetActivePeople returns all the people who are actively available for matching
func (p *People) GetActivePeople() []Person {
	var activeUsers []Person
	for _, person := range p.idToPerson {
		if person.Active {
			activeUsers = append(activeUsers, person)
		}
	}
	return activeUsers
}

//SplitActivePeopleIntoTwoRandomGroups separates the opted in people into two groups randomly.
//It is possible to have an odd person
func (p *People) SplitActivePeopleIntoTwoRandomGroups() (group1 []Person, group2 []Person, oddPerson *Person) {
	activePeople := p.GetActivePeople()

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
func (p *People) GetAliases(personIDs []int) []string {
	var result []string
	for _, personID := range personIDs {
		person := p.idToPerson[personID]
		result = append(result, person.Alias)
	}
	return result
}

//GetSortedIDs returns the list of personIDs sorted from lowest to highest
func (p *People) GetSortedIDs() []int {
	var personIDs []int
	for personID := range p.idToPerson {
		personIDs = append(personIDs, personID)
	}
	sort.Ints(personIDs)
	return personIDs
}
