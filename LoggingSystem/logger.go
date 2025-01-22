package main

import (
	"fmt"
	"logSys/constants"
	"logSys/message"
)

type Logger struct {
	Config    *LogConfig
	Namespace string
}

func NewLogger(config *LogConfig, namespace string) *Logger {
	return &Logger{
		Config:    config,
		Namespace: namespace,
	}
}

func (l *Logger) Debug(msg string) {
	l.log(constants.LogLevelDEBUG, msg)
}
func (l *Logger) Info(msg string) {
	l.log(constants.LogLevelINFO, msg)
}
func (l *Logger) Warn(msg string) {
	l.log(constants.LogLevelWARN, msg)
}
func (l *Logger) Error(msg string) {
	l.log(constants.LogLevelERROR, msg)
}
func (l *Logger) Fatal(msg string) {
	l.log(constants.LogLevelFATAL, msg)
}

func (l *Logger) log(logLevel int, msg string) {
	if l.Config.LogLevel <= logLevel {
		m := message.NewMessage(logLevel, msg)
		l.Config.SinkLocation.AddWait()
		l.Config.SinkLocation.LogAppender(m)
	}
}

func (l *Logger) Show() {
	l.Config.SinkLocation.Wait()
	fmt.Println(l.Namespace)
	l.Config.SinkLocation.ShowLogs()
}
