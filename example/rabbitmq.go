package main

import (
	"github.com/google/uuid"
	"github.com/maykonlf/pubsub"
	"github.com/maykonlf/pubsub/rabbitmq"
	"log"
	"time"
)

func main() {
	go publisher()
	subscriber()
}

func publisher() {
	publisher := rabbitmq.NewPublisher(&rabbitmq.PublisherOptions{
		ConnectionOptions: &rabbitmq.ConnectionOptions{
			URI:                    "amqp://guest:guest@localhost:5672/",
		},
	})

	for i := 0; i < 100; i++ {
		go func() {
			_ = publisher.Publish(rabbitmq.NewMessage().SetCorrelationID(uuid.New()),
				"my-exchange")
		}()
	}
}

func subscriber() {
	subscriber := rabbitmq.NewSubscriber(&rabbitmq.SubscriberOptions{
		ConnectionOptions: &rabbitmq.ConnectionOptions{
			URI:                    "amqp://guest:guest@localhost:5672/",
		},
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

	subscriber.Subscribe(func(m pubsub.Message) {
		log.Printf("consumed message %s", m.ID())
		time.Sleep(100 * time.Millisecond)
		if err := m.Ack(); err != nil {
			log.Println(err)
		}
	})
}
