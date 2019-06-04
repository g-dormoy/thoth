package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

// addHandler handler http request to add new service
func addHandler(m *sync.Mutex, sc *ServiceCollection) http.HandlerFunc {
	return HTTPMiddlewarePipe(
		addHandlerFunc(m, sc),
		PutMiddleware,
		JSONMiddleware)
}

// addHandlerFunc contains the real logic of the endpoint
func addHandlerFunc(m *sync.Mutex, sc *ServiceCollection) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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
		if err = service.Validate(); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// Add the newly created service to the services collection
		m.Lock()
		sc.Services = append(sc.Services, service)
		m.Unlock()

		w.WriteHeader(http.StatusCreated)
	}
}
