package handlers

import (
	"net/http"

	"github.com/gorilla/mux"

	"db-performance-project/internal/forum/delivery/models"
	"db-performance-project/internal/forum/service"
	"db-performance-project/internal/pkg"
)

type forumGetThreadsHandler struct {
	forumService service.ForumService
}

func NewForumGetThreadsHandler(s service.ForumService) pkg.Handler {
	return &forumGetThreadsHandler{
		s,
	}
}

func (h *forumGetThreadsHandler) Configure(r *mux.Router, mw *pkg.HTTPMiddleware) {
	r.HandleFunc("/forum/{slug}/threads", h.Action).
		Methods(http.MethodGet)
}

func (h *forumGetThreadsHandler) Action(w http.ResponseWriter, r *http.Request) {
	request := models.NewForumGetThreadsRequest()

	request.Bind(r)
	// err := request.Bind(r)
	// if err != nil {
	//	pkg.DefaultHandlerHTTPError(r.Context(), w, err)
	//	return
	// }

	threads, err := h.forumService.GetThreads(r.Context(), request.GetForum(), request.GetParams())
	if err != nil {
		pkg.DefaultHandlerHTTPError(r.Context(), w, err)
		return
	}

	response := models.NewForumGetThreadsResponse(threads)

	pkg.Response(r.Context(), w, http.StatusOK, response)
}
