package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var defaultData = []byte("{}")

func readDatastoreFile(filename string) map[string]string {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		// creates empty file with empty object inside {}
		err := ioutil.WriteFile(filename, defaultData, 0644)
		if err != nil {
			panic("fatal error creating datastore file")
		}
		data = defaultData
	}

	// parsing json
	datamap := map[string]string{}
	err = json.Unmarshal(data, &datamap)
	if err != nil {
		fmt.Println(err)
	}

	return datamap
}

func writeDatastoreFile(datamap map[string]string, filename string) {
	datamapJSON, _ := json.Marshal(datamap)
	err = ioutil.WriteFile(filename, datamapJSON, 0644)
	if err != nil {
		panic("fatal error creating datastore file")
	}
}

func clearDatastoreFile(filename string) {
	err := ioutil.WriteFile(filename, defaultData, 0644)
	if err != nil {
		panic("fatal error creating datastore file")
	}
}

func qwe() {

	datamap := readDatastoreFile(config.Vertex.Datastore)
	datamap["111"] = "okidoki"
	writeDatastoreFile(datamap, config.Vertex.Datastore)

	fmt.Println(config.Vertex.Datastore)
}
