package departmentAssigner

import (
	"appointment/constants"
	"appointment/tokenGenerator"
	"errors"
	"sync"
)

type DepartmentAssigner struct {
	Queue map[int][]uint64
	Lock  map[int]*sync.Mutex
}

func NewDepartmentAssigner() *DepartmentAssigner {
	return &DepartmentAssigner{
		Queue: make(map[int][]uint64),
		Lock:  make(map[int]*sync.Mutex),
	}
}

func (da *DepartmentAssigner) AddDepartmentQueue(departmentId int) {
	da.Queue[departmentId] = make([]uint64, 0)
	da.Lock[departmentId] = &sync.Mutex{}
}

func (da *DepartmentAssigner) AssignDepartment(task int, t *tokenGenerator.Token) error {
	if task == constants.DEPARTMENT_TYPE_LOANS || task == constants.DEPARTMENT_TYPE_INSURANCE || task == constants.DEPARTMENT_TYPE_ACCOUNTS {
		t.DepartmentType = task
		da.AddToken(t.TokenId, task)
	} else {
		return errors.New("invalid business requirement")
	}
	return nil
}

func (da *DepartmentAssigner) GetNextTaskWrtDepartment(departmentId int) *tokenGenerator.Token {
	da.Lock[departmentId].Lock()
	defer da.Lock[departmentId].Unlock()
	if len(da.Queue[departmentId]) > 0 {
		da.Queue[departmentId] = da.Queue[departmentId][1:]
		return tokenGenerator.GetTokenDetails(da.Queue[departmentId][0])
	}
	return nil
}

func (da *DepartmentAssigner) AddToken(tokenId uint64, departmentId int) {
	if _, ok := da.Queue[departmentId]; ok {
		da.Queue[departmentId] = append(da.Queue[departmentId], tokenId)
	}

}
