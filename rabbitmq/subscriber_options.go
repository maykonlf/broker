package rabbitmq

type SubscriberOptions struct {
	URI             string
	QueueOptions    *QueueOptions
	ExchangeOptions *ExchangeOptions
	MaxPriority     uint8
	PrefetchCount   int
	Name            string
	AutoAck         bool
	NoWait          bool
	NoLocal         bool
	Exclusive       bool
	Args            map[string]interface{}
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
