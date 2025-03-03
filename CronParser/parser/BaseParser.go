package parser

import (
	"cron/constants"
	"cron/segment"
	"errors"
	"regexp"
)

type IParser interface {
	GetPossibilities() ([]int, error)
}

func GetParser(sgmt segment.BaseSegment) (IParser, error) {
	if sgmt.GetExpression() == constants.ParserRegexALL {
		return NewAllParser(sgmt), nil
	}
	if sgmt.GetExpression() == constants.ParserRegexANY {
		if sgmt.GetName() == constants.SegmentTypeWEEKDAY || sgmt.GetName() == constants.SegmentTypeDAY {
			return NewAnyParser(sgmt), nil
		} else {
			return nil, errors.New("this expression is not valid for this segment")
		}
	}
	if matched, _ := regexp.MatchString(constants.ParserRegexLIST, sgmt.GetExpression()); matched {
		return NewListParser(sgmt), nil
	}
	if matched, _ := regexp.MatchString(constants.ParserRegexRANGE, sgmt.GetExpression()); matched {
		return NewRangeParser(sgmt), nil
	}
	if matched, _ := regexp.MatchString(constants.ParserRegexSTEP, sgmt.GetExpression()); matched {
		return NewStepParser(sgmt), nil
	}
	if matched, _ := regexp.MatchString(constants.ParserRegexEXACT, sgmt.GetExpression()); matched {
		return NewExactParser(sgmt), nil
	}
	return nil, errors.New("invalid segment expression")
}
