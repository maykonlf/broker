package rabbitmq

import "errors"

var (
	ErrMessageIsNotDelivery = errors.New("message is not delivery")
)
