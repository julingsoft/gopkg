package xoss

import (
	"context"
	"io"
	"reflect"
	"strings"
	"testing"

	"github.com/aliyun/alibabacloud-oss-go-sdk-v2/oss"
)

func TestPutObject(t *testing.T) {
	type args struct {
		ctx        context.Context
		objectName string
		body       io.Reader
	}
	tests := []struct {
		name    string
		args    args
		want    *oss.PutObjectResult
		wantErr bool
	}{
		{
			name: "case1",
			args: args{
				ctx:        context.TODO(),
				objectName: "test/abc.txt",
				body:       strings.NewReader("haha"),
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetInstance(tt.args.ctx).PutObject(tt.args.objectName, tt.args.body)
			if (err != nil) != tt.wantErr {
				t.Errorf("PutObject() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PutObject() got = %v, want %v", got, tt.want)
			}
		})
	}
}
