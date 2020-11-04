package main

import (
	"github.com/maykonlf/pubsub/pkg/rabbitmq"
	"log"
	"time"
)

func main() {
	go subscriber()
	publisher()
}

func publisher() {
	publisher := rabbitmq.NewPublisher(&rabbitmq.PublisherOptions{
		URI: "amqp://guest:guest@localhost:5672/",
	})

	for {
		for i := 0; i < 100; i++ {
			_ = publisher.Publish(&rabbitmq.PublishOptions{
				Exchange:    "my-exchange",
				RoutingKey:  "",
				IsMandatory: false,
				IsImmediate: false,
			}, rabbitmq.NewMessage())
		}

		time.Sleep(200 * time.Millisecond)
	}
}

func subscriber() {
	subscriber := rabbitmq.NewSubscriber(&rabbitmq.SubscriberOptions{
		URI: "amqp://guest:guest@localhost:5672/",
		QueueOptions: &rabbitmq.QueueOptions{
			Name:        "my-consumer-queue",
			Durable:     true,
			AutoDelete:  false,
			Exclusive:   false,
			NoWait:      false,
			MaxPriority: 5,
		},
		ExchangeOptions: &rabbitmq.ExchangeOptions{
			Name:          "my-exchange",
			Type:          rabbitmq.ExchangeTypeFanout,
			IsDurable:     true,
			IsAutoDeleted: false,
			IsInternal:    false,
			NoWait:        false,
		},
		PrefetchCount: 50,
		Name:          "my-consumer",
		AutoAck:       false,
		NoWait:        false,
		NoLocal:       false,
		Exclusive:     false,
	})

	subscriber.Subscribe(func(m *rabbitmq.Message) {
		log.Println("consumed message")
		time.Sleep(100 * time.Millisecond)
		if err := m.Ack(); err != nil {
			log.Println(err)
		}
	})
}
