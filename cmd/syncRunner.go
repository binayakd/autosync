package cmd

import (
	"log"
	"os/exec"
)

func RunSync() {

	log.Printf("********* Staring Sync *********")
	c := GetConf()
	for int, entry := range c.SyncPairs {
		log.Printf("Sync Pair %v - src: %v, dest: %v", int+1, entry.Src, entry.Dest)
		command := exec.Command("rclone", "sync", entry.Src, entry.Dest, "-P")
		out, err := command.CombinedOutput()
		log.Printf("Sync output:\n%s\n", string(out))
		if err != nil {
			log.Printf("Sync failed with %s\n", err)
		}

	}
	log.Printf("--------- Sync complete --------")
}
