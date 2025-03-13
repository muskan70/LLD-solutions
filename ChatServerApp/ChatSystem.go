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

func (c *ChatSystem) HandleDirectMessage(req sendDirectMessageRequest) error {
	sender, ok1 := c.Users[req.SenderId]
	receiver, ok2 := c.Users[req.ReceiverId]
	if !ok1 || !ok2 {
		return errors.New("this senderId or receiverId is not registered")
	}
	chatId := sender.GetChatId(req.ReceiverId)
	var chat *Chat
	if chatId == -1 {
		chat = NewDirectChat(req.SenderId, req.ReceiverId)
		c.Chats[chat.ChatId] = chat
		sender.AddChatId(req.ReceiverId, chat.ChatId)
		receiver.AddChatId(req.SenderId, chat.ChatId)
	} else {
		chat = c.Chats[chatId]
	}
	chat.AddMessage(NewMessage(req.ReceiverId, chat.ChatId, req.Content))
	return nil
}

func (c *ChatSystem) ReceiveMessages(userId int) (map[int][]Message, map[int][]Message, error) {
	usr, ok := c.Users[userId]
	if !ok {
		return nil, nil, errors.New("this userId doesn't exist")
	}
	directMsgs := make(map[int][]Message)
	for usr, chatId := range usr.DirectChats {
		if msgs := c.Chats[chatId].ReceiveMessage(userId); len(msgs) > 0 {
			directMsgs[usr] = msgs
		}
	}
	grpMsgs := make(map[int][]Message)
	for grp, chatId := range usr.GroupChats {
		if msgs := c.Chats[chatId].ReceiveMessage(userId); len(msgs) > 0 {
			grpMsgs[grp] = msgs
		}
	}
	return directMsgs, grpMsgs, nil
}

func (c *ChatSystem) ChatHistoryOfUser(userId int) (map[int][]Message, map[int][]Message, error) {
	usr, ok := c.Users[userId]
	if !ok {
		return nil, nil, errors.New("this userId doesn't exist")
	}
	directMsgs := make(map[int][]Message)
	for usr, chatId := range usr.DirectChats {
		if msgs := c.Chats[chatId].GetAllMessages(); len(msgs) > 0 {
			directMsgs[usr] = msgs
		}
	}
	grpMsgs := make(map[int][]Message)
	for grp, chatId := range usr.GroupChats {
		if msgs := c.Chats[chatId].GetAllMessages(); len(msgs) > 0 {
			grpMsgs[grp] = msgs
		}
	}
	return directMsgs, grpMsgs, nil
}

func (c *ChatSystem) CreateGroup(req CreateGroupReq) int {
	g := NewGroup(req.GroupName, req.AdminId)
	chat := NewGroupChat(req.AdminId)
	g.ChatId = chat.ChatId
	c.Chats[g.ChatId] = chat
	c.Groups[g.GroupId] = g
	return g.GroupId
}

func (c *ChatSystem) AddUserToGroup(req AddUserToGroupReq) error {
	grp, ok := c.Groups[req.GroupId]
	if !ok {
		return errors.New("this groupId doesn't exist")
	}
	usr, ok := c.Users[req.UserId]
	if !ok {
		return errors.New("this userId doesn't exist")
	}
	grp.AddParticipant(req.UserId)
	c.Chats[grp.ChatId].AddUserToGroupChat(req.UserId)
	usr.AddGroupChatId(req.GroupId, grp.ChatId)
	return nil
}

func (c *ChatSystem) RemoveUserFromGroup(req RemoveUserFromGroupReq) error {
	grp, ok := c.Groups[req.GroupId]
	if !ok {
		return errors.New("this groupId doesn't exist")
	}
	usr, ok := c.Users[req.UserId]
	if !ok {
		return errors.New("this userId doesn't exist")
	}
	grp.RemoveParticipant(req.UserId)
	usr.RemoveGroupChatId(req.GroupId)
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
