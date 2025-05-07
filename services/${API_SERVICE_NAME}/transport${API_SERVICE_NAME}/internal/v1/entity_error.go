package v1

import (
	"context"
	"net/http"

	"github.com/arquivei/foundationkit/apiutil"
	"github.com/arquivei/foundationkit/errors"
	kithttp "github.com/go-kit/kit/transport/http"
	"${PROJECT_GO_MODULE}/services/${API_SERVICE_NAME}"
)

const (
	// ErrCodeInvalidRequest is returned when the request is invalid due endpoint layer validation
	ErrCodeInvalidRequest = errors.Code("INVALID_REQUEST")
)

// HTTPResponseError é retornado em caso de erro. O campo `code` contém um código
// para ser usado no tratamento dos erros enquanto que o campo `message` contém um texto descritivo
// sobre o que aconteceu. O campo `message` não deve ser usado para comparação pois não há garantias
// sobre o formato da mensagem ou conteúdo.
//
// ```
// Error Code             | HTTP | Descrição
// =======================|======|==========
// INVALID_REQUEST        | 400  | Requisição inválida ou mal formatada
// INTERNAL_ERROR         | 500  | Alguma falha ocorreu e não foi tratada corretamente
// ```
//
// swagger:response HTTPResponseError
type HTTPResponseError struct {
	// in:body
	Body struct {
		Error apiutil.ErrorDescription `json:"error"`
	}
}

// GetHTTPServerOption returns a http server option with an error encoder
func GetHTTPServerOption() []kithttp.ServerOption {
	return []kithttp.ServerOption{
		kithttp.ServerErrorEncoder(
			apiutil.NewHTTPErrorJSONEncoder(
				getHTTPStatusCode,
				encodeHTTPError,
			),
		),
	}
}

func encodeHTTPError(ctx context.Context, err error) interface{} {
	return HTTPResponseError{
		Body: struct {
			Error apiutil.ErrorDescription `json:"error"`
		}{
			Error: apiutil.ParseError(err),
		},
	}.Body
}

func getHTTPStatusCode(err error) int {
	switch errors.GetCode(err) {
	case ${API_SERVICE_NAME}.ErrCodeInternalError:
		return http.StatusInternalServerError
	case ${API_SERVICE_NAME}.ErrCodeInvalidRequest:
		return http.StatusBadRequest
	}
	return apiutil.GetDefaultErrorHTTPStatusCode(err)
}
