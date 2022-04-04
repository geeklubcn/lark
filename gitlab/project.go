package gitlab

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

const (
	ProjectPath = "/api/v4/projects"
)

func (g *gitlab) Project(id string) (*ProjectResult, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s%s/%s", g.domain, ProjectPath, id), nil)
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
	var res ProjectResult
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
