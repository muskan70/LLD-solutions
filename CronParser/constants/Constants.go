package constants

const (
	ParserRegexALL   = "*"
	ParserRegexANY   = "?"
	ParserRegexSTEP  = ".*\\/.*"
	ParserRegexLIST  = ".*,.*"
	ParserRegexRANGE = "[0-9a-z]+-[0-9a-z]+"
	ParserRegexEXACT = "^[0-9a-z]+$"
	ParserString     = "^[a-z]+$"
	ParserNumber     = "^[0-9]+$"
)

const (
	SegmentTypeDAY     = "day"
	SegmentTypeMINUTE  = "minute"
	SegmentTypeHOUR    = "hour"
	SegmentTypeMONTH   = "month"
	SegmentTypeWEEKDAY = "weekday"
)

var weekdayMap = map[string]string{
	"mon": "1",
	"tue": "2",
	"wed": "3",
	"thr": "4",
	"fri": "5",
	"sat": "6",
	"sun": "7",
}

func UpdateWeekDayExpression(exp string) string {
	val, ok := weekdayMap[exp]
	if ok {
		return val
	}
	return exp
}
