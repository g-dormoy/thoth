package main

import "net/http"

// HTTPMiddleware defines a http middleware prototype
type HTTPMiddleware func(http.HandlerFunc) http.HandlerFunc

// HTTPMiddlewarePipe is an helper to execute middleware
func HTTPMiddlewarePipe(h http.HandlerFunc, m ...HTTPMiddleware) http.HandlerFunc {
	w := h

	for i := 0; i < len(m); i++ {
		w = m[i](w)
	}

	return w
}

// PutMiddleware ensure the Put Method is Used for an http.HandlerFunc
func PutMiddleware(f http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// If not a PUT request, return an error
		if r.Method != http.MethodPut {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		f.ServeHTTP(w, r)
	})
}

// GetMiddleware ensure the Get Method is used for and http.HandlerFunc
func GetMiddleware(f http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// If not a GET request, return an error
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		f.ServeHTTP(w, r)
	})
}

// JSONMiddleware ensure the Content-Type of the handler is application/json
func JSONMiddleware(f http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Content-Type") != "application/json" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		f.ServeHTTP(w, r)
	})
}
