package entity

import "sync/atomic"

var userId atomic.Uint64

type User struct {
	Id    uint64
	name  string
	email string
}

func NewUser(name, email string) *User {
	return &User{
		Id:    userId.Add(1),
		name:  name,
		email: email,
	}
}
