package department

import "appointment/constants"

type InsuranceDepartment struct {
	Department
}

var insuranceDept *InsuranceDepartment

func NewInsuranceDepartment(desks int) {
	insuranceDept = &InsuranceDepartment{Department: Department{
		Id:        constants.DEPARTMENT_TYPE_INSURANCE,
		NoofDesks: desks,
	}}

	for i := 0; i < desks; i++ {
		desk := NewDesk(constants.DEPARTMENT_TYPE_ACCOUNTS)
		accountDept.DeskMap[desk.Id] = desk
	}
}

// func (ins *InsuranceDepartment) CompleteTask(customer *CustomerRequirement) {
// 	customer.Status = 1
// 	ins.Queue = ins.Queue[1:]
// }

// func (ins *InsuranceDepartment) GetNextTaskId() uint64 {
// 	if len(ins.Queue) > 0 {
// 		return ins.Queue[0]
// 	}
// 	return 0
// }
