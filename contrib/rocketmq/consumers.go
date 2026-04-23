package rocketmq

import (
	"fmt"

	"github.com/apache/rocketmq-clients/golang/v5"
	"github.com/apache/rocketmq-clients/golang/v5/credentials"
)

func NewSimpleConsumer(cfg Config) (golang.SimpleConsumer, error) {
	return golang.NewSimpleConsumer(&golang.Config{
		Endpoint:      cfg.Endpoint,
		ConsumerGroup: cfg.ConsumerGroup,
		Credentials: &credentials.SessionCredentials{
			AccessKey:    cfg.AccessKey,
			AccessSecret: cfg.AccessSecret,
		},
	},
		golang.WithSimpleAwaitDuration(awaitDuration),
		golang.WithSimpleSubscriptionExpressions(map[string]*golang.FilterExpression{
			cfg.Topic: golang.SUB_ALL,
		}),
	)
}

func NewPushConsumer(cfg Config) (golang.PushConsumer, error) {
	return golang.NewPushConsumer(&golang.Config{
		Endpoint:      cfg.Endpoint,
		ConsumerGroup: cfg.ConsumerGroup,
		Credentials: &credentials.SessionCredentials{
			AccessKey:    cfg.AccessKey,
			AccessSecret: cfg.AccessSecret,
		},
		NameSpace: cfg.NameSpace,
	},
		golang.WithPushAwaitDuration(awaitDuration),
		golang.WithPushSubscriptionExpressions(map[string]*golang.FilterExpression{
			cfg.Topic: golang.SUB_ALL,
		}),
		golang.WithPushMessageListener(&golang.FuncMessageListener{
			Consume: func(mv *golang.MessageView) golang.ConsumerResult {
				fmt.Println(mv)
				// ack message
				return golang.SUCCESS
			},
		}),
		golang.WithPushConsumptionThreadCount(20),
		golang.WithPushMaxCacheMessageCount(1024),
	)
}
