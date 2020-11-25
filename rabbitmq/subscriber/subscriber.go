package subscriber

import (
	"github.com/maykonlf/pubsub"
	"github.com/maykonlf/pubsub/rabbitmq"
	"github.com/maykonlf/pubsub/rabbitmq/connection"
	"github.com/streadway/amqp"
)

type Subscriber interface {
	pubsub.Subscriber
	AddExchange(exchange *Exchange)
	SetQueue(queue *Queue)
	SetName(name string)
	SetPrefetchQos(qos *PrefetchQos)
}

func NewSubscriber(uri string, options ...Option) Subscriber {
	subscriber := &subscriber{
		name:                      "",
		active:                    true,
		disconnectionErrorChannel: make(chan error),
		queue:                     &Queue{},
		prefetchQos:               &PrefetchQos{},
		connectionOptions:         &connection.ConnectionOptions{URI: uri},
	}

	for _, optionFunction := range options {
		optionFunction(subscriber)
	}

	return subscriber
}

type subscriber struct {
	active                    bool
	disconnectionErrorChannel chan error
	subscriberHandler         func(message pubsub.Message)
	messageDeliveryChannel    <-chan amqp.Delivery
	conn                      connection.Connection
	connectionOptions         *connection.ConnectionOptions
	exchanges                 []*Exchange
	queue                     *Queue
	isAutoAck                 bool
	name                      string
	isExclusive               bool
	isNoLocal                 bool
	noWaitForRabbitResponse   bool
	consumerArgs              map[string]interface{}
	prefetchQos               *PrefetchQos
}

func (s *subscriber) Subscribe(handler func(message pubsub.Message)) {
	s.registerSubscriberHandler(handler)
	s.setupSubscriber()
	s.openConsumerChannel()
	s.startSubscriber()
}

func (s *subscriber) registerSubscriberHandler(handler func(message pubsub.Message)) {
	s.subscriberHandler = handler
}

func (s *subscriber) setupSubscriber() {
	s.setupChannelQos()
	s.setupExchange()
	s.setupQueue()
	s.bindQueueToExchange()
}

func (s *subscriber) openConsumerChannel() {
	delivery, err := s.getConnection().GetChannel().Consume(
		s.queue.Name,
		s.name,
		s.isAutoAck,
		s.isExclusive,
		s.isNoLocal,
		s.noWaitForRabbitResponse,
		s.consumerArgs)
	if err != nil {
		panic(err)
	}

	s.messageDeliveryChannel = delivery
}

func (s *subscriber) startSubscriber() {
	s.getConnection().SetReconnectHooks(s.reconnectSubscriber)
	for {
		s.handleConsume()
	}
}

func (s *subscriber) setupChannelQos() {
	err := s.getConnection().GetChannel().Qos(s.prefetchQos.Count, s.prefetchQos.Size, s.prefetchQos.IsGlobal)
	if err != nil {
		panic(err)
	}
}

func (s *subscriber) setupQueue() {
	queue, err := s.getConnection().GetChannel().QueueDeclare(
		s.queue.Name,
		s.queue.Durable,
		s.queue.AutoDelete,
		s.queue.Exclusive,
		s.queue.NoWait,
		s.queue.GetArgs())
	if err != nil {
		panic(err)
	}

	s.queue.Name = queue.Name
}

func (s *subscriber) reconnectSubscriber() {
	s.setupSubscriber()
	s.openConsumerChannel()
}

func (s *subscriber) handleConsume() {
	for delivery := range s.messageDeliveryChannel {
		message := rabbitmq.NewMessageFromDelivery(delivery)
		go s.handleDelivery(message)
	}
}

func (s *subscriber) handleDelivery(message pubsub.Message) {
	s.subscriberHandler(message)
}

func (s *subscriber) setupExchange() {
	for _, exchange := range s.exchanges {
		err := s.getConnection().GetChannel().ExchangeDeclare(
			exchange.Name,
			exchange.Type.String(),
			exchange.IsDurable,
			exchange.IsAutoDeleted,
			exchange.IsInternal,
			exchange.NoWait,
			exchange.Args)
		if err != nil {
			panic(err)
		}
	}
}

func (s *subscriber) bindQueueToExchange() {
	for _, exchange := range s.exchanges {
		err := s.conn.GetChannel().QueueBind(
			s.queue.Name,
			exchange.RoutingKey,
			exchange.Name,
			s.queue.NoWait,
			s.queue.QueueBindArgs)
		if err != nil {
			panic(err)
		}
	}
}

func (s *subscriber) AddExchange(exchange *Exchange) {
	s.exchanges = append(s.exchanges, exchange)
}

func (s *subscriber) SetQueue(queue *Queue) {
	s.queue = queue
}

func (s *subscriber) SetName(name string) {
	s.name = name
}

func (s *subscriber) SetPrefetchQos(qos *PrefetchQos) {
	s.prefetchQos = qos
}

func (s *subscriber) getConnection() connection.Connection {
	if s.conn == nil {
		s.conn = connection.NewConnection(s.connectionOptions)
	}

	return s.conn
}
