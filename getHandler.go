package main

import (
	"encoding/json"
	"net/http"
	"sync"
)

// getHandler handler http request to get a service
func getHandler(m *sync.Mutex, sc *ServiceCollection) http.HandlerFunc {
	return HTTPMiddlewarePipe(getHandlerFunc(m, sc), GetMiddleware)
}

// getHandlerFunc contains the real logic of the endpoint
func getHandlerFunc(m *sync.Mutex, sc *ServiceCollection) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Checkin the URL parameters
		// TO-DO Add clarity when name param is missing
		name := r.URL.Query().Get("name")
		if name == "" {
			w.WriteHeader(http.StatusBadRequest)
		}

		s := sc.Filter(func(service Service) bool { return service.Name == name })
		_ = json.NewEncoder(w).Encode(s)
	}
}
