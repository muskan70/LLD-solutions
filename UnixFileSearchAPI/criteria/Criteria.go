package criteria

import (
	"file-search/fileSystem"
)

const (
	Equal            = "="
	LessThanEqual    = "<="
	GreaterThanEqual = ">="
)

type NameCriteria struct {
	Name string
}

func NewNameCriteria(n string) ICriteria {
	return &NameCriteria{Name: n}
}

func (n *NameCriteria) IsSatisfiedBy(f fileSystem.File) bool {
	return n.Name == f.GetName()
}

type SizeCriteria struct {
	Size     int
	Operator string
}

func NewSizeCriteria(s int, op string) ICriteria {
	return &SizeCriteria{Size: s, Operator: op}
}

func (s *SizeCriteria) IsSatisfiedBy(f fileSystem.File) bool {
	if s.Operator == Equal {
		return s.Size == f.GetSize()
	} else if s.Operator == LessThanEqual {
		return f.GetSize() <= s.Size
	} else if s.Operator == GreaterThanEqual {
		return f.GetSize() >= s.Size
	}
	return false
}

type ExtensionCriteria struct {
	Ext string
}

func NewExtensionCriteria(e string) ICriteria {
	return &ExtensionCriteria{Ext: e}
}

func (e *ExtensionCriteria) IsSatisfiedBy(f fileSystem.File) bool {
	return e.Ext == f.GetExtension()
}
