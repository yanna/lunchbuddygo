package core

import (
	"encoding/json"
	"io/ioutil"
	"testing"
)

func unmarshalPeopleMatches(data []byte) (PeopleMatches, error) {
	var pm PeopleMatches
	err := json.Unmarshal(data, &pm)
	return pm, err
}

func getPeopleMatches(b *testing.B) PeopleMatches {
	peopleMatchesJSON, err := ioutil.ReadFile("testdata/peoplematches.json")

	if err != nil {
		b.Error(err)
	}
	pm, _ := unmarshalPeopleMatches(peopleMatchesJSON)
	return pm
}

func BenchmarkGetPreferences(b *testing.B) {

	// Setup
	pm := getPeopleMatches(b)

	b.Run("normal", func(b *testing.B) {
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			pm.GetPreferences(Different)
		}
	})

	b.Run("goroutines", func(b *testing.B) {
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			pm.GetPreferencesParallel(Different)
		}
	})
}
