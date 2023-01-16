package models

type Forum struct {
	ID      uint32
	Title   string
	User    string
	Slug    string
	Posts   uint32
	Threads uint32
}
