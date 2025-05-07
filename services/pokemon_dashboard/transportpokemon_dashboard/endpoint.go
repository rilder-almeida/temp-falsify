package transportpokemon_dashboard

import (
	"context"
	"reflect"
	
	"github.com/arquivei/foundationkit/errors"
	"github.com/go-kit/kit/endpoint"
	
	"gitlab.com/arquivei/stark/templates/api/services/pokemon_dashboard"
	v1 "gitlab.com/arquivei/stark/templates/api/services/pokemon_dashboard/transportpokemon_dashboard/internal/v1"
)

// MakeEndpoint returns the endpoint responsible for pokemon_dashboard
func MakeEndpoint(service pokemon_dashboard.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		const op = errors.Op("transportpokemon_dashboard.Endpoint")

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