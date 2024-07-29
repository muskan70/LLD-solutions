package main

import "fmt"

type Topic struct {
	name          string
	messages      []string
	msgReverseMap map[string]int
	subscribers   map[string][]int
}

func NewTopic(s string) *Topic {
	return &Topic{name: s, messages: []string{}, msgReverseMap: make(map[string]int), subscribers: make(map[string][]int)}
}

func (t *Topic) PublishMessage(msg string) {
	t.messages = append(t.messages, msg)
	t.msgReverseMap[msg] = len(t.messages) - 1
}

func (t *Topic) ConsumeMessage(subName string) {
	subscriberDetails, ok := t.subscribers[subName]
	if !ok {
		fmt.Println("This topic is not subscribed by subscriber")
		return
	}
	curIdx := subscriberDetails[1]
	if len(t.messages) > subscriberDetails[1] {
		t.subscribers[subName] = []int{1, curIdx + 1}
		fmt.Println(t.messages[curIdx])
	} else {
		fmt.Println("No new Message published in this topic")
	}
}

func (t *Topic) AddSubscriber(s string) {
	subscriberDetails, ok := t.subscribers[s]
	if ok {
		t.subscribers[s] = []int{1, subscriberDetails[1]}
	} else {
		t.subscribers[s] = []int{1, 0}
	}
}

func (t *Topic) RemoveSubscriber(s string) {
	subscriberDetails := t.subscribers[s]
	t.subscribers[s] = []int{0, subscriberDetails[1]}
}

func (t *Topic) GetMessageStatus(msg string) {
	msgIdx := t.msgReverseMap[msg]
	for sub, subDetails := range t.subscribers {
		if subDetails[0] == 1 {
			if msgIdx < subDetails[1] {
				fmt.Println("This message is read by subscriber:", sub)
			} else {
				fmt.Println("This message is not read by subscriber:", sub)
			}
		}
	}
}

var topics map[string]*Topic

func AddTopic(topicName string) {
	if topics == nil {
		topics = make(map[string]*Topic)
	}
	topics[topicName] = NewTopic(topicName)
}

func GetListOfTopics() {
	fmt.Println("List of present Topics:")
	for topicName := range topics {
		fmt.Println(topicName)
	}
}

func PublishMessageToTopic(topicName string, msg string) {
	if topic, ok := topics[topicName]; ok {
		topic.PublishMessage(msg)
	}
}

func GetMessageFromTopic(topicName, subName string) {
	topics[topicName].ConsumeMessage(subName)
}

func AddSubscriberToTopic(topicName string, subName string) {
	topics[topicName].AddSubscriber(subName)
}

func RemoveSubscriberFromTopic(topicName string, subName string) {
	topics[topicName].RemoveSubscriber(subName)
}

func GetMessageStatusOfTopic(topicName string, msgName string) {
	topics[topicName].GetMessageStatus(msgName)
}
