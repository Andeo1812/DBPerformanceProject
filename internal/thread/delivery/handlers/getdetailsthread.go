package handlers

import (
	"net/http"

	"github.com/gorilla/mux"

	"db-performance-project/internal/pkg"
	"db-performance-project/internal/thread/delivery/models"
	"db-performance-project/internal/thread/service"
)

type threadGetDetailsHandler struct {
	threadService service.ThreadService
}

func NewThreadGetDetailsHandler(s service.ThreadService) pkg.Handler {
	return &threadGetDetailsHandler{
		s,
	}
}

func (h *threadGetDetailsHandler) Configure(r *mux.Router, mw *pkg.HTTPMiddleware) {
	r.HandleFunc("/api/thread/{slug_or_id}/details", h.Action).Methods(http.MethodGet)
}

func (h *threadGetDetailsHandler) Action(w http.ResponseWriter, r *http.Request) {
	request := models.NewThreadGetDetailsRequest()

	request.Bind(r)
	// err := request.Bind(r)
	// if err != nil {
	//	pkg.DefaultHandlerHTTPError(r.Context(), w, err)
	//	return
	// }

	thread, err := h.threadService.GetDetailsThread(r.Context(), request.GetThread())
	if err != nil {
		pkg.DefaultHandlerHTTPError(r.Context(), w, err)
		return
	}

	response := models.NewThreadGetDetailsResponse(&thread)

	pkg.Response(r.Context(), w, http.StatusOK, response)
}
