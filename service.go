package main

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"gopkg.in/go-playground/validator.v9"
)

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

// LoadCollectionFromFile loads a service collection from a given file
func LoadCollectionFromFile(path string) (*ServiceCollection, error) {
	c, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	sc := &ServiceCollection{}
	err = json.Unmarshal(c, &sc.Services)
	if err != nil {
		return nil, err
	}

	return sc, nil
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

// Persist the collection in a file
func (sc ServiceCollection) Persist(p string) error {
	f, err := os.OpenFile(p, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		return err
	}

	err = json.NewEncoder(f).Encode(sc.Services)
	f.Close()
	return err
}
