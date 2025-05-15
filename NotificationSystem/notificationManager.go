package main

import (
	"errors"
	"time"
)

type NotificationManager struct {
	notifications         map[uint64]*Notification
	notificationsByUserId map[int][]uint64
}

func NewNotificationManager() *NotificationManager {
	return &NotificationManager{
		notifications:         make(map[uint64]*Notification),
		notificationsByUserId: make(map[int][]uint64),
	}
}

func (n *NotificationManager) CreateNotification(userId, channel, priority int, message string, scheduledTime time.Time) (uint64, error) {
	notification, err := NewNotification(userId, channel, priority, message, scheduledTime)
	if err != nil {
		return 0, err
	}
	n.notifications[notification.Id] = notification
	n.notificationsByUserId[userId] = append(n.notificationsByUserId[userId], notification.Id)
	return notification.Id, nil
}

func (n *NotificationManager) GetNotificationById(id uint64) (*Notification, error) {
	notification, ok := n.notifications[id]
	if !ok {
		return nil, errors.New("invalid notificationId")
	}
	return notification, nil
}

func (n *NotificationManager) GetScheduledNotifications() []*Notification {
	var scheduledNotifications []*Notification
	for _, notification := range n.notifications {
		if notification.Status == STATUS_NEW && notification.ScheduledTime.Before(time.Now()) {
			scheduledNotifications = append(scheduledNotifications, notification)
		}
	}
	return scheduledNotifications
}

func (n *NotificationManager) ListNotifications(userId int, status int) []*Notification {
	var userNotifications []*Notification
	_, ok := n.notificationsByUserId[userId]
	if !ok {
		return userNotifications
	}
	for _, id := range n.notificationsByUserId[userId] {
		notification, err := n.GetNotificationById(id)
		if err == nil && notification.Status == status {
			userNotifications = append(userNotifications, notification)
		}
	}
	return userNotifications
}
