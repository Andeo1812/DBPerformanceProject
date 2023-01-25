package handlers

import (
	"net/http"

	"github.com/gorilla/mux"

	"db-performance-project/internal/pkg"
	"db-performance-project/internal/post/delivery/models"
	"db-performance-project/internal/post/service"
)

type postUpdateHandler struct {
	postService service.PostService
}

func NewPostUpdateHandler(s service.PostService) pkg.Handler {
	return &postUpdateHandler{
		s,
	}
}

func (h *postUpdateHandler) Configure(r *mux.Router, mw *pkg.HTTPMiddleware) {
	r.HandleFunc("/api/post/{id}/details", h.Action).Methods(http.MethodPost)
}

func (h *postUpdateHandler) Action(w http.ResponseWriter, r *http.Request) {
	request := models.NewPostUpdateRequest()

	request.Bind(r)
	// err := request.Bind(r)
	// if err != nil {
	//	pkg.DefaultHandlerHTTPError(r.Context(), w, err)
	//	return
	// }

	post, err := h.postService.UpdatePost(r.Context(), request.GetPost())
	if err != nil {
		pkg.DefaultHandlerHTTPError(r.Context(), w, err)
		return
	}

	response := models.NewPostUpdateResponse(post)

	pkg.Response(r.Context(), w, http.StatusOK, response)
}
