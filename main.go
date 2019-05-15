package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"lunchbuddy/csv"

	"github.com/pkg/errors"
)

func main() {

	buddyCsvFilePath := flag.String("csv", "", "Location of the Lunch Buddy csv file")
	flag.Parse()

	if *buddyCsvFilePath == "" {
		log.Fatalln("csv flag is required.")
	}

	fmt.Println("Csv", *buddyCsvFilePath)

	personReader := csv.NewPersonReader(*buddyCsvFilePath)
	peopleMatches, dataError := personReader.GetData()

	if dataError != nil {
		fmt.Printf("Error:\n%+v\n", errors.Cause(dataError))
		return
	}

	// Be care: json.Marshal() will return null for var myslice []int and [] for initialized slice myslice := []int{}
	peopleJSON, _ := json.MarshalIndent(*peopleMatches, "", " ")
	fmt.Println(string(peopleJSON))
}
