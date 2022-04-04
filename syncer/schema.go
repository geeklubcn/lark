package syncer

import (
	"context"
	"encoding/json"
	"github.com/geeklubcn/lark/define"
	"github.com/larksuite/oapi-sdk-go/core"
	larkBitable "github.com/larksuite/oapi-sdk-go/service/bitable/v1"
)

type SchemaSyncer interface {
	SyncSchemaToRemoteBitable(ctx context.Context, def *define.Define) error
	FetchSchemaByRemoteBitable(context context.Context) (*define.Define, error)
	Dump(filename string) error
	WithMarshalFunc(marshal MarshalFunc)
	WithFileWriterFunc(write FileWriterFunc)
}

type MarshalFunc func(*define.Define) ([]byte, error)

type FileWriterFunc func(filename string, data []byte) error

func (s *syncer) WithMarshalFunc(marshal MarshalFunc) {
	s.marshal = marshal
}
func (s *syncer) WithFileWriterFunc(write FileWriterFunc) {
	s.write = write
}

func (s *syncer) Dump(outPath string) error {
	ctx := core.WrapContext(context.Background())
	d, err := s.FetchSchemaByRemoteBitable(ctx)
	if err != nil {
		return err
	}
	data, err := s.marshal(d)
	if err != nil {
		return err
	}
	if err = s.write(outPath, data); err != nil {
		return err
	}
	return nil
}

func (s *syncer) SyncSchemaToRemoteBitable(ctx context.Context, def *define.Define) error {
	if err := s.syncDefineTable(ctx, def); err != nil {
		return err
	}
	if err := s.syncDefineTableView(ctx, def); err != nil {
		return err
	}
	if err := s.syncDefineTableField(ctx, def); err != nil {
		return err
	}
	return nil
}

func (s *syncer) syncDefineTable(c context.Context, def *define.Define) error {
	ctx := core.WrapContext(c)
	larkTables, _ := s.b.ListTables(ctx)
	prepare := &larkBitable.AppTableBatchCreateReqBody{}
	for _, t := range def.App.Tables {
		if _, ok := larkTables[t.Name]; !ok {
			prepare.Tables = append(prepare.Tables, &larkBitable.ReqTable{Name: t.Name})
		}
	}
	if len(prepare.Tables) > 0 {
		if _, err := s.b.BatchCreateTable(ctx, prepare); err != nil {
			return err
		}
	}
	return nil
}

func (s *syncer) syncDefineTableView(c context.Context, def *define.Define) error {
	ctx := core.WrapContext(c)
	larkTables, err := s.b.ListTables(ctx)
	if err != nil {
		return err
	}

	defTableMap := def.App.GetTableMap()
	for _, lt := range larkTables {
		if defTable, ok := defTableMap[lt.Name]; ok {
			larkTableViews, _ := s.b.ListViews(ctx, lt.TableId)

			views := make([]*larkBitable.AppTableView, 0)

			for _, defView := range defTable.Views {
				if _, vOk := larkTableViews[defView.ViewName]; !vOk {
					views = append(views, &larkBitable.AppTableView{
						ViewName: defView.ViewName,
						ViewType: defView.ViewType,
					})
				}
			}

			if len(views) > 0 {
				_ = s.b.SyncViews(ctx, lt.TableId, views)
			}
		}
	}
	return nil
}

func (s *syncer) syncDefineTableField(c context.Context, def *define.Define) error {
	ctx := core.WrapContext(c)
	larkTables, err := s.b.ListTables(ctx)
	if err != nil {
		return err
	}

	defTableMap := def.App.GetTableMap()
	for _, lt := range larkTables {
		if defTable, ok := defTableMap[lt.Name]; ok {
			larkTableFields, _ := s.b.ListFields(ctx, lt.TableId)

			fields := make([]*larkBitable.AppTableField, 0)

			for _, defField := range defTable.Fields {
				if _, fOk := larkTableFields[defField.FieldName]; !fOk {
					f := &larkBitable.AppTableField{
						FieldName: defField.FieldName,
						Type:      defField.Type,
					}
					if defField.Property != nil {
						f.Property = &larkBitable.AppTableFieldProperty{
							Options: func(ops []*define.FieldPropertyOption) []*larkBitable.AppTableFieldPropertyOption {
								res := make([]*larkBitable.AppTableFieldPropertyOption, 0)
								for _, op := range ops {
									res = append(res, &larkBitable.AppTableFieldPropertyOption{
										Name: op.Name,
									})
								}
								return res
							}(defField.Property.Options),
							Formatter:  defField.Property.Formatter,
							DateFormat: defField.Property.DateFormat,
							TimeFormat: defField.Property.TimeFormat,
							AutoFill:   defField.Property.AutoFill,
							Multiple:   defField.Property.Multiple,
							Fields:     defField.Property.Fields,
						}
					}

					fields = append(fields, f)
				}
			}

			for _, f := range fields {
				_, _ = s.b.CreateField(ctx, lt.TableId, f)

			}
		}
	}
	return nil
}

func (s *syncer) FetchSchemaByRemoteBitable(context context.Context) (*define.Define, error) {
	ctx := core.WrapContext(context)
	d := &define.Define{}
	app, err := s.b.GetApp(ctx)
	if err != nil {
		return nil, err
	}

	d.App = &define.App{
		Name: app.Name,
	}
	tables, err := s.b.ListTables(ctx)
	if err != nil {
		return nil, err
	}

	for _, table := range tables {
		fields, err := s.fetchFields(ctx, table.TableId)
		if err != nil {
			return nil, err
		}
		views, err := s.fetchViews(ctx, table.TableId)
		if err != nil {
			return nil, err
		}

		d.App.Tables = append(d.App.Tables, &define.Table{
			Revision: table.Revision,
			Name:     table.Name,
			Fields:   fields,
			Views:    views,
		})
	}
	return d, nil
}

func (s *syncer) fetchFields(ctx *core.Context, tableId string) ([]*define.Field, error) {
	fields, err := s.b.ListFields(ctx, tableId)
	if err != nil {
		return nil, err
	}

	res := make([]*define.Field, 0)
	for _, field := range fields {
		var f define.Field
		j, _ := json.Marshal(field)
		_ = json.Unmarshal(j, &f)
		res = append(res, &f)
	}
	return res, nil
}

func (s *syncer) fetchViews(ctx *core.Context, tableId string) ([]*define.View, error) {
	views, err := s.b.ListViews(ctx, tableId)
	if err != nil {
		return nil, err
	}

	res := make([]*define.View, 0)
	for _, view := range views {
		var v define.View
		j, _ := json.Marshal(view)
		_ = json.Unmarshal(j, &v)
		res = append(res, &v)
	}
	return res, nil
}
