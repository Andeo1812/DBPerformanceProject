package models

import (
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mailru/easyjson"

	"db-performance-project/internal/models"
)

//go:generate easyjson -all -disallow_unknown_fields createthread.go

type ForumCreateThreadRequest struct {
	Slug    string
	Title   string `json:"title"`
	Author  string `json:"author"`
	Message string `json:"message"`
	Created string `json:"created"`
}

func NewForumCreateThreadRequest() *ForumCreateThreadRequest {
	return &ForumCreateThreadRequest{}
}

func (req *ForumCreateThreadRequest) Bind(r *http.Request) error {
	// if r.Header.Get("Content-Type") == "" {
	//	return pkg.ErrContentTypeUndefined
	// }
	//
	// if r.Header.Get("Content-Type") != pkg.ContentTypeJSON {
	//	return pkg.ErrUnsupportedMediaType
	// }

	vars := mux.Vars(r)

	req.Slug = vars["slug"]

	body, _ := io.ReadAll(r.Body)
	// if err != nil {
	//	return pkg.ErrBadBodyRequest
	// }
	// defer func() {
	//	err = r.Body.Close()
	//	if err != nil {
	//		logrus.Error(err)
	//	}
	// }()

	// if len(body) == 0 {
	//	return pkg.ErrEmptyBody
	// }

	easyjson.Unmarshal(body, req)
	// err = easyjson.Unmarshal(body, req)
	// if err != nil {
	//	return pkg.ErrJSONUnexpectedEnd
	// }

	return nil
}

func (req *ForumCreateThreadRequest) GetThread() *models.Thread {
	return &models.Thread{
		Slug:    req.Slug,
		Title:   req.Title,
		Author:  req.Author,
		Message: req.Message,
		Created: req.Created,
	}
}

type ForumCreateThreadResponse struct {
	ID      int64  `json:"id"`
	Title   string `json:"title"`
	Author  string `json:"author"`
	Forum   string `json:"forum"`
	Slug    string `json:"slug"`
	Message string `json:"message"`
	Created string `json:"created"`
	Votes   int64  `json:"votes"`
}

func NewForumCreateThreadResponse(forum *models.Thread) *ForumCreateThreadResponse {
	return &ForumCreateThreadResponse{
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
