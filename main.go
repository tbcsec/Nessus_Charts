package main

import (
	"tbconsulting.com/nessus-charts/sql"
)

func main() {
	// csvLocation := "./scan.csv"
	dbLocation := "./vulns.db"
	tableName := "Testing"
	// headers, records := csv.ProcessCSV(csvLocation)
	// sql.CreateTable(dbLocation, tableName, headers)
	// sql.InsertDB(dbLocation, tableName, headers, records)
	sql.RunQueries(dbLocation, tableName)
}
