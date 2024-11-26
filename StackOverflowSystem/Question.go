package main

import (
	"time"
)

type Question struct {
	Id           int
	Title        string
	AuthorId     int
	Content      string
	AnswerIds    []int
	Votes        Votes
	CommentIds   []int
	TagIds       []int
	CreationDate time.Time
}

var quesId = 0
var Questions = make(map[int]Question)

func AskQuestion(title, content string, authorId int) *Question {
	quesId++
	ques := Question{
		Id:           quesId,
		Title:        title,
		AuthorId:     authorId,
		Content:      content,
		CreationDate: time.Now(),
	}
	Questions[quesId] = ques
	return &ques
}

func (q *Question) AddAnswer(answerId int) {
	q.AnswerIds = append(q.AnswerIds, answerId)
}

func (q *Question) AddComment(commentId int) {
	q.CommentIds = append(q.CommentIds, commentId)
}

func (q *Question) AddTag(name string) {
	q.TagIds = append(q.TagIds, GetTagId(name))
}
