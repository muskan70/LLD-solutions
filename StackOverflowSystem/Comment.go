package main

import "time"

type Commentable interface {
	AddComment(authorId int, content string) *Comment
}

type Comment struct {
	Id           int
	ParentId     int
	ParentType   string
	AuthorId     int
	Content      string
	CreationDate time.Time
}

var commentId = 0
var Comments = make(map[int]*Comment)

func NewComment(parentId int, parentType string, authorId int, content string) *Comment {
	commentId++
	comm := &Comment{
		Id:           commentId,
		ParentId:     parentId,
		AuthorId:     authorId,
		ParentType:   parentType,
		Content:      content,
		CreationDate: time.Now(),
	}
	Comments[commentId] = comm
	return comm
}
