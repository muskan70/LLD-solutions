package main

import "sync"

type LogSink struct {
	Namespace string
	messages  []Message
	//WG       *sync.WaitGroup
}

func NewSink(namespace string) *LogSink {
	return &LogSink{
		Namespace: namespace,
		messages:  []Message{},
		//WG:       new(sync.WaitGroup),
	}
}

func (s *LogSink) AddLog(msg Message, wg *sync.WaitGroup) {
	//s.WG.Add(1)
	s.messages = append(s.messages, msg)
	wg.Done()
}

func (s *LogSink) ShowMessages() {
	for i := len(s.messages) - 1; i >= 0; i-- {
		s.messages[i].Show()
	}
}
