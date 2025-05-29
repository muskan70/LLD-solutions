package department

import "appointment/constants"

type AccountDepartment struct {
	Department
}

var accountDept *AccountDepartment

func NewAccountDepartment(desks int) {
	accountDept = &AccountDepartment{Department: Department{
		Id:        constants.DEPARTMENT_TYPE_ACCOUNTS,
		NoofDesks: desks,
	}}
	for i := 0; i < desks; i++ {
		desk := NewDesk(constants.DEPARTMENT_TYPE_ACCOUNTS)
		accountDept.DeskMap[desk.Id] = desk
	}

}
