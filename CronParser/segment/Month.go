package segment

import "cron/constants"

type Month struct {
	BaseSegment
}

func NewMonth(exp string) Month {
	return Month{
		BaseSegment: NewBaseSegment(constants.SegmentTypeMONTH, 1, 12, exp),
	}
}
