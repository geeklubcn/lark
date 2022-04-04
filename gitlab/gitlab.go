package gitlab

import "net/http"

type Domain string

type Gitlab interface {
	Issues() ([]*IssueResult, error)
	Project(id string) (*ProjectResult, error)
}

type gitlab struct {
	domain     string
	token      string
	client     *http.Client
	issueLabel string
}

func New(domain, token, issueLabel string) Gitlab {
	return &gitlab{
		domain:     domain,
		token:      token,
		client:     &http.Client{},
		issueLabel: issueLabel,
	}
}
