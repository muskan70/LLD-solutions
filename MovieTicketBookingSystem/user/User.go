package user

import "sync/atomic"

var userId atomic.Uint64

type User struct {
	UserId uint64
	Name   string
	Phone  int
	Email  string
}

func NewUser(name, email string, phone int) *User {
	return &User{
		UserId: userId.Add(1),
		Name:   name,
		Phone:  phone,
		Email:  email,
	}
}
