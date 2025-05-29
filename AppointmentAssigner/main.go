package main

import (
	"appointment/department"
	"appointment/departmentAssigner"
	"appointment/tokenGenerator"
	"math/rand"
)

//Token Manager -> DepartmentAssigner -->Department -> desk

func main() {
	tokenManager := tokenGenerator.NewTokenManager()
	deptAssign := departmentAssigner.NewDepartmentAssigner(tokenManager)

	insuranceDept := department.NewInsuranceDepartment(2)
	deptAssign.AddDepartmentQueue(int(insuranceDept.Id))

	accountDept := department.NewAccountDepartment(2)
	deptAssign.AddDepartmentQueue(int(accountDept.Id))

	loanDept := department.NewLoanDepartment(1)
	deptAssign.AddDepartmentQueue(int(loanDept.Id))

	for i := 0; i < 10; i++ {
		t := tokenManager.NewToken(i)
		deptAssign.AssignDepartment(rand.Intn(3)+1, t)
	}

	for _, desk := range insuranceDept.DeskMap {
		if task := deptAssign.GetNextTaskWrtDepartment(int(insuranceDept.Id)); task != nil {
			desk.CompleteTask(task)
		}
	}

	for _, desk := range accountDept.DeskMap {
		if task := deptAssign.GetNextTaskWrtDepartment(int(accountDept.Id)); task != nil {
			desk.CompleteTask(task)
		}
	}

	for _, desk := range loanDept.DeskMap {
		if task := deptAssign.GetNextTaskWrtDepartment(int(loanDept.Id)); task != nil {
			desk.CompleteTask(task)
		}
	}

}
