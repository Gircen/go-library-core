package conf

import (
	"github.com/Gircen/go-library-api/main/config"
	"github.com/Gircen/go-library-core/main/logs"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

var Conf *config.Config

func GetConfig() {

	confPath, found := os.LookupEnv("MS_CONFIG")
	var yamlFile []byte

	if found {
		yamlFile = GetYml(confPath)
	} else {
		yamlFile = GetYml("config/application.yaml")
	}

	err := yaml.Unmarshal(yamlFile, Conf)
	if err != nil {
		log.Fatalf("Error convert YAML to data: %v", err)
	}

	logs.Logger = logs.Log()

}

func GetYml(confPath string) []byte {
	yamlFile, err := os.ReadFile(confPath)
	if err != nil {
		log.Fatalf("Error reading YAML file: %v", err)
	}
	return yamlFile
}
