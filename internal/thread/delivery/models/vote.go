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

//go:generate easyjson -all -disallow_unknown_fields vote.go

type ThreadVoteRequest struct {
	SlugOrID string
	Nickname string `json:"nickname"`
	Voice    int32  `json:"voice"`
}

func NewThreadVoteRequest() *ThreadVoteRequest {
	return &ThreadVoteRequest{}
}

func (req *ThreadVoteRequest) Bind(r *http.Request) error {
	// if r.Header.Get("Content-Type") != "" {
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

	err = easyjson.Unmarshal(body, req)
	if err != nil {
		return pkg.ErrJSONUnexpectedEnd
	}

	return nil
}

func (req *ThreadVoteRequest) GetThread() *models.Thread {
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

func (req *ThreadVoteRequest) GetParams() *pkg.VoteParams {
	return &pkg.VoteParams{
		Nickname: req.Nickname,
		Voice:    req.Voice,
	}
}

type ThreadVoteResponse struct {
	ID      uint32 `json:"id"`
	Title   string `json:"title"`
	Author  string `json:"author"`
	Forum   string `json:"forum"`
	Slug    string `json:"slug"`
	Message string `json:"message"`
	Created string `json:"created"`
	Votes   int32  `json:"votes"`
}

func NewThreadVoteResponse(thread *models.Thread) *ThreadVoteResponse {
	return &ThreadVoteResponse{
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
