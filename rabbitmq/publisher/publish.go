package publisher

type PublishOptions struct {
	Exchange    string
	RoutingKey  string
	IsMandatory bool
	IsImmediate bool
}
