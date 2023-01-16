package handlers

import (
	"db-performanc-eproject/internal/forum/service"
	"net/http"

	"github.com/gorilla/mux"

	"db-performanc-eproject/internal/forum/delivery/models"
	"db-performanc-eproject/internal/pkg"
)

type forumCreateHandler struct {
	forumService service.ForumService
}

func NewForumCreateHandler(s service.ForumService) pkg.Handler {
	return &forumCreateHandler{
		s,
	}
}

func (h *forumCreateHandler) Configure(r *mux.Router, mw *pkg.HTTPMiddleware) {
	r.HandleFunc("/forum/create", h.Action).Methods(http.MethodPost)
}

func (h *forumCreateHandler) Action(w http.ResponseWriter, r *http.Request) {
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
