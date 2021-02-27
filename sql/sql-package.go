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
func InsertDB(dbLocation string, tableName string, headers []string, values [][]string) {
	// Open the SQLite DB at the provided location
	database, err := sql.Open("sqlite3", dbLocation)
	// Handle any errors opening the DB
	if err != nil {
		fmt.Printf("Error opening SQLite DB File. Error: %v\n", err)
		os.Exit(2)
	}

	// Build insert query
	insertValueQuery := "INSERT INTO %v (%v) VALUES (%v)"
	headersString := strings.Join(headers[:], "', '")
	headersString = "'" + headersString + "'"
	valueCount := len(headers)
	count := 1
	valueString := "?"
	for count < valueCount {
		valueString = valueString + ", ?"
		count++
	}
	insertValueQuery = fmt.Sprintf(insertValueQuery, tableName, headersString, valueString)

	// Insert the data
	fmt.Printf("Inserting data into database: %v table: %v\n", dbLocation, tableName)
	rows := 0
	for _, value := range values {
		// Convert slice to interface
		row := make([]interface{}, len(value))
		for id := range value {
			row[id] = value[id]
		}
		insertQuery, err := database.Prepare(insertValueQuery)
		if err != nil {
			fmt.Printf("Error in inserting data into DB. Error: %v\n", err)
			os.Exit(4)
		}
		defer insertQuery.Close()
		insertQuery.Exec(row...)
		rows++
	}
	fmt.Printf("%v rows have been inserted into the table: %v\n", rows, tableName)
}

// RunQueries runs all queries on the sql database and returns a map of results
func RunQueries(dbLocation string, tableName string) {
	// Open the SQLite DB at the provided location
	database, err := sql.Open("sqlite3", dbLocation)
	// Handle any errors opening the DB
	if err != nil {
		fmt.Printf("Error opening SQLite DB File. Error: %v\n", err)
		os.Exit(2)
	}
	vulnBySeverity(database, tableName)
	topTenVulnHosts(database, tableName)
	mostDangerousVulns(database, tableName)
}

func structureHeaders(headers []string) string {
	headersString := strings.Join(headers[:], "' TEXT, '")
	headersString = "'" + headersString + "' TEXT"
	headersString = strings.Replace(headersString, "'CVSS' TEXT", "'CVSS' NUMERIC", 1)
	return headersString
}

func vulnBySeverity(conn *sql.DB, tableName string) {
	// Run the first query
	var criticalTotal int
	var severeTotal int
	var highTotal int
	var mediumTotal int
	var lowTotal int
	query := `
	SELECT
	(SELECT COUNT(*) FROM !! WHERE CVSS = 10) AS Critical,
	(SELECT COUNT(*) FROM !! WHERE CVSS BETWEEN 9 AND 9.9) AS Severe,
	(SELECT COUNT(*) FROM !! WHERE CVSS BETWEEN 7 and 8.9) AS High,
	(SELECT COUNT(*) FROM !! WHERE CVSS BETWEEN 4 and 6.9) AS Medium,
	(SELECT COUNT(*) FROM !! WHERE CVSS BETWEEN 0 and 3.9) AS Low
	`
	query = strings.Replace(query, "!!", tableName, -1)
	rows, err := conn.Query(query)
	if err != nil {
		fmt.Printf("Error running SQL Query. Error: %v\n", err)
		os.Exit(4)
	}
	for rows.Next() {
		rows.Scan(&criticalTotal, &severeTotal, &highTotal, &mediumTotal, &lowTotal)
		fmt.Println(criticalTotal, severeTotal, highTotal, mediumTotal, lowTotal)
	}
}

func topTenVulnHosts(conn *sql.DB, tableName string) {
	// Run the second query
	var mostVulnHost string
	var cvssTotal int
	var criticalTotal int
	var severeTotal int
	var highTotal int
	var mediumTotal int
	var lowTotal int
	query := `
	SELECT Host, ROUND(SUM(CVSS)) AS CVSS_Total,
	SUM(CASE WHEN CVSS = 10 THEN 1 ELSE 0 END) AS Critical,
	SUM(CASE WHEN CVSS BETWEEN 9 AND 9.9 THEN 1 ELSE 0 END) AS Severe,
	SUM(CASE WHEN CVSS BETWEEN 7 AND 8.9 THEN 1 ELSE 0 END) AS High,
	SUM(CASE WHEN CVSS BETWEEN 4 AND 6.9 THEN 1 ELSE 0 END) AS Medium,
	SUM(CASE WHEN CVSS BETWEEN 0 AND 3.9 THEN 1 ELSE 0 END) AS Low
	FROM !! GROUP BY Host ORDER BY CVSS_Total DESC LIMIT 10
	`
	query = strings.Replace(query, "!!", tableName, -1)
	rows, err := conn.Query(query)
	if err != nil {
		fmt.Printf("Error running SQL Query. Error: %v\n", err)
		os.Exit(4)
	}
	for rows.Next() {
		rows.Scan(&mostVulnHost, &cvssTotal, &criticalTotal, &severeTotal, &highTotal, &mediumTotal, &lowTotal)
		fmt.Println(mostVulnHost, cvssTotal, criticalTotal, severeTotal, highTotal, mediumTotal, lowTotal)
	}
}

func mostDangerousVulns(conn *sql.DB, tableName string) {
	// Run the second query
	var vulnName string
	var cvssTotal int
	var vulnsTotal int
	query := `
	SELECT Name, CVSS, COUNT(*) AS Total
	FROM !!
	WHERE CVSS BETWEEN 7 AND 10
	GROUP BY Name
	ORDER BY Total DESC
	LIMIT 10
	`
	query = strings.Replace(query, "!!", tableName, -1)
	rows, err := conn.Query(query)
	if err != nil {
		fmt.Printf("Error running SQL Query. Error: %v\n", err)
		os.Exit(4)
	}
	for rows.Next() {
		rows.Scan(&vulnName, &cvssTotal, &vulnsTotal)
		fmt.Println(vulnName, cvssTotal, vulnsTotal)
	}
}
