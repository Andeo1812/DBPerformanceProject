package handlers

import (
	"net/http"

	"github.com/gorilla/mux"

	"db-performance-project/internal/forum/delivery/models"
	"db-performance-project/internal/forum/service"
	"db-performance-project/internal/pkg"
)

type forumCreateThreadHandler struct {
	forumService service.ForumService
}

func NewForumCreateThreadHandler(s service.ForumService) pkg.Handler {
	return &forumCreateThreadHandler{
		s,
	}
}

func (h *forumCreateThreadHandler) Configure(r *mux.Router, mw *pkg.HTTPMiddleware) {
	r.HandleFunc("/forum/{slug}/create", h.Action).Methods(http.MethodPost)
}

func (h *forumCreateThreadHandler) Action(w http.ResponseWriter, r *http.Request) {
	request := models.NewForumCreateThreadRequest()

	request.Bind(r)
	// err := request.Bind(r)
	// if err != nil {
	//	pkg.DefaultHandlerHTTPError(r.Context(), w, err)
	//	return
	// }

	thread, err := h.forumService.CreateThread(r.Context(), request.GetThread())
	if err != nil {
		pkg.DefaultHandlerHTTPError(r.Context(), w, err)
		return
	}

	response := models.NewForumCreateThreadResponse(thread)

	pkg.Response(r.Context(), w, http.StatusOK, response)
}
