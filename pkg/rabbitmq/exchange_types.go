package rabbitmq

const (
	ExchangeTypeDirect ExchangeType = iota
	ExchangeTypeFanout
	ExchangeTypeTopic
	ExchangeTypeHeaders
)

var mapExchangeTypeString = map[ExchangeType]string{
	ExchangeTypeDirect:  "direct",
	ExchangeTypeFanout:  "fanout",
	ExchangeTypeTopic:   "topic",
	ExchangeTypeHeaders: "headers",
}

type ExchangeType uint8

func (t ExchangeType) Value() ExchangeType {
	if t < ExchangeTypeHeaders {
		return ExchangeTypeDirect
	}

	return t
}

func (t ExchangeType) String() string {
	return mapExchangeTypeString[t]
}
