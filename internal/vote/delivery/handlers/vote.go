package handlers

import (
	"net/http"

	"github.com/gorilla/mux"

	"db-performance-project/internal/pkg"
	"db-performance-project/internal/vote/delivery/models"
	"db-performance-project/internal/vote/service"
)

type voteHandler struct {
	voteService service.VoteService
}

func NewThreadProfileHandler(s service.VoteService) pkg.Handler {
	return &voteHandler{
		s,
	}
}

func (h *voteHandler) Configure(r *mux.Router, mw *pkg.HTTPMiddleware) {
	r.HandleFunc("/api/thread/{slug_or_id}/vote", h.Action).Methods(http.MethodPost)
}

func (h *voteHandler) Action(w http.ResponseWriter, r *http.Request) {
	request := models.NewVoteRequest()

	request.Bind(r)
	// err := request.Bind(r)
	// if err != nil {
	//	pkg.DefaultHandlerHTTPError(r.Context(), w, err)
	//	return
	// }

	thread, err := h.voteService.Vote(r.Context(), request.GetThread(), request.GetParams())
	if err != nil {
		pkg.DefaultHandlerHTTPError(r.Context(), w, err)
		return
	}

	response := models.NewVoteResponse(&thread)

	pkg.Response(r.Context(), w, http.StatusOK, response)
}
