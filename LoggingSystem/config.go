package main

import "logSys/sink"

type LogConfig struct {
	LogLevel     int
	SinkType     string
	SinkLocation sink.Sink
}

func NewConfig(level int, sinkType string) *LogConfig {
	return &LogConfig{
		LogLevel:     level,
		SinkType:     sinkType,
		SinkLocation: sink.NewSink(sinkType),
	}
}

func (c *LogConfig) SetLogLevel(level int) {
	c.LogLevel = level
}
