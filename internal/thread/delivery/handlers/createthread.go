package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"

	"db-performance-project/internal/pkg"
	"db-performance-project/internal/thread/delivery/models"
	"db-performance-project/internal/thread/service"
)

type forumCreateThreadHandler struct {
	threadService service.ThreadService
}

func NewForumCreateThreadHandler(s service.ThreadService) pkg.Handler {
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

	thread, err := h.threadService.CreateThread(r.Context(), request.GetThread())
	if err != nil {
		if errors.Is(errors.Cause(err), pkg.ErrSuchThreadExist) {
			response := models.NewForumCreateThreadResponse(thread)

			pkg.Response(r.Context(), w, http.StatusConflict, response)

			return
		}

		pkg.DefaultHandlerHTTPError(r.Context(), w, err)

		return
	}

	response := models.NewForumCreateThreadResponse(thread)

	pkg.Response(r.Context(), w, http.StatusCreated, response)
}
