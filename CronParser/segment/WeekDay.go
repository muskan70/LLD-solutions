package segment

import (
	"cron/constants"
)

type Weekday struct {
	BaseSegment
}

func NewWeekday(exp string) Weekday {
	return Weekday{
		BaseSegment: NewBaseSegment(constants.SegmentTypeWEEKDAY, 1, 7, exp),
	}
}
