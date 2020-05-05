package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"lunchbuddy/core"
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
	matchModeFlag := flag.String("mode", "", "Different/Similar to indicate the type of match we want")
	flag.Parse()

	if *buddyCsvFilePath == "" {
		log.Fatalln("csv flag is required.")
	}

	matchMode := core.Different
	switch *matchModeFlag {
	case "":
		log.Fatalln("mode flag is required.")
	case "Similar":
		matchMode = core.Similar
	case "Different":
		matchMode = core.Different
	default:
		log.Fatalln("mode flag only expects Similar or Different.")
	}

	rand.Seed(time.Now().UnixNano())

	personReader := csv.NewPersonReader(*buddyCsvFilePath)
	peopleMatches, err := personReader.GetData()
	printErrorAndExit(err)
	if *verbose {
		print("People matches:", *peopleMatches)
	}

	group1, group2, oddPerson := peopleMatches.GetPreferences(matchMode)
	if *verbose {
		print("Group1:", group1)
		print("Group2:", group2)
	}

	stableMarriage, err := matching.NewStableMarriage(group1, group2)
	printErrorAndExit(err)

	matches := stableMarriage.CreateStablePairs()
	if *verbose {
		matchesJSON, _ := json.MarshalIndent(matches, "", " ")
		fmt.Println(string(matchesJSON))
	}

	matchesOutput := core.NewMatchesOutput(matches, peopleMatches, oddPerson)
	matchesOutput.Print()
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
