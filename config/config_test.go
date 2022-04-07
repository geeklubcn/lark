package config

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"
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

	t.Run("config default val", func(t *testing.T) {
		_ = os.Unsetenv(LarkGitlabIssueLabel)
		Load()
		assert.Equal(t, "lark", GetConfig().IssueLabel)
		assert.Equal(t, 60*time.Second, GetConfig().SyncPeriod)

		_ = os.Setenv(LarkSyncPeriod, "10m")
		defer func() { _ = os.Unsetenv(LarkSyncPeriod) }()
		Load()
		assert.Equal(t, 10*time.Minute, GetConfig().SyncPeriod)
	})
}
