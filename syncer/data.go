package syncer

import (
	"context"
	"github.com/geeklubcn/lark/gitlab"
	"strconv"
	"time"

	"github.com/geeklubcn/lark/define"
	"github.com/larksuite/oapi-sdk-go/core"
)

const (
	Project   = "Project"
	Milestone = "Milestone"
	Issue     = "Issue"
	Member    = "Member"
)

type DataSyncer interface {
	Sync(ctx context.Context) error
	SyncProjectToBitable(ctx context.Context, projects []*define.Project) error
	SyncIssueToBitable(ctx context.Context, issues []*define.Issue) error
	convertIssue(gitlabIssue *gitlab.IssueResult) *define.Issue
	convertProject(gitlabIssue *gitlab.ProjectResult) *define.Project
}

func (s *syncer) convertProject(p *gitlab.ProjectResult) *define.Project {
	res := &define.Project{
		ID:          strconv.Itoa(p.Id),
		Name:        p.Name,
		Description: p.Description,
		WebUrl:      p.WebUrl,
		TagList:     p.TagList,
		Topics:      p.Topics,
	}
	return res
}

func (s *syncer) convertIssue(i *gitlab.IssueResult) *define.Issue {
	res := &define.Issue{
		ID:           strconv.Itoa(i.Id),
		Title:        i.Title,
		Description:  i.Description,
		State:        i.State,
		CreatedAt:    i.CreatedAt,
		UpdatedAt:    i.UpdatedAt,
		HealthStatus: i.HealthStatus,
	}
	switch i.DueDate.(type) {
	case string:
		if dueData, err := time.Parse("2006-01-02", i.DueDate.(string)); err == nil {
			res.DueDate = dueData
		}
	}
	tableId := s.meta.GetProjectTableID()
	if tableId != "" {
		records, _ := s.b.ListRecords(core.WrapContext(context.Background()), tableId)
		for _, r := range records {
			if r.Fields["ID"].(string) == strconv.Itoa(i.ProjectId) {
				res.ProjectRef = &define.ProjectRef{
					Text:      r.Fields["Name"].(string),
					Type:      "text",
					TableID:   tableId,
					RecordIDs: []string{r.RecordId},
				}
			}
		}
	}
	return res
}

func (s *syncer) Sync(ctx context.Context) error {
	gitlabIssues, err := s.g.Issues()
	if err != nil {
		panic(err)
	}
	issues := make([]*define.Issue, 0)
	projects := make([]*define.Project, 0)
	for _, gitlabIssue := range gitlabIssues {
		issues = append(issues, s.convertIssue(gitlabIssue))
	}
	for _, gitlabIssue := range gitlabIssues {
		p, _ := s.g.Project(strconv.Itoa(gitlabIssue.ProjectId))
		projects = append(projects, s.convertProject(p))
	}
	// TODO diff & cache
	if err = s.SyncProjectToBitable(ctx, projects); err != nil {
		return err
	}
	if err = s.SyncIssueToBitable(ctx, issues); err != nil {
		return err
	}
	return nil
}

func (s *syncer) SyncIssueToBitable(ctx context.Context, issues []*define.Issue) error {
	c := core.WrapContext(ctx)
	tables, err := s.b.ListTables(c)
	if err != nil {
		return err
	}

	if t, ok := tables[Issue]; ok {
		fields := make([]map[string]interface{}, 0)
		for _, it := range issues {
			fields = append(fields, it.ToField())
		}
		return s.b.SyncRecords(c, t.TableId, fields)
	}
	return nil
}

func (s *syncer) SyncProjectToBitable(ctx context.Context, projects []*define.Project) error {
	c := core.WrapContext(ctx)
	tables, err := s.b.ListTables(c)
	if err != nil {
		return err
	}

	if t, ok := tables[Project]; ok {
		fields := make([]map[string]interface{}, 0)
		for _, p := range projects {
			fields = append(fields, p.ToField())
		}
		return s.b.SyncRecords(c, t.TableId, fields)
	}
	return nil
}
