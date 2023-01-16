package models

//go:generate easyjson -all -disallow_unknown_fields special.go

type StatusService struct {
	User   uint32 `json:"user"`
	Forum  uint32 `json:"forum"`
	Thread uint32 `json:"thread"`
	Post   uint32 `json:"post"`
}

type PostDetails struct {
	Post   Post   `json:"post"`
	Author User   `json:"author,omitempty"`
	Thread Thread `json:"thread,omitempty"`
	Forum  Forum  `json:"forum,omitempty"`
}
