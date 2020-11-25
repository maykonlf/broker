package subscriber

type Option func(s Subscriber)

func WithName(name string) Option {
	return func(s Subscriber) {
		s.SetName(name)
	}
}

func WithPrefetchQos(count, size int, isGlobal bool) Option {
	return func(s Subscriber) {
		s.SetPrefetchQos(&PrefetchQos{
			Count:    count,
			Size:     size,
			IsGlobal: isGlobal,
		})
	}
}

func WithDurableQueue(name string) Option {
	return func(s Subscriber) {
		s.SetQueue(&Queue{
			Name:    name,
			Durable: true,
		})
	}
}

func WithDurablePriorityQueue(name string, maxPriority uint8) Option {
	return func(s Subscriber) {
		s.SetQueue(&Queue{
			Name:        name,
			Durable:     true,
			MaxPriority: maxPriority,
		})
	}
}

func WithDurableFanoutExchange(name string) Option {
	return func(s Subscriber) {
		s.AddExchange(&Exchange{
			Name:      name,
			Type:      ExchangeTypeFanout,
			IsDurable: true,
		})
	}
}

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
