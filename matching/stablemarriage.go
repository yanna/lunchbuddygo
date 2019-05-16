package matching

import (
	"errors"
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

type matchData struct {
	// FemaleName -> [male1, male2...] where male1 is the most preferred of the opposite group
	femalePrefs map[string][]string
	// MaleName -> [female1, female2...] where female1 is the most preferred of the opposite group
	malePrefs map[string][]string
	// Name -> Name
	matches   map[string]string
	maleNames []string
}

func newMatchData(femalePrefs map[string][]string, malesPrefs map[string][]string) *matchData {
	// The algorithm will modify the females and males maps so need to do a copy
	return &matchData{
		femalePrefs: copyMap(femalePrefs),
		malePrefs:   copyMap(malesPrefs),
		matches:     make(map[string]string),
		maleNames:   getMapKeys(malesPrefs),
	}
}

func copyMap(originalMap map[string][]string) map[string][]string {
	newMap := make(map[string][]string)
	for k, v := range originalMap {
		newMap[k] = v
	}

	return newMap
}

func (data *matchData) getFreeMaleName() (string, error) {
	for _, name := range data.maleNames {
		if data.isPersonFree(name) {
			return name, nil
		}
	}

	return "", errors.New("No free male")
}

func (data *matchData) isPersonFree(name string) bool {
	val, ok := data.matches[name]
	return !ok || val == ""
}

func (data *matchData) engage(name1 string, name2 string) {
	data.matches[name1] = name2
	data.matches[name2] = name1
}

func (data *matchData) getMatch(name string) string {
	return data.matches[name]
}

func (data *matchData) clearMatch(name string) {
	data.matches[name] = ""
}

func (data *matchData) doesFemalePreferMaleXOverY(female string, maleX string, maleY string) bool {
	maleNames := data.femalePrefs[female]

	maleXIndex := len(maleNames) - 1
	maleYIndex := len(maleNames) - 1

	for i, name := range maleNames {
		if name == maleX {
			maleXIndex = i
			break
		} else if name == maleY {
			maleYIndex = i
			break
		}
	}

	return maleXIndex < maleYIndex
}

func (data *matchData) breakup(female string) {
	male := data.getMatch(female)
	data.clearMatch(female)
	data.clearMatch(male)
}

func (data *matchData) removeMalesFirstPref(male string) string {
	malePrefs := data.malePrefs[male]
	femaleName := malePrefs[0]
	data.malePrefs[male] = malePrefs[1:]
	return femaleName
}

// CreateStablePairs provides a pairing between the two groups based on the preferences
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

func getMapKeys(dict map[string][]string) []string {
	keys := make([]string, len(dict))

	i := 0
	for key := range dict {
		keys[i] = key
		i++
	}

	return keys
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
