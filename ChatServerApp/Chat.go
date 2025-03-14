package main

var chatId = 0

type Chat struct {
	ChatId      int
	Messages    []Message
	UsersOffset map[int]int
}

func NewDirectChat(userId1, userId2 int) *Chat {
	chatId++
	return &Chat{
		ChatId:   chatId,
		Messages: []Message{},
		UsersOffset: map[int]int{
			userId1: 0,
			userId2: 0,
		},
	}
}

func NewGroupChat(userId int) *Chat {
	chatId++
	return &Chat{
		ChatId:   chatId,
		Messages: []Message{},
		UsersOffset: map[int]int{
			userId: 0,
		},
	}
}

func (c *Chat) AddUserToGroupChat(userId int) {
	if _, ok := c.UsersOffset[userId]; !ok {
		c.UsersOffset[userId] = 0
	}
}

func (c *Chat) AddMessage(msg Message) {
	c.Messages = append(c.Messages, msg)
}

func (c *Chat) GetAllMessages() []Message {
	return c.Messages
}

func (c *Chat) ReceiveMessage(userId int) []Message {
	if c.UsersOffset[userId] == len(c.Messages) {
		return []Message{}
	}
	msgs := c.Messages[c.UsersOffset[userId]:]
	for i := range msgs {
		msgs[i].AddStatus(MessageStatusDelivered)
	}
	c.UsersOffset[userId] = len(c.Messages)
	return msgs
}
