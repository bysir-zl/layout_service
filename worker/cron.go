// 定时任务
package worker

import (
	"github.com/robfig/cron"
	"github.com/bysir-zl/bygo/log"
)

func startCheckUserLevelDown() {
	spec := "0 0 1 * * ? *"
	c := cron.New()
	c.AddFunc(spec, func() {

	})
	c.Start()

	log.Info("cron CheckUserLevelDown started, spec: " + spec)
}

func init() {
	startCheckUserLevelDown()
}
