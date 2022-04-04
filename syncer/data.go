package syncer

import (
	"context"

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
	SyncProjectToBitable(ctx context.Context, projects []*define.Project) error
	SyncIssueToBitable(ctx context.Context, issues []*define.Issue) error
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
