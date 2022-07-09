package model

type Comment struct {
	Model
	Message  string
	Mentions []string
}
