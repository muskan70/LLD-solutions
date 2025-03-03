package parser

import (
	"cron/segment"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type StepParser struct {
	exp segment.BaseSegment
}

func NewStepParser(seg segment.BaseSegment) IParser {
	return &StepParser{
		exp: seg,
	}
}

func (a *StepParser) GetPossibilities() ([]int, error) {
	steps := strings.Split(a.exp.GetExpression(), "/")

	if len(steps) != 2 {
		return nil, errors.New("step does not have valid expression")
	}

	stepSegments := make([]int, 2)
	var err error
	if steps[0] == "*" {
		stepSegments[0] = a.exp.GetMinimum()
	} else {
		if stepSegments[0], err = strconv.Atoi(steps[0]); err != nil {
			return nil, errors.New("step does not have valid expression")
		}
	}
	if stepSegments[1], err = strconv.Atoi(steps[1]); err != nil {
		return nil, errors.New("step does not have valid expression")
	}

	if stepSegments[1] > a.exp.GetMaximum() {
		fmt.Println(stepSegments[1], a.exp.GetMaximum())
		return nil, errors.New("step size is more than maximum value")
	}

	if stepSegments[0] > a.exp.GetMaximum() {
		return nil, errors.New("step start is more than maximum value")
	}

	var vals []int
	for i := stepSegments[0]; i <= a.exp.GetMaximum(); i += stepSegments[1] {
		vals = append(vals, i)
	}
	return vals, nil
}
