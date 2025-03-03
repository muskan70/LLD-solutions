package segment

import "cron/constants"

type Hour struct {
	BaseSegment
}

func NewHour(exp string) Hour {
	return Hour{
		BaseSegment: NewBaseSegment(constants.SegmentTypeHOUR, 0, 23, exp),
	}
}
