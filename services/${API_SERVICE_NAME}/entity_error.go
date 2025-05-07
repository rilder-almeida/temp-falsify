package ${API_SERVICE_NAME}

import "github.com/arquivei/foundationkit/errors"

const (
	// ErrCodeInternalError is returned when failed has occurred while getting response
	ErrCodeInternalError = errors.Code("INTERNAL_ERROR")

	// ErrCodeInvalidRequest is returned when the request is invalid due service layer validation
	ErrCodeInvalidRequest = errors.Code("INVALID_REQUEST")
)
