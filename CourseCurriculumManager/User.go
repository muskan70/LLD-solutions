package main

const (
	UserTypeAdmin             = "ADMIN"
	UserTypeCustomer          = "CUSTOMER"
	UserCourseStatusNew       = "NEW"
	UserCourseStatusCompleted = "COMPLETE"
)

var x = 0

type User struct {
	UserId   int
	Name     string
	Phone    string
	Email    string
	UserType string
}

func NewUser(name, phone, email, usertype string) User {
	x++
	usr := User{
		UserId:   x,
		Name:     name,
		Phone:    phone,
		Email:    email,
		UserType: usertype,
	}
	return usr
}
