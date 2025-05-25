package core

import (
	"log"
	"server/config"
	"server/utils"

	"gopkg.in/yaml.v3"
)

func InitConf() *config.Config {
	c := &config.Config{}
	yamlConf, err := utils.LoadYAML()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}
	if err = yaml.Unmarshal(yamlConf, c); err != nil {
		log.Fatalf("Failed to unmarshal YAML configutation: %v", err)
	}
	return c
}
