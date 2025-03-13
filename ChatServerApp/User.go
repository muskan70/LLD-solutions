package main

var userId = 0

type User struct {
	UserId int
	Name   string
	Email  string
	Phone  int
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
