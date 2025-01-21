package main

import (
	"fmt"
	"sync"
)

type LogSink struct {
	Namespace string
	messages  []Message
	WG        *sync.WaitGroup
}

func NewSink(namespace string) *LogSink {
	return &LogSink{
		Namespace: namespace,
		messages:  []Message{},
		WG:        new(sync.WaitGroup),
	}
}

func (s *LogSink) AddLog(msg Message) {
	s.messages = append(s.messages, msg)
	if msg.LogLevel >= logConfig[msg.Namespace].LogLevel {
		fmt.Println(logConfig[msg.Namespace].SinkType, GetLogLevel(msg.LogLevel), msg.Timestamp, msg.Content, msg.Namespace)
	} else {
		fmt.Println("This log level is below configured Loglevel:", GetLogLevel(msg.LogLevel))
	}
	s.WG.Done()
}

func (s *LogSink) ShowMessages() {
	for i := len(s.messages) - 1; i >= 0; i-- {
		s.messages[i].Show()
	}
}
