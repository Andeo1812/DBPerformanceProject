package models

//go:generate easyjson -all -disallow_unknown_fields forum.go

type Forum struct {
	ID      uint32 `json:"id,omitempty"`
	Title   string `json:"title,omitempty"`
	User    string `json:"user,omitempty"`
	Slug    string `json:"slug,omitempty"`
	Posts   uint32 `json:"posts,omitempty"`
	Threads uint32 `json:"threads,omitempty"`
}
