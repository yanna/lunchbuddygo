package core

import (
	"sort"
	"sync"
)

// PeopleMatches contains the people and past matches
type PeopleMatches struct {
	People
	Matches
}

//GetPreferences splits the opted in people into two groups randomly and
//returns their preferences of the other group.
//If there is an odd number of people, an odd person may also be returned
func (pm *PeopleMatches) GetPreferences() (group1Prefs map[string][]string, group2Prefs map[string][]string, oddPerson *Person) {
	group1, group2, oddPerson := pm.SplitActivePeopleIntoTwoRandomGroups()
	group1Prefs = pm.getPreferencesForFirstGroup(group1, group2)
	group2Prefs = pm.getPreferencesForFirstGroup(group2, group1)
	return
}

//GetPreferencesParallel splits the opted in people into two groups randomly and
//returns their preferences of the other group using goroutines
//If there is an odd number of people, an odd person may also be returned
func (pm *PeopleMatches) GetPreferencesParallel() (group1Prefs map[string][]string, group2Prefs map[string][]string, oddPerson *Person) {
	group1, group2, oddPerson := pm.SplitActivePeopleIntoTwoRandomGroups()

	group1PrefsChannel := make(chan map[string][]string)
	go pm.getPreferencesForFirstGroupParallel(group1, group2, group1PrefsChannel)

	group2PrefsChannel := make(chan map[string][]string)
	go pm.getPreferencesForFirstGroupParallel(group2, group1, group2PrefsChannel)

	group1Prefs = <-group1PrefsChannel
	group2Prefs = <-group2PrefsChannel

	return
}

func (pm *PeopleMatches) getPreferencesForFirstGroup(group1 []Person, group2 []Person) map[string][]string {

	// Key: Alias of the person
	// Value: List of aliases in preference order for that person
	group1Prefs := make(map[string][]string)

	group2IDs := getPersonIDsFromPeople(group2)
	group2IDsSet := NewSetFromInts(group2IDs)

	for _, person := range group1 {
		group1Prefs[person.Alias] = pm.getPreferencesForPerson(person.ID, group2IDsSet)
	}

	return group1Prefs
}

func (pm *PeopleMatches) getPreferencesForFirstGroupParallel(group1 []Person, group2 []Person, group1PrefsChannel chan map[string][]string) {

	// Key: Alias of the person
	// Value: List of aliases in preference order for that person
	group1Prefs := make(map[string][]string)

	group2IDs := getPersonIDsFromPeople(group2)
	group2IDsSet := NewSetFromInts(group2IDs)

	group1Len := len(group1)

	var wg sync.WaitGroup
	wg.Add(group1Len)

	for i := 0; i < group1Len; i++ {
		go func(i int) {
			defer wg.Done()
			person := group1[i]
			group1Prefs[person.Alias] = pm.getPreferencesForPerson(person.ID, group2IDsSet)
		}(i)
	}

	wg.Wait()

	group1PrefsChannel <- group1Prefs
}

func (pm *PeopleMatches) getPreferencesForPerson(personID int, inputPersonIDs *Set) []string {
	matches := pm.GetAllMatchesByPersonID(personID)
	matchedPersonIds := getPersonIDsFromMatches(matches)

	matchedPersonIdsInInput := inputPersonIDs.Intersect(NewSetFromInts(matchedPersonIds))
	matchedPersonIdsInInputSorted := pm.SortMatchesByDate(personID, matchedPersonIdsInInput.ToSlice())

	nonMatchedInInput := inputPersonIDs.Difference(matchedPersonIdsInInput).ToSlice()
	nonMatchedInInputSorted := pm.sortInOrderOfPreference(personID, nonMatchedInInput)

	prefPersonIds := []int{}
	prefPersonIds = append(prefPersonIds, nonMatchedInInputSorted...)
	prefPersonIds = append(prefPersonIds, matchedPersonIdsInInputSorted...)

	return pm.GetAliases(prefPersonIds)
}

func (pm *PeopleMatches) sortInOrderOfPreference(personID int, personIDsToSort []int) []int {
	person := pm.GetPerson(personID)
	personScores := make([]PersonScore, len(personIDsToSort))
	for i, personIDToSort := range personIDsToSort {
		personToSort := pm.GetPerson(personIDToSort)
		personScores[i] = PersonScore{
			personID: personIDToSort,
			score:    personToSort.GetScore(&person),
		}
	}

	// Sort based on rank. highest first.
	sort.Slice(personScores, func(i, j int) bool {
		return personScores[i].score > personScores[j].score
	})

	//fmt.Println(personScores)

	return getPersonIDsFromPersonScores(personScores)
}

func getPersonIDsFromPeople(people []Person) []int {
	personIDs := make([]int, len(people))
	for i, person := range people {
		personIDs[i] = person.ID
	}
	return personIDs
}

func getPersonIDsFromMatches(matches []Match) []int {
	personIDs := make([]int, len(matches))
	for i, match := range matches {
		personIDs[i] = match.PersonID
	}
	return personIDs
}

func getPersonIDsFromPersonScores(personScores []PersonScore) []int {
	personIDs := make([]int, len(personScores))
	for i, personScore := range personScores {
		personIDs[i] = personScore.personID
	}
	return personIDs
}
