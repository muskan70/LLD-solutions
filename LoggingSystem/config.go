package main

import "sync"

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

type LogConfig struct {
	LogLevel     int
	SinkType     string
	SinkLocation map[int]*LogSink
}

var logConfig LogConfig

func AddConfig(level int, sinkType string) LogConfig {
	logConfig = LogConfig{
		LogLevel:     level,
		SinkType:     sinkType,
		SinkLocation: CreateLogSink(),
	}
	return logConfig
}

func (c *LogConfig) SetLogLevel(level int, wg *sync.WaitGroup) {
	wg.Wait()
	c.LogLevel = level
}
