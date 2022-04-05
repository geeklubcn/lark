package syncer

import (
	"context"
	"github.com/geeklubcn/lark/bitable"
	"github.com/larksuite/oapi-sdk-go/core"
)

type Meta interface {
	GetProjectTableID() string
}

type meta struct {
	b bitable.Bitable
}

// GetProjectTableID TODO cache
func (m meta) GetProjectTableID() string {
	if tables, err := m.b.ListTables(core.WrapContext(context.Background())); err == nil {
		if table, ok := tables[Project]; ok {
			return table.TableId
		}
	}
	return ""
}
