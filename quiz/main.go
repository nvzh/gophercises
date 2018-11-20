package main

import (
	"flag"
)

func main() {

	csvFile := flag.String("csv", "problems.csv", "Specify a csv file with questions")
	flag.Parse()
	_ = csvFile
}
