package sink

import (
	"logSys/constants"
	"logSys/message"
	"sync"
)

type Sink interface {
	LogAppender(msg message.Message)
	AddWait()
	ShowLogs()
	Wait()
}

func NewSink(sinkType string) Sink {
	if sinkType == constants.SinkTypeCONSOLE {
		return &ConsoleSink{
			messages: []message.Message{},
			WG:       new(sync.WaitGroup),
		}
	} else if sinkType == constants.SinkTypeFILE {
		return &FileSink{
			messages: []message.Message{},
			WG:       new(sync.WaitGroup),
		}
	} else {
		return &DatabaseSink{
			messages: []message.Message{},
			WG:       new(sync.WaitGroup),
		}
	}
}
