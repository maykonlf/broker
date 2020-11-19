package rabbitmq

type SubscriberOptions struct {
	ConnectionOptions *ConnectionOptions
	QueueOptions      *QueueOptions
	ExchangeOptions   *ExchangeOptions
	PrefetchCount     int
	Name              string
	AutoAck           bool
	NoWait            bool
	NoLocal           bool
	Exclusive         bool
	Args              map[string]interface{}
}

type ExchangeOptions struct {
	Name          string
	Type          ExchangeType
	IsDurable     bool
	IsAutoDeleted bool
	IsInternal    bool
	NoWait        bool
	Args          map[string]interface{}
}
