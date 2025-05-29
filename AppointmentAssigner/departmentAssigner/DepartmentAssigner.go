package departmentAssigner

import (
	"appointment/tokenGenerator"
	"errors"
	"fmt"
	"sync"
)

type DepartmentAssigner struct {
	TokenManager *tokenGenerator.TokenManager
	Queue        map[int][]uint64
	Lock         map[int]*sync.Mutex
}

func NewDepartmentAssigner(tm *tokenGenerator.TokenManager) *DepartmentAssigner {
	return &DepartmentAssigner{
		TokenManager: tm,
		Queue:        make(map[int][]uint64),
		Lock:         make(map[int]*sync.Mutex),
	}
}

func (da *DepartmentAssigner) AddDepartmentQueue(departmentId int) {
	da.Queue[departmentId] = []uint64{}
	da.Lock[departmentId] = &sync.Mutex{}
}

func (da *DepartmentAssigner) AssignDepartment(task int, t *tokenGenerator.Token) error {
	fmt.Println("task vs token", task, t.TokenId)
	if _, ok := da.Queue[task]; ok {
		t.DepartmentType = task
		da.Queue[task] = append(da.Queue[task], t.TokenId)
	} else {
		return errors.New("invalid business requirement")
	}
	return nil
}

func (da *DepartmentAssigner) GetNextTaskWrtDepartment(departmentId int) *tokenGenerator.Token {
	da.Lock[departmentId].Lock()
	defer da.Lock[departmentId].Unlock()
	if len(da.Queue[departmentId]) > 0 {
		token := da.TokenManager.GetTokenDetails(da.Queue[departmentId][0])
		da.Queue[departmentId] = da.Queue[departmentId][1:]
		return token
	}
	return nil
}
