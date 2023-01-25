package handlers

import (
	"net/http"

	"github.com/gorilla/mux"

	"db-performance-project/internal/pkg"
	"db-performance-project/internal/service/service"
)

type serviceClearHandler struct {
	service service.Service
}

func NewServiceClearHandler(s service.Service) pkg.Handler {
	return &serviceClearHandler{
		s,
	}
}

func (h *serviceClearHandler) Configure(r *mux.Router, mw *pkg.HTTPMiddleware) {
	r.HandleFunc("/api/service/clear", h.Action).Methods(http.MethodPost)
}

func (h *serviceClearHandler) Action(w http.ResponseWriter, r *http.Request) {
	err := h.service.Clear(r.Context())
	if err != nil {
		pkg.DefaultHandlerHTTPError(r.Context(), w, err)
		return
	}

	pkg.NoBody(w, http.StatusOK)
}
