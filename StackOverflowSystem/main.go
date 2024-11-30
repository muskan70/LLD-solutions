package main

import (
	"fmt"
	"strings"
)

func main() {
	// create Users
	alice := NewUser("Alice", "alice@example.com")
	bob := NewUser("Bob", "bob@example.com")
	charlie := NewUser("Charlie", "charlie@example.com")

	// Alice asks a question
	javaQuestion := alice.AskQuestion("What is polymorphism in Java?", "Can someone explain polymorphism in Java with an example?")

	javaQuestion.AddTag("java")
	javaQuestion.AddTag("oop")

	// Bob answers Alice's question
	bobAnswer := bob.AnswerQuestion(javaQuestion.Id, "Polymorphism in Java is the ability of an object to take on many forms...")

	// Charlie comments on the question
	charlie.AddComment(javaQuestion, "Great question! I'm also interested in learning about this.")

	// Alice comments on Bob's answer
	alice.AddComment(bobAnswer, "Thanks for the explanation! Could you provide a code example?")

	// Charlie votes on the question and answer
	javaQuestion.Vote(charlie.Id, 1) // Upvote
	bobAnswer.Vote(charlie.Id, 1)    // Upvote

	// accept Bob's answer
	bobAnswer.MarkAsAccepted()

	// Bob asks another question
	pythonQuestion := bob.AskQuestion("How to use list comprehensions in Python?",
		"I'm new to Python and I've heard about list comprehensions. Can someone explain how to use them?")
	pythonQuestion.AddTag("python")
	pythonQuestion.AddTag("list-comprehension")

	// Alice answers Bob's question
	aliceAnswer := alice.AnswerQuestion(pythonQuestion.Id,
		"List comprehensions in Python provide a concise way to create lists...")

	// Charlie votes on Bob's question and Alice's answer
	pythonQuestion.Vote(charlie.Id, 1) // Upvote
	aliceAnswer.Vote(charlie.Id, 1)    // Upvote

	// Print out the current state
	fmt.Println("Question: " + javaQuestion.GetTitle())
	fmt.Println("Asked by: " + Users[javaQuestion.GetAuthor()].GetUsername())
	fmt.Println("Tags: " + strings.Join(javaQuestion.GetTags(), ","))
	fmt.Println("Votes: ", javaQuestion.GetVoteCount())
	fmt.Println("Comments: ", len(javaQuestion.GetComments()))
	fmt.Println("\nAnswer by " + Users[bobAnswer.GetAuthor()].GetUsername() + ":")
	fmt.Println(bobAnswer.GetContent())
	fmt.Println("Votes: ", bobAnswer.GetVoteCount())
	fmt.Println("Accepted: ", bobAnswer.IsAccepted)
	fmt.Println("Comments: ", len(bobAnswer.GetComments()))

	fmt.Println("\nUser Reputations:")
	fmt.Println("Alice: ", alice.GetReputation())
	fmt.Println("Bob: ", bob.GetReputation())
	fmt.Println("Charlie: ", charlie.GetReputation())

	// Demonstrate search functionality
	fmt.Println("\nSearch Results for 'java':")
	searchResults := SearchQuestions("java")
	for _, q := range searchResults {
		fmt.Println(q.GetTitle())
	}

	fmt.Println("\nSearch Results for 'python':")
	searchResults = SearchQuestions("python")
	for _, q := range searchResults {
		fmt.Println(q.GetTitle())
	}

	// Demonstrate getting questions by user
	fmt.Println("\nBob's Questions:")
	bobQuestions := bob.GetQuestions()
	for _, q := range bobQuestions {
		fmt.Println(q.GetTitle())
	}
}
