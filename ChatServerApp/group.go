package main

import "errors"

type message struct {
	Msg  string `json:"message"`
	User string `json:"userName"`
}

type Group struct {
	GroupName string
	messages  []message
	Users     map[string]bool
}

func NewGroup(gpName string) *Group {
	return &Group{GroupName: gpName, messages: []message{}, Users: make(map[string]bool)}
}

func (g *Group) AddUser(u string) {
	g.Users[u] = true
}

func (g *Group) AddMessage(usr string, msg string) error {
	if _, ok := g.Users[usr]; !ok {
		return errors.New("this user is not a member of group")
	}
	g.messages = append(g.messages, message{Msg: msg, User: usr})
	return nil
}

func (g *Group) GetAllMessages(u string) ([]message, error) {
	if _, ok := g.Users[u]; !ok {
		return nil, errors.New("this user is not a member of group")
	}
	return g.messages, nil
}
