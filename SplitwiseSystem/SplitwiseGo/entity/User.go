package entity

var x = 0

type User struct {
	UserId   int
	Name     string
	Balances map[string]Balance
}

type Balance struct {
	Amount float64
	Type   string
}

func NewUser(name string) User {
	x++
	return User{UserId: x, Name: name}
}
