package handlers

import (
	"net/http"

	"github.com/gorilla/mux"

	"db-performance-project/internal/forum/delivery/models"
	"db-performance-project/internal/forum/service"
	"db-performance-project/internal/pkg"
)

type forumGetUsersHandler struct {
	forumService service.ForumService
}

func NewForumGetUsersHandler(s service.ForumService) pkg.Handler {
	return &forumGetUsersHandler{
		s,
	}
}

func (h *forumGetUsersHandler) Configure(r *mux.Router, mw *pkg.HTTPMiddleware) {
	r.HandleFunc("/forum/{slug}/users", h.Action).
		Methods(http.MethodGet).
		Queries(
			"limit", "{limit}",
			"since", "{since}",
			"desc", "{desc}")
}

func (h *forumGetUsersHandler) Action(w http.ResponseWriter, r *http.Request) {
	request := models.NewForumGetUsersRequest()

	request.Bind(r)
	// err := request.Bind(r)
	// if err != nil {
	//	pkg.DefaultHandlerHTTPError(r.Context(), w, err)
	//	return
	// }

	users, err := h.forumService.GetUsers(r.Context(), request.GetForum(), request.GetParams())
	if err != nil {
		pkg.DefaultHandlerHTTPError(r.Context(), w, err)
		return
	}

	response := models.NewForumGetUsersResponse(users)

	pkg.Response(r.Context(), w, http.StatusOK, response)
}
