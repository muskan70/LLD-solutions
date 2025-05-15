package main

import (
	"errors"
	"sync/atomic"
	"time"
)

var notificationId atomic.Uint64

type Notification struct {
	Id            uint64
	UserId        int
	Channel       int
	Message       string
	ScheduledTime time.Time
	Priority      int
	Status        int
}

func NewNotification(userId, channel, priority int, message string, scheduledTime time.Time) (*Notification, error) {
	if scheduledTime.Before(time.Now()) {
		return nil, errors.New("Invalid scheduled time")
	}
	if len(message) == 0 {
		return nil, errors.New("empty message")
	}
	notification := &Notification{
		Id:            notificationId.Add(1),
		UserId:        userId,
		Channel:       channel,
		Message:       message,
		ScheduledTime: scheduledTime,
		Priority:      priority,
		Status:        STATUS_NEW,
	}
	return notification, nil
}

func (n *Notification) UpdateScheduledTime(newTime time.Time) error {
	if newTime.Before(time.Now()) {
		return errors.New("invalid scheduled time")
	}
	n.ScheduledTime = newTime
	return nil
}

func (n *Notification) UpdateMessage(message string) error {
	if len(message) == 0 {
		return errors.New("empty message")
	}
	n.Message = message
	return nil
}

func (n *Notification) CancelNotification() error {
	if n.ScheduledTime.Before(time.Now()) {
		return errors.New("notification already processed")
	}
	n.Status = STATUS_CANCELLED
	return nil
}

func (n *Notification) UpdateStatus(status int) {
	n.Status = status
}
