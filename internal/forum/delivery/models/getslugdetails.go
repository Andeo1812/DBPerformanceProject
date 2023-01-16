package models

import (
	"net/http"

	"github.com/gorilla/mux"

	"db-performanc-eproject/internal/models"
)

//go:generate easyjson -disallow_unknown_fields getslugthreads.go

type ForumGetSlugDetailsRequest struct {
	Slug string
}

func NewForumGetSlugDetailsRequest() *ForumGetSlugDetailsRequest {
	return &ForumGetSlugDetailsRequest{}
}

func (req *ForumGetSlugDetailsRequest) Bind(r *http.Request) error {
	// if r.Header.Get("Content-Type") != "" {
	//	return pkg.ErrUnsupportedMediaType
	// }

	vars := mux.Vars(r)

	req.Slug = vars["slug"]

	return nil
}

func (req *ForumGetSlugDetailsRequest) GetForum() *models.Forum {
	return &models.Forum{
		Slug: req.Slug,
	}
}

//easyjson:json
type ForumGetSlugDetailsResponse struct {
	Title   string `json:"title"`
	User    string `json:"user"`
	Slug    string `json:"slug"`
	Posts   uint32 `json:"posts"`
	Threads uint32 `json:"threads"`
}

func NewForumGetSlugDetailsResponse(forum *models.Forum) *ForumGetSlugDetailsResponse {
	return &ForumGetSlugDetailsResponse{
		Title:   forum.Title,
		User:    forum.User,
		Slug:    forum.Slug,
		Posts:   forum.Posts,
		Threads: forum.Threads,
	}
}
