package models

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"

	"db-performance-project/internal/models"
	"db-performance-project/internal/pkg"
)

//go:generate easyjson -disallow_unknown_fields getdetailspost.go

type PostGetDetailsRequest struct {
	ID      uint32
	Related []string
}

func NewPostGetDetailsRequest() *PostGetDetailsRequest {
	return &PostGetDetailsRequest{}
}

func (req *PostGetDetailsRequest) Bind(r *http.Request) error {
	// if r.Header.Get("Content-Type") != "" {
	//	return pkg.ErrUnsupportedMediaType
	// }

	vars := mux.Vars(r)

	param := vars["id"]

	value, _ := strconv.Atoi(param)
	// if err != nil {
	//	return pkg.ErrBadRequestParams
	// }

	req.ID = uint32(value)

	param = r.URL.Query().Get("related")

	req.Related = strings.Split(param, ",")
	// for _, arg := range args {
	//	if arg != "user" && arg != "forum" && arg != "thread" {
	//		return pkg.ErrBadRequestParams
	//	}
	// }

	return nil
}

func (req *PostGetDetailsRequest) GetPost() *models.Post {
	return &models.Post{
		ID: req.ID,
	}
}

func (req *PostGetDetailsRequest) GetParams() *pkg.PostDetailsParams {
	return &pkg.PostDetailsParams{
		Related: req.Related,
	}
}

//easyjson:json
type PostGetDetailsAuthorResponse struct {
	Nickname string `json:"nickname"`
	FullName string `json:"fullname"`
	About    string `json:"about"`
	Email    string `json:"email"`
}

//easyjson:json
type PostGetDetailsPostResponse struct {
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
type PostGetDetailsThreadResponse struct {
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
type PostGetDetailsForumResponse struct {
	Title   string `json:"title"`
	User    string `json:"user"`
	Slug    string `json:"slug"`
	Posts   uint32 `json:"posts"`
	Threads uint32 `json:"threads"`
}

//easyjson:json
type PostGetDetailsResponse struct {
	Post   PostGetDetailsPostResponse   `json:"post"`
	Thread PostGetDetailsThreadResponse `json:"thread"`
	Author PostGetDetailsAuthorResponse `json:"author"`
	Forum  PostGetDetailsForumResponse  `json:"forum"`
}

func NewPostDetailsResponse(postDetails *models.PostDetails) *PostGetDetailsResponse {
	return &PostGetDetailsResponse{
		Post: PostGetDetailsPostResponse{
			ID:       postDetails.Post.ID,
			Parent:   postDetails.Post.Parent,
			Author:   postDetails.Post.Author.Nickname,
			Forum:    postDetails.Post.Forum,
			Thread:   postDetails.Post.Thread,
			Message:  postDetails.Post.Message,
			Created:  postDetails.Post.Created,
			IsEdited: postDetails.Post.IsEdited,
		},
		Author: PostGetDetailsAuthorResponse{
			Nickname: postDetails.Author.Nickname,
			FullName: postDetails.Author.FullName,
			About:    postDetails.Author.About,
			Email:    postDetails.Author.Email,
		},
		Thread: PostGetDetailsThreadResponse{
			ID:      postDetails.Thread.ID,
			Title:   postDetails.Thread.Title,
			Author:  postDetails.Thread.Author,
			Forum:   postDetails.Thread.Forum,
			Slug:    postDetails.Thread.Slug,
			Message: postDetails.Thread.Message,
			Created: postDetails.Thread.Created,
			Votes:   postDetails.Thread.Votes,
		},
		Forum: PostGetDetailsForumResponse{
			Title:   postDetails.Forum.Title,
			User:    postDetails.Forum.User,
			Slug:    postDetails.Forum.Slug,
			Posts:   postDetails.Forum.Posts,
			Threads: postDetails.Forum.Threads,
		},
	}
}
