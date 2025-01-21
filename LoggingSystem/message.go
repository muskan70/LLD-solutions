package main

import (
	"fmt"
	"sync"
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

func (m *Message) Log(wg *sync.WaitGroup) {
	m.Timestamp = time.Now()
	logConfig[m.Namespace].SinkLocation.AddLog(*m, wg)
	if m.LogLevel >= logConfig[m.Namespace].LogLevel {
		fmt.Println(logConfig[m.Namespace].SinkType, GetLogLevel(m.LogLevel), m.Timestamp, m.Content, m.Namespace)
	} else {
		fmt.Println("This log level is below configured Loglevel:", GetLogLevel(m.LogLevel))
	}
}

func (m *Message) Show() {
	fmt.Println(m.Timestamp, GetLogLevel(m.LogLevel), m.Namespace, m.Content)
}
