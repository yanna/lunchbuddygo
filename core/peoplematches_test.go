package core

import (
	"encoding/json"
	"io/ioutil"
	"testing"
)

func unmarshalPeopleMatches(data []byte) (PeopleMatches, error) {
	var r PeopleMatches
	err := json.Unmarshal(data, &r)
	return r, err
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
			pm.GetPreferences()
		}
	})

	b.Run("goroutines", func(b *testing.B) {
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			pm.GetPreferencesParallel()
		}
	})
}
