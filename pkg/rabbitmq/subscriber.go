package rabbitmq

import (
	"github.com/streadway/amqp"
)

type Subscriber interface {
	Subscribe(handler func(body []byte) error)
}

func NewSubscriber(options *SubscriberOptions) Subscriber {
	conn := NewConnection(&ConnectionOptions{
		URI: options.URI,
	})

	return &subscriber{
		conn:                      conn,
		subscriberOptions:         options,
		active:                    true,
		disconnectionErrorChannel: make(chan error),
	}
}

type subscriber struct {
	active                    bool
	disconnectionErrorChannel chan error
	subscriberOptions         *SubscriberOptions
	subscriberHandler         func(body []byte) error
	messageDeliveryChannel    <-chan amqp.Delivery
	queue                     amqp.Queue
	conn                      Connection
}

func (s *subscriber) Subscribe(handler func(body []byte) error) {
	s.registerSubscriberHandler(handler)
	s.setupSubscriber()
	s.openConsumerChannel()
	s.startSubscriber()
}

func (s *subscriber) registerSubscriberHandler(handler func(body []byte) error) {
	s.subscriberHandler = handler
}

func (s *subscriber) setupSubscriber() {
	s.setupChannelQos()
	s.setupExchange()
	s.setupQueue()
	s.bindQueueToExchange()
}

func (s *subscriber) openConsumerChannel() {
	delivery, err := s.conn.GetChannel().Consume(
		s.queue.Name,
		s.subscriberOptions.Name,
		s.subscriberOptions.AutoAck,
		s.subscriberOptions.Exclusive,
		s.subscriberOptions.NoLocal,
		s.subscriberOptions.NoWait,
		s.subscriberOptions.Args)
	if err != nil {
		panic(err)
	}

	s.messageDeliveryChannel = delivery
}

func (s *subscriber) startSubscriber() {
	s.conn.SetReconnectHooks(s.reconnectSubscriber)
	for {
		s.handleConsume()
	}
}

func (s *subscriber) setupChannelQos() {
	err := s.conn.GetChannel().Qos(s.subscriberOptions.PrefetchCount, 0, false)
	if err != nil {
		panic(err)
	}
}

func (s *subscriber) setupQueue() {
	queue, err := s.conn.GetChannel().QueueDeclare(
		s.subscriberOptions.QueueOptions.Name,
		s.subscriberOptions.QueueOptions.Durable,
		s.subscriberOptions.QueueOptions.AutoDelete,
		s.subscriberOptions.QueueOptions.Exclusive,
		s.subscriberOptions.QueueOptions.NoWait,
		s.subscriberOptions.QueueOptions.GetArgs())
	if err != nil {
		panic(err)
	}

	s.queue = queue
}

func (s *subscriber) reconnectSubscriber() {
	s.setupSubscriber()
	s.openConsumerChannel()
}

func (s *subscriber) handleConsume() {
	for delivery := range s.messageDeliveryChannel {
		message := delivery
		go s.handleDelivery(&message)
	}
}

func (s *subscriber) handleDelivery(delivery *amqp.Delivery) {
	err := s.subscriberHandler(delivery.Body)
	if err != nil {
		_ = delivery.Nack(false, true)
		return
	}
	_ = delivery.Ack(false)
}

func (s *subscriber) setupExchange() {
	if s.subscriberOptions.ExchangeOptions == nil {
		return
	}

	err := s.conn.GetChannel().ExchangeDeclare(
		s.subscriberOptions.ExchangeOptions.Name,
		s.subscriberOptions.ExchangeOptions.Type.String(),
		s.subscriberOptions.ExchangeOptions.IsDurable,
		s.subscriberOptions.ExchangeOptions.IsAutoDeleted,
		s.subscriberOptions.ExchangeOptions.IsInternal,
		s.subscriberOptions.ExchangeOptions.NoWait,
		s.subscriberOptions.ExchangeOptions.Args)
	if err != nil {
		panic(err)
	}
}

func (s *subscriber) bindQueueToExchange() {
	if s.subscriberOptions.ExchangeOptions == nil {
		return
	}

	err := s.conn.GetChannel().QueueBind(
		s.queue.Name,
		s.subscriberOptions.QueueOptions.RoutingKey,
		s.subscriberOptions.ExchangeOptions.Name,
		s.subscriberOptions.QueueOptions.NoWait,
		s.subscriberOptions.QueueOptions.QueueBindArgs)
	if err != nil {
		panic(err)
	}
}
