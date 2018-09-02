package main

import (
	"context"
	"strconv"
	"strings"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

func CreateFilterOptions() *github.IssueListOptions {
	listOptions := &github.IssueListOptions{Filter: "all", ListOptions: github.ListOptions{PerPage: 100}}
	if *status == "open" || *status == "closed" {
		listOptions.State = *status
	} else {
		listOptions.State = "all"
	}
	if *labels != "" {
		listOptions.Labels = strings.Split(*labels, ",")
	}
	return listOptions
}

func findIssues(filterOptions *github.IssueListOptions) (issueList []*github.Issue, err error) {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: *accessToken},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)
	if *project != "" {
		issues, err := getIssueURLOfProject()
		if err != nil {
			return nil, err
		}
		for _, issue := range issues {
			splat := strings.Split(issue, "/")
			issueNumber, err := strconv.Atoi(splat[6])
			if err != nil {
				return nil, err
			}
			issueFromGit, _, err := client.Issues.Get(ctx, *organization, splat[4], issueNumber)
			if err != nil {
				return nil, err
			}
			issueList = append(issueList, issueFromGit)
		}
		return issueList, nil
	}

	i := 1
	for {
		filterOptions.ListOptions.Page = i
		issues, _, err := client.Issues.ListByOrg(ctx, *organization, filterOptions)
		if err != nil {
			return nil, err
		}
		if len(issues) == 0 {
			break
		}
		i++
		issueList = append(issueList, issues...)
	}
	return
}
