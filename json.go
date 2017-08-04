package main

import (
	"encoding/json"
	"fmt"

	"github.com/google/go-github/github"
)

func generateJSON(issues []*github.Issue) {
	records, err := generateJSONData(issues)
	if err != nil {
		fmt.Println("Centrifuge Error: Could not generate json", err)
		return
	}
	if *filename != "" {
		err = writeRecordsToFile(records)
		if err != nil {
			fmt.Println("Centrifuge Error: Could not write json to file", err)
			return
		}
	} else {
		fmt.Println(string(records))
	}
}

func generateJSONData(issues []*github.Issue) ([]byte, error) {
	list, err := createIssueStruct(issues)
	if err != nil {
		return nil, err
	}
	fmt.Println(list)
	listJSON, err := json.MarshalIndent(list, "", " ")
	if err != nil {
		return nil, err
	}
	return listJSON, nil
}

func writeRecordsToFile(records []byte) error {
	file, err := createFile()
	if err != nil {
		return err
	}
	_, err = file.Write(records)
	if err != nil {
		return err
	}
	fmt.Println("Data Successfully Written in: ", *filename)
	return nil
}
