package config

import "github.com/spf13/viper"

var cfg = &Config{}

const (
	LarkAppId           = "LARK_APP_ID"
	LarkAppSecret       = "LARK_APP_SECRET"
	LarkBitableAppToken = "LARK_BITABLE_APP_TOKEN"

	LarkGitlabDomain     = "LARK_GITLAB_DOMAIN"
	LarkGitlabToken      = "LARK_GITLAB_TOKEN"
	LarkGitlabIssueLabel = "LARK_GITLAB_ISSUE_LABEL"
)

type Config struct {
	AppId     string
	AppSecret string
	GitlabConfig
	BitableConfig
}

type BitableConfig struct {
	AppToken string
}
type GitlabConfig struct {
	Domain     string
	Token      string
	IssueLabel string
}

func Load() *Config {
	v := viper.New()
	v.AutomaticEnv()

	_ = v.BindEnv(LarkGitlabDomain)
	_ = v.BindEnv(LarkGitlabToken)
	_ = v.BindEnv(LarkGitlabIssueLabel)
	cfg.GitlabConfig.Domain = v.GetString(LarkGitlabDomain)
	cfg.GitlabConfig.Token = v.GetString(LarkGitlabToken)
	cfg.GitlabConfig.IssueLabel = v.GetString(LarkGitlabIssueLabel)

	_ = v.BindEnv(LarkAppId)
	_ = v.BindEnv(LarkAppSecret)
	_ = v.BindEnv(LarkBitableAppToken)
	cfg.AppId = v.GetString(LarkAppId)
	cfg.AppSecret = v.GetString(LarkAppSecret)
	cfg.BitableConfig.AppToken = v.GetString(LarkBitableAppToken)
	return cfg
}

func GetConfig() *Config {
	return cfg
}
