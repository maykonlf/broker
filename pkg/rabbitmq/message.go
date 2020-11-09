package rabbitmq

import (
	"fmt"
	"github.com/google/uuid"
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
	expiration      string
	messageType     string
	userID          string
	appID           string
	timestamp       time.Time
	delivery        amqp.Delivery
}

func NewMessage() *Message {
	return &Message{id: uuid.New()}
}

func newMessageFromDelivery(delivery amqp.Delivery) *Message {
	return &Message{
		id:              parseUUIDOrGetDefault(delivery.MessageId),
		correlationID:   parseUUIDOrGetDefault(delivery.CorrelationId),
		headers:         delivery.Headers,
		contentType:     delivery.ContentType,
		contentEncoding: delivery.ContentEncoding,
		body:            delivery.Body,
		deliveryMode:    delivery.DeliveryMode,
		priority:        delivery.Priority,
		replyTo:         delivery.ReplyTo,
		expiration:      delivery.Expiration,
		messageType:     delivery.Type,
		userID:          delivery.UserId,
		appID:           delivery.AppId,
		timestamp:       delivery.Timestamp,
		delivery:        delivery,
	}
}

func parseUUIDOrGetDefault(s string) uuid.UUID {
	v, err := uuid.Parse(s)
	if err != nil {
		return uuid.Nil
	}

	return v
}

func (m *Message) GetID() uuid.UUID {
	return m.id
}

func (m *Message) SetCorrelationID(id uuid.UUID) *Message {
	m.correlationID = id
	return m
}

func (m *Message) GetCorrelationID() uuid.UUID {
	return m.id
}

func (m *Message) SetHeader(key string, value interface{}) *Message {
	m.headers[key] = value
	return m
}

func (m *Message) GetHeader(key string) interface{} {
	return m.headers[key]
}

func (m *Message) GetHeaders() map[string]interface{} {
	return m.headers
}

func (m *Message) SetContentType(v string) *Message {
	m.contentType = v
	return m
}

func (m *Message) GetContentType() string {
	return m.contentType
}

func (m *Message) SetContentEncoding(v string) *Message {
	m.contentEncoding = v
	return m
}

func (m *Message) GetContentEncoding() string {
	return m.contentEncoding
}

func (m *Message) SetBody(body []byte) *Message {
	m.body = body
	return m
}

func (m *Message) GetBody() []byte {
	return m.body
}

func (m *Message) SetDeliveryModePersistent() *Message {
	m.deliveryMode = 2
	return m
}

func (m *Message) GetDeliveryMode() uint8 {
	return m.deliveryMode
}

func (m *Message) SetPriority(priority uint8) *Message {
	m.priority = priority
	return m
}

func (m *Message) GetPriority() uint8 {
	return m.priority
}

func (m *Message) SetReplyTo(v string) *Message {
	m.replyTo = v
	return m
}

func (m *Message) GetReplyTo() string {
	return m.replyTo
}

func (m *Message) SetExpiration(expiration time.Duration) *Message {
	if expiration > 0 {
		m.expiration = fmt.Sprintf("%d", expiration.Milliseconds())
	}

	return m
}

func (m *Message) GetExpiration() time.Duration {
	if m.expiration == "" {
		return 0
	}

	milliseconds, err := strconv.ParseInt(m.expiration, 10, 64)
	if err != nil {
		return 0
	}

	return time.Duration(milliseconds) * time.Millisecond
}

func (m *Message) GetExpirationString() string {
	return m.expiration
}

func (m *Message) SetType(v string) *Message {
	m.messageType = v
	return m
}

func (m *Message) GetType() string {
	return m.messageType
}

func (m *Message) SetUserID(useID string) *Message {
	m.userID = useID
	return m
}

func (m *Message) GetUserID() string {
	return m.userID
}

func (m *Message) SetAppID(appID string) *Message {
	m.appID = appID
	return m
}

func (m *Message) GetAppID() string {
	return m.appID
}

func (m *Message) SetTimestamp(timestamp time.Time) *Message {
	m.timestamp = timestamp
	return m
}

func (m *Message) GetTimestamp() time.Time {
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

func (m *Message) getPublishing() amqp.Publishing {
	return amqp.Publishing{
		Headers:         m.GetHeaders(),
		ContentType:     m.GetContentType(),
		ContentEncoding: m.GetContentEncoding(),
		DeliveryMode:    m.GetDeliveryMode(),
		Priority:        m.GetPriority(),
		CorrelationId:   m.GetCorrelationID().String(),
		ReplyTo:         m.GetReplyTo(),
		Expiration:      m.GetExpirationString(),
		MessageId:       m.GetID().String(),
		Timestamp:       m.GetTimestamp(),
		Type:            m.GetType(),
		UserId:          m.GetUserID(),
		AppId:           m.GetAppID(),
		Body:            m.GetBody(),
	}
}
