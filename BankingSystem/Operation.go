package main

type Operation struct {
	Function   string
	AccountId1 string
	AccountId2 string
	Balance    float64
}

func NewOperation(function, accountId1, accountId2 string, balance float64) *Operation {
	return &Operation{
		Function:   function,
		AccountId1: accountId1,
		AccountId2: accountId2,
		Balance:    balance,
	}
}
