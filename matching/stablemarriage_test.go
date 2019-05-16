package matching

import (
	"reflect"
	"testing"
)

func TestGetStablePairsUneven(t *testing.T) {
	malePrefs := make(map[string][]string)
	malePrefs["1"] = []string{"A", "B", "C"}

	femalePrefs := make(map[string][]string)
	femalePrefs["A"] = []string{"2", "3", "1"}
	femalePrefs["B"] = []string{"2", "1", "3"}

	_, error := NewStableMarriage(femalePrefs, malePrefs)
	if error == nil {
		t.Errorf("Expected an error")
	}
}

func TestGetStablePairsSamePrefs(t *testing.T) {
	malePrefs := make(map[string][]string)
	malePrefs["1"] = []string{"A", "B", "C"}
	malePrefs["2"] = []string{"A", "B", "C"}
	malePrefs["3"] = []string{"A", "B", "C"}

	femalePrefs := make(map[string][]string)
	femalePrefs["A"] = []string{"1", "2", "3"}
	femalePrefs["B"] = []string{"1", "2", "3"}
	femalePrefs["C"] = []string{"1", "2", "3"}

	expected := make(map[string]string)
	expected["1"] = "A"
	expected["2"] = "B"
	expected["3"] = "C"

	runGetStablePairsTest(t, malePrefs, femalePrefs, expected)
}

func TestGetStablePairs1(t *testing.T) {
	malePrefs := make(map[string][]string)
	malePrefs["1"] = []string{"A", "B", "C"}
	malePrefs["2"] = []string{"A", "B", "C"}
	malePrefs["3"] = []string{"A", "B", "C"}

	femalePrefs := make(map[string][]string)
	femalePrefs["A"] = []string{"3", "2", "1"}
	femalePrefs["B"] = []string{"1", "2", "3"}
	femalePrefs["C"] = []string{"2", "1", "3"}

	expected := make(map[string]string)
	expected["1"] = "B"
	expected["2"] = "C"
	expected["3"] = "A"

	runGetStablePairsTest(t, malePrefs, femalePrefs, expected)
}

func TestGetStablePairs2(t *testing.T) {
	malePrefs := make(map[string][]string)
	malePrefs["1"] = []string{"A", "B", "C"}
	malePrefs["2"] = []string{"B", "A", "C"}
	malePrefs["3"] = []string{"A", "B", "C"}

	femalePrefs := make(map[string][]string)
	femalePrefs["A"] = []string{"2", "3", "1"}
	femalePrefs["B"] = []string{"2", "1", "3"}
	femalePrefs["C"] = []string{"1", "2", "3"}

	expected := make(map[string]string)
	expected["1"] = "C"
	expected["2"] = "B"
	expected["3"] = "A"

	runGetStablePairsTest(t, malePrefs, femalePrefs, expected)
}

func runGetStablePairsTest(
	t *testing.T,
	malePrefs map[string][]string,
	femalePrefs map[string][]string,
	expected map[string]string) {
	algo, _ := NewStableMarriage(femalePrefs, malePrefs)

	got := algo.CreateStablePairs()
	eq := reflect.DeepEqual(got, expected)
	if !eq {
		t.Errorf("Want:\n%s\nExpected:\n%s", got, expected)
	}
}
