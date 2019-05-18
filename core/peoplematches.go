package core

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

		matchedPersonIdsInOtherGroup := group2IDsSet.Intersect(NewSetFromInts(matchedPersonIds))
		matchedPersonIdsInOtherGroupSorted := pm.SortMatchesByDate(person.ID, matchedPersonIdsInOtherGroup.ToSlice())

		nonMatchedInGroup2 := group2IDsSet.Difference(matchedPersonIdsInOtherGroup).ToSlice()

		prefPersonIds := []int{}
		prefPersonIds = append(prefPersonIds, nonMatchedInGroup2...)
		prefPersonIds = append(prefPersonIds, matchedPersonIdsInOtherGroupSorted...)

		personPrefs := pm.GetAliases(prefPersonIds)
		group1Prefs[person.Alias] = personPrefs
	}

	return group1Prefs
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
