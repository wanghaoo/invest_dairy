package service

import (
	"invest_dairy/common"

	"github.com/robfig/cron"
)

func InitCron() {
	var c = cron.New()
	// 每10秒钟调度任务
	err := c.AddFunc("*/3 * * * * ?", func() {
	})
	if err != nil {
		common.Mlog.Errorf("do cron 10second error: %s", err.Error())
	}
	c.Run()
}
