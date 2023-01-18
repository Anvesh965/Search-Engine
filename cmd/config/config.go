package config

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type Configure struct {
	Server struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
	} `yaml:"server"`
	Database struct {
		Host       string `yaml:"host"`
		Port       int    `yaml:"port"`
		Protocol   string `yaml:"protocol"`
		DBName     string `yaml:"dbname"`
		Collection string `yaml:"collection"`
	} `yaml:"database"`
}

var Config Configure

func LoadConfig() {

	// Read config.yaml file
	data, err := ioutil.ReadFile("config/config.yaml")
	if err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}

	// Unmarshal YAML into config struct
	err = yaml.Unmarshal(data, &Config)
	if err != nil {
		log.Fatalf("Error unmarshalling config: %s", err)
	}

	// Print the config to verify it was read correctly
	fmt.Printf("Server: %s:%d\n", Config.Server.Host, Config.Server.Port)
	fmt.Printf("Database: %s:%d/%s\n", Config.Database.Host, Config.Database.Port, Config.Database.Protocol)

}
