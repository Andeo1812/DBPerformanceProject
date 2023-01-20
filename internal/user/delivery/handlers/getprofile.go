package handlers

import (
	"net/http"

	"github.com/gorilla/mux"

	"db-performance-project/internal/pkg"
	"db-performance-project/internal/user/delivery/models"
	"db-performance-project/internal/user/service"
)

type userGetProfileHandler struct {
	userService service.UserService
}

func NewUserGetProfileHandler(s service.UserService) pkg.Handler {
	return &userGetProfileHandler{
		s,
	}
}

func (h *userGetProfileHandler) Configure(r *mux.Router, mw *pkg.HTTPMiddleware) {
	r.HandleFunc("/user/{nickname}/profile", h.Action).Methods(http.MethodGet)
}

func (h *userGetProfileHandler) Action(w http.ResponseWriter, r *http.Request) {
	request := models.NewProfileGetRequest()

	request.Bind(r)
	// err := request.Bind(r)
	// if err != nil {
	//	pkg.DefaultHandlerHTTPError(r.Context(), w, err)
	//	return
	// }

	user, err := h.userService.GetProfile(r.Context(), request.GetUser())
	if err != nil {
		pkg.DefaultHandlerHTTPError(r.Context(), w, err)
		return
	}

	response := models.NewProfileGetResponse(&user)

	pkg.Response(r.Context(), w, http.StatusOK, response)
}
