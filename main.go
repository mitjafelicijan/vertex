package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/robertkrimen/otto"
	"github.com/robertkrimen/otto/underscore"
)

var sha1ver string
var buildTime string
var err error

var vm *otto.Otto
var config Config

var restAPIFiles []string
var restAPIRoutes map[string]string

func main() {

	fmt.Println("=================================>")
	//fmt.Println("Build on", buildTime)
	//fmt.Println("SHA-1", sha1ver)

	// parsing config file
	config = parseConfigFile("vertex.yml")

	// init new js vm
	vm = otto.New()

	// embedding underscore.js lib
	underscore.Enable()

	// register fetch api
	vm.Set("fetch", fetch)

	// register local storage api's port
	vm.Set("localStorage", new(LocalStorage))

	// list all js files in rest folder
	err := filepath.Walk(config.Vertex.Rest, func(path string, info os.FileInfo, err error) error {
		restAPIFiles = append(restAPIFiles, path)
		return nil
	})
	if err != nil {
		panic(err)
	}

	// parsing available rest api endpoints
	restAPIRoutes = make(map[string]string)
	for _, file := range restAPIFiles {
		if strings.Contains(file, ".js") {
			routeName := strings.Replace(filepath.Base(file), ".js", "", -1)
			restAPIRoutes[routeName] = file
		}

	}

	// declaring router
	r := mux.NewRouter()

	r.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		restAPIRoutesJSON, _ := json.Marshal(restAPIRoutes)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(restAPIRoutesJSON))
		return
	})

	// dynamically mount routes
	for route := range restAPIRoutes {
		fmt.Println(fmt.Sprintf("Registering route /%s", route))

		r.HandleFunc(fmt.Sprintf("/api/%s", route), func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")

			// read file
			endpoint := strings.Replace(r.RequestURI, config.Vertex.Root, "", -1)
			source, err := ioutil.ReadFile(restAPIRoutes[endpoint])
			if err != nil {
				panic(err)
			}

			// replaces unsupported methods
			routeSource := transcode(string(source))

			// executes script
			response, err := vm.Run(routeSource)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			w.WriteHeader(http.StatusOK)
			w.Write([]byte(response.String()))
			return
		})
	}

	// server handler
	server := &http.Server{
		Handler:      r,
		Addr:         fmt.Sprintf(":%d", config.Vertex.Port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("Listening on", config.Vertex.Port)
	log.Fatal(server.ListenAndServe())

}
