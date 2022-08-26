package entity

import "testing"

func TestExpense_AddTransaction(t *testing.T) {
	type fields struct {
		ExpenseId          int
		ListOfTransactions []Transaction
		Users              map[string]map[string]Balance
	}
	type args struct {
		t *Transaction
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "test1",
			fields: fields{ExpenseId: 1,ListOfTransactions: []Transaction{},Users: make(map)},
		}
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Expense{
				ExpenseId:          tt.fields.ExpenseId,
				ListOfTransactions: tt.fields.ListOfTransactions,
				Users:              tt.fields.Users,
			}
			if err := e.AddTransaction(tt.args.t); (err != nil) != tt.wantErr {
				t.Errorf("Expense.AddTransaction() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
