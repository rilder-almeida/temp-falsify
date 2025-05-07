package v1

import (
	"context"

	ctxMap "github.com/arquivei/foundationkit/contextmap"
	"github.com/rs/zerolog"
)

// HTTPResponse represents the pokemon_dashboard response on the application layer
//
// swagger:response HTTPResponse
type HTTPResponse struct {
	// in:body
	Body BodyResponse
}

// BodyResponse represents the pokemon_dashboard body response on the application layer
type BodyResponse struct {
	JobID string `json:"jobid"`
}

// EnrichLog enriches log with service response
func (r HTTPResponse) EnrichLog(
	ctx context.Context,
	zctx zerolog.Context,
) (context.Context, zerolog.Context) {
	return ctx, zctx.
		Str("contextmap", ctxMap.Ctx(ctx).String())
}
