package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/pkg/errors"
)

// Person - Information about a person
type csvPerson struct {
	ID          int        `json:"id"`
	FullName    string     `json:"fullname"`
	Alias       string     `json:"alias"`
	Team        string     `json:"team"`
	Discipline  string     `json:"discipline"`
	OptIn       string     `json:"optin"`
	PastMatches []csvMatch `json:"past_matches"`
}

// Match - Information about a match
type csvMatch struct {
	Date  string `json:"date"`
	Alias string `json:"alias"`
}

func main() {

	buddyCsv := flag.String("csv", "", "Location of the Lunch Buddy csv file")
	flag.Parse()

	if *buddyCsv == "" {
		log.Fatalln("csv flag is required.")
	}

	fmt.Println("Csv", *buddyCsv)
	people, csvErr := getCsvPerson(*buddyCsv)

	if csvErr != nil {
		fmt.Printf("Error:\n%+v\n", errors.Cause(csvErr))
	}

	// Be care: json.Marshal() will return null for var myslice []int and [] for initialized slice myslice := []int{}
	peopleJSON, _ := json.MarshalIndent(people, "", " ")
	fmt.Println(string(peopleJSON))
}

// GetPerson reads the csv file and populates the Person object
func getCsvPerson(csvFilePath string) ([]csvPerson, error) {
	var people []csvPerson

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

	people = make([]csvPerson, allLinesLength-1)

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

func createCsvPerson(line []string, headers []string) (csvPerson, error) {
	//todo: don't hardcode the headers
	personIDStr := line[0]
	personID, err := strconv.Atoi(personIDStr)
	if err != nil {
		return csvPerson{}, errors.Wrapf(err, "Could not convert id %s to int", personIDStr)
	}
	// Expectation is that we have these fixed headers, then the rest will be columns of matches
	// The header has the date in the format yyyymm and the cell value is the alias
	const MatchesStartingIndex = 6
	person := csvPerson{
		ID:         personID,
		FullName:   line[1],
		Alias:      line[2],
		Team:       line[3],
		Discipline: line[4],
		OptIn:      line[5],
	}

	// Instantiate so JSON won't be null for the list
	matches := []csvMatch{}

	if matches == nil {
		return person, nil
	}

	matchValues := line[MatchesStartingIndex:]
	for i, matchValue := range matchValues {

		if matchValue == "" {
			break
		}

		matches = append(matches, csvMatch{
			Date:  headers[MatchesStartingIndex+i],
			Alias: matchValue,
		})
	}

	person.PastMatches = matches

	return person, nil
}
