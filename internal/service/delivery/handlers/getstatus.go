package handlers

import (
	"net/http"

	"github.com/gorilla/mux"

	"db-performance-project/internal/pkg"
	"db-performance-project/internal/service/delivery/models"
	"db-performance-project/internal/service/service"
)

type serviceGetStatusHandler struct {
	service service.Service
}

func NewServiceGetStatusHandler(s service.Service) pkg.Handler {
	return &serviceGetStatusHandler{
		s,
	}
}

func (h *serviceGetStatusHandler) Configure(r *mux.Router, mw *pkg.HTTPMiddleware) {
	r.HandleFunc("/api/service/status", h.Action).Methods(http.MethodGet)
}

func (h *serviceGetStatusHandler) Action(w http.ResponseWriter, r *http.Request) {
	status, err := h.service.GetStatus(r.Context())
	if err != nil {
		pkg.DefaultHandlerHTTPError(r.Context(), w, err)
		return
	}

	response := models.NewServiceGetStatusResponse(status)

	pkg.Response(r.Context(), w, http.StatusOK, response)
}
