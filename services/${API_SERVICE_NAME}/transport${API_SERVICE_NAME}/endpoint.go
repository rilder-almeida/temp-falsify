package transport${API_SERVICE_NAME}

import (
	"context"
	"reflect"
	
	"github.com/arquivei/foundationkit/errors"
	"github.com/go-kit/kit/endpoint"
	
	"${PROJECT_GO_MODULE}/services/${API_SERVICE_NAME}"
	v1 "${PROJECT_GO_MODULE}/services/${API_SERVICE_NAME}/transport${API_SERVICE_NAME}/internal/v1"
)

// MakeEndpoint returns the endpoint responsible for ${API_SERVICE_NAME}
func MakeEndpoint(service ${API_SERVICE_NAME}.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		const op = errors.Op("transport${API_SERVICE_NAME}.Endpoint")

		req, ok := request.(EndpointRequest)
		if !ok {
			return v1.HTTPResponse{}, errors.E(
				op,
				"invalid request type, must be EndpointRequest type",
				v1.ErrCodeInvalidRequest,
				errors.KV("type", reflect.TypeOf(req)),
			)
		}

		data, err := service.Process(ctx, endpointRequestToServiceRequest(req))
		return serviceResponseToEndpointResponse(data), err
	}
}