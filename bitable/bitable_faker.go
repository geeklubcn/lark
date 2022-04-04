package bitable

import (
	"encoding/json"
	"github.com/larksuite/oapi-sdk-go/core"
	larkBitable "github.com/larksuite/oapi-sdk-go/service/bitable/v1"
)

var FakerBitable = &fakerBitable{}

type fakerBitable struct{}

func (f *fakerBitable) GetApp(ctx *core.Context) (*larkBitable.App, error) {
	var res larkBitable.App
	_ = json.Unmarshal([]byte(`
{
	"app": {
		"app_token": "appbcbWCzen6D8dezhoCH2RpMAh",
		"name": "mybitable",
		"revision": 1
	}
}
`), &res)
	return &res, nil
}

func (f *fakerBitable) ListTables(ctx *core.Context) (map[string]*larkBitable.AppTable, error) {
	var res larkBitable.AppTable
	_ = json.Unmarshal([]byte(`
{
	"has_more": false,
	"page_token": "tblKz5D60T4JlfcT",
	"total": 1,
	"items": [
		{
			"table_id": "tblKz5D60T4JlfcT",
			"revision": 1,
			"name": "数据表1"
		}
	]
}
`), &res)
	return map[string]*larkBitable.AppTable{
		res.Name: &res,
	}, nil
}

func (f *fakerBitable) CreateTable(ctx *core.Context, body *larkBitable.AppTableCreateReqBody) (string, error) {
	return "tblKz5D60T4JlfcT", nil
}

func (f *fakerBitable) BatchCreateTable(ctx *core.Context, body *larkBitable.AppTableBatchCreateReqBody) ([]string, error) {
	return []string{"tblKz5D60T4JlfcT"}, nil
}

func (f *fakerBitable) ListViews(ctx *core.Context, tableId string) (map[string]*larkBitable.AppTableView, error) {
	var res larkBitable.AppTableView
	_ = json.Unmarshal([]byte(`
{
	"items": [
		{
			"view_id": "vewTpR1urY",
			"view_name": "甘特视图1",
			"view_type": "gantt"
		}
	],
	"page_token": "vewdHB3HyE",
	"has_more": false,
	"total": 1
}
`), &res)
	return map[string]*larkBitable.AppTableView{
		res.ViewName: &res,
	}, nil
}

func (f *fakerBitable) SyncViews(ctx *core.Context, tableId string, body []*larkBitable.AppTableView) error {
	return nil
}

func (f *fakerBitable) ListFields(ctx *core.Context, tableId string) (map[string]*larkBitable.AppTableField, error) {
	var res larkBitable.AppTableField
	_ = json.Unmarshal([]byte(`
[
	{
		"field_id": "fldjX7dUj5",
		"field_name": "多行文本",
		"property": null,
		"type": 1
	},
	{
		"field_id": "fldoMnnvIR",
		"field_name": "数字",
		"property": {
			"formatter": "0.00"
		},
		"type": 2
	}
]
`), &res)
	return map[string]*larkBitable.AppTableField{
		res.FieldName: &res,
	}, nil
}

func (f *fakerBitable) CreateField(ctx *core.Context, tableId string, body *larkBitable.AppTableField) (string, error) {
	return "fldjX7dUj5", nil
}
