package message

import (
	"time"
)

type Message struct {
	Content   string
	Timestamp time.Time
	LogLevel  int
}

func NewMessage(logLevel int, content string) Message {
	return Message{
		Content:   content,
		LogLevel:  logLevel,
		Timestamp: time.Now(),
	}
}
