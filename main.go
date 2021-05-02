package main

import (
	"encoding/csv"
	"flag"
	"log"
	"os"
)

var (
	timeLimit = flag.Int("limit", 20, "time limit for the quiz in seconds")
)

func main() {

	csvFilename := flag.String("csv", "problems.csv", "a csv file in the formato of 'question, answer'")
	flag.Parse()

	file, err := os.Open(*csvFilename)
	errorHandling(err)

	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	errorHandling(err)
	// fmt.Println(lines)
	gq(lines)

}

func errorHandling(err error) {
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
