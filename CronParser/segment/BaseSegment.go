package segment

import "strconv"

type BaseSegment struct {
	segmentName    string
	minimumValue   int
	maximumValue   int
	expression     string
	PossibleValues []int
}

func (b *BaseSegment) GetName() string {
	return b.segmentName
}

func NewBaseSegment(name string, min int, max int, exp string) BaseSegment {
	return BaseSegment{
		segmentName:  name,
		minimumValue: min,
		maximumValue: max,
		expression:   exp,
	}
}

func (b *BaseSegment) GetMinimum() int {
	return b.minimumValue
}

func (b *BaseSegment) GetMaximum() int {
	return b.maximumValue
}

func (b *BaseSegment) GetExpression() string {
	return b.expression
}

func (b *BaseSegment) GetPossibleValuesString() string {
	vals := ""
	for i := range b.PossibleValues {
		vals += strconv.FormatInt(int64(b.PossibleValues[i]), 10) + " "
	}
	return vals
}
