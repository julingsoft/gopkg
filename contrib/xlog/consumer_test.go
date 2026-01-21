package xlog

import (
	"reflect"
	"testing"
	"time"

	sls "github.com/aliyun/aliyun-log-go-sdk"
)

func TestConsumer_GetLogs(t *testing.T) {
	type fields struct {
		instance sls.ClientInterface
		project  string
		logStore string
		topic    string
		source   string
	}
	type args struct {
		req *sls.GetLogRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *sls.GetLogsResponse
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "case1",
			fields: fields{
				instance: ClientInstance(Config{
					Endpoint:        "cn-shanghai.log.aliyuncs.com",
					AccessKeyID:     "",
					AccessKeySecret: "",
					ProjectName:     "logs",
					LogStoreName:    "log",
					Topic:           "",
					Source:          "",
				}),
				project:  "logs",
				logStore: "log",
				topic:    "",
				source:   "",
			},
			args: args{
				req: &sls.GetLogRequest{
					From:   time.Now().Unix() - 3600, // 最近1小时
					To:     time.Now().Unix(),
					Lines:  1,
					Offset: 0,
					Query:  "prices/query",
				},
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Consumer{
				instance: tt.fields.instance,
				project:  tt.fields.project,
				logStore: tt.fields.logStore,
				topic:    tt.fields.topic,
				source:   tt.fields.source,
			}
			got, err := c.GetLogs(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetLogs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetLogs() got = %v, want %v", got, tt.want)
			}
		})
	}
}
