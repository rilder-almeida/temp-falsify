package transportpokemon_dashboard

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/arquivei/foundationkit/apiutil"
	"github.com/arquivei/foundationkit/errors"
	"github.com/go-kit/kit/endpoint"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"

	v1 "gitlab.com/arquivei/stark/templates/api/services/pokemon_dashboard/transportpokemon_dashboard/internal/v1"
)

// NewHTTPHandler returns a new HTTP handler for the pokemon_dashboard service.
func NewHTTPHandler(endpoint endpoint.Endpoint) (http.Handler, error) {
	httpHandlerV1 := kithttp.NewServer(
		endpoint,
		decodeRequestV1,
		encodeResponseV1,
		v1.GetHTTPServerOption()...,
	)
	r := mux.NewRouter()

	// swagger:route /pokedex_verse/fetch_pokemon POST Request
	//
	// [Alpha] Service title
	//
	// Sevice long description
	//
	// ---
	// produces:
	// - application/json
	//
	// responses:
	//   default: Response
	r.Handle("/pokedex_verse/pokemon_dashboard/fetch_pokemon", httpHandlerV1).Methods("POST")
	return r, nil
}

func decodeRequestV1(
	_ context.Context,
	r *http.Request,
) (any, error) {
	const op = errors.Op("decodeRequestV1")

	var httpRequest v1.HTTPRequest

	defer r.Body.Close()

	err := json.NewDecoder(r.Body).Decode(&httpRequest.Body)

	if err != nil {
		return nil, errors.E(op, err, errors.SeverityInput)
	}

	err = httpRequest.Validate()
	if err != nil {
		return nil, errors.E(op, err, errors.SeverityInput)
	}

	return httpRequestToEndpointRequest(httpRequest), nil
}

func encodeResponseV1(
	ctx context.Context,
	w http.ResponseWriter,
	r any,
) error {
	const op = errors.Op("encodeResponseV1")

	endpointResponse, ok := r.(EndpointResponse)
	if !ok {
		return errors.E(op, "fail to cast serviceResponse to EndpointResponseV1", errors.KV("serviceResponse", r))
	}

	httpResponse := endpointResponseToHTTPResponse(endpointResponse)

	err := apiutil.EncodeJSONResponse(ctx, w, httpResponse.Body)

	if err != nil {
		return errors.E(op, err)
	}

	return nil
}  