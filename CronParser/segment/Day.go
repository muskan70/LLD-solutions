package segment

import (
	"cron/constants"
)

type Day struct {
	BaseSegment
}

func NewDay(exp string) Day {
	return Day{
		BaseSegment: NewBaseSegment(constants.SegmentTypeDAY, 1, 31, exp),
	}
}
