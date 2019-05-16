package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"lunchbuddy/matching"
)

func main() {

	buddyCsvFilePath := flag.String("csv", "", "Location of the Lunch Buddy csv file")
	flag.Parse()

	if *buddyCsvFilePath == "" {
		log.Fatalln("csv flag is required.")
	}

	fmt.Println("Csv", *buddyCsvFilePath)
	/*
		personReader := csv.NewPersonReader(*buddyCsvFilePath)
		peopleMatches, dataError := personReader.GetData()

		if dataError != nil {
			fmt.Printf("Error getting data:\n%+v\n", errors.Cause(dataError))
			return
		}

		// Be care: json.Marshal() will return null for var myslice []int and [] for initialized slice myslice := []int{}
		peopleJSON, _ := json.MarshalIndent(*peopleMatches, "", " ")
		fmt.Println(string(peopleJSON))
	*/
	males := make(map[string][]string)
	males["0"] = []string{"7", "5", "6", "4"}
	males["1"] = []string{"5", "4", "6", "7"}
	males["2"] = []string{"4", "5", "6", "7"}
	males["3"] = []string{"4", "5", "6", "7"}

	females := make(map[string][]string)
	females["4"] = []string{"0", "1", "2", "3"}
	females["5"] = []string{"0", "1", "2", "3"}
	females["6"] = []string{"0", "1", "2", "3"}
	females["7"] = []string{"0", "1", "2", "3"}

	stableMarriage, error := matching.NewStableMarriage(females, males)
	if error != nil {
		fmt.Printf("Error getting data:\n%+v\n", error)
		return
	}

	matches := stableMarriage.CreateStablePairs()
	matchesJSON, _ := json.MarshalIndent(matches, "", " ")
	fmt.Println(string(matchesJSON))
}
