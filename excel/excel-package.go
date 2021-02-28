// Package excel performs excel function
package excel

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

// WriteData function writes the query values to the appropriate pages
func WriteData(fileLocation string, cvssBySeverity []string, topTenVulnHosts [][]string, mostDangerousVulns [][]string, vulnByType []string, countCVSSYear [][]string, rawHeaders []string, rawRecords [][]string) {
	// Open the Excel Doc at the provided location
	file, err := excelize.OpenFile(fileLocation)
	// Handle any errors opening the DB
	if err != nil {
		fmt.Printf("Error opening Excel File. Error: %v\n", err)
		os.Exit(2)
	}
	writeRow(file, "CVSS By Severity", 2, cvssBySeverity)
	writeMultipleRow(file, "Top Vulnerable Hosts", topTenVulnHosts)
	writeMultipleRow(file, "Most Common Vulnerabilities", mostDangerousVulns)
	writeRow(file, "Vulnerabilities By Type", 2, vulnByType)
	writeMultipleRow(file, "Vulnerabilities By Year", countCVSSYear)
	writeRow(file, "Scan Data", 1, rawHeaders)
	writeMultipleRow(file, "Scan Data", rawRecords)
	newFile := ""
	if filepath.Dir(fileLocation) == "." {
		newFile = "Populated_" + filepath.Base(fileLocation)
	} else {
		newFile = filepath.Dir(fileLocation) + "Populated_" + filepath.Base(fileLocation)
	}
	if err := file.SaveAs(newFile); err != nil {
		fmt.Println(err)
	}
}

func writeRow(file *excelize.File, sheet string, rowID int, values []string) {
	for id, value := range values {
		cell := toCharStr(id+1) + strconv.Itoa(rowID)
		file.SetCellDefault(sheet, cell, value)
	}
}

func writeMultipleRow(file *excelize.File, sheet string, values [][]string) {
	for rowID, row := range values {
		adjRowID := rowID + 2
		writeRow(file, sheet, adjRowID, row)
	}
}

func toCharStr(i int) string {
	abc := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	return abc[i-1 : i]
}
