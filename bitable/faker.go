package bitable

import (
	"encoding/json"
	"github.com/larksuite/oapi-sdk-go/core"
	larkBitable "github.com/larksuite/oapi-sdk-go/service/bitable/v1"
)

var FakerBitable = &fakerBitable{}

type fakerBitable struct{}

func (f *fakerBitable) GetApp(_ *core.Context) (*larkBitable.App, error) {
	var res larkBitable.App
	_ = json.Unmarshal([]byte(`
{
	"app_token": "appbcbWCzen6D8dezhoCH2RpMAh",
	"name": "mybitable",
	"revision": 1
}
`), &res)
	return &res, nil
}

func (f *fakerBitable) ListTables(_ *core.Context) (map[string]*larkBitable.AppTable, error) {
	var res map[string]*larkBitable.AppTable
	_ = json.Unmarshal([]byte(`
{
    "数据表1": {
        "table_id": "tblKz5D60T4JlfcT",
        "revision": 1,
        "name": "数据表1"
    }
}
`), &res)
	return res, nil
}

func (f *fakerBitable) CreateTable(_ *core.Context, _ *larkBitable.AppTableCreateReqBody) (string, error) {
	return "tblKz5D60T4JlfcT", nil
}

func (f *fakerBitable) BatchCreateTable(_ *core.Context, _ *larkBitable.AppTableBatchCreateReqBody) ([]string, error) {
	return []string{"tblKz5D60T4JlfcT"}, nil
}

func (f *fakerBitable) ListViews(_ *core.Context, _ string) (map[string]*larkBitable.AppTableView, error) {
	var res map[string]*larkBitable.AppTableView
	_ = json.Unmarshal([]byte(`
{
    "甘特视图1": {
        "view_id": "vewTpR1urY",
        "view_name": "甘特视图1",
        "view_type": "gantt"
    }
}
`), &res)
	return res, nil
}

func (f *fakerBitable) SyncViews(_ *core.Context, _ string, _ []*larkBitable.AppTableView) error {
	return nil
}

func (f *fakerBitable) ListFields(_ *core.Context, _ string) (map[string]*larkBitable.AppTableField, error) {
	var res map[string]*larkBitable.AppTableField
	_ = json.Unmarshal([]byte(`
{
	"多行文本": {
		"field_id": "fldjX7dUj5",
		"field_name": "多行文本",
		"property": null,
		"type": 1
	},
	"数字": {
		"field_id": "fldoMnnvIR",
		"field_name": "数字",
		"property": {
			"formatter": "0.00"
		},
		"type": 2
	}
}
`), &res)
	return res, nil
}

func (f *fakerBitable) CreateField(_ *core.Context, _ string, _ *larkBitable.AppTableField) (string, error) {
	return "fldjX7dUj5", nil
}

func (f *fakerBitable) UpdateField(_ *core.Context, _ string, _ *larkBitable.AppTableField) error {
	return nil
}

func (f *fakerBitable) ListRecords(_ *core.Context, _ string) (map[string]*larkBitable.AppTableRecord, error) {
	var res map[string]*larkBitable.AppTableRecord
	_ = json.Unmarshal([]byte(`
{
    "recY5PsAPc": {
        "id": "recY5PsAPc",
        "record_id": "recY5PsAPc",
        "fields": {
            "Description": "用于测试的Gitlab项目",
            "WebUrl": {
                "text": "https://gitlab.com/wangyuheng77/integration",
                "link": "https://gitlab.com/wangyuheng77/integration"
            },
            "ID": "a",
            "Name": "Gitlab项目"
        }
    },
    "rec6KXl9eq": {
        "id": "rec6KXl9eq",
        "record_id": "rec6KXl9eq",
        "fields": {
            "WebUrl": {
                "text": "https://gitlab.com/wangyuheng77/integration",
                "link": "https://gitlab.com/wangyuheng77/integration"
            },
            "ID": "b",
            "Name": "Gitlab项目",
            "Description": "用于测试的Gitlab项目"
        }
    }
}
`), &res)
	return res, nil
}

func (f *fakerBitable) BatchCreateRecord(_ *core.Context, _ string, _ *larkBitable.AppTableRecordBatchCreateReqBody) ([]string, error) {
	return []string{"fldjX7dUj5"}, nil
}

func (f *fakerBitable) BatchUpdateRecord(_ *core.Context, _ string, _ *larkBitable.AppTableRecordBatchUpdateReqBody) error {
	return nil
}

func (f *fakerBitable) SyncRecords(_ *core.Context, _ string, _ []map[string]interface{}) error {
	return nil
}
