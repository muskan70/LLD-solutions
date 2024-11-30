package main

import (
	"time"
)

type Answer struct {
	Id           int
	QuestionId   int
	AuthorId     int
	Content      string
	IsAccepted   bool
	Votes        Votes
	CommentIds   []int
	CreationDate time.Time
}

var answerId = 0
var Answers = make(map[int]Answer)

func NewAnswer(questionId int, content string, authorId int) *Answer {
	answerId++
	ans := Answer{
		Id:           answerId,
		QuestionId:   questionId,
		AuthorId:     authorId,
		IsAccepted:   false,
		Content:      content,
		CreationDate: time.Now(),
		Votes:        make(Votes),
	}
	Answers[answerId] = ans
	return &ans
}

func (a *Answer) AddComment(authorId int, content string) *Comment {
	comment := NewComment(a.Id, "Answer", authorId, content)
	a.CommentIds = append(a.CommentIds, comment.Id)
	return comment
}

func (a *Answer) Vote(userId int, val int) {
	a.Votes[userId] = val
	Users[a.AuthorId].UpdateReputation(val * REPUTATION_VOTE)
}

func (a *Answer) MarkAsAccepted() {
	a.IsAccepted = true
	Users[a.AuthorId].UpdateReputation(REPUTATION_ACCEPTED_ANSWER)
}

func (a *Answer) GetVoteCount() int {
	sum := 0
	for _, val := range a.Votes {
		sum += val
	}
	return sum
}

func (a *Answer) GetComments() []Comment {
	var comments []Comment
	for i := range a.CommentIds {
		comments = append(comments, *Comments[a.CommentIds[i]])
	}
	return comments
}

func (a *Answer) GetContent() string {
	return a.Content
}

func (a *Answer) GetAuthor() int {
	return a.AuthorId
}
