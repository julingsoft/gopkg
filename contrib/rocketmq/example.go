package rocketmq

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/apache/rocketmq-clients/golang/v5"
	"github.com/apache/rocketmq-clients/golang/v5/credentials"
)

func example(cfg *Config) {
	// log to console
	os.Setenv("mq.consoleAppender.enabled", "true")
	golang.ResetLogger()

	// new producer instance
	producer, err := golang.NewProducer(&golang.Config{
		Endpoint:      cfg.Endpoint,
		NameSpace:     cfg.NameSpace,
		ConsumerGroup: cfg.ConsumerGroup,
		Credentials: &credentials.SessionCredentials{
			AccessKey:     cfg.AccessKey,
			AccessSecret:  cfg.SecretKey,
			SecurityToken: cfg.SecurityToken,
		},
	},
		golang.WithTopics(cfg.Topic),
	)
	if err != nil {
		log.Fatal(err)
	}

	// start producer
	err = producer.Start()
	if err != nil {
		log.Fatal(err)
	}

	// gracefule stop producer
	defer producer.GracefulStop()
	for i := 0; i < 10; i++ {
		// new a message
		msg := &golang.Message{
			Topic: cfg.Topic,
			Body:  []byte("this is a message : " + strconv.Itoa(i)),
		}
		// set keys and tag
		msg.SetKeys("a", "b")
		msg.SetTag("ab")
		// send message in async
		producer.SendAsync(context.TODO(), msg, func(ctx context.Context, resp []*golang.SendReceipt, err error) {
			if err != nil {
				log.Fatal(err)
			}
			for i := 0; i < len(resp); i++ {
				fmt.Printf("%#v\n", resp[i])
			}
		})
		// wait a moment
		time.Sleep(time.Second * 1)
	}
}
