package main

const (
	REPUTATION_QUESTION        = 5
	REPUTATION_ANSWER          = 10
	REPUTATION_COMMENT         = 2
	REPUTATION_ACCEPTED_ANSWER = 15
	REPUTATION_VOTE            = 10
)

type User struct {
	Id              int
	Name            string
	Email           string
	Reputation      int
	ListOfQuestions []int
	ListOfAnswers   []int
	ListOfComments  []int
}

var userId = 0
var Users = make(map[int]*User)

func NewUser(name, email string) *User {
	userId++
	user := &User{
		Id:         userId,
		Name:       name,
		Email:      email,
		Reputation: 0,
	}
	Users[userId] = user
	return user
}

func (u *User) UpdateReputation(score int) {
	u.Reputation += score
	if u.Reputation < 0 {
		u.Reputation = 0
	}
}

func (u *User) AskQuestion(title, content string) *Question {
	ques := NewQuestion(title, content, u.Id)
	u.ListOfQuestions = append(u.ListOfQuestions, ques.Id)
	u.UpdateReputation(REPUTATION_QUESTION)
	return ques
}

func (u *User) AnswerQuestion(questionId int, content string) *Answer {
	ans := Questions[questionId].AddAnswer(content, u.Id)
	u.ListOfAnswers = append(u.ListOfAnswers, ans.Id)
	u.UpdateReputation(REPUTATION_ANSWER)
	return ans
}

func (u *User) AddComment(commentable Commentable, content string) {
	comment := commentable.AddComment(u.Id, content)
	u.ListOfComments = append(u.ListOfComments, comment.Id)
	u.UpdateReputation(REPUTATION_COMMENT)
}

func (u *User) GetReputation() int {
	return u.Reputation
}

func (u *User) GetQuestions() []Question {
	var questions []Question
	for i := range u.ListOfQuestions {
		questions = append(questions, *Questions[u.ListOfQuestions[i]])
	}
	return questions
}

func (u *User) GetAnswers() []Answer {
	var answers []Answer
	for i := range u.ListOfAnswers {
		answers = append(answers, Answers[u.ListOfAnswers[i]])
	}
	return answers
}

func (u *User) GetComments() []Comment {
	var comments []Comment
	for i := range u.ListOfComments {
		comments = append(comments, *Comments[u.ListOfComments[i]])
	}
	return comments
}

func (u *User) GetUsername() string {
	return u.Name
}
