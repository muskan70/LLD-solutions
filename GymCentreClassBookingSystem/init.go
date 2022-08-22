package main

var centres map[string]*Centre
var users map[string]*User

func initMain() {
	centres = make(map[string]*Centre)
	users = make(map[string]*User)
}
