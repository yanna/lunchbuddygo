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

var globalOddPerson *Person

func BenchmarkGetPreferences(b *testing.B) {

	peopleMatchesJSON, err := ioutil.ReadFile("testdata/peoplematches.json")

	if err != nil {
		println(err)
		return
	}

	pm, _ := unmarshalPeopleMatches(peopleMatchesJSON)
	var oddPerson *Person
	for n := 0; n < b.N; n++ {
		_, _, oddPerson = pm.GetPreferences()
	}

	// Make sure the function call is not optimized out.
	globalOddPerson = oddPerson
}
