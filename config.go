package main

import (
	"fmt"
	"io/ioutil"
	"os"

	yaml "gopkg.in/yaml.v2"
)

// Config struct
type Config struct {
	Vertex struct {
		Host      string `yaml:"host"`
		Port      int    `yaml:"port"`
		Prefix    string `yaml:"prefix"`
		Datastore string `yaml:"datastore"`
		Static    string `yaml:"static"`
		Endpoints string `yaml:"endpoints"`
	}
}

var configTemplate string = `vertex:

  # server settings
  host: 0.0.0.0
  port: 4001
  
  # api endpoint root
  prefix: /api/
  
  # database file location
  datastore: data.json
  
  # folder for html, css, js
  static: ./static
  
  # folder for rest api endpoints
  endpoints: ./endpoints`

func parseConfigFile(filename string) Config {
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {

		err = ioutil.WriteFile(filename, []byte(configTemplate), 0644)
		if err != nil {
			panic("Fatal error creating config file\n")
			os.Exit(1)
		}

		fmt.Println("\nNew config file successfully created.")
		fmt.Println("Open vertex.yml and change settings to your needs.")
		fmt.Println("Relaunch with ./vertex\n")
		os.Exit(1)
	}

	var config Config
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		fmt.Println("Error parsing config file\n")
		os.Exit(1)
	}

	return config
}
