package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
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

	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.TrimLeadingSpace = true

	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	for _, columns := range records {
		if len(columns) > 1 {
			identifier := columns[0]
			assetType := columns[1]

			if identifier == "identifier" {
				continue
			}

			file, err := os.OpenFile(strings.ToLower(assetType)+"_targets.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				fmt.Println("Error opening file:", err)
				return
			}
			defer file.Close()

			_, err = file.WriteString(identifier + "\n")
			if err != nil {
				fmt.Println("Error appending to file:", err)
				return
			}
		} else {
			fmt.Println("Invalid input format")
		}
	}
}
