package main

var userId = 0

type User struct {
	UserId      int
	Name        string
	Email       string
	Phone       int
	DirectChats map[int]int
}

func NewUser(name, email string, phone int) *User {
	userId++
	return &User{
		UserId: userId,
		Name:   name,
		Email:  email,
		Phone:  phone,
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
