package models

import (
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/mailru/easyjson"
	"github.com/sirupsen/logrus"

	"db-performance-project/internal/models"
	"db-performance-project/internal/pkg"
)

//go:generate easyjson -disallow_unknown_fields createthread.go

//easyjson:json
type PostRequest struct {
	Parent  uint32 `json:"parent"`
	Author  string `json:"author"`
	Message string `json:"message"`
}

//easyjson:json
type PostsRequestList []PostRequest

type ThreadCreateRequest struct {
	SlugOrID string
	Posts    PostsRequestList
}

func NewThreadCreateRequest() *ThreadCreateRequest {
	return &ThreadCreateRequest{}
}

func (req *ThreadCreateRequest) Bind(r *http.Request) error {
	// if r.Header.Get("Content-Type") == "" {
	//	return pkg.ErrContentTypeUndefined
	// }
	//
	// if r.Header.Get("Content-Type") != pkg.ContentTypeJSON {
	//	return pkg.ErrUnsupportedMediaType
	// }

	vars := mux.Vars(r)

	req.SlugOrID = vars["slug_or_id"]

	body, err := io.ReadAll(r.Body)
	if err != nil {
		return pkg.ErrBadBodyRequest
	}
	defer func() {
		err = r.Body.Close()
		if err != nil {
			logrus.Error(err)
		}
	}()

	// if len(body) == 0 {
	//	return pkg.ErrEmptyBody
	// }

	err = easyjson.Unmarshal(body, &req.Posts)
	if err != nil {
		return pkg.ErrJSONUnexpectedEnd
	}

	return nil
}

func (req *ThreadCreateRequest) GetPosts() []*models.Post {
	res := make([]*models.Post, len(req.SlugOrID))

	for idx, value := range req.Posts {
		res[idx] = &models.Post{
			Parent:  value.Parent,
			Message: value.Message,
			Author:  value.Author,
		}
	}

	return res
}

func (req *ThreadCreateRequest) GetThread() *models.Thread {
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

//easyjson:json
type PostResponse struct {
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
type PostsResponseList []PostResponse

func NewThreadCreateResponse(posts []*models.Post) PostsResponseList {
	res := make([]PostResponse, len(posts))

	for idx, value := range posts {
		res[idx] = PostResponse{
			ID:       value.ID,
			Parent:   value.Parent,
			Author:   value.Author,
			Forum:    value.Forum,
			IsEdited: value.IsEdited,
			Message:  value.Message,
			Created:  value.Created,
			Thread:   value.Thread,
		}
	}

	return res
}
