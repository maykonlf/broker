package pubsub

type Publisher interface {
	// Publish publishes a message to a topic (and optionally to a routing key).
	Publish(message Message, topic string, routingKey ...string) error
}
