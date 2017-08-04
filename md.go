package main

import (
	"fmt"
	"strings"

	"github.com/google/go-github/github"
)

func generateMD(issues []*github.Issue) {
	mdData := generateMDData(issues)
	if *filename != "" {
		err := writeRecordsToFile([]byte(mdData))
		if err != nil {
			fmt.Println("Centrifuge Error: Could not write md to file", err)
			return
		}
	} else {
		fmt.Println(mdData)
	}
}

func generateMDData(issues []*github.Issue) string {
	records := generateIssueArray(issues)
	var issuesMD string
	for _, record := range records {
		if issuesMD != "" {
			issuesMD = fmt.Sprintf("%s\n%s", issuesMD, strings.Join(record, "|"))
		} else {
			issuesMD = fmt.Sprintf("### Organization: %s\n### Total Count: %d\n", *organization, len(records)-1)

			issuesMD = fmt.Sprintf("%s%s\n--------|---------|-----------|----------|----------|------------|-------", issuesMD, strings.Join(record, "|"))
		}
	}
	return issuesMD + "\n"
}
