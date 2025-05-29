package department

import (
	"appointment/tokenGenerator"
	"sync/atomic"
)

var deskId atomic.Uint64

type Department struct {
	Id        uint64
	NoofDesks int
	DeskMap   map[uint64]*Desk
}

type Desk struct {
	Id           uint64
	DepartmentId uint64
}

func NewDesk(deptId uint64) *Desk {
	return &Desk{
		Id:           deskId.Add(1),
		DepartmentId: deptId,
	}
}

func (d *Desk) CompleteTask(t *tokenGenerator.Token) {
	t.MarkComplete()
}
