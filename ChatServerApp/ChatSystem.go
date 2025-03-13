package main

import "errors"

type ChatSystem struct {
	Chats  map[int]*Chat
	Groups map[int]*Group
	Users  map[int]*User
}

func NewChatSystem() *ChatSystem {
	return &ChatSystem{
		Chats:  make(map[int]*Chat),
		Groups: make(map[int]*Group),
		Users:  make(map[int]*User),
	}
}

func (c *ChatSystem) HandleDirectMessage(req sendRequest) error {
	sender, ok1 := c.Users[req.SenderId]
	receiver, ok2 := c.Users[req.ReceiverId]
	if !ok1 || !ok2 {
		return errors.New("this senderId or receiverId is not registered")
	}
	chatId := sender.GetChatId(req.ReceiverId)
	var chat *Chat
	if chatId == -1 {
		chat = NewChat()
		c.Chats[chat.ChatId] = chat
		sender.AddChatId(req.ReceiverId, chat.ChatId)
		receiver.AddChatId(req.SenderId, chat.ChatId)
	} else {
		chat = c.Chats[chatId]
	}
	chat.AddMessage(NewMessage(req.ReceiverId, chat.ChatId, req.Content))
	return nil
}

func (c *ChatSystem) CreateGroup(req CreateGroupReq) int {
	g := NewGroup(req.GroupName, req.AdminId)
	chat := NewChat()
	g.ChatId = chat.ChatId
	c.Chats[g.ChatId] = chat
	c.Groups[g.GroupId] = g
	return g.GroupId
}

func (c *ChatSystem) AddUserToGroup(req AddUserToGroupReq) error {
	if _, ok := c.Groups[req.GroupId]; !ok {
		return errors.New("this groupId doesn't exist")
	}
	if _, ok := c.Users[req.UserId]; !ok {
		return errors.New("this userId doesn't exist")
	}
	c.Groups[req.GroupId].AddParticipant(req.UserId)
	return nil
}

func (c *ChatSystem) RemoveUserFromGroup(req RemoveUserFromGroupReq) error {
	if _, ok := c.Groups[req.GroupId]; !ok {
		return errors.New("this groupId doesn't exist")
	}
	if _, ok := c.Users[req.UserId]; !ok {
		return errors.New("this userId doesn't exist")
	}
	c.Groups[req.GroupId].RemoveParticipant(req.UserId)
	return nil
}

func (c *ChatSystem) HandleGroupMessage(req AddMessageToGroupReq) error {
	if _, ok := c.Users[req.UserId]; !ok {
		return errors.New("this userId doesn't exist")
	}
	group, ok := c.Groups[req.GroupId]
	if !ok {
		return errors.New("this groupId doesn't exist")
	}
	if chatId, err := group.GetChatId(req.UserId); err != nil {
		return err
	} else {
		chat := c.Chats[chatId]
		chat.AddMessage(NewMessage(req.UserId, chat.ChatId, req.Content))
	}
	return nil
}

func (c *ChatSystem) GetGroupChatHistory(req GetMessagesFromGroupRequest) ([]Message, error) {
	if _, ok := c.Users[req.UserId]; !ok {
		return nil, errors.New("this userId doesn't exist")
	}
	group, ok := c.Groups[req.GroupId]
	if !ok {
		return nil, errors.New("this groupId doesn't exist")
	}
	if chatId, err := group.GetChatId(req.UserId); err != nil {
		return nil, err
	} else {
		chat := c.Chats[chatId]
		return chat.GetAllMessages(), nil
	}
}
