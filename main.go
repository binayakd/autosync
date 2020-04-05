package main

import (
	"github.com/binayakd/autosync/cmd"
)

func main() {
	c := cmd.GetConf()
	cmd.RunSchedulerWithCrontab(&c)
	for {
	}
}
