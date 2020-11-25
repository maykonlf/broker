package subscriber

type Exchange struct {
	Name          string
	Type          ExchangeType
	IsDurable     bool
	IsAutoDeleted bool
	IsInternal    bool
	NoWait        bool
	Args          map[string]interface{}
	RoutingKey    string
}
