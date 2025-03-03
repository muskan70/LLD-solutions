package segment

import "cron/constants"

type Minute struct {
	BaseSegment
}

func NewMinute(exp string) Minute {
	return Minute{
		BaseSegment: NewBaseSegment(constants.SegmentTypeMINUTE, 0, 59, exp),
	}
}
