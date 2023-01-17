package models

type StatusService struct {
	User   uint32
	Forum  uint32
	Thread uint32
	Post   uint32
}

type PostDetails struct {
	Post   Post
	Author User
	Thread Thread
	Forum  Forum
}
