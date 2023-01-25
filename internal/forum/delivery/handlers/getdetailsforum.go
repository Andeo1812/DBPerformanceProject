package handlers

import (
	"net/http"

	"github.com/gorilla/mux"

	"db-performance-project/internal/forum/delivery/models"
	"db-performance-project/internal/forum/service"
	"db-performance-project/internal/pkg"
)

type forumGetDetailsHandler struct {
	forumService service.ForumService
}

func NewForumGetDetailsHandler(s service.ForumService) pkg.Handler {
	return &forumGetDetailsHandler{
		s,
	}
}

func (h *forumGetDetailsHandler) Configure(r *mux.Router, mw *pkg.HTTPMiddleware) {
	r.HandleFunc("/api/forum/{slug}/details", h.Action).Methods(http.MethodGet)
}

func (h *forumGetDetailsHandler) Action(w http.ResponseWriter, r *http.Request) {
	request := models.NewForumGetDetailsRequest()

	request.Bind(r)
	// err := request.Bind(r)
	// if err != nil {
	//	pkg.DefaultHandlerHTTPError(r.Context(), w, err)
	//	return
	// }

	forum, err := h.forumService.GetDetailsForum(r.Context(), request.GetForum())
	if err != nil {
		pkg.DefaultHandlerHTTPError(r.Context(), w, err)
		return
	}

	response := models.NewForumGetDetailsResponse(forum)

	pkg.Response(r.Context(), w, http.StatusOK, response)
}
