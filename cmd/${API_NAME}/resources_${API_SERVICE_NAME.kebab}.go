package main

import (
	"net/http"

	"github.com/arquivei/foundationkit/contextmap"
	"github.com/arquivei/foundationkit/gokitmiddlewares/loggingmiddleware"
	"github.com/arquivei/foundationkit/gokitmiddlewares/metricsmiddleware"
	"github.com/go-kit/kit/endpoint"
	"github.com/rs/zerolog/log"

	"${PROJECT_GO_MODULE}/services/${API_SERVICE_NAME}"
	"${PROJECT_GO_MODULE}/services/${API_SERVICE_NAME}/adapters${API_SERVICE_NAME}/noopadapter"
	"${PROJECT_GO_MODULE}/services/${API_SERVICE_NAME}/transport${API_SERVICE_NAME}"
)

var (
	${API_SERVICE_NAME}HTTPHandler http.Handler
	${API_SERVICE_NAME}Endpoint endpoint.Endpoint
	${API_SERVICE_NAME}Service ${API_SERVICE_NAME}.Service
	${API_SERVICE_NAME}Repository ${API_SERVICE_NAME}.Repository
)

func mustInit${API_SERVICE_NAME}() {
	registerHandler("/${API_NAME}/${API_SERVICE_NAME}", mustGet${API_SERVICE_NAME}Handler)
}

func get${API_SERVICE_NAME}Adapter() ${API_SERVICE_NAME}.Repository {
	if ${API_SERVICE_NAME}Repository == nil {
		${API_SERVICE_NAME}Repository = noopadapter.NewNoOpAdapter()
	}
	return ${API_SERVICE_NAME}Repository
}

func get${API_SERVICE_NAME}Service() ${API_SERVICE_NAME}.Service {
	if ${API_SERVICE_NAME}Service == nil {
		${API_SERVICE_NAME}Service = ${API_SERVICE_NAME}.NewService(get${API_SERVICE_NAME}Adapter())
	}
	return ${API_SERVICE_NAME}Service
}

func get${API_SERVICE_NAME}Endpoint() endpoint.Endpoint {
	if ${API_SERVICE_NAME}Endpoint == nil {
		loggingConfig := loggingmiddleware.NewDefaultConfig("${API_NAME}")
		metricsConfig := metricsmiddleware.NewDefaultConfig("${API_NAME}", "${API_SERVICE_NAME}")

		${API_SERVICE_NAME}Endpoint = endpoint.Chain(
			contextmap.NewEndpointMiddleware(),
			loggingmiddleware.MustNew(loggingConfig),
			metricsmiddleware.MustNew(metricsConfig),
			)(transport${API_SERVICE_NAME}.MakeEndpoint(get${API_SERVICE_NAME}Service()))
	}
	return ${API_SERVICE_NAME}Endpoint
}

func mustGet${API_SERVICE_NAME}Handler() http.Handler {
	if ${API_SERVICE_NAME}HTTPHandler == nil {
		var err error
		${API_SERVICE_NAME}HTTPHandler, err = transport${API_SERVICE_NAME}.NewHTTPHandler(
			get${API_SERVICE_NAME}Endpoint(),
		)
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to create new server")
		}
	}
	return ${API_SERVICE_NAME}HTTPHandler
}

