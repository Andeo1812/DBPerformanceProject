package models

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"db-performanc-eproject/internal/models"
	"db-performanc-eproject/internal/pkg"
)

//go:generate easyjson -disallow_unknown_fields getslugdetails.go

type ForumGetSlugThreadsRequest struct {
	Slug  string
	Limit uint32
	Since string
	Desc  bool
}

func NewForumGetSlugThreadsRequest() *ForumGetSlugThreadsRequest {
	return &ForumGetSlugThreadsRequest{}
}

func (req *ForumGetSlugThreadsRequest) Bind(r *http.Request) error {
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
	} else {
		req.Limit = 100
	}

	// if err != nil {
	//	return pkg.ErrConvertQueryType
	// }

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

func (req *ForumGetSlugThreadsRequest) GetForum() *models.Forum {
	return &models.Forum{
		Slug: req.Slug,
	}
}

func (req *ForumGetSlugThreadsRequest) GetParams() *pkg.GetSlugThreadsParams {
	return &pkg.GetSlugThreadsParams{
		Limit: req.Limit,
		Since: req.Since,
		Desc:  req.Desc,
	}
}

//easyjson:json
type ForumGetSlugThreadsResponse struct {
	ID      uint32 `json:"id"`
	Title   string `json:"title"`
	Author  string `json:"author"`
	Forum   string `json:"forum"`
	Slug    string `json:"slug"`
	Message string `json:"message"`
	Created string `json:"created"`
	Votes   int32  `json:"votes"`
}

//easyjson:json
type ThreadsList []ForumGetSlugThreadsResponse

func NewForumGetSlugThreadsResponse(threads []*models.Thread) ThreadsList {
	res := make([]ForumGetSlugThreadsResponse, len(threads))

	for idx, value := range threads {
		res[idx] = ForumGetSlugThreadsResponse{
			ID:      value.ID,
			Title:   value.Title,
			Author:  value.Author,
			Forum:   value.Forum,
			Slug:    value.Slug,
			Message: value.Message,
			Created: value.Created,
			Votes:   value.Votes,
		}
	}

	return res
}
