package main

import (
	"fmt"
	"time"
)

type NotificationDispatcher struct {
	manager              *NotificationManager
	LastChecked          time.Time
	CheckWindowInSeconds int
}

func NewNotificationDispatcher(m *NotificationManager, window int) *NotificationDispatcher {
	return &NotificationDispatcher{
		manager:     m,
		LastChecked: time.Now(),
		//CheckWindowInSeconds: window,
	}
}

func (dispatcher *NotificationDispatcher) DispatchScheduledNotifications() error {
	// if dispatcher.LastChecked.Add(time.Duration(dispatcher.CheckWindowInSeconds)).After(time.Now()) {
	// 	return errors.New("notifications dispactched for current window")
	// }
	notifications := dispatcher.manager.GetScheduledNotifications()
	dispatcher.LastChecked = time.Now()

	fmt.Println(time.Now(), "notifications to dispatch", len(notifications))

	for i := range notifications {
		notifier := GetNotifier(notifications[i].Channel)
		if notifier == nil {
			fmt.Println("notification channel not valid :", notifications[i].Id)
			continue
		}
		notifications[i].UpdateStatus(STATUS_IN_TRANSIT)
		if err := notifier.SendNotification(notifications[i]); err != nil {
			fmt.Println(err)
			notifications[i].UpdateStatus(STATUS_FAILED)
		} else {
			notifications[i].UpdateStatus(STATUS_SENT)
		}
	}
	return nil
}

func (dispatcher *NotificationDispatcher) StartDispatcher() {
	for {
		if err := dispatcher.DispatchScheduledNotifications(); err != nil {
			fmt.Println(err.Error())
		}
		time.Sleep(5 * time.Second)
	}
}
