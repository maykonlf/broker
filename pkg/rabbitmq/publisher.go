package rabbitmq

import (
	"sync"
)

type Publisher interface {
	Publish(options *PublishOptions, message *Message) error
}

func NewPublisher(options *PublisherOptions) Publisher {
	conn := NewConnection(&ConnectionOptions{
		URI: options.URI,
	})
	return &publisher{
		options: options,
		conn:    conn,
		mutex:   &sync.Mutex{},
	}
}

type publisher struct {
	options *PublisherOptions
	conn    Connection
	mutex   *sync.Mutex
}

func (p *publisher) Publish(options *PublishOptions, message *Message) error {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.conn.GetChannel().Publish(
		options.Exchange,
		options.RoutingKey,
		options.IsMandatory,
		options.IsImmediate,
		message.getPublishing())
}
