package entity

type Transaction struct {
	Type            string
	ListOfPaidUsers []string
	PaidByUser      string
	Amount          float64
}

func NewTransaction(typ string, paidUsers []string, paidBy string, amount float64) Transaction {
	return Transaction{Type: typ, ListOfPaidUsers: paidUsers, PaidByUser: paidBy, Amount: amount}
}
