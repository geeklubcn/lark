package gitlab

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

const (
	IssuesPath   = "/api/v4/issues"
	SyncAllLabel = "SYNC_ALL"
)

func (g *gitlab) Issues() (map[string]*IssueResult, error) {
	url := fmt.Sprintf("%s%s", g.domain, IssuesPath)
	if g.issueLabel != SyncAllLabel {
		url = fmt.Sprintf("%s?labels=%s", url, g.issueLabel)
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("PRIVATE-TOKEN", g.token)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 300 {
		return nil, errors.New(fmt.Sprintf("req fail status:%s", resp.Status))
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(body))
	var issues = make([]*IssueResult, 0)
	err = json.Unmarshal(body, &issues)
	if err != nil {
		return nil, err
	}

	res := make(map[string]*IssueResult, 0)
	for _, it := range issues {
		res[strconv.Itoa(it.Id)] = it
	}

	return res, nil
}
