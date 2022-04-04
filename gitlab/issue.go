package gitlab

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

const (
	IssuesPath = "/api/v4/issues"
)

func (g *gitlab) Issues() ([]*IssueResult, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s%s?labels=%s", g.domain, IssuesPath, g.issueLabel), nil)
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
	var res = make([]*IssueResult, 0)
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
