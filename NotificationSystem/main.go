package main

import (
	"encoding/json"
	"fmt"
	"time"
)

func main() {
	manager := NewNotificationManager()
	dispatcher := NewNotificationDispatcher(manager, 10)

	notifyId1, err := manager.CreateNotification(1, CHANNEL_EMAIL, PRIORITY_MEDIUM, "hello I am muskan", time.Now().Add(time.Second*20))
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("notificationId created:", notifyId1)
	}

	notifyId2, err := manager.CreateNotification(1, CHANNEL_SMS, PRIORITY_LOW, "nice to meet you", time.Now().Add(time.Second*25))
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("notificationId created:", notifyId2)
	}

	notifyId3, err := manager.CreateNotification(2, CHANNEL_PUSH, PRIORITY_MEDIUM, "what happened", time.Now().Add(time.Second*35))
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("notificationId created:", notifyId3)
	}

	notifyId4, err := manager.CreateNotification(2, CHANNEL_EMAIL, PRIORITY_HIGH, "any requirements", time.Now().Add(time.Second*30))
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("notificationId created:", notifyId4)
	}

	notifyId5, err := manager.CreateNotification(2, CHANNEL_SMS, PRIORITY_HIGH, "any requirements", time.Now().Add(time.Second*20))
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("notificationId created:", notifyId5)
	}

	notifyId6, err := manager.CreateNotification(2, CHANNEL_PUSH, PRIORITY_HIGH, "any requirements", time.Now().Add(time.Second*37))
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("notificationId created:", notifyId6)
	}

	notifyId7, err := manager.CreateNotification(2, CHANNEL_EMAIL, PRIORITY_HIGH, "any requirements", time.Now().Add(time.Second*45))
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("notificationId created:", notifyId7)
	}

	notifyId8, err := manager.CreateNotification(2, CHANNEL_EMAIL, PRIORITY_HIGH, "any requirements", time.Now().Add(time.Second*26))
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("notificationId created:", notifyId8)
	}

	notification, _ := manager.GetNotificationById(3)
	notification.CancelNotification()

	notification, _ = manager.GetNotificationById(2)
	notification.UpdateScheduledTime(time.Now().Add(time.Minute))

	go dispatcher.StartDispatcher()
	time.Sleep(time.Second * 50)

	notifications := manager.ListNotifications(1, STATUS_SENT)
	notifyStr, _ := json.Marshal(notifications)
	fmt.Println(string(notifyStr))

	notifications = manager.ListNotifications(2, STATUS_SENT)
	notifyStr, _ = json.Marshal(notifications)
	fmt.Println(string(notifyStr))

}
