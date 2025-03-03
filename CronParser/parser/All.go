package parser

import "cron/segment"

type AllParser struct {
	exp segment.BaseSegment
}

func NewAllParser(seg segment.BaseSegment) IParser {
	return &AllParser{
		exp: seg,
	}
}

func (a *AllParser) GetPossibilities() ([]int, error) {
	var vals []int
	for i := a.exp.GetMinimum(); i <= a.exp.GetMaximum(); i++ {
		vals = append(vals, i)
	}
	return vals, nil
}
