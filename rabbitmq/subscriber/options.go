package subscriber

// Option is a subscriber option used to customize the consumer.
type Option func(s Subscriber)

// WithName set consumer name.
func WithName(name string) Option {
	return func(s Subscriber) {
		s.SetName(name)
	}
}

// WithPrefetch set prefetch count, and size.
func WithPrefetch(count int) Option {
	return func(s Subscriber) {
		s.SetPrefetchQos(&PrefetchQos{
			Count:    count,
			Size:     0,
			IsGlobal: false,
		})
	}
}

// WithDurableQueue defines a named durable queue for consumer.
func WithDurableQueue(name string) Option {
	return func(s Subscriber) {
		s.SetQueue(&Queue{
			Name:    name,
			Durable: true,
		})
	}
}

// WithDurablePriorityQueue defines a named durable queue for consumer with priority.
func WithDurablePriorityQueue(name string, maxPriority uint8) Option {
	return func(s Subscriber) {
		s.SetQueue(&Queue{
			Name:        name,
			Durable:     true,
			MaxPriority: maxPriority,
		})
	}
}

// WithDurableFanoutExchange declare a fanout exchange an bind it to the consumer queue.
func WithDurableFanoutExchange(name string) Option {
	return func(s Subscriber) {
		s.AddExchange(&Exchange{
			Name:      name,
			Type:      ExchangeTypeFanout,
			IsDurable: true,
		})
	}
}

// WithDurableTopicExchange declare a topic exchange and bind it to the consumer queue.
func WithDurableTopicExchange(name, routingKey string) Option {
	return func(s Subscriber) {
		s.AddExchange(&Exchange{
			Name:       name,
			Type:       ExchangeTypeTopic,
			IsDurable:  true,
			RoutingKey: routingKey,
		})
	}
}
