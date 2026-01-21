package xlog

import (
	sls "github.com/aliyun/aliyun-log-go-sdk"
)

type Consumer struct {
	instance sls.ClientInterface
	project  string
	logStore string
	topic    string
	source   string
}

func ClientInstance(config Config) sls.ClientInterface {
	credentialsProvider := sls.NewStaticCredentialsProvider(config.AccessKeyID, config.AccessKeySecret, "")
	return sls.CreateNormalInterfaceV2(config.Endpoint, credentialsProvider)
}

func NewConsumer(config Config) *Consumer {
	instance := ClientInstance(config)
	return &Consumer{
		instance: instance,
		project:  config.ProjectName,
		logStore: config.LogStoreName,
		topic:    config.Topic,
		source:   config.Source,
	}
}

func (c *Consumer) GetLogs(req *sls.GetLogRequest) (*sls.GetLogsResponse, error) {
	return c.instance.GetLogsV2(c.project, c.logStore, req)
}
