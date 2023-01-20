package pkg

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Handler is an interface for universal handlers.
type Handler interface {
	Action(http.ResponseWriter, *http.Request)
	Configure(*mux.Router, *HTTPMiddleware)
}
