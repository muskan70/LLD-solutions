package main

import (
	"time"
)

type Answer struct {
	Id           int
	QuestionId   int
	AuthorId     int
	Content      string
	Votes        Votes
	CommentIds   []int
	CreationDate time.Time
}

var answerId = 0
var answers = make(map[int]Answer)

func AnswerQuestion(questionId int, content string, authorId int) *Answer {
	answerId++
	ans := Answer{
		Id:           answerId,
		QuestionId:   questionId,
		AuthorId:     authorId,
		Content:      content,
		CreationDate: time.Now(),
	}
	answers[answerId] = ans
	return &ans
}

func (a *Answer) AddComment(commentId int) {
	a.CommentIds = append(a.CommentIds, commentId)
}
