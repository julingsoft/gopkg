package rocketmq

import "time"

type Config struct {
	Topic         string `json:"topic"`
	Endpoint      string `json:"endpoint"`
	NameSpace     string `json:"nameSpace"`
	ConsumerGroup string `json:"consumerGroup"`
	AccessKey     string `json:"accessKey"`
	AccessSecret  string `json:"accessSecret"`
	EnableSsl     bool   `json:"enableSsl"`
}

var (
	// maximum waiting time for receive func
	awaitDuration = time.Second * 5
	// maximum number of messages received at one time
	maxMessageNum int32 = 16
	// invisibleDuration should > 20s
	invisibleDuration = time.Second * 20
	// receive messages in a loop
)
