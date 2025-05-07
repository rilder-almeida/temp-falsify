package transportpokemon_dashboard

import (
	"gitlab.com/arquivei/stark/templates/api/services/pokemon_dashboard"
	v1 "gitlab.com/arquivei/stark/templates/api/services/pokemon_dashboard/transportpokemon_dashboard/internal/v1"
)

// EndpointRequest is the request for pokemon_dashboard endpoint layer
type EndpointRequest struct {
	pokemon_dashboard.ServiceRequest
}

// endpointRequestToServiceRequest transforms a request from the endpoint layer to a request of the service layer
func endpointRequestToServiceRequest(request EndpointRequest) pokemon_dashboard.ServiceRequest {
	return pokemon_dashboard.ServiceRequest{
		Namespace: request.Namespace,
	}
}

// httpRequestToEndpointRequest transforms a request from the http layer to a request of the endpoint layer
func httpRequestToEndpointRequest(request v1.HTTPRequest) EndpointRequest {
	return EndpointRequest{
		ServiceRequest: pokemon_dashboard.ServiceRequest{
			Namespace: request.Body.Namespace,
		},
	}
}
