package parser

import (
	"cron/constants"
	"cron/segment"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type RangeParser struct {
	exp segment.BaseSegment
}

func NewRangeParser(seg segment.BaseSegment) IParser {
	return &RangeParser{
		exp: seg,
	}
}

func (a *RangeParser) GetPossibilities() ([]int, error) {
	rangeLimitStrs := strings.Split(a.exp.GetExpression(), "-")

	if len(rangeLimitStrs) != 2 {
		return nil, errors.New("invalid segment expression")
	}
	var err error
	rangeLimits := make([]int, 2)
	if a.exp.GetName() == constants.SegmentTypeWEEKDAY {
		fmt.Println(rangeLimitStrs)
		rangeLimitStrs[0] = constants.UpdateWeekDayExpression(rangeLimitStrs[0])
		rangeLimitStrs[1] = constants.UpdateWeekDayExpression(rangeLimitStrs[1])
	}

	if rangeLimits[0], err = strconv.Atoi(rangeLimitStrs[0]); err != nil {
		return nil, errors.New("invalid segment expression")
	}
	if rangeLimits[1], err = strconv.Atoi(rangeLimitStrs[1]); err != nil {
		return nil, errors.New("invalid segment expression")
	}

	if rangeLimits[1] < rangeLimits[0] {
		return nil, errors.New("range minimum/maximum are in wrong order")
	}

	if rangeLimits[0] < a.exp.GetMinimum() {
		return nil, errors.New("range minimum is not valid")
	}

	if rangeLimits[0] > a.exp.GetMaximum() {
		return nil, errors.New("range minimum is not valid")
	}

	if rangeLimits[1] > a.exp.GetMaximum() {
		return nil, errors.New("range maximum is not valid")
	}
	var vals []int
	for i := rangeLimits[0]; i <= rangeLimits[1]; i++ {
		vals = append(vals, i)
	}
	return vals, nil
}
