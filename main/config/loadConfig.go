package config

import (
	"github.com/Gircen/go-library-api/main/config"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

func GetConfig() *config.Config {

	confPath, found := os.LookupEnv("MS_CONFIG")
	var yamlFile []byte

	if found {
		yamlFile = GetYml(confPath)
	} else {
		yamlFile = GetYml("config/application.yaml")
	}

	var cfg config.Config
	err := yaml.Unmarshal(yamlFile, &cfg)
	if err != nil {
		log.Fatalf("Error convert YAML to data: %v", err)
	}
	return &cfg
}

func GetYml(confPath string) []byte {
	yamlFile, err := os.ReadFile(confPath)
	if err != nil {
		log.Fatalf("Error reading YAML file: %v", err)
	}
	return yamlFile
}
