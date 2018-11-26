package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"strings"
)

// curl -X GET 'http://localhost:4001/api/products'
// curl -X GET 'http://localhost:4001/api/products?id=1&sub=4'
func getHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	urlParts := strings.Split(r.RequestURI, "?")
	endpoint := strings.Replace(urlParts[0], config.Vertex.Prefix, "", -1)

	// extracts query params
	requestQueryParams := make(map[string]string)
	for param := range r.URL.Query() {
		if reflect.TypeOf(param).Kind() == reflect.String {
			requestQueryParams[param] = r.URL.Query().Get(param)
		}
	}

	// convert query params into json
	requestQueryParamsJSON, err := json.Marshal(r.URL.Query())
	if err != nil {
		log.Println(err)
	}

	// read file
	source, err := ioutil.ReadFile(restAPIRoutes[endpoint+":get"].Filepath)
	if err != nil {
		panic(err)
	}

	// replaces unsupported methods
	routeSource := transcode(string(source))

	// executes script
	vm.Set("queryParams", string(requestQueryParamsJSON))
	response, err := vm.Run(routeSource)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response.String()))
	return
}

// curl -X POST 'http://localhost:4001/api/products' --data '{"username":"xyz","password":"xyz"}'
// curl -X PUT 'http://localhost:4001/api/products' --data '{"username":"xyz","password":"xyz"}'
// curl -X DELETE 'http://localhost:4001/api/products' --data '{"username":"xyz","password":"xyz"}'
func postPutDeleteHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	urlParts := strings.Split(r.RequestURI, "?")
	endpoint := strings.Replace(urlParts[0], config.Vertex.Prefix, "", -1)

	// extracts query params
	requestQueryParams := make(map[string]string)
	for param := range r.URL.Query() {
		if reflect.TypeOf(param).Kind() == reflect.String {
			requestQueryParams[param] = r.URL.Query().Get(param)
		}
	}

	// convert query params into json
	requestQueryParamsJSON, err := json.Marshal(r.URL.Query())
	if err != nil {
		log.Println(err)
	}

	// read file
	source, err := ioutil.ReadFile(restAPIRoutes[endpoint+":post"].Filepath)
	if err != nil {
		panic(err)
	}

	// replaces unsupported methods
	routeSource := transcode(string(source))

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// executes script
	vm.Set("body", string(body))
	vm.Set("queryParams", string(requestQueryParamsJSON))
	response, err := vm.Run(routeSource)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response.String()))

	return
}
