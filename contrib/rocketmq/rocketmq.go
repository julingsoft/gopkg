package rocketmq

import (
	"sync"

	"github.com/apache/rocketmq-clients/golang/v5"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

var (
	producer           golang.Producer
	producerOnce       sync.Once
	simpleConsumer     golang.SimpleConsumer
	simpleConsumerOnce sync.Once
	pushConsumer       golang.PushConsumer
	pushConsumerOnce   sync.Once
	err                error
)

func GetProducer(cfg Config) golang.Producer {
	producerOnce.Do(func() {
		producer, err = NewProducer(cfg)
		if err != nil {
			panic(err)
		}
		golang.EnableSsl = cfg.EnableSsl
		if err = producer.Start(); err != nil {
			g.Log().Fatal(gctx.New(), "RocketMQ Producer 启动失败", err)
		}
		g.Log().Info(gctx.New(), "RocketMQ Producer 启动成功")
	})
	return producer
}

func GetSimpleConsumer(cfg Config) golang.SimpleConsumer {
	simpleConsumerOnce.Do(func() {
		simpleConsumer, err = NewSimpleConsumer(cfg)
		if err != nil {
			panic(err)
		}
		golang.EnableSsl = cfg.EnableSsl
		if err = simpleConsumer.Start(); err != nil {
			g.Log().Fatal(gctx.New(), "RocketMQ Consumer 启动失败", cfg, err)
		}
		g.Log().Info(gctx.New(), "RocketMQ SimpleConsumer 启动成功")
	})
	return simpleConsumer
}

func GetPushConsumer(cfg Config) golang.PushConsumer {
	pushConsumerOnce.Do(func() {
		pushConsumer, err = NewPushConsumer(cfg)
		if err != nil {
			panic(err)
		}
		golang.EnableSsl = cfg.EnableSsl
		if err = pushConsumer.Start(); err != nil {
			g.Log().Fatal(gctx.New(), "RocketMQ PushConsumer 启动失败", err)
		}
		g.Log().Info(gctx.New(), "RocketMQ PushConsumer 启动成功")
	})
	return pushConsumer
}

func Shutdown() {
	if producer != nil {
		producer.GracefulStop()
		g.Log().Info(gctx.New(), "RocketMQ Producer 已优雅关闭")
	}
	if simpleConsumer != nil {
		simpleConsumer.GracefulStop()
		g.Log().Info(gctx.New(), "RocketMQ SimpleConsumer 已优雅关闭")
	}
	if pushConsumer != nil {
		pushConsumer.GracefulStop()
		g.Log().Info(gctx.New(), "RocketMQ PushConsumer 已优雅关闭")
	}
}
