package handlers

import (
	"net/http"

	"github.com/gorilla/mux"

	"db-performance-project/internal/pkg"
	"db-performance-project/internal/post/delivery/models"
	"db-performance-project/internal/post/service"
)

type postGetDetailsHandler struct {
	postService service.PostService
}

func NewPostGetDetailsHandler(s service.PostService) pkg.Handler {
	return &postGetDetailsHandler{
		s,
	}
}

func (h *postGetDetailsHandler) Configure(r *mux.Router, mw *pkg.HTTPMiddleware) {
	r.HandleFunc("/post/{id}/details", h.Action).
		Methods(http.MethodGet)
}

func (h *postGetDetailsHandler) Action(w http.ResponseWriter, r *http.Request) {
	request := models.NewPostGetDetailsRequest()

	request.Bind(r)
	// err := request.Bind(r)
	// if err != nil {
	//	pkg.DefaultHandlerHTTPError(r.Context(), w, err)
	//	return
	// }

	postDetails, err := h.postService.GetDetailsPost(r.Context(), request.GetPost(), request.GetParams())
	if err != nil {
		pkg.DefaultHandlerHTTPError(r.Context(), w, err)
		return
	}

	response := models.NewPostDetailsResponse(postDetails)

	pkg.Response(r.Context(), w, http.StatusOK, response)
}
