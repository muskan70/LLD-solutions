package main

import (
	"time"
)

var messageId = 0

const (
	MessageStatusSent = iota + 1
	MessageStatusDelivered
	MessageStatusSeen
)

type Message struct {
	MessageID int
	UserId    int    `json:"userId"`
	ChatId    int    `json:"chatId"`
	Content   string `json:"message"`
	StatusMap map[int]time.Time
}

func NewMessage(userId, chatId int, content string) Message {
	messageId++
	return Message{
		MessageID: messageId,
		UserId:    userId,
		ChatId:    chatId,
		Content:   content,
		StatusMap: map[int]time.Time{
			MessageStatusSent: time.Now(),
		}}
}

func (m *Message) AddStatus(status int) {
	m.StatusMap[status] = time.Now()
}
