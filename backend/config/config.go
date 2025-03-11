package config

import (
	"gopkg.in/yaml.v2"
	"log"
	"os"
	"xiaoyun/backend/types/system_types"
)

var (
	configFile = "config.yaml"
	Config     system_types.Config
	Version    = "v1.0.8"
)

func init() {
	data, err := os.ReadFile(configFile)
	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}
	err = yaml.Unmarshal(data, &Config)
	if err != nil {
		log.Fatalf("Error unmarshalling config: %v", err)
	}

}
