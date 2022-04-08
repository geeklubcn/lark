package main

import (
	"context"
	"github.com/geeklubcn/lark/bitable"
	"github.com/geeklubcn/lark/config"
	"github.com/geeklubcn/lark/define"
	"github.com/geeklubcn/lark/gitlab"
	"github.com/geeklubcn/lark/syncer"
	"github.com/sirupsen/logrus"
	"time"
)

func main() {
	cfg := config.Load()
	logrus.SetLevel(cfg.LogLevel)
	g := gitlab.NewGitlabClient(cfg.Domain, cfg.Token, cfg.IssueLabel)
	b := bitable.NewBitable(cfg.AppId, cfg.AppSecret, cfg.AppToken)
	s := syncer.NewSyncer(b, g)

	logrus.Info("============LARK RUNNING============")

	dsl, _ := define.NewParser().Parse("./define.yaml")
	_ = s.SyncSchemaToRemoteBitable(context.Background(), &dsl)

	c := make(chan int, 1)
	q := make(chan int, 1)
	go listen(c)
	go func() {
		for {
			select {
			case <-time.Tick(cfg.SyncPeriod):
				logrus.Debug("============start sync============")
				ctx := context.Background()
				if err := s.Sync(ctx); err != nil {
					logrus.WithContext(ctx).Error("sync fail!")
				}
				logrus.Debug("============end sync============")
			}
		}
	}()
	<-q
}

func listen(_ chan int) {
	//TODO handle hook
}
