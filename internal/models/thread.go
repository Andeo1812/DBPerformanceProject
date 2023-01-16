package models

type Thread struct {
	ID      uint32
	Title   string
	Author  string
	Forum   string
	Slug    string
	Message string
	Created string
	Votes   int32
}
