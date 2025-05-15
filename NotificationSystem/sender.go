package main

import "fmt"

type INotifier interface {
	SendNotification(n *Notification) error
}

func GetNotifier(channel int) INotifier {
	if channel == CHANNEL_EMAIL {
		return &EmailNotifier{}
	} else if channel == CHANNEL_PUSH {
		return &PushNotifier{}
	} else if channel == CHANNEL_SMS {
		return &SMSNotifier{}
	}
	return nil
}

type EmailNotifier struct {
}

func (e *EmailNotifier) SendNotification(n *Notification) error {
	fmt.Println("notification sent via email")
	return nil
}

type SMSNotifier struct {
}

func (e *SMSNotifier) SendNotification(n *Notification) error {
	fmt.Println("notification sent via sms")
	return nil
}

type PushNotifier struct {
}

func (e *PushNotifier) SendNotification(n *Notification) error {
	fmt.Println("notification sent via push notifier")
	return nil
}
