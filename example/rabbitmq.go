package main

import (
	"github.com/google/uuid"
	"github.com/maykonlf/pubsub"
	"github.com/maykonlf/pubsub/rabbitmq"
	publisher2 "github.com/maykonlf/pubsub/rabbitmq/publisher"
	subscriber2 "github.com/maykonlf/pubsub/rabbitmq/subscriber"
	"log"
	"time"
)

func main() {
	go publisher()
	subscriber()
}

func publisher() {
	publisher := publisher2.NewPublisher("amqp://guest:guest@localhost:5672/")

	for {
		_ = publisher.Publish(rabbitmq.NewMessage().SetCorrelationID(uuid.New()),
			"my-topic", "my-key")
	}
}

func subscriber() {
	subs := subscriber2.NewSubscriber("amqp://guest:guest@localhost:5672/",
		subscriber2.WithName("my consumer name"),
		subscriber2.WithDurableTopicExchange("my-topic", "my-key"),
		subscriber2.WithDurableFanoutExchange("my-fanout"),
		subscriber2.WithDurablePriorityQueue("my-priority-queue", 5))

	subs.Subscribe(func(m pubsub.Message) {
		log.Printf("consumed message %s", m.ID())
		time.Sleep(100 * time.Millisecond)
		if err := m.Ack(); err != nil {
			log.Println(err)
		}
	})
}
