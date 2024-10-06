package criteria

import (
	"file-search/fileSystem"
)

type ICriteria interface {
	IsSatisfiedBy(f fileSystem.File) bool
}

type AndCriteria struct {
	Criterias []ICriteria
}

func NewAndCriteria(c []ICriteria) ICriteria {
	return &AndCriteria{Criterias: c}
}

func (a *AndCriteria) IsSatisfiedBy(f fileSystem.File) bool {
	for _, criteria := range a.Criterias {
		if !criteria.IsSatisfiedBy(f) {
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

func (o *OrCriteria) IsSatisfiedBy(f fileSystem.File) bool {
	for _, criteria := range o.Criterias {
		if criteria.IsSatisfiedBy(f) {
			return true
		}
	}
	return false
}
