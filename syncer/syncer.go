package syncer

import (
	"github.com/geeklubcn/lark/gitlab"
	"io/ioutil"

	"github.com/geeklubcn/lark/bitable"
	"github.com/geeklubcn/lark/define"
	"gopkg.in/yaml.v3"
)

type Syncer interface {
	DataSyncer
	SchemaSyncer
}

type syncer struct {
	b       bitable.Bitable
	g       gitlab.Gitlab
	meta    Meta
	marshal MarshalFunc
	write   FileWriterFunc
}

func NewSyncer(b bitable.Bitable, g gitlab.Gitlab) Syncer {
	return &syncer{
		b:    b,
		g:    g,
		meta: &meta{b: b},
		marshal: func(d *define.Define) ([]byte, error) {
			return yaml.Marshal(d)
		},
		write: func(filename string, data []byte) error {
			return ioutil.WriteFile(filename, data, 0644)
		},
	}
}
