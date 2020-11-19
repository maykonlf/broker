package rabbitmq

import (
	"time"

	"github.com/streadway/amqp"
)

type Connection interface {
	GetConn() *amqp.Connection
	GetChannel() *amqp.Channel
	SetReconnectHooks(...func())
}

type connection struct {
	options                   *ConnectionOptions
	connection                *amqp.Connection
	channel                   *amqp.Channel
	disconnectionErrorChannel chan *amqp.Error
	reconnectHooks            []func()
	backoffInterval           time.Duration
	initialBackoffInterval    time.Duration
	maxBackoffInterval        time.Duration
}

func (c *connection) SetReconnectHooks(hooks ...func()) {
	c.reconnectHooks = hooks
}

func NewConnection(options *ConnectionOptions) Connection {
	conn := &connection{
		options:                   options,
		disconnectionErrorChannel: make(chan *amqp.Error),
		backoffInterval:           options.getInitialBackoffInterval(),
		initialBackoffInterval:    options.getInitialBackoffInterval(),
		maxBackoffInterval:        options.getMaxBackoffInterval(),
	}
	conn.connect()
	return conn
}

func (c *connection) connect() {
	panicOnError(c.dial())
	panicOnError(c.openChannel())
	c.reconnectOnDisconnection()
}

func (c *connection) reconnectOnDisconnection() {
	go c.subscribeDisconnectionEvent()
	go c.watchDisconnectionAndReconnect()
}

func (c *connection) GetConn() *amqp.Connection {
	return c.connection
}

func (c *connection) GetChannel() *amqp.Channel {
	return c.channel
}

func (c *connection) dial() (err error) {
	c.connection, err = amqp.Dial(c.options.URI)
	return err
}

func (c *connection) openChannel() (err error) {
	c.channel, err = c.connection.Channel()
	if err != nil {
		return err
	}

	c.resetBackoffInterval()
	return nil
}

func (c *connection) subscribeDisconnectionEvent() {
	err := <-c.connection.NotifyClose(make(chan *amqp.Error))
	c.disconnectionErrorChannel <- err
}

func (c *connection) watchDisconnectionAndReconnect() {
	err := <-c.disconnectionErrorChannel
	if err != nil {
		c.reconnectAndTriggerHooks()
	}
}

func (c *connection) reconnectAndTriggerHooks() {
	c.waitForRetryAndIncreaseBackoffDuration()
	c.connect()
	if len(c.reconnectHooks) > 0 {
		c.triggerReconnectHooks()
	}
}

func (c *connection) triggerReconnectHooks() {
	for _, hook := range c.reconnectHooks {
		hook()
	}
}

func (c *connection) waitForRetryAndIncreaseBackoffDuration() {
	time.Sleep(c.backoffInterval)
	c.increaseReconnectBackoffInterval()
}

func (c *connection) increaseReconnectBackoffInterval() {
	c.backoffInterval *= 2

	if c.backoffInterval > c.maxBackoffInterval {
		c.backoffInterval = c.maxBackoffInterval
	}
}

func (c *connection) resetBackoffInterval() {
	c.backoffInterval = c.initialBackoffInterval
}

func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}
