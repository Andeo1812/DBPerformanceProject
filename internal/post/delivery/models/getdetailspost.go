package models

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"

	"db-performance-project/internal/models"
	"db-performance-project/internal/pkg"
)

//go:generate easyjson -omit_empty -disallow_unknown_fields getdetailspost.go

type PostGetDetailsRequest struct {
	ID      int64
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

	req.ID = int64(value)

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
	Nickname string `json:"nickname,omitempty"`
	FullName string `json:"fullname,omitempty"`
	About    string `json:"about,omitempty"`
	Email    string `json:"email,omitempty"`
}

//easyjson:json
type PostGetDetailsPostResponse struct {
	ID       int64  `json:"id,omitempty"`
	Parent   int64  `json:"parent,omitempty"`
	Author   string `json:"author,omitempty"`
	Message  string `json:"message,omitempty"`
	IsEdited bool   `json:"isEdited,omitempty"`
	Forum    string `json:"forum,omitempty"`
	Thread   int64  `json:"thread,omitempty"`
	Created  string `json:"created,omitempty"`
}

//easyjson:json
type PostGetDetailsThreadResponse struct {
	ID      int64  `json:"id,omitempty"`
	Title   string `json:"title,omitempty"`
	Author  string `json:"author,omitempty"`
	Forum   string `json:"forum,omitempty"`
	Slug    string `json:"slug,omitempty"`
	Message string `json:"message,omitempty"`
	Created string `json:"created,omitempty"`
	Votes   int64  `json:"votes,omitempty"`
}

//easyjson:json
type PostGetDetailsForumResponse struct {
	Title   string `json:"title,omitempty"`
	User    string `json:"user,omitempty"`
	Slug    string `json:"slug,omitempty"`
	Posts   int64  `json:"posts,omitempty"`
	Threads int64  `json:"threads,omitempty"`
}

//easyjson:json
type PostGetDetailsResponse struct {
	Post   *PostGetDetailsPostResponse   `json:"post,omitempty"`
	Thread *PostGetDetailsThreadResponse `json:"thread,omitempty"`
	Author *PostGetDetailsAuthorResponse `json:"author,omitempty"`
	Forum  *PostGetDetailsForumResponse  `json:"forum,omitempty"`
}

func NewPostDetailsResponse(postDetails *models.PostDetails) PostGetDetailsResponse {
	post := PostGetDetailsPostResponse{
		ID:       postDetails.Post.ID,
		Parent:   postDetails.Post.Parent,
		Author:   postDetails.Post.Author.Nickname,
		Forum:    postDetails.Post.Forum,
		Thread:   postDetails.Post.Thread,
		Message:  postDetails.Post.Message,
		Created:  postDetails.Post.Created,
		IsEdited: postDetails.Post.IsEdited,
	}

	author := PostGetDetailsAuthorResponse{
		Nickname: postDetails.Author.Nickname,
		FullName: postDetails.Author.FullName,
		About:    postDetails.Author.About,
		Email:    postDetails.Author.Email,
	}

	thread := PostGetDetailsThreadResponse{
		ID:      postDetails.Thread.ID,
		Title:   postDetails.Thread.Title,
		Author:  postDetails.Thread.Author,
		Forum:   postDetails.Thread.Forum,
		Slug:    postDetails.Thread.Slug,
		Message: postDetails.Thread.Message,
		Created: postDetails.Thread.Created,
		Votes:   postDetails.Thread.Votes,
	}

	forum := PostGetDetailsForumResponse{
		Title:   postDetails.Forum.Title,
		User:    postDetails.Forum.User,
		Slug:    postDetails.Forum.Slug,
		Posts:   postDetails.Forum.Posts,
		Threads: postDetails.Forum.Threads,
	}

	return PostGetDetailsResponse{
		Post:   &post,
		Author: &author,
		Thread: &thread,
		Forum:  &forum,
	}
}
