package subscriber

import (
	"github.com/streadway/amqp"
)

type Queue struct {
	Name          string
	Durable       bool
	AutoDelete    bool
	Exclusive     bool
	NoWait        bool
	MaxPriority   uint8
	RoutingKey    string
	QueueBindArgs map[string]interface{}
}

func (o *Queue) GetArgs() amqp.Table {
	args := map[string]interface{}{}
	if o.MaxPriority > 0 {
		args["x-max-priority"] = o.MaxPriority
	}

	return args
}
