package main

import (
	"fmt"

	"tbconsulting.com/nessus-charts/csv"
)

func main() {
	csvLocation := "./scan.csv"
	headers, records := csv.ProcessCSV(csvLocation)
	fmt.Printf("Headers: %v\n", headers)
	fmt.Printf("Records: %v\n", records)
}
