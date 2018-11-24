package main

import "github.com/robertkrimen/otto"

// js: localStorage.setItem('lastname, 'Smith);
func localStorageSetItem(call otto.FunctionCall) otto.Value {
	result, _ := vm.ToValue(1)
	return result
}

// js: var user = localStorage.getItem('user');
func localStorageGetItem(call otto.FunctionCall) otto.Value {
	result, _ := vm.ToValue(1)
	return result
}

// js: localStorage.removeItem('user');
func localStorageRemoveItem(call otto.FunctionCall) otto.Value {
	result, _ := vm.ToValue(1)
	return result
}

// js: localStorage.clear();
func localStorageClear(call otto.FunctionCall) otto.Value {
	result, _ := vm.ToValue(1)
	return result
}