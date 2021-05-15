package config

import (
	"log"
	"os"

	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v2"
)

// Config is to store the configuration for the service.
var Config *AntConfig

// InitConfig is to load the config file and environment variables
func InitConfig(fileName string) {
	loadLocalConfig(fileName)
	loadEnvConfig()
}

/////////////////////////////////////////////
// Load config files
////////////////////////////////////////////
func loadLocalConfig(fileName string) {
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Could not load config: %+v", err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&Config)
	if err != nil {
		log.Fatalf("Could not decode config: %+v", err)
	}
}

func loadEnvConfig() {
	err := envconfig.Process("", Config)
	if err != nil {
		log.Fatalf("Could not decode env config: %+v", err)
	}
}
