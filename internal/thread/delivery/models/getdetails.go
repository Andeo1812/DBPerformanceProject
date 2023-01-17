package models

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"db-performance-project/internal/models"
)

//go:generate easyjson -disallow_unknown_fields getdetails.go

type ThreadGetDetailsRequest struct {
	SlugOrID string
}

func NewThreadGetDetailsRequest() *ThreadGetDetailsRequest {
	return &ThreadGetDetailsRequest{}
}

func (req *ThreadGetDetailsRequest) Bind(r *http.Request) error {
	// if r.Header.Get("Content-Type") != "" {
	//	return pkg.ErrUnsupportedMediaType
	// }

	vars := mux.Vars(r)

	req.SlugOrID = vars["slug_or_id"]

	return nil
}

func (req *ThreadGetDetailsRequest) GetThread() *models.Thread {
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
type ThreadGetDetailsResponse struct {
	ID      uint32 `json:"id"`
	Title   string `json:"title"`
	Author  string `json:"author"`
	Forum   string `json:"forum"`
	Slug    string `json:"slug"`
	Message string `json:"message"`
	Created string `json:"created"`
	Votes   int32  `json:"votes"`
}

func NewThreadGetDetailsResponse(thread *models.Thread) *ThreadGetDetailsResponse {
	return &ThreadGetDetailsResponse{
		ID:      thread.ID,
		Title:   thread.Title,
		Author:  thread.Author,
		Forum:   thread.Forum,
		Slug:    thread.Slug,
		Message: thread.Message,
		Created: thread.Created,
		Votes:   thread.Votes,
	}
}
