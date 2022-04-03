package define

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParser(t *testing.T) {
	definePath := "/tmp/define.yaml"
	defineContent := []byte(`
app:
  name: 项目管理甘特图
  tables:
    - revision: 3
      name: "\U0001F4A1  使用说明"
      fields:
        - field_name: 模块
          type: 1
        - field_name: 标签
          property:
            options:
              - name: 使用指南
                id: optbJNdSZZ
          type: 3
        - field_name: 附件
          type: 17
        - field_name: 描述
          type: 1
      views:
        - view_name: 使用指南
          view_type: gallery
`)
	if err := ioutil.WriteFile(definePath, defineContent, 0644); err != nil {
		t.FailNow()
	}
	define, err := NewParser().Parse(definePath)
	assert.NoError(t, err)
	assert.Equal(t, "项目管理甘特图", define.App.Name)
	assert.Equal(t, 3, define.App.Tables[0].Revision)
	assert.Equal(t, "gallery", define.App.Tables[0].Views[0].ViewType)
	assert.Equal(t, 4, len(define.App.Tables[0].Fields))
}
