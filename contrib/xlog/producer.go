package xlog

import (
	"encoding/json"
	"os"
	"os/signal"
	"time"

	sls "github.com/aliyun/aliyun-log-go-sdk"
	"github.com/aliyun/aliyun-log-go-sdk/producer"
	"github.com/gogf/gf/v2/util/gconv"
)

type Producer struct {
	instance *producer.Producer
	project  string
	logStore string
	topic    string
	source   string
}

func NewProducer(config Config) *Producer {
	cfg := producer.GetDefaultProducerConfig()
	cfg.Endpoint = config.Endpoint
	cfg.CredentialsProvider = sls.NewStaticCredentialsProvider(config.AccessKeyID, config.AccessKeySecret, "")
	instance, err := producer.NewProducer(cfg)
	if err != nil {
		panic(err)
	}

	ch := make(chan os.Signal)
	signal.Notify(ch, os.Kill, os.Interrupt)
	instance.Start()

	return &Producer{
		instance: instance,
		project:  config.ProjectName,
		logStore: config.LogStoreName,
		topic:    config.Topic,
		source:   config.Source,
	}
}

func (p *Producer) Instance() *producer.Producer {
	return p.instance
}

func (p *Producer) Write(data []byte) (n int, err error) {
	// 容器信息
	containerName := "N/A"
	if hostName, err := os.Hostname(); err == nil {
		containerName = hostName
	}

	// 创建日志对象
	logContents := map[string]string{
		"ContainerName": containerName,
	}

	var logItems map[string]interface{}
	if err = json.Unmarshal(data, &logItems); err != nil {
		logContents["Message"] = gconv.String(p)
	} else {
		if logContent, ok := logItems["Content"].(string); ok {
			var contents map[string]interface{}
			if err = json.Unmarshal([]byte(logContent), &contents); err == nil {
				for k, v := range contents {
					logItems[k] = v
				}
				delete(logItems, "Content")
			}
		}
		for k, v := range logItems {
			logContents[k] = gconv.String(v)
		}
	}

	// 发送日志
	log := producer.GenerateLog(uint32(time.Now().Unix()), logContents)
	if err = p.Instance().SendLog(p.project, p.logStore, p.topic, p.source, log); err != nil {
		return 0, err
	}

	return len(data), nil
}

func (p *Producer) Close() {
	if p.Instance() != nil {
		p.Instance().SafeClose()
	}
}
