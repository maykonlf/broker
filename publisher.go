package pubsub

type Publisher interface {
	Publish(message Message, topic string, routingKey ...string) error
}
