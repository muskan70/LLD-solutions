package main

import "errors"

type Account struct {
	AccountId string
	Balance   float64
}

var accounts map[string]*Account

func NewAccount(accountId string, balance float64) {
	account := &Account{
		AccountId: accountId,
		Balance:   balance,
	}
	accounts[accountId] = account
}

func Deposit(accountId string, balance float64) (float64, string) {
	if account, ok := accounts[accountId]; ok {
		account.Balance += balance
		return account.Balance, "an account with this identifier already exists"
	} else {
		NewAccount(accountId, balance)
		return balance, ""
	}
}

func Withdraw(accountId string, balance float64) (float64, error) {
	if account, ok := accounts[accountId]; ok {
		if account.Balance < balance {
			return 0, errors.New("withdrawal amount exceeds the account balance")
		}
		account.Balance -= balance
		return account.Balance, nil
	} else {
		return 0, errors.New("an account with this identifier doesn't exist")
	}
}

func Transfer(accountId1, accountId2 string, balance float64) (float64, error) {
	if accountId1 == accountId2 {
		return 0, errors.New("both accountIds are same")
	}
	account1, ok1 := accounts[accountId1]
	account2, ok2 := accounts[accountId2]
	if ok1 && ok2 {
		if account1.Balance < balance {
			return 0, errors.New("this account has insufficient funds")
		}
		account1.Balance -= balance
		account2.Balance += balance
		return account1.Balance, nil
	} else {
		return 0, errors.New("an account with this identifier doesn't exist")
	}
}
