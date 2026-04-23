package xhttp

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gclient"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/julingsoft/gopkg/contrib/xlog"
)

type HttpClient struct {
	BaseUrl string
	Client  *gclient.Client
}

var (
	httpClientMap = make(map[string]*HttpClient)
	mu            sync.RWMutex
)

// New 获取或创建 HttpClient 单例（每个 baseUrl 一个单例）
func New(baseUrl string, timeouts ...time.Duration) *HttpClient {
	mu.RLock()
	if c, ok := httpClientMap[baseUrl]; ok {
		mu.RUnlock()
		return c
	}
	mu.RUnlock()

	mu.Lock()
	defer mu.Unlock()

	// 二次检查防止竞态
	if c, ok := httpClientMap[baseUrl]; ok {
		return c
	}

	var timeout = 5 * time.Second
	if len(timeouts) > 0 {
		timeout = timeouts[0]
	}

	var (
		hostname, _ = os.Hostname()
		clientAgent = fmt.Sprintf(`%s`, hostname)
	)
	c := &HttpClient{
		BaseUrl: baseUrl,
		Client: gclient.New().
			SetHeader("User-Agent", clientAgent).
			SetTimeout(timeout),
	}
	httpClientMap[baseUrl] = c
	return c
}

func (c *HttpClient) SetTimeout(timeout time.Duration) *HttpClient {
	// 由于是单例模式，为了防止在此处修改超时时间影响该 BaseUrl 的其他并发请求，
	// 我们返回一个克隆的客户端副本。Clone() 会共享底层的连接池（Transport）。
	return &HttpClient{
		BaseUrl: c.BaseUrl,
		Client:  c.Client.Clone().SetTimeout(timeout),
	}
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
