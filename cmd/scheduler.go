package cmd

import (
	"time"

	"github.com/go-co-op/gocron"
)

func RunScheduler() {

	s1 := gocron.NewScheduler(time.Local)
	s1.Every(10).Seconds().At("00:00").Do(RunSync)
	<-s1.Start()
}
