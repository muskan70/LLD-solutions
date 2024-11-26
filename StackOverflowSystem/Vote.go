package main

type Votable interface {
	Vote(userId int, val int)
	GetVoteCount() int
}

type Votes map[int]int //map of userId to vote value
