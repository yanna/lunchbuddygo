package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {

	buddyCsv := flag.String("csv", "", "Location of the Lunch Buddy csv file")
	flag.Parse()

	if *buddyCsv == "" {
		log.Fatal("csv flag is required.")
	}

	fmt.Println("Csv", *buddyCsv)

}
