package handlers

import (
	"net/http"

	"github.com/gorilla/mux"

	"db-performance-project/internal/pkg"
	"db-performance-project/internal/thread/delivery/models"
	"db-performance-project/internal/thread/service"
)

type threadCreatePostsHandler struct {
	threadService service.ThreadService
}

func NewThreadCreatePostsHandler(s service.ThreadService) pkg.Handler {
	return &threadCreatePostsHandler{
		s,
	}
}

func (h *threadCreatePostsHandler) Configure(r *mux.Router, mw *pkg.HTTPMiddleware) {
	r.HandleFunc("/thread/{slug_or_id}/create", h.Action).Methods(http.MethodPost)
}

func (h *threadCreatePostsHandler) Action(w http.ResponseWriter, r *http.Request) {
	request := models.NewThreadCreatePostsRequest()

	request.Bind(r)
	// err := request.Bind(r)
	// if err != nil {
	//	pkg.DefaultHandlerHTTPError(r.Context(), w, err)
	//	return
	// }

	posts, err := h.threadService.CreatePosts(r.Context(), request.GetThread(), request.GetPosts())
	if err != nil {
		pkg.DefaultHandlerHTTPError(r.Context(), w, err)
		return
	}

	response := models.NewThreadCreatePostsResponse(posts)

	pkg.Response(r.Context(), w, http.StatusOK, response)
}
