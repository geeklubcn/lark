package config

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestConfig(t *testing.T) {

	t.Run("should get cfg val after Load", func(t *testing.T) {
		assert.Equal(t, "", GetConfig().Token)
		_ = os.Setenv(LarkGitlabToken, "abc")
		defer func() { _ = os.Unsetenv(LarkGitlabToken) }()
		assert.Equal(t, "", GetConfig().Token)
		Load()
		assert.Equal(t, "abc", GetConfig().Token)
	})

}
