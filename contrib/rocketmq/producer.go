package rocketmq

import (
	"github.com/apache/rocketmq-clients/golang/v5"
	"github.com/apache/rocketmq-clients/golang/v5/credentials"
)

func NewProducer(cfg Config) (golang.Producer, error) {
	return golang.NewProducer(&golang.Config{
		Endpoint: cfg.Endpoint,
		Credentials: &credentials.SessionCredentials{
			AccessKey:    cfg.AccessKey,
			AccessSecret: cfg.AccessSecret,
		}},
		golang.WithTopics(cfg.Topic),
	)
}
