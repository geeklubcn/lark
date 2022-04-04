package syncer

import (
	"context"
	"encoding/json"
	"github.com/geeklubcn/lark/bitable"
	"github.com/geeklubcn/lark/define"
	"github.com/larksuite/oapi-sdk-go/core"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type Syncer interface {
	SchemaSyncer
}

type SchemaSyncer interface {
	FetchSchemaByRemoteBitable(context context.Context) (*define.Define, error)
	Dump(filename string) error
	WithMarshalFunc(marshal MarshalFunc)
	WithFileWriterFunc(write FileWriterFunc)
}

type syncer struct {
	b       bitable.Bitable
	marshal MarshalFunc
	write   FileWriterFunc
}

func NewSyncer(b bitable.Bitable) Syncer {
	return &syncer{
		b: b,
		marshal: func(d *define.Define) ([]byte, error) {
			return yaml.Marshal(d)
		},
		write: func(filename string, data []byte) error {
			return ioutil.WriteFile(filename, data, 0644)
		},
	}
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
