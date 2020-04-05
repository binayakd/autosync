package cmd

import (
	"log"

	"github.com/mileusna/crontab"
)

// func RunSchedulerWithGoCron() {

// 	s1 := gocron.NewScheduler(time.Local)
// 	s1.Every(10).Seconds().At("00:00").Do(RunSync)
// 	<-s1.Start()
// }

func RunSchedulerWithCrontab(c *Conf) {
	log.Printf("Starting Scheduler with crontab: %s", c.SyncCron)
	cron := crontab.New()
	cron.MustAddJob(c.SyncCron, RunSync, c)

}
