package gitlab

import "net/http"

type Domain string

type Gitlab interface {
	Issues() (map[string]*IssueResult, error)
	Project(id string) (*ProjectResult, error)
}

type gitlab struct {
	domain     string
	token      string
	client     *http.Client
	issueLabel string
}

func NewGitlabClient(domain, token, issueLabel string) Gitlab {
	return &gitlab{
		domain:     domain,
		token:      token,
		client:     &http.Client{},
		issueLabel: issueLabel,
	}
}
