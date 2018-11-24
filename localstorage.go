package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/robertkrimen/otto"
)

// LocalStorage object
type LocalStorage struct{}

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

// SetItem -> JS usage: localStorage.setItem('lastname', 'Smith');
func (ls LocalStorage) SetItem(call otto.FunctionCall) otto.Value {

	key := call.Argument(0).String()
	value := call.Argument(1).String()

	datamap := readDatastoreFile(config.Vertex.Datastore)
	datamap[key] = value
	writeDatastoreFile(datamap, config.Vertex.Datastore)

	result, _ := vm.ToValue(true)
	return result
}

// GetItem -> JS usage: localStorage.getItem('lastname');
func (ls LocalStorage) GetItem(call otto.FunctionCall) otto.Value {
	key := call.Argument(0).String()
	datamap := readDatastoreFile(config.Vertex.Datastore)
	result, _ := vm.ToValue(datamap[key])
	return result
}

// RemoveItem -> JS usage: localStorage.removeItem('lastname');
func (ls LocalStorage) RemoveItem(call otto.FunctionCall) otto.Value {
	key := call.Argument(0).String()

	datamap := readDatastoreFile(config.Vertex.Datastore)
	delete(datamap, key)
	writeDatastoreFile(datamap, config.Vertex.Datastore)

	result, _ := vm.ToValue(datamap[key])
	return result
}

// Clear -> JS usage: localStorage.clear();
func (ls LocalStorage) Clear(call otto.FunctionCall) otto.Value {
	clearDatastoreFile(config.Vertex.Datastore)
	result, _ := vm.ToValue(true)
	return result
}
