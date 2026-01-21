package xcache

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
)

func New() (cache *gcache.Cache) {
	return GetInstance()
}

func GetInstance() (cache *gcache.Cache) {
	return gcache.NewWithAdapter(gcache.NewAdapterRedis(g.Redis("cache")))
}
