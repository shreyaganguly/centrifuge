package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// TODO: Project handle incomplete results later
type Project struct {
	IncompleteResults bool `json:"incomplete_results"`
	Items             []struct {
		HTMLURL string `json:"html_url"`
	} `json:"items"`
}

func getIssueURLOfProject() ([]string, error) {
	var issues []string
	var projectResp *Project
	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.github.com/search/issues?q=org:%s+project:%s", *organization, *project), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("token %s", *accessToken))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &projectResp)
	if err != nil {
		return nil, err
	}
	for _, issue := range projectResp.Items {
		issues = append(issues, issue.HTMLURL)
	}

	defer resp.Body.Close()
	return issues, nil
}
