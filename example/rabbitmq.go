package main

import (
	"github.com/google/uuid"
	"github.com/maykonlf/pubsub"
	"github.com/maykonlf/pubsub/rabbitmq"
	"github.com/maykonlf/pubsub/rabbitmq/publisher"
	"github.com/maykonlf/pubsub/rabbitmq/subscriber"
	"log"
	"time"
)

func main() {
	go startPublisher()
	startSubscriber()
}

func startPublisher() {
	pub := publisher.NewPublisher("amqp://guest:guest@localhost:5672/")

	for {
		_ = pub.Publish(rabbitmq.NewMessage().SetCorrelationID(uuid.New()),
			"my-topic", "my-key")
	}
}

func startSubscriber() {
	subs := subscriber.NewSubscriber("amqp://guest:guest@localhost:5672/",
		subscriber.WithName("my consumer name"),
		subscriber.WithDurableTopicExchange("my-topic", "my-key"),
		subscriber.WithDurableFanoutExchange("my-fanout"),
		subscriber.WithDurablePriorityQueue("my-priority-queue", 5),
		subscriber.WithPrefetch(20))

	subs.Subscribe(func(m pubsub.Message) {
		log.Printf("consumed message %s", m.ID())
		time.Sleep(100 * time.Millisecond)
		if err := m.Ack(); err != nil {
			log.Println(err)
		}
	})
}
