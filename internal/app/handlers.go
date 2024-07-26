package app

import (
	"github.com/jvisser-di/httpservices/internal/logging"
	"net/http"
)

func handleHelloWorld(logger *logging.Logger) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			logger.Info(r.Context(), "Hello World", "handler", "handleHelloWorld")
		})
}
