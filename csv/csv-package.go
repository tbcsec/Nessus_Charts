// Package csv performs CSV operations
package csv

import (
	"encoding/csv"
	"fmt"
	"os"
)

// ProcessCSV takes the csv and returns the values
func ProcessCSV(csvLocation string) ([]string, [][]string) {
	// Open the CSV at the provided location
	file, err := os.Open(csvLocation)
	// Handle any errors opening the file
	if err != nil {
		fmt.Printf("Error opening CSV file. Error: %v\n", err)
		os.Exit(2)
	}
	// Read the CSV
	reader := csv.NewReader(file)
	records, _ := reader.ReadAll()
	// These fields are required for the built-in SQL queries to work
	// Check for the presence of each required field and log any missing fields
	requiredFields := []string{
		"Plugin ID",
		"CVE",
		"CVSS",
		"Risk",
		"Host",
		"Protocol",
		"Port",
		"Name",
		"Synopsis",
		"Description",
		"Solution",
		"See Also",
		"Plugin Output",
	}
	for _, requiredField := range requiredFields {
		if checkRequiredFields(requiredField, records[0]) != true {
			fmt.Printf("Missing required field: %v\n", requiredField)
			os.Exit(1)
		}
	}
	// Strip down the CSV to only include required fields
	// Return the headers and CSV values
	return requiredFields, records
}

func checkRequiredFields(requiredField string, headers []string) bool {
	for _, header := range headers {
		if requiredField == header {
			return true
		}
	}
	return false
}
