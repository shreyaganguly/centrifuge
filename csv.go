package main

import (
	"encoding/csv"
	"fmt"
	"strings"

	"github.com/google/go-github/github"
)

func generateCSV(issues []*github.Issue) {
	records := generateIssueArray(issues)
	if *filename != "" {
		err := writeCSVToFIle(records)
		if err != nil {
			fmt.Println("Centrifuge Error: Could not write csv to file", err)
			return
		}
	} else {
		for _, v := range records {
			fmt.Println(strings.Join(v, ","))
		}
	}
}

func writeCSVToFIle(records [][]string) error {
	csvfile, err := createFile()
	if err != nil {
		return err
	}
	defer csvfile.Close()
	writer := csv.NewWriter(csvfile)
	for _, record := range records {
		err := writer.Write(record)
		if err != nil {
			return err
		}
	}
	fmt.Println("Data Successfully Written in: ", *filename)
	writer.Flush()
	return nil
}
