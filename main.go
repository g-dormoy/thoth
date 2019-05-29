package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"sync"
)

// Service Struct represent a service
type Service struct {
	Name string `json:"name"`
	IP   string `json:"ip"`
	Port string `json:"port"`
	Desc string `json:"desc"`
	Type string `json:"type"`
}

// addHandler handler http request to add new service
func addHandler(mutex *sync.Mutex, services []Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// If not a PUT request, return an error
		if r.Method != http.MethodPut {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		// Create a new service from request query
		// TO-DO add service spec validation
		service := Service{}
		body, _ := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		json.Unmarshal(body, &service)

		// Add the newly created service to the services collection
		mutex.Lock()
		services = append(services, service)
		mutex.Unlock()
	}
}

func main() {
	var services []Service
	mutex := &sync.Mutex{}
	http.HandleFunc("/add", addHandler(mutex, services))
	http.ListenAndServe(":4242", nil)
}
