// Package sql performs SQL operations
package sql

import (
	"database/sql"
	"fmt"
	"os"

	// Import the sqlite3 SQL driver
	_ "github.com/mattn/go-sqlite3"
)

// CreateTable creates the table from the supplied values
func CreateTable(dbLocation string, tableName string, values []string) {
	// Open the SQLite DB at the provided location
	database, err := sql.Open("sqlite3", dbLocation)
	// Handle any errors opening the DB
	if err != nil {
		fmt.Printf("Error opening SQLite DB File. Error: %v\n", err)
		os.Exit(2)
	}
	// Create the table
	createTableQuery := `CREATE TABLE IF NOT EXISTS ` + tableName + `
	(
		"Plugin ID" TEXT, CVE TEXT, CVSS NUMERIC, Risk TEXT, Host TEXT, Protocol TEXT,
		Port TEXT, Name TEXT, Synopsis TEXT, Description TEXT, Solution TEXT,
		"See Also" TEXT, "Plugin Output" TEXT
	)`
	tableQuery, err := database.Prepare(createTableQuery)
	if err != nil {
		fmt.Printf("Improper SQL Query. Error: %v\n", err)
		os.Exit(3)
	}
	defer tableQuery.Close()
	tableQuery.Exec()
}

// InsertDB inserts data into the database
func InsertDB(dbLocation string, tableName string, values [][]string) {

	// Insert the data
	for _, value := range values {
		insertValueQuery := "INSERT INTO " + tableName + "(firstname, lastname) VALUES (?, ?)"
		insertQuery, err := database.Prepare(insertValueQuery)
		if err != nil {
			fmt.Printf("Error in inserting data into DB. Error: %v\n", err)
			os.Exit(4)
		}
		defer insertQuery.Close()
		insertQuery.Exec(value)
	}
}
