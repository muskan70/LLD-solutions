package parser

import "cron/segment"

type AnyParser struct {
	exp segment.BaseSegment
}

func NewAnyParser(seg segment.BaseSegment) IParser {
	return &AnyParser{
		exp: seg,
	}
}

func (a *AnyParser) GetPossibilities() ([]int, error) {
	var vals []int
	for i := a.exp.GetMinimum(); i <= a.exp.GetMaximum(); i++ {
		vals = append(vals, i)
	}
	return vals, nil
}
