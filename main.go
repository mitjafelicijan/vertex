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

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

var sha1ver string
var buildTime string
var err error

var vm *otto.Otto
var config Config

var restAPIFiles []string
var restAPIRoutes map[string]string

func main() {

	fmt.Printf("%s\n\n", strings.Repeat("-", 70))
	fmt.Printf("Version:    Vertex - Mock REST API's the easy way\n")
	fmt.Printf("Repository: https://github.com/mitjafelicijan/vertex\n\n")

	fmt.Printf("Version:  %v\n", version)
	fmt.Printf("Built at: %v\n", date)
	fmt.Printf("SHA-1:    %v\n\n", commit)

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
	err := filepath.Walk(config.Vertex.Endpoints, func(path string, info os.FileInfo, err error) error {
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

	// api info about datastore and mounted routes
	r.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		response := make(map[string]interface{})
		response["routes"] = restAPIRoutes
		response["datastore"] = readDatastoreFile(config.Vertex.Datastore)

		responseJSON, _ := json.Marshal(response)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(responseJSON))
		return
	})

	// dynamically mount routes
	fmt.Printf("Mounting routes:\n")
	for route := range restAPIRoutes {
		fmt.Printf(" ↳ Registering route: `%s%s`\n", config.Vertex.Prefix, route)

		r.HandleFunc(fmt.Sprintf("%s%s", config.Vertex.Prefix, route), func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")

			// read file
			endpoint := strings.Replace(r.RequestURI, config.Vertex.Prefix, "", -1)
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

	// sandbox console
	r.HandleFunc("/sandbox", func(w http.ResponseWriter, r *http.Request) {
		style := "*{font: 14px Arial;}"
		body := ""
		//links := []string{}

		for route := range restAPIRoutes {
			fmt.Println(route)
			body += fmt.Sprintf("<li><a href='/sandbox/%s'>%s</a></li>", route, route)
		}
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf("%s <style>%s</style>", body, style)))
		return
	})

	// sandbox console
	r.HandleFunc("/sandbox/{endpoint}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		endpoint := vars["endpoint"]

		// read file
		//endpoint := strings.Replace(r.RequestURI, config.Vertex.Prefix, "", -1)
		source, err := ioutil.ReadFile(restAPIRoutes[endpoint])
		if err != nil {
			panic(err)
		}

		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf("<script>%s</script>", string(source))))
		return
	})

	// static file server
	r.Handle("/{url:.*}", http.FileServer(http.Dir(config.Vertex.Static)))

	// server handler
	server := &http.Server{
		Handler:      r,
		Addr:         fmt.Sprintf(":%d", config.Vertex.Port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Printf("\nWeb application: http://%s:%d", config.Vertex.Host, config.Vertex.Port)
	fmt.Printf("\nAPI endpoints:   http://%s:%d%s", config.Vertex.Host, config.Vertex.Port, config.Vertex.Prefix[:(len(config.Vertex.Prefix)-1)])
	fmt.Printf("\nSandbox env:     http://%s:%d/sandbox\n\n", config.Vertex.Host, config.Vertex.Port)
	log.Fatal(server.ListenAndServe())

}
