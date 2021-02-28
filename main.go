package main

import (
	"flag"
	"fmt"
	"time"

	"tbconsulting.com/nessus-charts/csv"
	"tbconsulting.com/nessus-charts/excel"
	"tbconsulting.com/nessus-charts/sql"
)

func main() {
	year, month, day := time.Now().Date()
	dateString := fmt.Sprint(day, month, year)
	csvLocation := flag.String("csv", "scan.csv", "Provide the filepath to the CSV output from Nessus.")
	dbLocation := flag.String("db", "vulns.db", "Provide the filepath to open the existing SQLite database, or where you would like the new one to be saved.")
	tableName := flag.String("table", dateString, "Provide the name of the table that you would like to use.")
	excelLocation := flag.String("excel", "Chart_Template.xlsx", "Provide the filepath to the Excel Chart Template.")
	flag.Parse()
	// csvLocation := "./scan.csv"
	// dbLocation := "./vulns.db"
	// tableName := "Lowell_1"
	// excelLocation := "./Chart_Template.xlsx"
	headers, records := csv.ProcessCSV(*csvLocation)
	sql.CreateTable(*dbLocation, *tableName, headers)
	sql.InsertDB(*dbLocation, *tableName, headers, records)
	vulnBySeverity, topTenVulnHosts, mostDangerousVulns, vulnByType, countCVSSYear := sql.RunQueries(*dbLocation, *tableName)
	excel.WriteData(*excelLocation, vulnBySeverity, topTenVulnHosts, mostDangerousVulns, vulnByType, countCVSSYear, headers, records)
}
