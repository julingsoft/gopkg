package rocketmq

import (
	"sync"

	rmqclient "github.com/apache/rocketmq-clients/golang/v5"
	"github.com/apache/rocketmq-clients/golang/v5/credentials"
	"github.com/gogf/gf/v2/errors/gerror"
)

var (
	consumerInstance rmqclient.SimpleConsumer
	consumerOnce     sync.Once
)

func NewConsumer(c Config, opts ...rmqclient.SimpleConsumerOption) (rmqclient.SimpleConsumer, error) {
	var err error
	consumerOnce.Do(func() {
		config := &rmqclient.Config{
			Endpoint:      c.Endpoint,
			NameSpace:     c.NameSpace,
			ConsumerGroup: c.ConsumerGroup,
			Credentials: &credentials.SessionCredentials{
				AccessKey:     c.AccessKey,
				AccessSecret:  c.AccessSecret,
				SecurityToken: c.SecurityToken,
			},
		}

		simpleConsumer, e := rmqclient.NewSimpleConsumer(config, opts...)
		if e != nil {
			err = gerror.Wrap(e, `RocketMQ NewSimpleConsumer`)
			return
		}
		consumerInstance = simpleConsumer
	})

	return consumerInstance, err
}
