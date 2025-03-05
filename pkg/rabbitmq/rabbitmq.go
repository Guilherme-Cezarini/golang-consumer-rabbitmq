package rabbitmq

import (
	"github.com/streadway/amqp"
)

func Connect(uri string) (*amqp.Connection, error) {
	return amqp.Dial(uri)
}