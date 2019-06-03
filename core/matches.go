package core

import "sort"

// Matches stores the past matches for each person
type Matches struct {
	PersonIDToMatches map[int][]Match
}

// GetAllMatchesByPersonID returns previous matches
func (m *Matches) GetAllMatchesByPersonID(personID int) []Match {
	matches := m.PersonIDToMatches[personID]
	return matches
}

// SortMatchesByDate expects a subset of the person's matched ids and returns them sorted by date
// of the match. The more recent the match, the later it is in the returned list.
func (m *Matches) SortMatchesByDate(personID int, matchedPersonIDs []int) []int {
	allMatches := m.GetAllMatchesByPersonID(personID)

	// Sort by oldest first. The data is assumed to be in a sortable format
	sort.Slice(allMatches, func(i, j int) bool {
		return allMatches[i].Date < allMatches[j].Date
	})

	sortedPersonIDs := []int{}
	for _, match := range allMatches {
		if isIntInSlice(match.PersonID, matchedPersonIDs) {
			sortedPersonIDs = append(sortedPersonIDs, match.PersonID)
		}
	}
	return sortedPersonIDs
}

// HaveBeenMatched returns true if personID1 have ever been matched with personID2
func (m *Matches) HaveBeenMatched(personID1 int, personID2 int) bool {
	allMatches := m.GetAllMatchesByPersonID(personID1)

	haveBeenMatched := false
	for _, match := range allMatches {
		if personID2 == match.PersonID {
			haveBeenMatched = true
			break
		}
	}

	return haveBeenMatched
}
