package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"sync"
)

// ServiceCollection represent a collection for Services
type ServiceCollection []Service

// Filter return a new collection based of the filtering function result
func (sc ServiceCollection) Filter(ff func(Service) bool) []Service {
	var nsc []Service

	for _, service := range sc {
		if ff(service) == true {
			nsc = append(nsc, service)
		}
	}

	return nsc
}

// Service Struct represent a service
type Service struct {
	Name string `json:"name"`
	IP   string `json:"ip"`
	Port string `json:"port"`
	Desc string `json:"desc,omitempty"`
	Type string `json:"type"`
}

// getHandler handler http request to add new service
func getHandler(mutex *sync.Mutex, services ServiceCollection) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// If not a PUT request, return an error
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		// Checkin the URL parameters
		// TO-DO Add clarity when name param is missing
		name := r.URL.Query().Get("name")
		if name != "" {
			w.WriteHeader(http.StatusBadRequest)
		}

		s := services.Filter(func(service Service) bool { return service.Name == name })
		_ = json.NewEncoder(w).Encode(s)
	}
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

		w.WriteHeader(http.StatusCreated)
	}
}

func main() {
	var services []Service
	mutex := &sync.Mutex{}
	http.HandleFunc("/add", addHandler(mutex, services))
	http.HandleFunc("/get", getHandler(mutex, services))
	http.ListenAndServe(":4242", nil)
}
