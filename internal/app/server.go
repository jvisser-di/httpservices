package app

import (
	"github.com/jvisser-di/httpservices/internal/logging"
	"net/http"
)

func NewServer(
	logger *logging.Logger,
) http.Handler {
	mux := http.NewServeMux()
	addRoutes(mux, logger)
	var handler http.Handler = mux
	return handler
}
