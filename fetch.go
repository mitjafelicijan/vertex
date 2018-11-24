package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/robertkrimen/otto"
)

// fetches remote url and returns contents
func fetch(call otto.FunctionCall) otto.Value {

	// TODO: add validation of url
	// TODO: add checkup for method

	//method := call.Argument(0).String()
	url := call.Argument(1).String()

	// TODO: based on method do appropriate method
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	result, _ := vm.ToValue(string(body))
	return result
}
