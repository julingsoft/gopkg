package xcos

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/tencentyun/cos-go-sdk-v5"
)

type Client struct {
	Config *Config
	Client *cos.Client
}

var (
	cosClient  *Client
	clientOnce sync.Once
)

func NewClient(config *Config) *Client {
	clientOnce.Do(func() {
		u, _ := url.Parse(fmt.Sprintf("%s.cos.%s.myqcloud.com", config.Bucket, config.Region))
		b := &cos.BaseURL{BucketURL: u}
		c := cos.NewClient(b, &http.Client{
			Transport: &cos.AuthorizationTransport{
				SecretID:  config.SecretId,
				SecretKey: config.SecretKey,
			},
		})
		cosClient = &Client{
			Config: config,
			Client: c,
		}
	})
	return cosClient
}

// 获取预签名 URL
func (c *Client) GetCosURL(ctx context.Context, name string, expired ...time.Duration) (string, error) {
	expire := time.Hour
	if len(expired) > 0 {
		expire = expired[0]
	}
	presignedURL, err := c.Client.Object.GetPresignedURL(ctx, http.MethodGet, name, c.Config.SecretId, c.Config.SecretKey, expire, nil)
	if err != nil {
		g.Log().Error(ctx, err)
		return "", err
	}
	return presignedURL.String(), nil
}

func GetCosKeyFromURLSafe(rawURL string) string {
	parsed, err := url.Parse(rawURL)
	if err != nil {
		return rawURL
	}

	key := strings.TrimPrefix(parsed.Path, "/")

	// 解码 URL 编码字符
	key, err = url.QueryUnescape(key)
	if err != nil {
		return rawURL
	}

	return key
}
