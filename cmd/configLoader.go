package cmd

import (
	"io/ioutil"
	"log"
	"os"
	"path"

	"gopkg.in/yaml.v2"
)

type Conf struct {
	SyncCron  string     `yaml:"syncCron"`
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

	configPath := os.Getenv("CONFIG_PATH")
	if len(configPath) == 0 {
		configPath = "configs"
	}

	log.Printf("ENV: " + env + "...")
	confFilePath := path.Join(configPath, env) + ".yaml"
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
