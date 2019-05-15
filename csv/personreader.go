package csv

import (
	"bufio"
	"encoding/csv"
	"lunchbuddy/core"
	"os"
	"strconv"

	"github.com/pkg/errors"
)

// PersonReader retrieves Person data from the CSV
type PersonReader struct {
	csvFilePath string
}

// NewPersonReader creates a new PersonReader
func NewPersonReader(csvFilePath string) *PersonReader {
	return &PersonReader{csvFilePath: csvFilePath}
}

// GetData retrieves the people and matches from the csv
func (reader PersonReader) GetData() (*core.PeopleMatches, error) {

	csvPeople, error := getCsvPeople(reader.csvFilePath)
	if error != nil {
		return nil, error
	}

	people := createPeople(csvPeople)
	matches, error := createMatches(csvPeople, people)
	if error != nil {
		return nil, error
	}

	return &core.PeopleMatches{People: *people, Matches: *matches}, nil
}

func getCsvPeople(csvFilePath string) ([]person, error) {
	var people []person

	csvFile, openErr := os.Open(csvFilePath)
	if openErr != nil {
		return people, errors.Wrap(openErr, "Couldn't open the csv file")
	}

	reader := csv.NewReader(bufio.NewReader(csvFile))
	allLines, readErr := reader.ReadAll()
	if readErr != nil {
		return people, errors.Wrap(readErr, "Problem encountered reading the csv file")
	}

	allLinesLength := len(allLines)
	if allLinesLength <= 1 {
		return people, errors.New("Expecting at least two rows. First row is the headers, followed by people data")
	}

	people = make([]person, allLinesLength-1)

	headers := allLines[0]
	lines := allLines[1:]
	for i, line := range lines {
		newPerson, personErr := createCsvPerson(line, headers)
		if personErr != nil {
			return people, personErr
		}

		people[i] = newPerson
	}

	return people, nil
}

func convertPersonID(personIDStr string) (int, error) {
	personID, err := strconv.Atoi(personIDStr)
	if err != nil {
		return -1, errors.Wrapf(err, "Could not convert id %s to int", personIDStr)
	}

	return personID, nil
}

func createCsvPerson(line []string, headers []string) (person, error) {
	//todo: don't hardcode the headers
	personID, error := convertPersonID(line[0])
	if error != nil {
		return person{}, error
	}
	// Expectation is that we have these fixed headers, then the rest will be columns of matches
	// The header has the date in the format yyyymm and the cell value is the alias
	const MatchesStartingIndex = 6
	person := person{
		ID:         personID,
		FullName:   line[1],
		Alias:      line[2],
		Team:       line[3],
		Discipline: line[4],
		OptIn:      line[5],
	}

	matches := []match{}

	matchValues := line[MatchesStartingIndex:]
	for i, matchValue := range matchValues {

		if matchValue == "" {
			continue
		}

		matches = append(matches, match{
			Date:  headers[MatchesStartingIndex+i],
			Alias: matchValue,
		})
	}

	person.PastMatches = matches

	return person, nil
}
