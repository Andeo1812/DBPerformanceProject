package models

//go:generate easyjson -all -disallow_unknown_fields user.go

type User struct {
	ID       uint32 `json:"id,omitempty"`
	Nickname string `json:"nickname,omitempty"`
	FullName string `json:"fullname,omitempty"`
	About    string `json:"about,omitempty"`
	Email    string `json:"email,omitempty"`
}
