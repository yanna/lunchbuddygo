package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"lunchbuddy/csv"
	"lunchbuddy/matching"
	"math/rand"
	"os"
	"time"

	"github.com/pkg/errors"
)

func main() {

	buddyCsvFilePath := flag.String("csv", "", "Location of the Lunch Buddy csv file")
	verbose := flag.Bool("verbose", false, "Lots of logging")
	flag.Parse()

	if *buddyCsvFilePath == "" {
		log.Fatalln("csv flag is required.")
	}

	rand.Seed(time.Now().UnixNano())

	personReader := csv.NewPersonReader(*buddyCsvFilePath)
	peopleMatches, err := personReader.GetData()
	printErrorAndExit(err)

	if *verbose {
		print("People matches:", *peopleMatches)
	}

	group1, group2, oddPerson := peopleMatches.GetPreferences()
	if *verbose {
		print("Group1:", group1)
		print("Group2:", group2)
	}

	stableMarriage, err := matching.NewStableMarriage(group1, group2)
	printErrorAndExit(err)

	matches := stableMarriage.CreateStablePairs()

	if oddPerson != nil {
		fmt.Println(fmt.Sprintf("Odd Person: %s", oddPerson.Alias))
	}

	matchesJSON, _ := json.MarshalIndent(matches, "", " ")
	fmt.Println(string(matchesJSON))
}

func printErrorAndExit(err error) {
	if err != nil {
		fmt.Printf("Error: %T %v\n", errors.Cause(err), errors.Cause(err))
		fmt.Printf("Stack trace:\n%+v\n", err)
		os.Exit(1)
	}
}

func print(preText string, data interface{}) {
	dataJSON, _ := json.MarshalIndent(data, "", " ")
	fmt.Println(preText + "\n" + string(dataJSON))
}
