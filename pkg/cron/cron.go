package cron

import (
	"github.com/robfig/cron"
	"goMian/pkg/log"
	"log/slog"
)

func Init() error {
	c := cron.New()
	if err := c.AddFunc("@daily", func() {
		if err := log.Init(); err != nil {
			slog.Error(err.Error())
		}
	}); err != nil {
		return err
	}
	c.Start()
	return nil
}
