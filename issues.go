package main

import (
	"strings"

	"github.com/google/go-github/github"
)

type Issue struct {
	RepositoryURL string `json:"repository_url"`
	Title         string `json:"title"`
	Status        string `json:"status"`
	Author        string `json:"author"`
	Milestone     string `json:"milestone"`
	Assignees     string `json:"assignees"`
	Labels        string `json:"labels"`
}
type List struct {
	TotalCount   int     `json:"total_count"`
	Organization string  `json:"organization"`
	Issues       []Issue `json:"issues"`
}

func createIssueStruct(issues []*github.Issue) (List, error) {

	list := List{
		Organization: *organization,
	}
	var issuesJSON []Issue
	for _, issue := range issues {
		var (
			issueJson       Issue
			names           []string
			labels          []string
			resultMilestone string
		)
		for _, label := range issue.Labels {
			labels = append(labels, label.GetName())
		}
		for _, assignee := range issue.Assignees {
			if assignee != nil {
				names = append(names, assignee.GetLogin())
			}
		}

		if issue.Milestone != nil {
			resultMilestone = issue.Milestone.GetTitle()
		}

		if (*milestone == "") || (*milestone != "" && *milestone == resultMilestone) {
			issueJson = Issue{
				RepositoryURL: issue.GetHTMLURL(),
				Title:         issue.GetTitle(),
				Status:        issue.GetState(),
				Author:        issue.User.GetLogin(),
				Milestone:     resultMilestone,
				Assignees:     strings.Join(names, ","),
				Labels:        strings.Join(labels, ","),
			}
			issuesJSON = append(issuesJSON, issueJson)
		}

	}
	list.Issues = issuesJSON
	list.Organization = *organization
	list.TotalCount = len(issuesJSON)
	return list, nil
}

func generateIssueArray(issues []*github.Issue) (records [][]string) {
	records = append(records, []string{"Repository URL", "Title", "Status", "Author", "Milestone", "Assignee", "Labels"})
	for _, issue := range issues {
		var (
			names           []string
			labels          []string
			resultMilestone string
		)
		for _, label := range issue.Labels {
			labels = append(labels, label.GetName())
		}
		for _, assignee := range issue.Assignees {
			if assignee != nil {
				names = append(names, assignee.GetLogin())
			}
		}

		if issue.Milestone != nil {
			resultMilestone = issue.Milestone.GetTitle()
		}

		if (*milestone == "") || (*milestone != "" && *milestone == resultMilestone) {
			records = append(records, []string{issue.GetHTMLURL(), issue.GetTitle(), issue.GetState(), issue.User.GetLogin(), resultMilestone, strings.Join(names, ","), strings.Join(labels, ",")})
		}

	}
	return
}
