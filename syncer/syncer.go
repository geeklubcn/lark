package syncer

import (
	"github.com/geeklubcn/lark/bitable"
	"github.com/geeklubcn/lark/define"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type Syncer interface {
	SchemaSyncer
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
