package app

import (
	"github.com/jvisser-di/httpservices/internal/logging"
	"net/http"
)

func addRoutes(
	mux *http.ServeMux,
	logger *logging.Logger,
) {
	mux.Handle("/helloworld", handleHelloWorld(logger))
	mux.Handle("/", http.NotFoundHandler())
}
