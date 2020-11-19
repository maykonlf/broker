package rabbitmq

import (
	"github.com/google/uuid"
	"github.com/maykonlf/pubsub"
	"github.com/streadway/amqp"
	"strconv"
	"time"
)

type Message struct {
	id              uuid.UUID
	correlationID   uuid.UUID
	headers         map[string]interface{}
	contentType     string
	contentEncoding string
	body            []byte
	deliveryMode    uint8
	priority        uint8
	replyTo         string
	expiration      time.Duration
	messageType     string
	userID          string
	appID           string
	timestamp       time.Time
	delivery        amqp.Delivery
}

func NewMessage() pubsub.Message {
	return &Message{id: uuid.New()}
}

func newMessageFromDelivery(delivery amqp.Delivery) pubsub.Message {
	return &Message{
		id:              uuid.MustParse(delivery.MessageId),
		correlationID:   uuid.MustParse(delivery.CorrelationId),
		headers:         delivery.Headers,
		contentType:     delivery.ContentType,
		contentEncoding: delivery.ContentEncoding,
		body:            delivery.Body,
		deliveryMode:    delivery.DeliveryMode,
		priority:        delivery.Priority,
		replyTo:         delivery.ReplyTo,
		expiration:      parseDurationStringToTimeDuration(delivery.Expiration),
		messageType:     delivery.Type,
		userID:          delivery.UserId,
		appID:           delivery.AppId,
		timestamp:       delivery.Timestamp,
		delivery:        delivery,
	}
}

func parseDurationStringToTimeDuration(s string) time.Duration {
	milliseconds, _ := strconv.ParseInt(s, 10, 64)
	return time.Duration(milliseconds) * time.Millisecond
}

func (m *Message) ID() uuid.UUID {
	return m.id
}

func (m *Message) SetCorrelationID(id uuid.UUID) pubsub.Message {
	m.correlationID = id
	return m
}

func (m *Message) CorrelationID() uuid.UUID {
	return m.id
}

func (m *Message) SetHeader(key string, value interface{}) pubsub.Message {
	m.headers[key] = value
	return m
}

func (m *Message) GetHeader(key string) interface{} {
	return m.headers[key]
}

func (m *Message) Headers() map[string]interface{} {
	return m.headers
}

func (m *Message) SetHeaders(headers map[string]interface{}) pubsub.Message {
	m.headers = headers
	return m
}

func (m *Message) SetContentType(v string) pubsub.Message {
	m.contentType = v
	return m
}

func (m *Message) ContentType() string {
	return m.contentType
}

func (m *Message) SetContentEncoding(v string) pubsub.Message {
	m.contentEncoding = v
	return m
}

func (m *Message) ContentEncoding() string {
	return m.contentEncoding
}

func (m *Message) SetBody(body []byte) pubsub.Message {
	m.body = body
	return m
}

func (m *Message) Body() []byte {
	return m.body
}

func (m *Message) SetDeliveryModePersistent() pubsub.Message {
	m.deliveryMode = 2
	return m
}

func (m *Message) DeliveryMode() uint8 {
	return m.deliveryMode
}

func (m *Message) SetPriority(priority uint8) pubsub.Message {
	m.priority = priority
	return m
}

func (m *Message) Priority() uint8 {
	return m.priority
}

func (m *Message) SetReplyTo(v string) pubsub.Message {
	m.replyTo = v
	return m
}

func (m *Message) ReplyTo() string {
	return m.replyTo
}

func (m *Message) SetExpiration(expiration time.Duration) pubsub.Message {
	m.expiration = expiration
	return m
}

func (m *Message) Expiration() time.Duration {
	return m.expiration
}

func (m *Message) SetType(v string) pubsub.Message {
	m.messageType = v
	return m
}

func (m *Message) Type() string {
	return m.messageType
}

func (m *Message) SetUserID(useID string) pubsub.Message {
	m.userID = useID
	return m
}

func (m *Message) UserID() string {
	return m.userID
}

func (m *Message) SetAppID(appID string) pubsub.Message {
	m.appID = appID
	return m
}

func (m *Message) AppID() string {
	return m.appID
}

func (m *Message) SetTimestamp(timestamp time.Time) pubsub.Message {
	m.timestamp = timestamp
	return m
}

func (m *Message) Timestamp() time.Time {
	return m.timestamp
}

func (m *Message) Ack() error {
	return m.delivery.Ack(false)
}

func (m *Message) Nack() error {
	return m.delivery.Nack(false, true)
}

func (m *Message) Reject() error {
	return m.delivery.Reject(false)
}
