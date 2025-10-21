package rocketmq

import (
	"context"

	rmqclient "github.com/apache/rocketmq-clients/golang/v5"
	"github.com/gogf/gf/v2/frame/g"
)

func (p *Producer) Normal(ctx context.Context, tag string, keys []string, messages []*rmqclient.Message) error {
	for _, message := range messages {
		message.SetTag(tag)
		message.SetKeys(keys...)

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
