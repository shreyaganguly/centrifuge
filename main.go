package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	accessToken  = flag.String("token", "", "Access Token of your github account, if this flag is NOT set it will read from environemnt variable GITHUB_TOKEN")
	organization = flag.String("org", "", "Organization for which issues are to be searched")
	status       = flag.String("status", "all", "Filtering based on the status of the github issue(value can be either open or closed or all)(DEFAULT:all)")
	milestone    = flag.String("milestone", "", "Filtering based on the milestone assigned to the issues")
	labels       = flag.String("labels", "", "Filtering based on the labels marked to the issues(give comma-separated values)")
	filename     = flag.String("name", "", "file to save the extarcted issues, if empty it will print to stdout")
	format       = flag.String("format", "json", "Format to store after extracting issue details(json|csv|html|md)")
)

func main() {
	flag.Parse()
	if *accessToken == "" {
		*accessToken = os.Getenv("GITHUB_TOKEN")
		if *accessToken == "" {
			fmt.Println("Centrifuge Error: Please provide the github token in -token flag or set it by environment variable GITHUB_TOKEN ")
			return
		}
	}
	filterOptions := CreateFilterOptions()
	issues, err := findIssues(filterOptions)
	if err != nil {
		fmt.Println("Centrifuge Error: Error in listing by organization ", err)
		return
	}
	switch flagFormat := *format; flagFormat {
	case "csv":
		generateCSV(issues)
	case "json":
		generateJSON(issues)
	case "html":
		generateHTML(issues)
	case "md":
		generateMD(issues)
	default:
		fmt.Printf("Sorry %s fileformat is not supported. Supported formats are csv,html,md,json", flagFormat)
	}
}
