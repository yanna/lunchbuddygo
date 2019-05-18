package core

import (
	"sort"
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
	group1, group2, oddPerson := pm.SplitOptedInPeopleIntoTwoGroups()
	group1Prefs = pm.getPreferencesForFirstGroup(group1, group2)
	group2Prefs = pm.getPreferencesForFirstGroup(group2, group1)
	return
}

func (pm *PeopleMatches) getPreferencesForFirstGroup(group1 []Person, group2 []Person) map[string][]string {

	group1Prefs := make(map[string][]string)

	group2IDs := getPersonIDsFromPeople(group2)
	group2IDsSet := NewSetFromInts(group2IDs)

	for _, person := range group1 {
		matches := pm.GetAllMatchesByPersonID(person.ID)
		matchedPersonIds := getPersonIDsFromMatches(matches)

		matchedPersonIdsInGroup2 := group2IDsSet.Intersect(NewSetFromInts(matchedPersonIds))
		matchedPersonIdsInGroup2Sorted := pm.SortMatchesByDate(person.ID, matchedPersonIdsInGroup2.ToSlice())

		nonMatchedInGroup2 := group2IDsSet.Difference(matchedPersonIdsInGroup2).ToSlice()
		nonMatchedInGroup2Sorted := pm.sortInOrderOfPreference(person.ID, nonMatchedInGroup2)

		prefPersonIds := []int{}
		prefPersonIds = append(prefPersonIds, nonMatchedInGroup2Sorted...)
		prefPersonIds = append(prefPersonIds, matchedPersonIdsInGroup2Sorted...)

		group1Prefs[person.Alias] = pm.GetAliases(prefPersonIds)
	}

	return group1Prefs
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
