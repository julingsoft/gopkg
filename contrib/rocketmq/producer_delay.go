package rocketmq

import (
	"context"
	"time"

	rmqclient "github.com/apache/rocketmq-clients/golang/v5"
	"github.com/gogf/gf/v2/frame/g"
)

func (p *Producer) Delay(ctx context.Context, tag string, keys []string, messages []*rmqclient.Message, deliveryTimestamp time.Time) error {
	for _, message := range messages {
		message.SetTag(tag)
		message.SetKeys(keys...)
		message.SetDelayTimestamp(deliveryTimestamp)

		// send message in async
		resp, err := p.Client.Send(ctx, message)
		if err != nil {
			g.Log().Error(ctx, err)
		}

		for i := 0; i < len(resp); i++ {
			g.Log().Infof(ctx, "send message success. %#v\n", resp[i])
		}
	}

	return nil
}
