// Package sql performs SQL operations
package sql

import (
	"database/sql"
	"fmt"
	"os"
	"strings"

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
	// Structure the headers
	headers := structureHeaders(values)
	// Create the table
	createTableQuery := `CREATE TABLE IF NOT EXISTS ` + tableName + `(` + headers + `)`
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
	// Open the SQLite DB at the provided location
	database, err := sql.Open("sqlite3", dbLocation)
	// Handle any errors opening the DB
	if err != nil {
		fmt.Printf("Error opening SQLite DB File. Error: %v\n", err)
		os.Exit(2)
	}
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

func structureHeaders(headers []string) string {
	headersString := strings.Join(headers[:], "' TEXT, '")
	headersString = "'" + headersString + "' TEXT"
	headersString = strings.Replace(headersString, "'CVSS' TEXT", "'CVSS' NUMERIC", 1)
	return headersString
}
