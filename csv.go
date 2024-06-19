package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

func hackerone_csv() {
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
			eligible_for_submission := columns[4]

			if eligible_for_submission != "true" {
				assetType = "out-of-scope"
			}

			if identifier == "identifier" {
				// skip the csv header
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