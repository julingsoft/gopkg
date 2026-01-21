package amqp

import (
	"context"

	rmq "github.com/rabbitmq/rabbitmq-amqp-go-client/pkg/rabbitmqamqp"
)

type Amqp struct {
	config *Config
}

func NewAmqp(config *Config) *Amqp {
	return &Amqp{
		config: config,
	}
}

func (a *Amqp) GetEnv() *rmq.Environment {
	var env *rmq.Environment
	if len(a.config.Address) == 1 {
		env = rmq.NewEnvironment(a.config.Address[0], &rmq.AmqpConnOptions{})
	} else {
		endpoints := make([]rmq.Endpoint, 0, len(a.config.Address))
		for _, addr := range a.config.Address {
			endpoints = append(endpoints, rmq.Endpoint{
				Address: addr,
				Options: &rmq.AmqpConnOptions{},
			})
		}
		env = rmq.NewClusterEnvironment(endpoints)
	}
	return env
}

func (a *Amqp) Connection() (*rmq.AmqpConnection, error) {
	return a.GetEnv().NewConnection(context.Background())
}
