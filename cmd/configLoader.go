package cmd

import (
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Conf struct {
	SyncPairs []SyncPair `yaml:"syncPairs"`
}
type SyncPair struct {
	Src  string `yaml:"src"`
	Dest string `yamlx:"dest"`
}

func GetConf() Conf {
	log.Printf("Loading config....")

	env := os.Getenv("ENV")
	if len(env) == 0 {
		env = "local"
	}
	log.Printf("ENV: " + env + "...")
	confFilePath := "configs/" + env + ".yaml"
	log.Printf("Reading config file " + confFilePath + "....")
	c := Conf{}
	configFile, err := ioutil.ReadFile(confFilePath)
	if err != nil {
		log.Printf("jsonFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal([]byte(configFile), &c)
	if err != nil {
		log.Fatalf("Unmarshal err: %v", err)
	}
	return c
}
