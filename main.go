package main

import (
	"fmt"
	"io/ioutil"

	"github.com/robertkrimen/otto"
	"github.com/robertkrimen/otto/underscore"
)

var vm *otto.Otto

func main() {

	fmt.Println("=================================>")

	source, err := ioutil.ReadFile("examples/general.js")
	if err != nil {
		panic(err)
	}

	vm = otto.New()

	underscore.Enable()

	// register fetch api
	vm.Set("fetch", fetch)

	// register local storage api's port
	vm.Set("localStorage.setItem", localStorageSetItem)
	vm.Set("localStorage.getItem", localStorageGetItem)
	vm.Set("localStorage.removeItem", localStorageRemoveItem)
	vm.Set("localStorage.clear", localStorageClear)

	value, err := vm.Run(string(source))
	if err != nil {
		fmt.Println(value)
	}
}
