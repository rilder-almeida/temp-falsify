package transport${API_SERVICE_NAME}

import (
	"${PROJECT_GO_MODULE}/services/${API_SERVICE_NAME}"
	v1 "${PROJECT_GO_MODULE}/services/${API_SERVICE_NAME}/transport${API_SERVICE_NAME}/internal/v1"
)

// EndpointRequest is the request for ${API_SERVICE_NAME} endpoint layer
type EndpointRequest struct {
	${API_SERVICE_NAME}.ServiceRequest
}

// endpointRequestToServiceRequest transforms a request from the endpoint layer to a request of the service layer
func endpointRequestToServiceRequest(request EndpointRequest) ${API_SERVICE_NAME}.ServiceRequest {
	return ${API_SERVICE_NAME}.ServiceRequest{
		Namespace: request.Namespace,
	}
}

// httpRequestToEndpointRequest transforms a request from the http layer to a request of the endpoint layer
func httpRequestToEndpointRequest(request v1.HTTPRequest) EndpointRequest {
	return EndpointRequest{
		ServiceRequest: ${API_SERVICE_NAME}.ServiceRequest{
			Namespace: request.Body.Namespace,
		},
	}
}
