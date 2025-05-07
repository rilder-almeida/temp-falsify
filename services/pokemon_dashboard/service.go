package pokemon_dashboard

import "context"

// Service provides an pokemon_dashboard gateway for endpoint layer
type Service interface {
	Process(context.Context, ServiceRequest) (ServiceResponse, error)
}

type service struct{
	Repository Repository
}

// NewService instantiate a new pokemon_dashboard service
func NewService(repository Repository) Service {
	return &service{
		Repository: repository,
	}
}

// Process is responsible for return related data from service adapters
func (s *service) Process(context.Context, ServiceRequest) (ServiceResponse, error) {
	return ServiceResponse{}, nil
}
