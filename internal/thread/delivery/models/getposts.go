package models

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"db-performance-project/internal/models"
	"db-performance-project/internal/pkg"
)

//go:generate easyjson -disallow_unknown_fields getposts.go

type ThreadGetPostsRequest struct {
	SlugOrID string
	Limit    uint32
	Since    string
	Desc     bool
	Sort     string
}

func NewForumGetSlugThreadsRequest() *ThreadGetPostsRequest {
	return &ThreadGetPostsRequest{}
}

func (req *ThreadGetPostsRequest) Bind(r *http.Request) error {
	// if r.Header.Get("Content-Type") != "" {
	//	return pkg.ErrUnsupportedMediaType
	// }

	vars := mux.Vars(r)

	req.SlugOrID = vars["slug_or_id"]

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

	req.Sort = r.FormValue("sort")
	if req.Sort == "" {
		req.Sort = "flat"
	}

	return nil
}

func (req *ThreadGetPostsRequest) GetThread() *models.Thread {
	id, err := strconv.Atoi(req.SlugOrID)
	if err != nil {
		res := uint32(id)

		return &models.Thread{
			ID: res,
		}
	}

	return &models.Thread{
		Slug: req.SlugOrID,
	}
}

func (req *ThreadGetPostsRequest) GetParams() *pkg.GetPostsParams {
	return &pkg.GetPostsParams{
		Limit: req.Limit,
		Since: req.Since,
		Desc:  req.Desc,
		Sort:  req.Sort,
	}
}

//easyjson:json
type ThreadGetPostsResponse struct {
	ID       uint32 `json:"id"`
	Parent   uint32 `json:"parent"`
	Author   string `json:"author"`
	Message  string `json:"message"`
	IsEdited bool   `json:"isEdited"`
	Forum    string `json:"forum"`
	Thread   uint32 `json:"thread"`
	Created  string `json:"created"`
}

//easyjson:json
type PostsList []ThreadGetPostsResponse

func NewThreadGetPostsResponse(posts []*models.Post) PostsList {
	res := make([]ThreadGetPostsResponse, len(posts))

	for idx, value := range posts {
		res[idx] = ThreadGetPostsResponse{
			ID:       value.ID,
			Parent:   value.Parent,
			Author:   value.Author.Nickname,
			Forum:    value.Forum,
			Thread:   value.Thread,
			Message:  value.Message,
			Created:  value.Created,
			IsEdited: value.IsEdited,
		}
	}

	return res
}
