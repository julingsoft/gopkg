package xcache

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestCache(t *testing.T) {
	cache := New()

	cacheKey := "haha"
	cacheVal := "world."

	ctx := context.Background()
	err2 := cache.Set(ctx, cacheKey, cacheVal, time.Second*10)
	if err2 != nil {
		panic(err2)
	}

	s := cache.MustGet(ctx, cacheKey).String()
	fmt.Println(s)
}
