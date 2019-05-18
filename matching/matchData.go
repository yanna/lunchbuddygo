package matching

import "errors"

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

func copyMap(originalMap map[string][]string) map[string][]string {
	newMap := make(map[string][]string)
	for k, v := range originalMap {
		newMap[k] = v
	}

	return newMap
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
