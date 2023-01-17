package handlers

import (
	"net/http"

	"github.com/gorilla/mux"

	"db-performance-project/internal/pkg"
	"db-performance-project/internal/thread/delivery/models"
	"db-performance-project/internal/thread/service"
)

type threadVoteHandler struct {
	threadService service.ThreadService
}

func NewThreadProfileHandler(s service.ThreadService) pkg.Handler {
	return &threadVoteHandler{
		s,
	}
}

func (h *threadVoteHandler) Configure(r *mux.Router, mw *pkg.HTTPMiddleware) {
	r.HandleFunc("/thread/{slug_or_id}/vote", h.Action).Methods(http.MethodPost)
}

func (h *threadVoteHandler) Action(w http.ResponseWriter, r *http.Request) {
	request := models.NewThreadVoteRequest()

	request.Bind(r)
	// err := request.Bind(r)
	// if err != nil {
	//	pkg.DefaultHandlerHTTPError(r.Context(), w, err)
	//	return
	// }

	thread, err := h.threadService.Vote(r.Context(), request.GetThread(), request.GetParams())
	if err != nil {
		pkg.DefaultHandlerHTTPError(r.Context(), w, err)
		return
	}

	response := models.NewThreadVoteResponse(thread)

	pkg.Response(r.Context(), w, http.StatusOK, response)
}
