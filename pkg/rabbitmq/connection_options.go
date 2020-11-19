package rabbitmq

import "time"

const (
	defaultInitialBackoffInterval = 1 * time.Second
	defaultMaxBackoffInterval     = 1 * time.Minute
)

type ConnectionOptions struct {
	URI                    string
	InitialBackoffInterval time.Duration
	MaxBackoffInterval     time.Duration
}

func (c *ConnectionOptions) getInitialBackoffInterval() time.Duration {
	if c.InitialBackoffInterval == 0 {
		return defaultInitialBackoffInterval
	}

	return c.InitialBackoffInterval
}

func (c *ConnectionOptions) getMaxBackoffInterval() time.Duration {
	if c.MaxBackoffInterval == 0 {
		return defaultMaxBackoffInterval
	}

	return c.MaxBackoffInterval
}
