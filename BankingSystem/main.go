package main

import (
	"fmt"
)

func main() {
	var operations []*Operation
	operations = append(operations, NewOperation("DEPOSIT", "account1", "", 1000))
	operations = append(operations, NewOperation("DEPOSIT", "account1", "", 500))
	operations = append(operations, NewOperation("DEPOSIT", "account2", "", 1000))
	operations = append(operations, NewOperation("WITHDRAW", "non-existing", "", 2700))
	operations = append(operations, NewOperation("WITHDRAW", "account1", "", 2000))
	operations = append(operations, NewOperation("WITHDRAW", "account1", "", 500))
	operations = append(operations, NewOperation("TRANSFER", "account1", "account2", 1001))
	operations = append(operations, NewOperation("TRANSFER", "account1", "account2", 200))

	accounts = make(map[string]*Account)

	for i := 0; i < len(operations); i++ {
		if operations[i].Function == "DEPOSIT" {
			if balance, message := Deposit(operations[i].AccountId1, operations[i].Balance); len(message) > 0 {
				fmt.Println(balance)
			} else {
				fmt.Println(true)
			}
		} else if operations[i].Function == "WITHDRAW" {
			if balance, err := Withdraw(operations[i].AccountId1, operations[i].Balance); err != nil {
				fmt.Println()
			} else {
				fmt.Println(balance)
			}
		} else {
			if balance, err := Transfer(operations[i].AccountId1, operations[i].AccountId2, operations[i].Balance); err != nil {
				fmt.Println()
			} else {
				fmt.Println(balance)
			}
		}
	}
}
