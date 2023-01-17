package handlers

import (
	"net/http"

	"github.com/gorilla/mux"

	"db-performance-project/internal/forum/delivery/models"
	"db-performance-project/internal/forum/service"
	"db-performance-project/internal/pkg"
)

type forumCreateSlugHandler struct {
	forumService service.ForumService
}

func NewForumCreateSlugHandler(s service.ForumService) pkg.Handler {
	return &forumCreateSlugHandler{
		s,
	}
}

func (h *forumCreateSlugHandler) Configure(r *mux.Router, mw *pkg.HTTPMiddleware) {
	r.HandleFunc("/forum/create", h.Action).Methods(http.MethodPost)
}

func (h *forumCreateSlugHandler) Action(w http.ResponseWriter, r *http.Request) {
	request := models.NewForumCreateRequest()

	err := request.Bind(r)
	if err != nil {
		pkg.DefaultHandlerHTTPError(r.Context(), w, err)
		return
	}

	forum, err := h.forumService.CreateForum(r.Context(), request.GetForum())
	if err != nil {
		pkg.DefaultHandlerHTTPError(r.Context(), w, err)
		return
	}

	response := models.NewForumCreateResponse(forum)

	pkg.Response(r.Context(), w, http.StatusOK, response)
}
