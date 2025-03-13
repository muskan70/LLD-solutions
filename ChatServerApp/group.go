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
	groupId++
	users := make(map[int]bool)
	users[adminId] = true
	return &Group{GroupId: groupId, GroupName: gpName, Participants: users, AdminId: adminId}
}

func (g *Group) AddParticipant(userId int) {
	g.Participants[userId] = true
}

func (g *Group) RemoveParticipant(userId int) {
	delete(g.Participants, userId)
}

func (g *Group) GetChatId(userId int) (int, error) {
	if _, ok := g.Participants[userId]; !ok {
		return -1, errors.New("this user is not a member of group")
	}
	return g.ChatId, nil
}
