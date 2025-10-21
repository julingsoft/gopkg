package xredis

import (
	"context"
	"fmt"

	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

func New(typeName ...string) *gredis.Redis {
	redisType := "default"
	if len(typeName) > 0 {
		redisType = typeName[0]
	}

	return g.Redis(redisType)
}

func GetLock(ctx context.Context, key string, exp int64) (bool, error) {
	redis := New()
	if redis == nil {
		return false, fmt.Errorf("redis not found")
	}

	result, err := redis.Set(ctx, key, 1, gredis.SetOption{
		TTLOption: gredis.TTLOption{
			EX: gconv.PtrInt64(exp), // 设置锁的有效期
		},
		NX:  true,  // 只有key不存在时才会成功设置
		Get: false, // 不需要获取原始值
	})

	if err != nil {
		return false, fmt.Errorf("failed to get lock: %v", err)
	}

	// 判断是否获取到锁
	if result.Val() == "OK" {
		return true, nil // 成功获取锁
	}

	return false, nil // 未获取到锁
}

func DelLock(ctx context.Context, key string) (bool, error) {
	redis := New()
	if redis == nil {
		return false, fmt.Errorf("redis not found")
	}

	_, err := redis.Del(ctx, key)

	if err != nil {
		return false, fmt.Errorf("failed to del lock: %v", err)
	}

	return true, nil // 未获取到锁
}
