package main

import "time"

var userId = 0

const (
	UserOnline  = true
	UserOffline = false
)

type User struct {
	UserId         int
	Name           string
	Email          string
	Phone          int
	DirectChats    map[int]int
	GroupChats     map[int]int
	PresenceStatus bool
	LastSeen       time.Time
}

func NewUser(name, email string, phone int) *User {
	userId++
	return &User{
		UserId:      userId,
		Name:        name,
		Email:       email,
		Phone:       phone,
		DirectChats: make(map[int]int),
		GroupChats:  make(map[int]int),
	}
}

func (u *User) GetChatId(userId int) int {
	_, ok := u.DirectChats[userId]
	if !ok {
		return -1
	}
	return u.DirectChats[userId]
}

func (u *User) AddChatId(userId int, chatId int) {
	u.DirectChats[userId] = chatId
}

func (u *User) AddGroupChatId(grpId int, chatId int) {
	u.GroupChats[grpId] = chatId
}

func (u *User) RemoveGroupChatId(grpId int) {
	delete(u.GroupChats, grpId)
}

func (u *User) UpdatePresence(status bool) {
	u.PresenceStatus = status
	if !status {
		u.LastSeen = time.Now()
	} else {
		u.LastSeen = time.Time{}
	}
	// TODO: notify observers
}
