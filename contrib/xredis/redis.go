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

// GetLockWithValue 使用唯一标识符，防止误删其他实例的锁
func GetLockWithValue(ctx context.Context, key string, value string, exp int64) (bool, error) {
	redis := New()
	if redis == nil {
		return false, fmt.Errorf("redis not found")
	}

	result, err := redis.Set(ctx, key, value, gredis.SetOption{
		TTLOption: gredis.TTLOption{
			EX: gconv.PtrInt64(exp),
		},
		NX:  true,
		Get: false,
	})

	if err != nil {
		return false, fmt.Errorf("failed to get lock: %v", err)
	}

	return result.Val() == "OK", nil
}

// DelLockWithValue 只删除自己持有的锁（使用 Lua 脚本保证原子性）
func DelLockWithValue(ctx context.Context, key string, value string) (bool, error) {
	redis := New()
	if redis == nil {
		return false, fmt.Errorf("redis not found")
	}

	script := `
		if redis.call("get", KEYS[1]) == ARGV[1] then
			return redis.call("del", KEYS[1])
		else
			return 0
		end
	`

	result, err := redis.Eval(ctx, script, 1, []string{key}, []any{value})
	if err != nil {
		return false, fmt.Errorf("failed to del lock: %v", err)
	}

	return result.Int() == 1, nil
}
