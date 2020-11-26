package connection

import "time"

const (
	defaultInitialBackoffInterval = 1 * time.Second
	defaultMaxBackoffInterval     = 1 * time.Minute
)

// Options contains the RabbitMQ connection and reconnection params
type Options struct {
	/*
		URI is the RabbitMQ connection string with format: "amqp://username:password@localhost:5672"
	*/
	URI string

	/*
		InitialBackoffInterval is the initial delay between reconnections.

		Each time the client reconnects, it waits for last(or initial) backoff interval * 2 until and
		this value increases (x2) on each retry until reach the max configured interval.

		Once the client successful reconnects the backoff interval is set back to the initial value.
	*/
	InitialBackoffInterval time.Duration

	/*
		MaxBackoffInterval is the max interval between reconnections.

		Each time the client reconnects, it waits for last(or initial) backoff interval * 2 until and
		this value increases (x2) on each retry until reach the max configured interval.

		Once the client successful reconnects the backoff interval is set back to the initial value.
	*/
	MaxBackoffInterval time.Duration
}

func (c *Options) getInitialBackoffInterval() time.Duration {
	if c.InitialBackoffInterval == 0 {
		return defaultInitialBackoffInterval
	}

	return c.InitialBackoffInterval
}

func (c *Options) getMaxBackoffInterval() time.Duration {
	if c.MaxBackoffInterval == 0 {
		return defaultMaxBackoffInterval
	}

	return c.MaxBackoffInterval
}
