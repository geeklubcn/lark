package syncer

import (
	"context"
	"fmt"
	"github.com/geeklubcn/lark/bitable"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSchemaSyncer(t *testing.T) {
	ctx := context.Background()
	s := NewSyncer(bitable.FakerBitable)

	t.Run("fetch remote", func(t *testing.T) {
		def, _ := s.FetchSchemaByRemoteBitable(ctx)
		assert.NotNil(t, def.App)
		assert.NotNil(t, def.App.Tables)
		assert.NotNil(t, def.App.Tables[0].Views)
		assert.NotNil(t, def.App.Tables[0].Fields)
	})

	t.Run("dump", func(t *testing.T) {
		s.WithFileWriterFunc(func(filename string, data []byte) error {
			if data == nil {
				assert.Fail(t, "data must be not nil")
			}
			fmt.Println(string(data))
			return nil
		})
		err := s.Dump("")
		assert.Nil(t, err)
	})

}
