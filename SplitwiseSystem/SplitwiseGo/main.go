package main

import (
	"fmt"
	"splitGo/entity"
	"splitGo/utils"
)

func main() {
	e := entity.NewExpense([]string{"muskan", "vipul", "yash", "manju"})
	t := entity.NewTransaction(utils.TransactionTypeEQUAL, []string{"manju", "vipul", "muskan"}, "yash", 20000)
	e.AddTransaction(&t)
	for user1, val := range e.Users {
		for user2, bal := range val {
			fmt.Println(user1, bal.Type, user2, bal.Amount)
		}
	}
}
