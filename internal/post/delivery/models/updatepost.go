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

//go:generate easyjson -all -disallow_unknown_fields updatepost.go

type PostUpdateRequest struct {
	ID      uint32
	Message string `json:"message"`
}

func NewPostUpdateRequest() *PostUpdateRequest {
	return &PostUpdateRequest{}
}

func (req *PostUpdateRequest) Bind(r *http.Request) error {
	// if r.Header.Get("Content-Type") == "" {
	//	return pkg.ErrContentTypeUndefined
	// }
	//
	// if r.Header.Get("Content-Type") != pkg.ContentTypeJSON {
	//	return pkg.ErrUnsupportedMediaType
	// }

	vars := mux.Vars(r)

	param := vars["id"]

	value, _ := strconv.Atoi(param)
	// if err != nil {
	//	return pkg.ErrBadRequestParams
	// }

	req.ID = uint32(value)

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

	err = easyjson.Unmarshal(body, req)
	if err != nil {
		return pkg.ErrJSONUnexpectedEnd
	}

	return nil
}

func (req *PostUpdateRequest) GetPost() *models.Post {
	return &models.Post{
		ID:      req.ID,
		Message: req.Message,
	}
}

type PostUpdateResponse struct {
	ID       uint32 `json:"id"`
	Parent   uint32 `json:"parent"`
	Author   string `json:"author"`
	Message  string `json:"message"`
	IsEdited bool   `json:"isEdited"`
	Forum    string `json:"forum"`
	Thread   uint32 `json:"thread"`
	Created  string `json:"created"`
}

func NewPostUpdateResponse(post *models.Post) *PostUpdateResponse {
	return &PostUpdateResponse{
		ID:       post.ID,
		Parent:   post.Parent,
		Author:   post.Author.Nickname,
		Forum:    post.Forum,
		Thread:   post.Thread,
		Message:  post.Message,
		Created:  post.Created,
		IsEdited: post.IsEdited,
	}
}
