package main

var chatId = 0

type ChatUserMapping struct {
	ChatId int
	UserId int
}

type Chat struct {
	ChatId   int
	Messages []Message
}

func NewChat() *Chat {
	chatId++
	return &Chat{
		ChatId:   chatId,
		Messages: []Message{},
	}
}

func (c *Chat) AddMessage(msg Message) {
	c.Messages = append(c.Messages, msg)
}

func (c *Chat) GetAllMessages() []Message {
	return c.Messages
}
