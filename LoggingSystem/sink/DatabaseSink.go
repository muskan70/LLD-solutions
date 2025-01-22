package sink

import (
	"fmt"
	"logSys/constants"
	"logSys/message"
	"sync"
)

type DatabaseSink struct {
	messages []message.Message
	WG       *sync.WaitGroup
}

func (s *DatabaseSink) LogAppender(msg message.Message) {
	s.messages = append(s.messages, msg)
	s.WG.Done()
}

func (s *DatabaseSink) ShowLogs() {
	for i := len(s.messages) - 1; i >= 0; i-- {
		m := s.messages[i]
		fmt.Println(m.Timestamp, constants.GetLogLevel(m.LogLevel), m.Content)
	}
}

func (s *DatabaseSink) AddWait() {
	s.WG.Add(1)
}

func (s *DatabaseSink) Wait() {
	s.WG.Wait()
}
