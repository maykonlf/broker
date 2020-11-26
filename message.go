package pubsub

import (
	"github.com/google/uuid"
	"time"
)

// Message represents the consumed/published message
type Message interface {
	/*
		Ack delegates an acknowledgement that the client or server has finished work with the delivered message.

		All consumed messages should be acknowledged. If the consumer has autoAck enabled then the server will
		automatically ack each message and this method should not be called.

		An error will indicate that the acknowledge could not be delivered to the
		channel it was sent from.

		Either Ack, Reject or Nack must be called for every delivered message that is not automatically acknowledged.
	*/
	Ack() error

	/*
		Nack negatively acknowledge a delivered message then request the server to deliver this message to another
		consumer.

		This method must not be used to select or requeue message the client wished not to handle, rather it is to
		inform that the consumer is incapable of handling the message this time.

		Either Ack, Reject or Nack must be called for every delivered message that is not automatically acknowledged.
	*/
	Nack() error

	/*
		Reject negatively acknowledge a delivered message then request the server to drop this message.

		Once the message is rejected it cannot be recovered and wont be delivered to another consumer and should be
		used when the consumer cannot handle the message now or ever. Usually because the message is not valid.

		Either Ack, Reject or Nack must be called for every delivered message that is not automatically acknowledged.
	*/
	Reject() error

	// ID returns the unique message ID.
	ID() uuid.UUID

	/*
		CorrelationID returns the message correlation id when is set.

		Useful when you need to track which message, request or event generated the given message.
	*/
	CorrelationID() uuid.UUID

	/*
		SetCorrelationID set a correlation id for message.

		Useful when you need to track which message, request or event generated the given message.
	*/
	SetCorrelationID(id uuid.UUID) Message

	/*
		Headers returns the message headers.

		Useful when you need to send/receive additional information of with a message, but dont want to put it on
		the message body.
	*/
	Headers() map[string]interface{}

	/*
		SetHeaders set message headers.

		Useful when you need to send/receive additional information of with a message, but dont want to put it on
		the message body.
	*/
	SetHeaders(headers map[string]interface{}) Message

	/*
		GetHeader returns the value of a specific message header.

		Useful when you need to send/receive additional information of with a message, but dont want to put it on
		the message body.
	*/
	GetHeader(key string) interface{}

	/*
		SetHeader put some value to a specific message header.

		Useful when you need to send/receive additional information of with a message, but dont want to put it on
		the message body.
	*/
	SetHeader(key string, value interface{}) Message

	/*
		ContentType returns the message content type (returns "plain/text" if not set).

		Useful when you need to tell the consumer which content type is being delivered, so the consumer can handle
		the message properly.
	*/
	ContentType() string

	/*
		SetContentType set message content type.

		Useful when you need to tell the consumer which content type is being delivered, so the consumer can handle
		the message properly.
	*/
	SetContentType(v string) Message

	/*
		ContentEncoding returns the message content encoding.

		Useful when you need to tell the consumer which encoding should be used to read the delivered message.
	*/
	ContentEncoding() string

	/*
		SetContentEncoding set the message content encoding.

		Useful when you need to tell the consumer which encoding should be used to read the delivered message.
	*/
	SetContentEncoding(v string) Message

	/*
		Body returns the message content (body).
	*/
	Body() []byte

	/*
		SetBody put a content (body) on message.
	*/
	SetBody(body []byte) Message

	/*
		DeliveryMode returns message delivery mode:
		1 - non-persistent (default)
		2 - persistent

		A persistent delivery means that if the broker server fails the message will not be lost (for rabbitmq only) but
		it will make the pubsub operations for this slower.
	*/
	DeliveryMode() uint8

	/*
		SetDeliveryModePersistent set message delivery mode to persistent.

		A persistent delivery means that if the broker server fails the message will not be lost (for rabbitmq only) but
		it will make the pubsub operations for this slower.
	*/
	SetDeliveryModePersistent() Message

	/*
		Priority returns message configured priority (default: 0)
	*/
	Priority() uint8

	/*
		SetPriority defines message priority.

		Messages without a priority are treated as if their priority were 0. If published on a non-priority queue
		this value is ignored. Is recommended to use values between 1 and 10.
	*/
	SetPriority(priority uint8) Message

	/*
		ReplyTo returns the address to reply to (application usage only).
	*/
	ReplyTo() string

	/*
		SetReplyTo set the address the subscriber should reply to (application usage only).
	*/
	SetReplyTo(v string) Message

	/*
		Expiration returns message expiration time.
	*/
	Expiration() time.Duration

	/*
		SetExpiration set message expiration time in duration.

		If set the message will be automatically dropped when time in queue reaches the expiration time.
	*/
	SetExpiration(expiration time.Duration) Message

	/*
		Type returns message type (application usage only).
	*/
	Type() string

	/*
		SetType defines message type (application usage only).
	*/
	SetType(v string) Message

	/*
		UserID returns message user id (application usage only).
	*/
	UserID() string

	/*
		SetUserID defines message user id (application usage  only).
	*/
	SetUserID(useID string) Message

	/*
		AppID returns message app id (application usage only).
	*/
	AppID() string

	/*
		SetAppID set message app id (application usage only).
	*/
	SetAppID(appID string) Message

	/*
		Timestamp returns message timestamp (application usage only).
	*/
	Timestamp() time.Time

	/*
		SetTimestamp set message timestamp (application usage only).
	*/
	SetTimestamp(timestamp time.Time) Message
}
