package subscriber

// Exchange contains exchange settings.
type Exchange struct {
	// Name is the exchange name.
	Name string

	// Type is the exchange type. See ExchangeType for details.
	Type ExchangeType

	// IsDurable defines if is a durable queue.
	IsDurable bool

	// IsAutoDeleted defines if exchange is auto deleted.
	IsAutoDeleted bool

	// IsInternal defines if is a internal exchange.
	IsInternal bool

	// NoWait defines if is a noWait queue.
	NoWait bool

	// Args exchange creation args.
	Args map[string]interface{}

	// RoutingKey is the routingKey used to bind exchange to consumer queue (for topic exchanges).
	RoutingKey string
}
