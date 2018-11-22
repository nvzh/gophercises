package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// Define variable for csv file specifying
	csvFile := flag.String("csv", "problems.csv", "Specify a csv file with questions")
	flag.Parse()
	//_ = csvFile

	// The file parsing
	file, err := os.Open(*csvFile)
	if err != nil {
		exit(fmt.Sprintf("Unble to open the file. Check the path or so. There is no any \"%s\"\n", *csvFile))
	}
	_ = file
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
