package entity

import (
	"errors"
	"splitwise/utils"
)

type Expense struct {
	ExpenseId          int
	ListOfTransactions []Transaction
	Users              map[string]map[string]Balance
}

var y = 0

func NewExpense(users []string) Expense {
	y++
	e := Expense{Users: make(map[string]map[string]Balance), ExpenseId: y, ListOfTransactions: []Transaction{}}
	for _, user := range users {
		e.Users[user] = make(map[string]Balance)
	}
	return e
}

func (e *Expense) AddTransaction(t *Transaction) error {
	if _, ok := e.Users[t.PaidByUser]; !ok {
		return errors.New("invalid paid by user")
	}
	for _, user := range t.ListOfPaidUsers {
		if _, ok := e.Users[user]; !ok {
			return errors.New("invalid paid users")
		}
	}
	e.ListOfTransactions = append(e.ListOfTransactions, *t)

	amount := t.Amount
	if t.Type == utils.TransactionTypeEQUAL {
		amount = t.Amount / (float64(len(t.ListOfPaidUsers)) + 1.0)
	}
	for _, user := range t.ListOfPaidUsers {
		//Add debit balance
		if _, ok := e.Users[user][t.PaidByUser]; !ok {
			e.Users[user][t.PaidByUser] = Balance{Amount: amount, Type: utils.BalanceTypeDEBIT}
		} else {
			bal := e.Users[user][t.PaidByUser]
			if bal.Type == utils.BalanceTypeDEBIT {
				bal.Amount += amount
			} else {
				if bal.Amount >= amount {
					bal.Amount -= amount
				} else {
					bal.Amount = amount - bal.Amount
					bal.Type = utils.BalanceTypeDEBIT
				}
			}
			e.Users[user][t.PaidByUser] = bal
		}

		//add credit balance
		if _, ok := e.Users[t.PaidByUser][user]; !ok {
			e.Users[t.PaidByUser][user] = Balance{Amount: amount, Type: utils.BalanceTypeCREDIT}
		} else {
			bal := e.Users[t.PaidByUser][user]
			if bal.Type == utils.BalanceTypeCREDIT {
				bal.Amount += amount
			} else {
				if bal.Amount >= amount {
					bal.Amount -= amount
				} else {
					bal.Amount = amount - bal.Amount
					bal.Type = utils.BalanceTypeCREDIT
				}
			}
			e.Users[t.PaidByUser][user] = bal
		}
	}
	return nil
}
