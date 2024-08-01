package main

import "log"

type MessageList struct {
	msg []string
}

func NewMessageList() MessageList {
	return MessageList{msg: []string{}}
}

func (m *MessageList) AddMessage(curMsg string) {
	m.msg = append(m.msg, curMsg)
	log.Println(m.msg)
}

func (m *MessageList) GetAllMessages() string {
	msgs := ""
	for i := range m.msg {
		msgs = msgs + m.msg[i] + "\n"
	}
	return msgs
}
