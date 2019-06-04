package main

import "gopkg.in/go-playground/validator.v9"

// ServiceCollection represent a collection for Services
type ServiceCollection struct {
	Services []Service
}

// Service Struct represent a service
type Service struct {
	Name string `json:"name" validate:"required"`
	IP   string `json:"ip" validate:"required,ip"`
	Port string `json:"port"`
	Desc string `json:"desc,omitempty"`
	Type string `json:"type"`
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

// Validate validate the structure
func (s Service) Validate() error {
	validate := validator.New()

	return validate.Struct(&s)
}
