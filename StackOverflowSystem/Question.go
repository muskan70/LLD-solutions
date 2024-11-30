package main

import (
	"strings"
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
var Questions = make(map[int]*Question)

func NewQuestion(title, content string, authorId int) *Question {
	quesId++
	ques := &Question{
		Id:           quesId,
		Title:        title,
		AuthorId:     authorId,
		Content:      content,
		CreationDate: time.Now(),
		Votes:        make(Votes),
	}
	Questions[quesId] = ques
	return ques
}

func SearchQuestions(query string) []Question {
	var questions []Question
	for _, ques := range Questions {
		if strings.Contains(ques.Title, query) || strings.Contains(ques.Content, query) || strings.Contains(strings.Join(ques.GetTags(), " "), query) {
			questions = append(questions, *ques)
		}
	}
	return questions
}

func (q *Question) AddAnswer(content string, authorId int) *Answer {
	ans := NewAnswer(q.Id, content, q.AuthorId)
	q.AnswerIds = append(q.AnswerIds, ans.Id)
	return ans
}

func (q *Question) AddComment(authorId int, content string) *Comment {
	comment := NewComment(q.Id, "Question", authorId, content)
	q.CommentIds = append(q.CommentIds, comment.Id)
	return comment
}

func (q *Question) AddTag(name string) {
	q.TagIds = append(q.TagIds, GetTagId(name))
}

func (q *Question) Vote(userId int, val int) {
	q.Votes[userId] = val
	Users[q.AuthorId].UpdateReputation(val * REPUTATION_VOTE)
}

func (q *Question) GetVoteCount() int {
	sum := 0
	for _, val := range q.Votes {
		sum += val
	}
	return sum
}

func (q *Question) GetAnswers() []Answer {
	var answers []Answer
	for i := range q.AnswerIds {
		answers = append(answers, Answers[q.AnswerIds[i]])
	}
	return answers
}

func (q *Question) GetComments() []Comment {
	var comments []Comment
	for i := range q.CommentIds {
		comments = append(comments, *Comments[q.CommentIds[i]])
	}
	return comments
}

func (q *Question) GetTags() []string {
	var tags []string
	for i := range q.TagIds {
		tags = append(tags, GetTagName(q.TagIds[i]))
	}
	return tags
}

func (q *Question) GetTitle() string {
	return q.Title
}

func (q *Question) GetContent() string {
	return q.Content
}

func (q *Question) GetAuthor() int {
	return q.AuthorId
}
