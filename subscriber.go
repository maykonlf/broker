package pubsub

type Subscriber interface {
	Subscribe(handler func(message Message))
}
