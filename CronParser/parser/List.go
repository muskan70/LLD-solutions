package parser

import (
	"cron/segment"
	"sort"
	"strings"
)

type ListParser struct {
	exp segment.BaseSegment
}

func NewListParser(seg segment.BaseSegment) IParser {
	return &ListParser{
		exp: seg,
	}
}

func (a *ListParser) GetPossibilities() ([]int, error) {
	list := strings.Split(a.exp.GetExpression(), ",")

	uniqueVals := make(map[int]bool)
	for i := range list {
		if prsr, err := GetParser(segment.NewBaseSegment(a.exp.GetName(), a.exp.GetMinimum(), a.exp.GetMaximum(), list[i])); err != nil {
			return nil, err
		} else if possibleValues, err := prsr.GetPossibilities(); err != nil {
			return nil, err
		} else {
			for j := range possibleValues {
				uniqueVals[possibleValues[j]] = true
			}
		}
	}
	var vals []int
	for key := range uniqueVals {
		vals = append(vals, key)
	}
	sort.Ints(vals)
	return vals, nil
}
