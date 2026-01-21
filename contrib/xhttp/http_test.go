package xhttp

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/gogf/gf/v2/net/gclient"
)

func TestHttpClient_Get(t *testing.T) {
	client := New("https://www.qq.com")
	client = client.SetTimeout(10 * time.Millisecond)

	type args struct {
		ctx context.Context
		url string
	}
	tests := []struct {
		name    string
		fields  *HttpClient
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name:   "case1",
			fields: client,
			args: args{
				ctx: context.Background(),
				url: "/",
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &HttpClient{
				BaseUrl: tt.fields.BaseUrl,
				Client:  tt.fields.Client,
			}
			got, err := c.Get(tt.args.ctx, tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHttpClient_Post(t *testing.T) {
	client := New("https://www.qq.com")
	client = client.SetTimeout(10 * time.Millisecond)

	type fields struct {
		BaseUrl string
		Client  *gclient.Client
	}
	type args struct {
		ctx  context.Context
		url  string
		data string
	}
	tests := []struct {
		name    string
		fields  *HttpClient
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name:   "case1",
			fields: client,
			args: args{
				ctx:  context.Background(),
				url:  "/",
				data: "{\"hallo\":\"walold\"}",
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &HttpClient{
				BaseUrl: tt.fields.BaseUrl,
				Client:  tt.fields.Client,
			}
			got, err := c.Post(tt.args.ctx, tt.args.url, tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("Post() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Post() got = %v, want %v", got, tt.want)
			}
		})
	}
}
