package subscriber

const (
	/*
		ExchangeTypeDirect represents a direct exchange type.

		This exchange type is the default and will forward package directly to a queue which is the same
		as the given routing key.
	*/
	ExchangeTypeDirect ExchangeType = iota

	/*
		ExchangeTypeFanout represents a fanout exchange type.

		When a message is published to a fanout exchange the message will be forwarded to all binded queues.
	*/
	ExchangeTypeFanout

	/*
		ExchangeTypeTopic represents a topic exchange type.

		When a message is published to this type of exchange the message will only be forwarded to
		binded queues that matches to the publishing routingKey.
	*/
	ExchangeTypeTopic

	/*
		ExchangeTypeHeaders represents a header exchange type.

		When a message is published to this type of exchange the message will only be forwarded to
		binded queues that matches to publishing headers.
	*/
	ExchangeTypeHeaders
)

var mapExchangeTypeString = map[ExchangeType]string{
	ExchangeTypeDirect:  "direct",
	ExchangeTypeFanout:  "fanout",
	ExchangeTypeTopic:   "topic",
	ExchangeTypeHeaders: "headers",
}

// ExchangeType represents the RabbitMQ exchange type.
type ExchangeType uint8

// Value check if exchange is valid, and if valid returns the current value otherwise returns 'direct'
// exchange as default
func (t ExchangeType) Value() ExchangeType {
	if t < ExchangeTypeHeaders {
		return ExchangeTypeDirect
	}

	return t
}

/*
	String returns the ExchangeType as string:
		ExchangeTypeDirect  => "direct"
		ExchangeTypeFanout  => "fanout"
		ExchangeTypeTopic   => "topic"
		ExchangeTypeHeaders => "headers"
*/
func (t ExchangeType) String() string {
	return mapExchangeTypeString[t]
}
