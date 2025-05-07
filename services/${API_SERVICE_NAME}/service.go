package ${API_SERVICE_NAME}

import "context"

// Service provides an ${API_SERVICE_NAME} gateway for endpoint layer
type Service interface {
	Process(context.Context, ServiceRequest) (ServiceResponse, error)
}

type service struct{
	Repository Repository
}

// NewService instantiate a new ${API_SERVICE_NAME} service
func NewService(repository Repository) Service {
	return &service{
		Repository: repository,
	}
}

// Process is responsible for return related data from service adapters
func (s *service) Process(context.Context, ServiceRequest) (ServiceResponse, error) {
	return ServiceResponse{}, nil
}
