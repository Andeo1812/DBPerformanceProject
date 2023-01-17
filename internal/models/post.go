package models

type Post struct {
	ID       uint32
	Parent   uint32
	Author   User
	Message  string
	IsEdited bool
	Forum    string
	Thread   uint32
	Created  string
}
