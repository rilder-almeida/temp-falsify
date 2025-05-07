package transport${API_SERVICE_NAME}

import (
	"${PROJECT_GO_MODULE}/services/${API_SERVICE_NAME}"
	v1 "${PROJECT_GO_MODULE}/services/${API_SERVICE_NAME}/transport${API_SERVICE_NAME}/internal/v1"
)

// EndpointResponse is the response for ${API_SERVICE_NAME} endpoint layer
type EndpointResponse struct {
	${API_SERVICE_NAME}.ServiceResponse
}

// serviceResponseToEndpointResponse transforms a response from service layer to a response of the endpoint layer
func serviceResponseToEndpointResponse(response ${API_SERVICE_NAME}.ServiceResponse) EndpointResponse {
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