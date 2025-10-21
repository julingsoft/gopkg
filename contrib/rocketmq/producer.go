package rocketmq

import (
	"context"
	"sync"

	rmqclient "github.com/apache/rocketmq-clients/golang/v5"
	"github.com/apache/rocketmq-clients/golang/v5/credentials"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

type Producer struct {
	Client rmqclient.Producer
}

var (
	producerInstance *Producer
	producerOnce     sync.Once
)

// GetProducer 获取单例 Producer 实例
func GetProducer(c Config, topic string, opts ...rmqclient.ProducerOption) (*Producer, error) {
	var err error
	producerOnce.Do(func() {
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

		opts = append(opts, rmqclient.WithTopics(topic))
		producer, e := rmqclient.NewProducer(config, opts...)
		if e != nil {
			err = gerror.Wrap(e, `RocketMQ NewProducer`)
			return
		}

		// start producer
		if e = producer.Start(); e != nil {
			err = gerror.Wrap(e, `RocketMQ Producer Start`)
			return
		}

		// graceful stop producer
		defer func(ctx context.Context, producer rmqclient.Producer) {
			if err := producer.GracefulStop(); err != nil {
				g.Log().Error(ctx, `RocketMQ Producer GracefulStop`, err)
			}
		}(context.Background(), producer)

		producerInstance = &Producer{Client: producer}
	})

	return producerInstance, err
}
