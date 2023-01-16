package models

import (
	"io"
	"net/http"

	"github.com/mailru/easyjson"
	"github.com/sirupsen/logrus"

	"db-performance-project/internal/models"
	"db-performance-project/internal/pkg"
)

//go:generate easyjson -all -disallow_unknown_fields createslug.go

type ForumSlugCreateRequest struct {
	Title   string `json:"title"`
	Author  string `json:"author"`
	Message string `json:"message"`
	Created string `json:"created"`
}

func NewForumSlugCreateRequest() *ForumSlugCreateRequest {
	return &ForumSlugCreateRequest{}
}

func (req *ForumSlugCreateRequest) Bind(r *http.Request) error {
	// if r.Header.Get("Content-Type") == "" {
	//	return pkg.ErrContentTypeUndefined
	// }
	//
	// if r.Header.Get("Content-Type") != pkg.ContentTypeJSON {
	//	return pkg.ErrUnsupportedMediaType
	// }

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

func (req *ForumSlugCreateRequest) GetThread() *models.Thread {
	return &models.Thread{
		Title:   req.Title,
		Author:  req.Author,
		Message: req.Message,
		Created: req.Created,
	}
}

type ForumSlugCreateResponse struct {
	ID      uint32 `json:"id"`
	Title   string `json:"title"`
	Author  string `json:"author"`
	Forum   string `json:"forum"`
	Slug    string `json:"slug"`
	Message string `json:"message"`
	Created string `json:"created"`
	Votes   int32  `json:"votes"`
}

func NewForumSlugCreateResponse(forum *models.Thread) *ForumSlugCreateResponse {
	return &ForumSlugCreateResponse{
		ID:      forum.ID,
		Title:   forum.Title,
		Author:  forum.Author,
		Forum:   forum.Forum,
		Slug:    forum.Slug,
		Message: forum.Message,
		Created: forum.Created,
		Votes:   forum.Votes,
	}
}
