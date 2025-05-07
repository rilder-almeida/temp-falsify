package main

import (
	"net/http"

	"github.com/arquivei/foundationkit/contextmap"
	"github.com/arquivei/foundationkit/gokitmiddlewares/loggingmiddleware"
	"github.com/arquivei/foundationkit/gokitmiddlewares/metricsmiddleware"
	"github.com/go-kit/kit/endpoint"
	"github.com/rs/zerolog/log"

	"gitlab.com/arquivei/stark/templates/api/services/pokemon_dashboard"
	"gitlab.com/arquivei/stark/templates/api/services/pokemon_dashboard/adapterspokemon_dashboard/noopadapter"
	"gitlab.com/arquivei/stark/templates/api/services/pokemon_dashboard/transportpokemon_dashboard"
)

var (
	pokemon_dashboardHTTPHandler http.Handler
	pokemon_dashboardEndpoint endpoint.Endpoint
	pokemon_dashboardService pokemon_dashboard.Service
	pokemon_dashboardRepository pokemon_dashboard.Repository
)

func mustInitpokemon_dashboard() {
	registerHandler("/pokedex_verse/pokemon_dashboard", mustGetpokemon_dashboardHandler)
}

func getpokemon_dashboardAdapter() pokemon_dashboard.Repository {
	if pokemon_dashboardRepository == nil {
		pokemon_dashboardRepository = noopadapter.NewNoOpAdapter()
	}
	return pokemon_dashboardRepository
}

func getpokemon_dashboardService() pokemon_dashboard.Service {
	if pokemon_dashboardService == nil {
		pokemon_dashboardService = pokemon_dashboard.NewService(getpokemon_dashboardAdapter())
	}
	return pokemon_dashboardService
}

func getpokemon_dashboardEndpoint() endpoint.Endpoint {
	if pokemon_dashboardEndpoint == nil {
		loggingConfig := loggingmiddleware.NewDefaultConfig("pokedex_verse")
		metricsConfig := metricsmiddleware.NewDefaultConfig("pokedex_verse", "pokemon_dashboard")

		pokemon_dashboardEndpoint = endpoint.Chain(
			contextmap.NewEndpointMiddleware(),
			loggingmiddleware.MustNew(loggingConfig),
			metricsmiddleware.MustNew(metricsConfig),
			)(transportpokemon_dashboard.MakeEndpoint(getpokemon_dashboardService()))
	}
	return pokemon_dashboardEndpoint
}

func mustGetpokemon_dashboardHandler() http.Handler {
	if pokemon_dashboardHTTPHandler == nil {
		var err error
		pokemon_dashboardHTTPHandler, err = transportpokemon_dashboard.NewHTTPHandler(
			getpokemon_dashboardEndpoint(),
		)
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to create new server")
		}
	}
	return pokemon_dashboardHTTPHandler
}

