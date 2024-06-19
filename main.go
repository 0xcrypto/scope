package main

import (
	"flag"
	"os"
	"scope/csv"
)

var inputFile string

func init() {
	flag.StringVar(&inputFile, "input", "", "input file path")
	flag.Parse()

	if len(os.Args) > 1 {
		inputFile = os.Args[1]
	}
}

func main() {

	hackerone_csv(inputFile)

}
