package pubsub

type Subscriber interface {
	// Subscribe start consuming and delivery every consumed message to the given function.
	Subscribe(handler func(message Message))
}
