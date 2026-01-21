package xhttp

import (
	"context"
	"net/http"
	"time"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gclient"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/julingsoft/gogf/contrib/xlog"
)

type HttpClient struct {
	BaseUrl string
	Client  *gclient.Client
}

func New(baseUrl string, timeouts ...time.Duration) *HttpClient {
	var timeout = 5 * time.Second
	if len(timeouts) > 0 {
		timeout = timeouts[0]
	}

	var client = gclient.New().
		// 超时时间
		SetTimeout(timeout)
	// 重试次数和重试间隔
	// Retry(3, timeout*2)

	return &HttpClient{
		BaseUrl: baseUrl,
		Client:  client,
	}
}

func (c *HttpClient) SetTimeout(timeout time.Duration) *HttpClient {
	c.Client.SetTimeout(timeout)
	return c
}

func (c *HttpClient) Get(ctx context.Context, url string) ([]byte, error) {
	startTime := time.Now()

	url = gstr.TrimRight(c.BaseUrl, "/") + "/" + gstr.TrimLeft(url, "/")
	r, err := c.Client.Get(ctx, url)
	if err != nil {
		return nil, err
	}
	defer r.Close()

	respBytes := r.ReadAll()
	duration := time.Since(startTime).Milliseconds()
	g.Log().Info(ctx, xlog.LogData{
		Method:   http.MethodGet,
		Url:      url,
		Response: string(respBytes),
		Status:   r.Response.StatusCode,
		Duration: duration,
	})

	return respBytes, nil
}

func (c *HttpClient) Post(ctx context.Context, url string, data string) ([]byte, error) {
	startTime := time.Now()

	url = gstr.TrimRight(c.BaseUrl, "/") + "/" + gstr.TrimLeft(url, "/")
	r, err := c.Client.Post(ctx, url, data)
	if err != nil {
		return nil, gerror.Wrapf(err, "req: %s", data)
	}
	defer r.Close()

	respBytes := r.ReadAll()
	duration := time.Since(startTime).Milliseconds()
	g.Log().Info(ctx, xlog.LogData{
		Method:   http.MethodPost,
		Url:      url,
		Request:  data,
		Response: string(respBytes),
		Status:   r.Response.StatusCode,
		Duration: duration,
	})

	return respBytes, nil
}
