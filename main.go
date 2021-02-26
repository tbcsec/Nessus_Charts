package main

import (
	"fmt"

	"tbconsulting.com/nessus-charts/csv"
	"tbconsulting.com/nessus-charts/sql"
)

func main() {
	csvLocation := "./scan.csv"
	dbLocation := "./vulns.db"
	tableName := "Testing"
	headers, records := csv.ProcessCSV(csvLocation)
	fmt.Printf("Headers: %v\n", headers)
	fmt.Printf("Records: %v\n", records)
	sql.CreateTable(dbLocation, tableName, headers)
}
