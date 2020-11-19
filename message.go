package pubsub

import (
	"github.com/google/uuid"
	"time"
)

type Message interface {
	Ack() error
	Nack() error
	Reject() error

	ID() uuid.UUID

	CorrelationID() uuid.UUID
	SetCorrelationID(id uuid.UUID) Message

	Headers() map[string]interface{}
	SetHeaders(headers map[string]interface{}) Message
	GetHeader(key string) interface{}
	SetHeader(key string, value interface{}) Message

	ContentType() string
	SetContentType(v string) Message

	ContentEncoding() string
	SetContentEncoding(v string) Message

	Body() []byte
	SetBody(body []byte) Message

	DeliveryMode() uint8
	SetDeliveryModePersistent() Message

	Priority() uint8
	SetPriority(priority uint8) Message

	ReplyTo() string
	SetReplyTo(v string) Message

	Expiration() time.Duration
	SetExpiration(expiration time.Duration) Message

	Type() string
	SetType(v string) Message

	UserID() string
	SetUserID(useID string) Message

	AppID() string
	SetAppID(appID string) Message

	Timestamp() time.Time
	SetTimestamp(timestamp time.Time) Message
}
