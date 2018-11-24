package main

import (
	"fmt"

	"github.com/robertkrimen/otto"
	"github.com/robertkrimen/otto/underscore"
)

var sha1ver string
var buildTime string
var err error
var vm *otto.Otto
var config Config

func main() {

	//fmt.Println("=================================>")
	//fmt.Println("Build on", buildTime)
	//fmt.Println("SHA-1", sha1ver)

	//dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(dir)

	// parsing config file
	config = parseConfigFile("vertex.yml")

	fmt.Println(config)

	fmt.Println("=================================>")

	//source, err := ioutil.ReadFile("examples/console.js")
	//if err != nil {
	//	panic(err)
	//}

	vm = otto.New()

	underscore.Enable()

	// register fetch api
	vm.Set("fetch", fetch)

	// register local storage api's port
	vm.Set("localStorage.setItem", localStorageSetItem)
	vm.Set("localStorage.getItem", localStorageGetItem)
	vm.Set("localStorage.removeItem", localStorageRemoveItem)
	vm.Set("localStorage.clear", localStorageClear)

	//value, err := vm.Run(string(source))
	//if err != nil {
	//	fmt.Println(value)
	//}

	qwe()
}
