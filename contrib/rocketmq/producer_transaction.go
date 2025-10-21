package rocketmq

import (
	"context"

	rmqclient "github.com/apache/rocketmq-clients/golang/v5"
	"github.com/gogf/gf/v2/frame/g"
)

func (p *Producer) Transaction(ctx context.Context, tag string, keys []string, messages []*rmqclient.Message) error {
	for _, message := range messages {
		message.SetTag(tag)
		message.SetKeys(keys...)

		// send message in async
		transaction := p.Client.BeginTransaction()
		resp, err := p.Client.SendWithTransaction(ctx, message, transaction)
		if err != nil {
			g.Log().Error(ctx, err)
		}

		for i := 0; i < len(resp); i++ {
			g.Log().Infof(ctx, "send message success. %#v\n", resp[i])
		}

		// commit transaction message
		err = transaction.Commit()
		if err != nil {
			g.Log().Error(ctx, err)
		}
	}

	return nil
}
