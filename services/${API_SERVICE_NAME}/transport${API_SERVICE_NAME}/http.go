package transport${API_SERVICE_NAME}

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/arquivei/foundationkit/apiutil"
	"github.com/arquivei/foundationkit/errors"
	"github.com/go-kit/kit/endpoint"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"

	v1 "${PROJECT_GO_MODULE}/services/${API_SERVICE_NAME}/transport${API_SERVICE_NAME}/internal/v1"
)

// NewHTTPHandler returns a new HTTP handler for the ${API_SERVICE_NAME} service.
func NewHTTPHandler(endpoint endpoint.Endpoint) (http.Handler, error) {
	httpHandlerV1 := kithttp.NewServer(
		endpoint,
		decodeRequestV1,
		encodeResponseV1,
		v1.GetHTTPServerOption()...,
	)
	r := mux.NewRouter()

	// swagger:route /${API_NAME}/${API_SERVICE_ROUTE} ${API_SERVICE_METHOD} Request
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
	r.Handle("/${API_NAME}/${API_SERVICE_NAME}/${API_SERVICE_ROUTE}", httpHandlerV1).Methods("${API_SERVICE_METHOD}")
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