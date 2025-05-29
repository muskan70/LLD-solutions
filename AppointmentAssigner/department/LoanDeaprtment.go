package department

import "appointment/constants"

type LoansDepartment struct {
	Department
}

func NewLoanDepartment(desks int) *LoansDepartment {
	loanDept := &LoansDepartment{Department: Department{
		Id:        constants.DEPARTMENT_TYPE_LOANS,
		NoofDesks: desks,
		DeskMap:   make(map[uint64]*Desk),
	}}

	for i := 0; i < desks; i++ {
		desk := NewDesk(constants.DEPARTMENT_TYPE_ACCOUNTS)
		loanDept.DeskMap[desk.Id] = desk
	}
	return loanDept
}
