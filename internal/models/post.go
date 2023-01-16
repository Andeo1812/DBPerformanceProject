package models

//go:generate easyjson -all -disallow_unknown_fields post.go

type Post struct {
	ID       uint32 `json:"id,omitempty"`
	Parent   uint32 `json:"parent,omitempty"`
	Author   string `json:"author,omitempty"`
	Message  string `json:"message,omitempty"`
	IsEdited bool   `json:"isEdited,omitempty"`
	Forum    string `json:"forum,omitempty"`
	Thread   uint32 `json:"thread,omitempty"`
	Created  string `json:"created,omitempty"`
}
