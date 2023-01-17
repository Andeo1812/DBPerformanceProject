package handlers

import (
	"net/http"

	"github.com/gorilla/mux"

	"db-performance-project/internal/pkg"
	"db-performance-project/internal/thread/delivery/models"
	"db-performance-project/internal/thread/service"
)

type threadGetPostsHandler struct {
	threadService service.ThreadService
}

func NewThreadGetPostsHandler(s service.ThreadService) pkg.Handler {
	return &threadGetPostsHandler{
		s,
	}
}

func (h *threadGetPostsHandler) Configure(r *mux.Router, mw *pkg.HTTPMiddleware) {
	r.HandleFunc("/thread/{slug_or_id}/posts", h.Action).
		Methods(http.MethodGet).
		Queries(
			"limit", "{limit}",
			"since", "{since}",
			"desc", "{desc}",
			"sort", "{sort}")
}

func (h *threadGetPostsHandler) Action(w http.ResponseWriter, r *http.Request) {
	request := models.NewThreadGetPostsRequest()

	request.Bind(r)
	// err := request.Bind(r)
	// if err != nil {
	//	pkg.DefaultHandlerHTTPError(r.Context(), w, err)
	//	return
	// }

	posts, err := h.threadService.GetPosts(r.Context(), request.GetThread(), request.GetParams())
	if err != nil {
		pkg.DefaultHandlerHTTPError(r.Context(), w, err)
		return
	}

	response := models.NewThreadGetPostsResponse(posts)

	pkg.Response(r.Context(), w, http.StatusOK, response)
}
