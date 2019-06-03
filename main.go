package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

// ServiceCollection represent a collection for Services
type ServiceCollection struct {
	Services []Service
}

// Filter return a new collection based of the filtering function result
func (sc ServiceCollection) Filter(ff func(Service) bool) ServiceCollection {
	nsc := ServiceCollection{}

	for _, service := range sc.Services {
		if ff(service) == true {
			nsc.Services = append(nsc.Services, service)
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
func getHandler(mutex *sync.Mutex, sc *ServiceCollection) func(w http.ResponseWriter, r *http.Request) {
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

		s := sc.Filter(func(service Service) bool { return service.Name == name })
		_ = json.NewEncoder(w).Encode(s)
	}
}

// addHandler handler http request to add new service
func addHandler(mutex *sync.Mutex, sc *ServiceCollection) func(w http.ResponseWriter, r *http.Request) {
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

		if string(body) == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err := json.Unmarshal(body, &service)
		if err != nil {
			// TO-DO handle logging better
			log.Printf("%s", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// Add the newly created service to the services collection
		mutex.Lock()
		sc.Services = append(sc.Services, service)
		mutex.Unlock()

		w.WriteHeader(http.StatusCreated)
		return
	}
}

func main() {
	services := &ServiceCollection{}
	mutex := &sync.Mutex{}
	http.HandleFunc("/add", addHandler(mutex, services))
	http.HandleFunc("/get", getHandler(mutex, services))
	http.ListenAndServe(":4242", nil)
}
