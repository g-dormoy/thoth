package main

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
