package main

import "time"

type Comment struct {
	Id           int
	ParentId     int
	ParentType   string
	AuthorId     int
	Content      string
	CreationDate time.Time
}

var commentId = 0

func NewComment(parentId int, parentType string, authorId int, content string) *Comment {
	commentId++
	return &Comment{
		Id:           commentId,
		ParentId:     parentId,
		AuthorId:     authorId,
		ParentType:   parentType,
		Content:      content,
		CreationDate: time.Now(),
	}
}
