package handlers

import (
	"net/http"

	"github.com/gorilla/mux"

	"db-performance-project/internal/pkg"
	"db-performance-project/internal/thread/delivery/models"
	"db-performance-project/internal/thread/service"
)

type threadUpdateDetailsHandler struct {
	threadService service.ThreadService
}

func NewThreadUpdateDetailsHandler(s service.ThreadService) pkg.Handler {
	return &threadUpdateDetailsHandler{
		s,
	}
}

func (h *threadUpdateDetailsHandler) Configure(r *mux.Router, mw *pkg.HTTPMiddleware) {
	r.HandleFunc("/thread/{slug_or_id}/details", h.Action).Methods(http.MethodPost)
}

func (h *threadUpdateDetailsHandler) Action(w http.ResponseWriter, r *http.Request) {
	request := models.NewThreadUpdateDetailsRequest()

	request.Bind(r)
	// err := request.Bind(r)
	// if err != nil {
	//	pkg.DefaultHandlerHTTPError(r.Context(), w, err)
	//	return
	// }

	thread, err := h.threadService.UpdateThread(r.Context(), request.GetThread())
	if err != nil {
		pkg.DefaultHandlerHTTPError(r.Context(), w, err)
		return
	}

	response := models.NewThreadUpdateDetailsResponse(thread)

	pkg.Response(r.Context(), w, http.StatusOK, response)
}
