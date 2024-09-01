package criteria

import "file-search/file"

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

func (n *NameCriteria) IsSatifiedBy(f *file.File) bool {
	return n.Name == f.GetName()
}

type SizeCriteria struct {
	Size     int
	Operator string
}

func NewSizeCriteria(s int, op string) ICriteria {
	return &SizeCriteria{Size: s, Operator: op}
}

func (s *SizeCriteria) IsSatifiedBy(f *file.File) bool {
	if s.Operator == Equal {
		return s.Size == f.GetSize()
	} else if s.Operator == LessThanEqual {
		return s.Size <= f.GetSize()
	} else if s.Operator == GreaterThanEqual {
		return s.Size >= f.GetSize()
	}
	return false
}

type ExtensionCriteria struct {
	Ext string
}

func NewExtensionCriteria(e string) ICriteria {
	return &ExtensionCriteria{Ext: e}
}

func (e *ExtensionCriteria) IsSatifiedBy(f *file.File) bool {
	return e.Ext == f.GetExtension()
}
