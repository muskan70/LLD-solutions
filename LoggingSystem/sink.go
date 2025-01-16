package main

type LogSink struct {
	LogLevel int
	messages []Message
	//WG       *sync.WaitGroup
}

func NewSink(level int) *LogSink {
	return &LogSink{
		LogLevel: level,
		messages: []Message{},
		//WG:       new(sync.WaitGroup),
	}
}

func CreateLogSink() map[int]*LogSink {
	return map[int]*LogSink{
		LogLevelFATAL: NewSink(LogLevelFATAL),
		LogLevelERROR: NewSink(LogLevelERROR),
		LogLevelWARN:  NewSink(LogLevelWARN),
		LogLevelINFO:  NewSink(LogLevelINFO),
		LogLevelDEBUG: NewSink(LogLevelDEBUG),
	}
}

func (s *LogSink) AddLog(msg Message) {
	//s.WG.Add(1)
	s.messages = append(s.messages, msg)
	//s.WG.Done()
}

func (s *LogSink) ShowMessages() {
	for i := len(s.messages) - 1; i >= 0; i-- {
		s.messages[i].Show()
	}
}
