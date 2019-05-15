package csv

import (
	"lunchbuddy/core"
)

func createPeople(csvPeople []person) *core.People {
	var people []core.Person

	for _, csvPerson := range csvPeople {
		newPerson := createPerson(&csvPerson)
		people = append(people, *newPerson)
	}

	return &core.People{Data: people}
}

func createPerson(csvPerson *person) *core.Person {
	output := core.NewPerson(
		csvPerson.ID,
		csvPerson.FullName,
		csvPerson.Alias,
		csvPerson.Team,
		csvPerson.Discipline,
		csvPerson.OptIn)
	return output
}

func createMatches(csvPeople []person, idProvider core.IPersonIDProvider) (*core.Matches, error) {

	idToMatches := make(map[int][]core.Match)

	for _, csvPerson := range csvPeople {
		if len(csvPerson.PastMatches) == 0 {
			continue
		}

		var matches []core.Match
		for _, pastMatch := range csvPerson.PastMatches {
			personID, error := idProvider.GetPersonIDByAlias(pastMatch.Alias)
			if error != nil {
				return nil, error
			}
			matches = append(matches, *core.NewMatch(pastMatch.Date, pastMatch.Alias, personID))
		}

		idToMatches[csvPerson.ID] = matches
	}

	return &core.Matches{PersonIDToMatches: idToMatches}, nil
}
