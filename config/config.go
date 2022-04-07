package config

import (
	"github.com/spf13/viper"
	"time"
)

var cfg = &Config{}

const (
	LarkAppId           = "LARK_APP_ID"
	LarkAppSecret       = "LARK_APP_SECRET"
	LarkBitableAppToken = "LARK_BITABLE_APP_TOKEN"

	LarkGitlabDomain     = "LARK_GITLAB_DOMAIN"
	LarkGitlabToken      = "LARK_GITLAB_TOKEN"
	LarkGitlabIssueLabel = "LARK_GITLAB_ISSUE_LABEL"

	LarkSyncPeriod = "LARK_SYNC_PERIOD"
)

type Config struct {
	AppId      string
	AppSecret  string
	SyncPeriod time.Duration
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

	_ = v.BindEnv(LarkAppId)
	_ = v.BindEnv(LarkAppSecret)
	_ = v.BindEnv(LarkSyncPeriod)
	_ = v.BindEnv(LarkBitableAppToken)
	_ = v.BindEnv(LarkGitlabDomain)
	_ = v.BindEnv(LarkGitlabToken)
	_ = v.BindEnv(LarkGitlabIssueLabel)

	cfg.AppId = v.GetString(LarkAppId)
	cfg.AppSecret = v.GetString(LarkAppSecret)
	cfg.SyncPeriod = v.GetDuration(LarkSyncPeriod)
	cfg.BitableConfig.AppToken = v.GetString(LarkBitableAppToken)
	cfg.GitlabConfig.Domain = v.GetString(LarkGitlabDomain)
	cfg.GitlabConfig.Token = v.GetString(LarkGitlabToken)
	cfg.GitlabConfig.IssueLabel = v.GetString(LarkGitlabIssueLabel)

	return cfg
}

func GetConfig() *Config {
	return cfg
}
