package main

import (
	"time"
)

var messageId = 0

type Message struct {
	MessageID   int
	UserId      int    `json:"userId"`
	ChatId      int    `json:"chatId"`
	Content     string `json:"message"`
	SentAt      time.Time
	DeliveredAt map[int]time.Time
	SeenAt      map[int]time.Time
}

func NewMessage(userId, chatId int, content string) Message {
	messageId++
	return Message{
		MessageID: messageId,
		UserId:    userId,
		ChatId:    chatId,
		Content:   content,
		SentAt:    time.Now(),
	}
}
