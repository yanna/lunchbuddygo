package matching

import (
	"fmt"
)

// StableMarriage contains the data to implement the Gale-Shapley Stable Marriage algorithm.
// The Stable Marriage Problem states that given N men and N women, where each person has ranked all members of the opposite sex in order of preference,
// marry the men and women together such that there are no two people of opposite sex who would both rather have each other than their current partners.
// If there are no such people, all the marriages are “stable”.
//
// Note that the two groups are called male and female for the purposes of the algorithm only.
// It is not necessarily the case that everyone in female group is a female etc.
type StableMarriage struct {
	femalePrefs map[string][]string
	malePrefs   map[string][]string
}

// NewStableMarriage creates an object that contains the preferences of the two groups of people
func NewStableMarriage(femalePrefs map[string][]string, malePrefs map[string][]string) (*StableMarriage, error) {
	femaleCount := len(femalePrefs)
	maleCount := len(malePrefs)
	if femaleCount != maleCount {
		return nil, fmt.Errorf("Number of females (%d) and males (%d) are not the same", femaleCount, maleCount)
	}

	newObj := StableMarriage{
		femalePrefs: femalePrefs,
		malePrefs:   malePrefs}

	return &newObj, nil
}

// CreateStablePairs provides a pairing between the two groups based on the preferences. Returns a map where key and value are aliases.
func (algo *StableMarriage) CreateStablePairs() map[string]string {
	/*Gale-Shapley Stable Marriage algorithm
		https://en.wikipedia.org/wiki/Stable_marriage_problem
		function stableMatching {
			Initialize all m ∈ M and w ∈ W to free
			while ∃ free man m who still has a woman w to propose to {
			w = first woman on m’s list to whom m has not yet proposed
			if w is free
				(m, w) become engaged
			else some pair (m', w) already exists
				if w prefers m to m'
					m' becomes free
				(m, w) become engaged
				else
				(m', w) remain engaged
	    }
	*/
	matchData := newMatchData(algo.femalePrefs, algo.malePrefs)

	freeMalesCount := len(algo.malePrefs)
	for freeMalesCount > 0 {
		maleName, _ := matchData.getFreeMaleName()
		femaleName := matchData.removeMalesFirstPref(maleName)
		if matchData.isPersonFree(femaleName) {
			matchData.engage(maleName, femaleName)
			freeMalesCount--
		} else {
			currentMaleName := matchData.getMatch(femaleName)
			if matchData.doesFemalePreferMaleXOverY(femaleName, maleName, currentMaleName) {
				matchData.breakup(femaleName)
				matchData.engage(maleName, femaleName)
			}
		}
	}

	return makeMapWithKeys(matchData.matches, matchData.maleNames)
}

func makeMapWithKeys(dict map[string]string, keysToKeep []string) map[string]string {
	output := make(map[string]string)
	for _, key := range keysToKeep {
		if val, ok := dict[key]; ok {
			output[key] = val
		}
	}
	return output
}
