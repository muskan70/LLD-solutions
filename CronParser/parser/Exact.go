package parser

import (
	"cron/constants"
	"cron/segment"
	"errors"
	"strconv"
)

type ExactParser struct {
	exp segment.BaseSegment
}

func NewExactParser(seg segment.BaseSegment) IParser {
	return &ExactParser{
		exp: seg,
	}
}

func (a *ExactParser) GetPossibilities() ([]int, error) {
	exp := a.exp.GetExpression()
	if a.exp.GetName() == constants.SegmentTypeWEEKDAY {
		exp = constants.UpdateWeekDayExpression(exp)
	}
	value, err := strconv.Atoi(exp)
	if err != nil {
		return nil, errors.New("invalid segment expression")
	}
	if value > a.exp.GetMaximum() {
		return nil, errors.New("the value for segment is more than the maximum allowed")
	}

	if value < a.exp.GetMinimum() {
		return nil, errors.New("the value for segment is less than the minimum allowed")
	}

	return []int{value}, nil
}
