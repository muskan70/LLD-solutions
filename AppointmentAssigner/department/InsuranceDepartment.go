package department

import (
	"appointment/constants"
)

type InsuranceDepartment struct {
	Department
}

func NewInsuranceDepartment(desks int) *InsuranceDepartment {
	insuranceDept := &InsuranceDepartment{Department: Department{
		Id:        constants.DEPARTMENT_TYPE_INSURANCE,
		NoofDesks: desks,
		DeskMap:   make(map[uint64]*Desk),
	}}

	for i := 0; i < desks; i++ {
		desk := NewDesk(constants.DEPARTMENT_TYPE_ACCOUNTS)
		insuranceDept.DeskMap[desk.Id] = desk
	}
	return insuranceDept
}
