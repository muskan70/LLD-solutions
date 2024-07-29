package main

func main() {
	// 	Test Case
	// Create Topic
	// Vegetables, Fruits

	// Subscribers
	// Rahul
	// Karan
	// Mohit

	// Publish
	//  Vegetables -> “carrot”, “potato”, “tomato”
	//  Fruits -> “apple”, “banana”,

	// Subscribe
	// Rahul -> Vegetables
	// Karan -> Fruits
	// Mohit -> Vegetables, Fruits

	// Rahul . Consume Vegetables
	// Karan . Consume Vegetables
	// Mohit . Consume Vegetables

	// Rahul . Consume fruits
	// Karan . Consume fruits
	// Mohit . Consume fruits

	// Rahul . Consume Vegetables
	// Karan . Consume fruits
	// Mohit . Consume Vegetables

	// UnSubscribe
	// Mohit -> fruits

	// Karan . Consume fruits
	// Mohit . Consume Fruits

	AddTopic("Vegetables")
	AddTopic("Fruits")

	c1 := NewSubscriber("Rahul")
	c2 := NewSubscriber("Karan")
	c3 := NewSubscriber("Mohit")

	PublishMessageToTopic("Vegetables", "carrot")
	PublishMessageToTopic("Vegetables", "potato")
	PublishMessageToTopic("Vegetables", "tomato")

	PublishMessageToTopic("Fruits", "apple")
	PublishMessageToTopic("Fruits", "banana")

	c1.SubscribeTopic("Vegetables")
	c2.SubscribeTopic("Fruits")
	c3.SubscribeTopic("Vegetables")
	c3.SubscribeTopic("Fruits")

	c1.ConsumeMessageFromTopic("Vegetables")
	c2.ConsumeMessageFromTopic("Vegetables")
	c3.ConsumeMessageFromTopic("Vegetables")

	c1.ConsumeMessageFromTopic("Fruits")
	c2.ConsumeMessageFromTopic("Fruits")
	c3.ConsumeMessageFromTopic("Fruits")

	c1.ConsumeMessageFromTopic("Vegetables")
	c2.ConsumeMessageFromTopic("Fruits")
	c3.ConsumeMessageFromTopic("Vegetables")

	c3.UnsubscribeTopic("Fruits")
	c2.ConsumeMessageFromTopic("Fruits")
	c3.ConsumeMessageFromTopic("Fruits")

	c3.SubscribeTopic("Fruits")
	c3.ConsumeMessageFromTopic("Fruits")

	GetListOfTopics()
	GetMessageStatusOfTopic("Fruits", "banana")

}
