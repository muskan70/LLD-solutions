package main

import "fmt"

type Subscriber struct {
	name string
}

func NewSubscriber(n string) *Subscriber {
	return &Subscriber{name: n}
}

func (s *Subscriber) SubscribeTopic(topicName string) {
	AddSubscriberToTopic(topicName, s.name)
}

func (s *Subscriber) UnsubscribeTopic(topicName string) {
	RemoveSubscriberFromTopic(topicName, s.name)
}

func (s *Subscriber) ConsumeMessageFromTopic(topicName string) {
	fmt.Println("Subscriber Name", s.name, "Topic Name", topicName)
	GetMessageFromTopic(topicName, s.name)
}
