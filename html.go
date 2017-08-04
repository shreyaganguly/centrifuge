package main

import (
	"fmt"
	"os"
	"text/template"

	"github.com/google/go-github/github"
)

func generateHTML(issues []*github.Issue) {
	htmltmpl := generateHTMLTemplate()

	issuesHTML, err := createIssueStruct(issues)
	if err != nil {
		fmt.Println("Centrifuge Error: Could not create issue structure", err)
		return
	}
	if *filename != "" {
		err = writeHTMLToFile(issuesHTML, htmltmpl)
		if err != nil {
			fmt.Println("Centrifuge Error: Could not write html to file", err)
			return
		}
	} else {
		err = htmltmpl.Execute(os.Stdout, issuesHTML)
		if err != nil {
			fmt.Println("Centrifuge Error: Could not write html to stdout", err)
			return
		}
	}
}
func generateHTMLTemplate() *template.Template {

	return template.Must(template.New("issuetmpl").Parse(issuetmpl))
}

func writeHTMLToFile(issuesHTML List, tmpl *template.Template) error {
	htmlfile, err := createFile()
	if err != nil {
		return err
	}
	defer htmlfile.Close()
	err = tmpl.Execute(htmlfile, issuesHTML)
	if err != nil {
		return err
	}
	return nil
}
