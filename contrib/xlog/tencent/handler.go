package tencent

import (
	"context"
	"encoding/json"
	"os"

	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/util/gconv"
	tencentcloud_cls_sdk_go "github.com/tencentcloud/tencentcloud-cls-sdk-go"
)

type Handler struct {
	client *tencentcloud_cls_sdk_go.AsyncProducerClient
	config Config
}

func NewHandler(config Config) *Handler {
	producerConfig := tencentcloud_cls_sdk_go.GetDefaultAsyncProducerClientConfig()
	producerConfig.Endpoint = config.EndPoint
	producerConfig.AccessKeyID = config.SecretId
	producerConfig.AccessKeySecret = config.SecretKey

	client, err := tencentcloud_cls_sdk_go.NewAsyncProducerClient(producerConfig)
	if err != nil {
		panic(err)
	}

	client.Start()

	return &Handler{
		client: client,
		config: config,
	}
}

// Handler is the GoFrame v2 logging handler for Tencent Cloud CLS.
func (h *Handler) Handler(ctx context.Context, input *glog.HandlerInput) {
	// 容器信息
	containerName := "N/A"
	if hostName, err := os.Hostname(); err == nil {
		containerName = hostName
	}

	// 创建日志对象
	logContents := map[string]string{
		"ContainerName": containerName,
		"Level":         input.LevelFormat,
		"Time":          input.Time.Format("2006-01-02 15:04:05.000"),
		"TraceId":       gctx.CtxId(ctx),
	}

	// 尝试解析 JSON 内容
	var contents map[string]any
	if err := json.Unmarshal([]byte(input.ValuesContent()), &contents); err == nil {
		for k, v := range contents {
			logContents[k] = gconv.String(v)
		}
	} else {
		logContents["Content"] = input.ValuesContent()
	}

	// 处理内容
	if input.Content != "" {
		// 尝试解析 JSON 内容
		var contents map[string]any
		if err := json.Unmarshal([]byte(input.Content), &contents); err == nil {
			for k, v := range contents {
				logContents[k] = gconv.String(v)
			}
		} else {
			logContents["Content"] = input.Content
		}
	}

	// 附加堆栈信息
	if input.Stack != "" {
		logContents["Stack"] = input.Stack
	}

	// 发送日志
	log := tencentcloud_cls_sdk_go.NewCLSLog(input.Time.Unix(), logContents)
	if err := h.client.SendLog(h.config.TopicID, log, nil); err != nil {
		input.Logger.Print(ctx, "Tencent CLS SendLog error:", err)
	}

	// 继续执行下一个 Handler (如果有)
	input.Next(ctx)
}

func (h *Handler) Close() {
	if h.client != nil {
		h.client.Close(60000) // 等待最多 60s
	}
}
