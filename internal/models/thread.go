package models

//go:generate easyjson -all -disallow_unknown_fields thread.go

type Thread struct {
	ID      uint32 `json:"id,omitempty"`
	Title   string `json:"title,omitempty"`
	Author  string `json:"author,omitempty"`
	Forum   string `json:"forum,omitempty"`
	Slug    string `json:"slug,omitempty"`
	Message string `json:"message,omitempty"`
	Created string `json:"created,omitempty"`
	Votes   int32  `json:"votes,omitempty"`
}
