package handlers

import (
	"net/http"

	"github.com/gorilla/mux"

	"db-performance-project/internal/pkg"
	"db-performance-project/internal/user/delivery/models"
	"db-performance-project/internal/user/service"
)

type userUpdateProfileHandler struct {
	userService service.UserService
}

func NewUserUpdateProfileHandler(s service.UserService) pkg.Handler {
	return &userUpdateProfileHandler{
		s,
	}
}

func (h *userUpdateProfileHandler) Configure(r *mux.Router, mw *pkg.HTTPMiddleware) {
	r.HandleFunc("/api/user/{nickname}/profile", h.Action).Methods(http.MethodPost)
}

func (h *userUpdateProfileHandler) Action(w http.ResponseWriter, r *http.Request) {
	request := models.NewProfileUpdateRequest()

	request.Bind(r)
	// err := request.Bind(r)
	// if err != nil {
	//	pkg.DefaultHandlerHTTPError(r.Context(), w, err)
	//	return
	// }

	user, err := h.userService.UpdateProfile(r.Context(), request.GetUser())
	if err != nil {
		pkg.DefaultHandlerHTTPError(r.Context(), w, err)
		return
	}

	response := models.NewProfileUpdateResponse(&user)

	pkg.Response(r.Context(), w, http.StatusOK, response)
}
