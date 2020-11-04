package rabbitmq

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/streadway/amqp"
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
}

func NewMessage() *Message {
	return &Message{id: uuid.New()}
}

func (m *Message) GetID() uuid.UUID {
	return m.id
}

func (m *Message) CorrelationID(id uuid.UUID) *Message {
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

func (m *Message) ContentType(v string) *Message {
	m.contentType = v
	return m
}

func (m *Message) GetContentType() string {
	return m.contentType
}

func (m *Message) ContentEncoding(v string) *Message {
	m.contentEncoding = v
	return m
}

func (m *Message) GetContentEncoding() string {
	return m.contentEncoding
}

func (m *Message) Body(body []byte) *Message {
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

func (m *Message) Priority(priority uint8) *Message {
	m.priority = priority
	return m
}

func (m *Message) GetPriority() uint8 {
	return m.priority
}

func (m *Message) ReplyTo(v string) *Message {
	m.replyTo = v
	return m
}

func (m *Message) GetReplyTo() string {
	return m.replyTo
}

func (m *Message) Expiration(expiration time.Duration) *Message {
	m.expiration = expiration
	return m
}

func (m *Message) GetExpiration() time.Duration {
	return m.expiration
}

func (m *Message) Type(v string) *Message {
	m.messageType = v
	return m
}

func (m *Message) GetType() string {
	return m.messageType
}

func (m *Message) UserID(useID string) *Message {
	m.userID = useID
	return m
}

func (m *Message) GetUserID() string {
	return m.userID
}

func (m *Message) AppID(appID string) *Message {
	m.appID = appID
	return m
}

func (m *Message) GetAppID() string {
	return m.appID
}

func (m *Message) Timestamp(timestamp time.Time) *Message {
	m.timestamp = timestamp
	return m
}

func (m *Message) GetTimestamp() time.Time {
	return m.timestamp
}

func (m *Message) getPublishing() amqp.Publishing {
	return amqp.Publishing{
		Headers:         m.headers,
		ContentType:     m.GetContentType(),
		ContentEncoding: m.GetContentEncoding(),
		DeliveryMode:    m.GetDeliveryMode(),
		Priority:        m.GetPriority(),
		CorrelationId:   m.GetCorrelationID().String(),
		ReplyTo:         m.GetReplyTo(),
		Expiration:      fmt.Sprintf("%d", m.GetExpiration().Milliseconds()),
		MessageId:       m.GetID().String(),
		Timestamp:       m.GetTimestamp(),
		Type:            m.GetType(),
		UserId:          m.GetUserID(),
		AppId:           m.GetAppID(),
		Body:            m.GetBody(),
	}
}
