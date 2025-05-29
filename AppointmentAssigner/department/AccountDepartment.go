package department

import "appointment/constants"

type AccountDepartment struct {
	Department
}

func NewAccountDepartment(desks int) *AccountDepartment {
	accountDept := &AccountDepartment{Department: Department{
		Id:        constants.DEPARTMENT_TYPE_ACCOUNTS,
		NoofDesks: desks,
		DeskMap:   make(map[uint64]*Desk),
	}}
	for i := 0; i < desks; i++ {
		desk := NewDesk(constants.DEPARTMENT_TYPE_ACCOUNTS)
		accountDept.DeskMap[desk.Id] = desk
	}
	return accountDept
}
