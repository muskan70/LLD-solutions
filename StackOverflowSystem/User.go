package main

type User struct {
	Id              int
	Name            string
	Email           string
	Reputation      int
	ListOfQuestions []int
	ListOfAnswers   []int
	ListOfComments  []int
}
