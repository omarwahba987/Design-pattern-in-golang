package main

import (
	"fmt"
	mq "patternsample"
	"time"
)

func main() {

	// New
	mque := mq.New()

	// Subscribe topic
	mque.Subscribe("Notify", func(m *mq.Message) {
		fmt.Println("Received, Topic: Notify")
	})

	// Subscribe topic
	mque.Subscribe("Test2", func(m *mq.Message) {
		fmt.Println("Received, Topic: Test2")
	})

	// Publish value into topic
	mque.Publish(
		mq.Message{
			Topic: "Notify",
			Value: "blablabla",
		},
	)

	// ;)
	time.Sleep(2 * time.Second)

	// Publish value into topic
	mque.Publish(
		mq.Message{
			Topic: "Test2",
			Value: "blablabla",
		},
	)

	// Publish value into topic
	mque.Publish(
		mq.Message{
			Topic: "NotFound",
			Value: "blablabla",
		},
	)

	// ;)
	time.Sleep(1 * time.Second)

}
