package constants

const (
	LogLevelFATAL    = 5
	LogLevelERROR    = 4
	LogLevelWARN     = 3
	LogLevelINFO     = 2
	LogLevelDEBUG    = 1
	SinkTypeFILE     = "Text File"
	SinkTypeDATABASE = "Database"
	SinkTypeCONSOLE  = "Console"
)

func GetLogLevel(level int) string {
	switch level {
	case LogLevelFATAL:
		return "FATAL"
	case LogLevelERROR:
		return "ERROR"
	case LogLevelWARN:
		return "WARN"
	case LogLevelINFO:
		return "INFO"
	case LogLevelDEBUG:
		return "DEBUG"
	}
	return ""
}
