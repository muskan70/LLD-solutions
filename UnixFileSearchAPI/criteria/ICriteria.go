package criteria

import (
	"file-search/fileSystem"
)

type ICriteria interface {
	IsSatifiedBy(f fileSystem.File) bool
}

type AndCriteria struct {
	Criterias []ICriteria
}

func NewAndCriteria(c []ICriteria) ICriteria {
	return &AndCriteria{Criterias: c}
}

func (a *AndCriteria) IsSatifiedBy(f fileSystem.File) bool {
	for _, criteria := range a.Criterias {
		if !criteria.IsSatifiedBy(f) {
			return false
		}
	}
	return true
}

type OrCriteria struct {
	Criterias []ICriteria
}

func NewOrCriteria(c []ICriteria) ICriteria {
	return &OrCriteria{Criterias: c}
}

func (o *OrCriteria) IsSatifiedBy(f fileSystem.File) bool {
	for _, criteria := range o.Criterias {
		if criteria.IsSatifiedBy(f) {
			return true
		}
	}
	return false
}
