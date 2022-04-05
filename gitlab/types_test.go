package gitlab

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTypes(t *testing.T) {
	issues, _ := FakerGitlab.Issues()
	for _, issue := range issues {
		assert.NotNil(t, issue.Convert())
		assert.NotEmpty(t, issue.Convert().DueDate)
	}
}
