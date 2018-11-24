package main

import (
	"io/ioutil"
	"log"
	"os"

	yaml "gopkg.in/yaml.v2"
)

// Config struct
type Config struct {
	Vertex struct {
		Host      string `yaml:"host"`
		Port      int    `yaml:"port"`
		Root      string `yaml:"root"`
		Datastore string `yaml:"datastore"`
		Static    string `yaml:"static"`
		Rest      string `yaml:"rest"`
	}
}

func parseConfigFile(filename string) Config {
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Println("Error reading config file")
		os.Exit(1)
	}

	var config Config
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		log.Println("Error parsing config file")
		os.Exit(1)
	}

	return config
}
