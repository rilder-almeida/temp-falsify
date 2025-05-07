package main

import (
	"net/http"

	"github.com/arquivei/foundationkit/httpmiddlewares/enrichloggingmiddleware"
	"github.com/arquivei/foundationkit/httpmiddlewares/trackingmiddleware"
	"github.com/arquivei/go-app"
	"github.com/gorilla/mux"
)

var (
	httpServer *http.Server
	router     *mux.Router

	handlers map[string]http.Handler
)

func mustInitResources() {
	handlers = make(map[string]http.Handler, 0)
	mustInitpokemon_dashboard()
	getRouter()
	getHTTPServer()
}

func getRouter() *mux.Router {
	if router == nil {
		router = mux.NewRouter()
		for pathPrefix, handler := range handlers {
			router.PathPrefix(pathPrefix).Handler(handler)
		}
		router.Use(trackingmiddleware.New, enrichloggingmiddleware.New)
	}
	return router
}

func getHTTPServer() *http.Server {
	if httpServer == nil {
		httpAddr := ":" + config.HTTP.Port
		httpServer = &http.Server{Addr: httpAddr, Handler: router}
		app.RegisterShutdownHandler(
			&app.ShutdownHandler{
				Name:     "http_server",
				Priority: app.ShutdownPriority(100),
				Handler:  httpServer.Shutdown,
				Policy:   app.ErrorPolicyAbort,
			},
		)
	}
	return httpServer
}

func registerHandler(pathPrefix string, handlerFn func() http.Handler) {
	handlers[pathPrefix] = handlerFn()
}
