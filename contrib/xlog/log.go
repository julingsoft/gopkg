package xlog

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
	"time"

	sls "github.com/aliyun/aliyun-log-go-sdk"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogo/protobuf/proto"
)

type SLSWriter struct {
	debug    bool
	logger   *glog.Logger
	client   sls.ClientInterface
	project  string
	logStore string
	logChan  chan *sls.LogGroup
	wg       sync.WaitGroup
	mu       sync.Mutex // 用于保护 logCh
}

type LogData struct {
	Method   string
	Url      string
	Headers  map[string]string
	Request  string
	Response interface{}
	Status   int
	Duration int64
	LogTime  time.Time
}

func NewSLSWriter(cfg Config) *SLSWriter {
	credentialsProvider := sls.NewStaticCredentialsProvider(cfg.AccessKeyID, cfg.AccessKeySecret, "")
	client := sls.CreateNormalInterfaceV2(cfg.Endpoint, credentialsProvider)

	writer := &SLSWriter{
		debug:    cfg.Debug,
		client:   client,
		project:  cfg.ProjectName,
		logStore: cfg.LogStoreName,
		logChan:  make(chan *sls.LogGroup, 50000), // 设置缓冲通道
	}

	// 启动 goroutine 来处理日志
	go writer.processLogs()

	return writer
}

func (w *SLSWriter) processLogs() {
	var batchLog []*sls.Log
	var batchSize = 20

	if w.debug {
		batchSize = 1
	}

	// 一次提交20条日志
	for logData := range w.logChan {
		batchLog = append(batchLog, logData.Logs...)

		// 如果日志数量达到20条，则提交
		if len(batchLog) >= batchSize {
			batchLogGroup := &sls.LogGroup{
				Logs: batchLog,
			}

			w.wg.Add(1)
			go w.Send(batchLogGroup)

			// Reset batch after sending
			batchLog = make([]*sls.Log, 0, batchSize)
		}
	}

	// 程序退出时， 如果还有未提交的日志，则提交
	if len(batchLog) > 0 {
		batchLogGroup := &sls.LogGroup{
			Logs: batchLog,
		}

		w.wg.Add(1)
		go w.Send(batchLogGroup)
	}
}

// Send 发送日志
func (w *SLSWriter) Send(batchLogGroup *sls.LogGroup) {
	defer w.wg.Done()
	if err := w.client.PutLogs(w.project, w.logStore, batchLogGroup); err != nil {
		// resend
		if err := w.client.PutLogs(w.project, w.logStore, batchLogGroup); err != nil {
			fmt.Printf("Failed to resend log batch: %v\n", err)
		}
	}
}

func (w *SLSWriter) Write(p []byte) (n int, err error) {
	log := sls.Log{
		Time:     proto.Uint32(uint32(time.Now().Unix())),
		Contents: []*sls.LogContent{},
	}

	// 容器信息
	containerName := "N/A"
	if hostName, err := os.Hostname(); err == nil {
		containerName = hostName
	}
	log.Contents = append(log.Contents, &sls.LogContent{
		Key:   proto.String("ContainerName"),
		Value: proto.String(gconv.String(containerName)),
	})

	var logItems map[string]interface{}
	if err = json.Unmarshal(p, &logItems); err != nil {
		log.Contents = append(log.Contents, &sls.LogContent{
			Key:   proto.String("message"),
			Value: proto.String(gconv.String(p)),
		})
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
			log.Contents = append(log.Contents, &sls.LogContent{
				Key:   proto.String(k),
				Value: proto.String(gconv.String(v)),
			})
		}
	}

	logGroup := &sls.LogGroup{
		Logs: []*sls.Log{&log},
	}

	// 将日志数据发送到通道
	w.mu.Lock()
	defer w.mu.Unlock()
	w.logChan <- logGroup

	return len(p), nil
}

func (w *SLSWriter) Close() {
	close(w.logChan) // 关闭通道，通知 goroutine 停止
	w.wg.Wait()      // 等待所有 goroutine 完成
}
