package xcache

import (
	"sync"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
)

var (
	cacheInstance *gcache.Cache
	cacheOnce     sync.Once
)

func New() (cache *gcache.Cache) {
	return GetInstance()
}

func GetInstance() (cache *gcache.Cache) {
	cacheOnce.Do(func() {
		cacheInstance = gcache.NewWithAdapter(gcache.NewAdapterRedis(g.Redis("cache")))
	})

	return cacheInstance
}
