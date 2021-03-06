package core

import (
	"fmt"
	"math/rand"
	"sort"
)

//PersonTuple represents a pair of Person objects
type PersonTuple struct {
	person1 Person
	person2 Person
}

type nameSorter []PersonTuple

func (a nameSorter) Len() int           { return len(a) }
func (a nameSorter) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a nameSorter) Less(i, j int) bool { return a[i].person1.FullName < a[j].person1.FullName }

// MatchesOutput is reponsible for printing all the match output
type MatchesOutput struct {
	*PeopleMatches
	matches   []PersonTuple
	oddPerson *Person
}

// NewMatchesOutput creats MatchesOutput
func NewMatchesOutput(aliasMatches map[string]string, peopleMatches *PeopleMatches, oddPerson *Person) *MatchesOutput {
	var matches []PersonTuple
	for group1Alias, group2Alias := range aliasMatches {
		person1, _ := peopleMatches.GetPersonByAlias(group1Alias)
		person2, _ := peopleMatches.GetPersonByAlias(group2Alias)
		matches = append(matches, PersonTuple{
			person1: person1,
			person2: person2,
		})
	}

	sort.Sort(nameSorter(matches))

	return &MatchesOutput{
		matches:       matches,
		PeopleMatches: peopleMatches,
		oddPerson:     oddPerson,
	}
}

// Print prints the matching output
func (m *MatchesOutput) Print() {
	fmt.Println()
	m.printFullNames()
	fmt.Println("\nAliases: ")
	m.printAliases()
	fmt.Println("\nExcel column: ")
	m.printExcelColumn()
}

//PrintFullNames prints the full names of the matches
func (m *MatchesOutput) printFullNames() {

	indexToMatchOddPerson := -1

	if m.oddPerson != nil {
		foundMatch := false
		for i := 0; i < 1000; i++ {
			indexToMatchOddPerson = rand.Intn(len(m.matches))
			pairToJoin := m.matches[indexToMatchOddPerson]
			if m.HaveBeenMatched(pairToJoin.person1.ID, m.oddPerson.ID) {
				continue
			}

			if m.HaveBeenMatched(pairToJoin.person2.ID, m.oddPerson.ID) {
				continue
			}

			foundMatch = true
			break
		}

		if !foundMatch {
			fmt.Print("Error: after many iterations still didn't find a match for the odd person")
		}
	}

	for i, personTuple := range m.matches {
		p1 := personTuple.person1
		p2 := personTuple.person2
		fmt.Print(p1.FullName + " and " + p2.FullName)

		if m.HaveBeenMatched(p1.ID, p2.ID) {
			fmt.Print(" <--- Previously matched!! **********")
		}

		if indexToMatchOddPerson == i {
			fmt.Print(" and " + m.oddPerson.FullName)
		}

		fmt.Println()
	}
}

// PrintExcelColumn prints new column to enter
func (m *MatchesOutput) printExcelColumn() {
	for _, personID := range m.GetSortedIDs() {
		personToFindMatchFor := m.GetPerson(personID)
		//fmt.Print(personID)
		//fmt.Print(personToFindMatchFor.FullName + "\t")

		if !personToFindMatchFor.Active {
			fmt.Println("---") //Just printing "" caused some rows to be missing?! I don't really understand.
			continue
		}

		if personMatch, err := m.getMatchedPerson(personToFindMatchFor); err == nil {
			fmt.Println(personMatch.Alias)
		} else {
			fmt.Println("ERROR!" + err.Error())
		}
	}
}

func (m *MatchesOutput) printAliases() {
	for _, personTuple := range m.matches {
		p1 := personTuple.person1
		p2 := personTuple.person2
		fmt.Print(p1.Alias + " and " + p2.Alias)
		fmt.Println()
	}
}

func (m *MatchesOutput) getMatchedPerson(personToFindMatchFor Person) (Person, error) {

	for _, personTuple := range m.matches {
		p1 := personTuple.person1
		p2 := personTuple.person2
		if personToFindMatchFor.ID == p1.ID {
			return p2, nil
		}

		if personToFindMatchFor.ID == p2.ID {
			return p1, nil
		}
	}

	return Person{}, fmt.Errorf("Can't find person named %s with id %d", personToFindMatchFor.FullName, personToFindMatchFor.ID)
}
