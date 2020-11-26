package publisher

import (
	"fmt"
	"github.com/maykonlf/pubsub"
	"github.com/maykonlf/pubsub/rabbitmq/connection"
	"github.com/streadway/amqp"
	"sync"
	"time"
)

type Publisher interface {
	pubsub.Publisher
}

// NewPublisher returns a new RabbitMQ publisher.
func NewPublisher(uri string, options ...Option) Publisher {
	publisher := &publisher{
		mutex:             &sync.Mutex{},
		connectionOptions: &connection.Options{URI: uri},
	}

	for _, optionFunction := range options {
		optionFunction(publisher)
	}

	return publisher
}

type publisher struct {
	mutex             *sync.Mutex
	connectionOptions *connection.Options
	conn              connection.Connection
}

// Publish publishes a message to a topic (and optionally to a routing key).
func (p *publisher) Publish(m pubsub.Message, topic string, routingKey ...string) error {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	return p.getConnection().GetChannel().Publish(
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

func (p *publisher) getConnection() connection.Connection {
	if p.conn == nil {
		p.conn = connection.NewConnection(p.connectionOptions)
	}

	return p.conn
}
