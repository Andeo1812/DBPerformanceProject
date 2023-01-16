package models

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"db-performanc-eproject/internal/models"
	"db-performanc-eproject/internal/pkg"
)

//go:generate easyjson -disallow_unknown_fields getslugusers.go

type ForumGetSlugUsersRequest struct {
	Slug  string
	Limit uint32
	Since string
	Desc  bool
}

func NewForumGetSlugUsersRequest() *ForumGetSlugUsersRequest {
	return &ForumGetSlugUsersRequest{}
}

func (req *ForumGetSlugUsersRequest) Bind(r *http.Request) error {
	// if r.Header.Get("Content-Type") != "" {
	//	return pkg.ErrUnsupportedMediaType
	// }

	vars := mux.Vars(r)

	req.Slug = vars["slug"]

	param := ""

	param = r.FormValue("limit")
	if param != "" {
		value, _ := strconv.Atoi(param)
		// if err != nil {
		//	return pkg.ErrConvertQueryType
		// }

		req.Limit = uint32(value)

		// if req.Limit > 10000 || req.Limit < 1 {
		//	return pkg.ErrBadRequestParams
		// }
	} else {
		req.Limit = 100
	}

	req.Since = r.FormValue("since")
	// if req.Since == "" {
	//	return pkg.ErrBadRequestParamsEmptyRequiredFields
	// }

	param = r.FormValue("desc")
	// if param == "" {
	//	return pkg.ErrBadRequestParamsEmptyRequiredFields
	// } else if param == "true" {
	//	req.Desc = true
	// } else if param == "false" {
	//	req.Desc = false
	// } else {
	//	return pkg.ErrBadRequestParams
	// }

	if param == "true" {
		req.Desc = true
	}

	return nil
}

func (req *ForumGetSlugUsersRequest) GetForum() *models.Forum {
	return &models.Forum{
		Slug: req.Slug,
	}
}

func (req *ForumGetSlugUsersRequest) GetParams() *pkg.GetSlugUsersParams {
	return &pkg.GetSlugUsersParams{
		Limit: req.Limit,
		Since: req.Since,
		Desc:  req.Desc,
	}
}

//easyjson:json
type ForumGetSlugUsersResponse struct {
	Nickname string `json:"nickname"`
	FullName string `json:"fullname"`
	About    string `json:"about"`
	Email    string `json:"email"`
}

//easyjson:json
type UsersList []ForumGetSlugUsersResponse

func NewForumGetSlugUsersResponse(users []*models.User) UsersList {
	res := make([]ForumGetSlugUsersResponse, len(users))

	for idx, value := range users {
		res[idx] = ForumGetSlugUsersResponse{
			Nickname: value.Nickname,
			FullName: value.FullName,
			About:    value.About,
			Email:    value.Email,
		}
	}

	return res
}
