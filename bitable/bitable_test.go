package bitable

import (
	"context"
	"github.com/larksuite/oapi-sdk-go/core"
	"testing"
)

func TestFaker(t *testing.T) {
	ctx := core.WrapContext(context.Background())
	_, _ = FakerBitable.GetApp(ctx)
	_, _ = FakerBitable.ListTables(ctx)
	_, _ = FakerBitable.CreateTable(ctx, nil)
	_, _ = FakerBitable.BatchCreateTable(ctx, nil)
	_, _ = FakerBitable.ListViews(ctx, "")
	_ = FakerBitable.SyncViews(ctx, "", nil)
	_, _ = FakerBitable.ListFields(ctx, "")
	_, _ = FakerBitable.CreateField(ctx, "", nil)
}
