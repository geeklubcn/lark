package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"time"
)

var cfg = &Config{}

const (
	LarkAppId      = "LARK_APP_ID"
	LarkAppSecret  = "LARK_APP_SECRET"
	LarkSyncPeriod = "LARK_SYNC_PERIOD"
	LarkLogLevel   = "LARK_LOG_LEVEL"

	LarkBitableAppToken = "LARK_BITABLE_APP_TOKEN"

	LarkGitlabDomain     = "LARK_GITLAB_DOMAIN"
	LarkGitlabToken      = "LARK_GITLAB_TOKEN"
	LarkGitlabIssueLabel = "LARK_GITLAB_ISSUE_LABEL"
)

type Config struct {
	AppId      string
	AppSecret  string
	SyncPeriod time.Duration
	LogLevel   logrus.Level
	GitlabConfig
	BitableConfig
}

type BitableConfig struct {
	AppToken   string
	SyncSchema bool
}
type GitlabConfig struct {
	Domain     string
	Token      string
	IssueLabel string
}

func Load() *Config {
	v := viper.New()
	v.AutomaticEnv()
	v.SetDefault(LarkSyncPeriod, "60s")
	v.SetDefault(LarkGitlabIssueLabel, "lark")
	v.SetDefault(LarkLogLevel, logrus.InfoLevel.String())

	_ = v.BindEnv(LarkAppId)
	_ = v.BindEnv(LarkAppSecret)
	_ = v.BindEnv(LarkSyncPeriod)
	_ = v.BindEnv(LarkLogLevel)
	_ = v.BindEnv(LarkBitableAppToken)
	_ = v.BindEnv(LarkGitlabDomain)
	_ = v.BindEnv(LarkGitlabToken)
	_ = v.BindEnv(LarkGitlabIssueLabel)

	cfg.AppId = v.GetString(LarkAppId)
	cfg.AppSecret = v.GetString(LarkAppSecret)
	cfg.SyncPeriod = v.GetDuration(LarkSyncPeriod)
	if l, err := logrus.ParseLevel(v.GetString(LarkLogLevel)); err == nil {
		cfg.LogLevel = l
	}
	cfg.BitableConfig.AppToken = v.GetString(LarkBitableAppToken)
	cfg.GitlabConfig.Domain = v.GetString(LarkGitlabDomain)
	cfg.GitlabConfig.Token = v.GetString(LarkGitlabToken)
	cfg.GitlabConfig.IssueLabel = v.GetString(LarkGitlabIssueLabel)

	return cfg
}

func GetConfig() *Config {
	return cfg
}
