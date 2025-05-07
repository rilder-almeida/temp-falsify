package transportpokemon_dashboard

import (
	"gitlab.com/arquivei/stark/templates/api/services/pokemon_dashboard"
	v1 "gitlab.com/arquivei/stark/templates/api/services/pokemon_dashboard/transportpokemon_dashboard/internal/v1"
)

// EndpointResponse is the response for pokemon_dashboard endpoint layer
type EndpointResponse struct {
	pokemon_dashboard.ServiceResponse
}

// serviceResponseToEndpointResponse transforms a response from service layer to a response of the endpoint layer
func serviceResponseToEndpointResponse(response pokemon_dashboard.ServiceResponse) EndpointResponse {
	return EndpointResponse{
		ServiceResponse: response,
	}
}

// endpointResponseToHTTPResponse transforms a response from endopoint layer to a response of the http layer
func endpointResponseToHTTPResponse(response EndpointResponse) v1.HTTPResponse {
	return v1.HTTPResponse{
		Body: v1.BodyResponse{
			JobID: response.JobID,
		},
	}
}