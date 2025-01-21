package main

import (
	"fmt"
	"time"
)

type Message struct {
	LogLevel  int
	Content   string
	Namespace string
	Timestamp time.Time
}

func NewMessage(logLevel int, content, namespace string) Message {
	return Message{
		LogLevel:  logLevel,
		Content:   content,
		Namespace: namespace,
	}
}

func (m *Message) Log() {
	m.Timestamp = time.Now()
	logConfig[m.Namespace].SinkLocation.WG.Add(1)
	go logConfig[m.Namespace].SinkLocation.AddLog(*m)
}

func (m *Message) Show() {
	fmt.Println(m.Timestamp, GetLogLevel(m.LogLevel), m.Namespace, m.Content)
}
