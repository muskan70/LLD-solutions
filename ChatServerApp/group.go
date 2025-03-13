package main

import "errors"

var groupId = 0

type Group struct {
	GroupId      int
	ChatId       int
	GroupName    string
	ProfileURL   string
	Participants map[int]bool
	AdminId      int
}

func NewGroup(gpName string, adminId int) *Group {
	users := make(map[int]bool)
	users[adminId] = true
	return &Group{GroupName: gpName, ChatHistory: NewChat(), Participants: users, AdminId: adminId}
}

func (g *Group) AddParticipant(userId int) {
	g.Participants[u] = true
}

func (g *Group) RemoveParticipant(userId int) {
	delete(g.Participants, u)
}

func (g *Group) AddMessage(userId int, content string) error {
	if _, ok := g.Participants[userId]; !ok {
		return errors.New("this user is not a member of group")
	}
	chats[g.ChatId].AddMessage(NewMessage(userId, g.ChatId, content))
	return nil
}

func (g *Group) GetChatHistory(userId Int) ([]Message, error) {
	if _, ok := g.Participants[userId]; !ok {
		return nil, errors.New("this user is not a member of group")
	}
	return chats[g.ChatId].GetAllMessages(), nil
}
