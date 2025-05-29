package department

import "appointment/constants"

type LoansDepartment struct {
	Department
}

var loanDept *LoansDepartment

func NewLoanDepartment(desks int) {
	loanDept = &LoansDepartment{Department: Department{
		Id:        constants.DEPARTMENT_TYPE_LOANS,
		NoofDesks: desks,
	}}

	for i := 0; i < desks; i++ {
		desk := NewDesk(constants.DEPARTMENT_TYPE_ACCOUNTS)
		accountDept.DeskMap[desk.Id] = desk
	}
}

// func (lo *LoansDepartment) CompleteTask(customer *CustomerRequirement) {
// 	customer.Status = 1
// 	lo.Queue = lo.Queue[1:]
// }

// func (lo *LoansDepartment) GetNextTaskId() uint64 {
// 	if len(lo.Queue) > 0 {
// 		return lo.Queue[0]
// 	}
// 	return 0
// }
