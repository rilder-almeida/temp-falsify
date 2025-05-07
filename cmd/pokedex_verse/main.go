package main

import (
	"context"
	"time"

	"github.com/arquivei/go-app"
)

func main() {
	defer app.Recover()

	app.Bootstrap(version, &config)
	mustInitResources()

	otelShutdown := startOpenTelemetryExporter()
	app.RegisterShutdownHandler(&app.ShutdownHandler{
		Name:     "opentelemetry",
		Handler:  otelShutdown,
		Policy:   app.ErrorPolicyWarn,
		Priority: app.ShutdownPriority(0),
		Timeout:  time.Second,
	})

	app.RegisterShutdownHandler(
		&app.ShutdownHandler{
			Name:     "http_server",
			Priority: app.ShutdownPriority(100),
			Handler:  httpServer.Shutdown,
			Policy:   app.ErrorPolicyAbort,
		})

	app.RunAndWait(func(_ context.Context) error {
		return httpServer.ListenAndServe()
	})
}
