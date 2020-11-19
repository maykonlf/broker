package rabbitmq

import (
	"fmt"
	"github.com/maykonlf/pubsub"
	"github.com/streadway/amqp"
	"sync"
	"time"
)

func NewPublisher(options *PublisherOptions) pubsub.Publisher {
	conn := NewConnection(options.ConnectionOptions)
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

func (p *publisher) Publish(m pubsub.Message, topic string, routingKey ...string) error {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	return p.conn.GetChannel().Publish(
		topic,
		p.extractFirstRoutingKeyOrDefault(routingKey),
		false,
		false,
		amqp.Publishing{
			Headers:         m.Headers(),
			ContentType:     m.ContentType(),
			ContentEncoding: m.ContentEncoding(),
			DeliveryMode:    m.DeliveryMode(),
			Priority:        m.Priority(),
			CorrelationId:   m.CorrelationID().String(),
			ReplyTo:         m.ReplyTo(),
			Expiration:      p.getExpirationStringInMillisecondsOrDefault(m.Expiration()),
			MessageId:       m.ID().String(),
			Timestamp:       m.Timestamp(),
			Type:            m.Type(),
			UserId:          m.UserID(),
			AppId:           m.AppID(),
			Body:            m.Body(),
		})
}

func (p *publisher) getExpirationStringInMillisecondsOrDefault(expiration time.Duration) string {
	if expiration > 0 {
		return fmt.Sprintf("%d", expiration.Milliseconds())
	}

	return ""
}

func (p *publisher) extractFirstRoutingKeyOrDefault(key []string) string {
	if len(key) > 0 {
		return key[0]
	}

	return ""
}
