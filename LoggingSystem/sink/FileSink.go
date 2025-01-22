package sink

import (
	"fmt"
	"logSys/constants"
	"logSys/message"
	"sync"
)

type FileSink struct {
	messages []message.Message
	WG       *sync.WaitGroup
}

func (s *FileSink) LogAppender(msg message.Message) {
	s.messages = append(s.messages, msg)
	s.WG.Done()
}

func (s *FileSink) ShowLogs() {
	for i := len(s.messages) - 1; i >= 0; i-- {
		m := s.messages[i]
		fmt.Println(m.Timestamp, constants.GetLogLevel(m.LogLevel), m.Content)
	}
}

func (s *FileSink) AddWait() {
	s.WG.Add(1)
}

func (s *FileSink) Wait() {
	s.WG.Wait()
}
