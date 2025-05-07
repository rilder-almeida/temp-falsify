package v1

import (
	"context"
	"fmt"
	"regexp"

	ctxMap "github.com/arquivei/foundationkit/contextmap"
	"github.com/rs/zerolog"
)

// HTTPRequest represents the ${API_SERVICE_NAME} request on the application layer
//
// swagger:parameters HTTPRequest
type HTTPRequest struct {
	// in:body
	Body BodyRequest
}

// BodyRequest represents the ${API_SERVICE_NAME} body request on the application layer
type BodyRequest struct {
	// Namespace é o namespace do usuário na arquitetura multi-tenancy
	// required: true
	Namespace string `json:"namespace"`
	// FlowIdentifier é o identificador da aplicação que está a realizar a request. Deve estar no formato system_subsystem_flow.
	// required: true
	FlowIdentifier string `json:"flowidentifier"`
}

// Validate check has a valid flow identidier and a valid namespace
func (r HTTPRequest) Validate() error {
	if !isValidFlowIdentifier(r.Body.FlowIdentifier) {
		return fmt.Errorf("invalid flow identifier: %s", r.Body.FlowIdentifier)
	}
	if !isValidNamespace(r.Body.Namespace) {
		return fmt.Errorf("namespace is empty")
	}
	return nil
}

// isValidNamespace checks if namespace is not empty
func isValidNamespace(namespace string) bool {
	return len(namespace) > 0
}

// isValidFlowIdentifier checks is the flow identidier is in a format system_subsystem_flow
func isValidFlowIdentifier(queryIdentifier string) bool {
	queryIdentifierRegex := regexp.MustCompile(`^[a-zA-Z0-9]+_[a-zA-Z0-9]+_[a-zA-Z0-9]+$`)
	return queryIdentifierRegex.MatchString(queryIdentifier)
}

// EnrichLog enriches log with service request
func (r HTTPRequest) EnrichLog(
	ctx context.Context,
	zctx zerolog.Context,
) (context.Context, zerolog.Context) {
	zctx = zctx.
		Str("flow_identifier", string(r.Body.FlowIdentifier)).
		Str("namespace", string(r.Body.Namespace)).
		Str("contextmap", ctxMap.Ctx(ctx).String())
	return ctx, zctx
}
