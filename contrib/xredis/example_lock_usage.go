package xredis

import (
	"context"
	"fmt"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/google/uuid"
)

// ExampleOrderLock 示例：订单处理中使用分布式锁
func ExampleOrderLock(ctx context.Context, orderId string) error {
	// 1. 生成唯一的锁值（建议使用：实例ID + 时间戳 + UUID）
	lockValue := fmt.Sprintf("%s-%d-%s",
		gconv.String(g.Cfg().MustGet(ctx, "server.instance_id", "default")),
		time.Now().UnixNano(),
		uuid.New().String(),
	)

	// 2. 定义锁的key和过期时间
	lockKey := fmt.Sprintf("order:lock:%s", orderId)
	lockExp := int64(30) // 30秒过期

	// 3. 尝试获取锁（带重试机制）
	maxRetries := 3
	retryInterval := 500 * time.Millisecond
	hasLock := false
	var err error

	for i := 0; i < maxRetries; i++ {
		hasLock, err = GetLockWithValue(ctx, lockKey, lockValue, lockExp)
		if err != nil {
			g.Log().Errorf(ctx, "获取锁失败，第 %d 次尝试: %v", i+1, err)
			if i == maxRetries-1 {
				return fmt.Errorf("获取订单锁失败: %w", err)
			}
			time.Sleep(retryInterval)
			continue
		}

		if hasLock {
			g.Log().Infof(ctx, "成功获取订单锁: %s, 锁值: %s", lockKey, lockValue)
			break
		}

		if i < maxRetries-1 {
			g.Log().Infof(ctx, "订单锁被占用，等待后重试，第 %d 次", i+1)
			time.Sleep(retryInterval)
		}
	}

	if !hasLock {
		return fmt.Errorf("订单锁被占用，请稍后重试")
	}

	// 4. 确保锁一定会被释放
	defer func() {
		released, err := DelLockWithValue(ctx, lockKey, lockValue)
		if err != nil {
			g.Log().Errorf(ctx, "释放订单锁失败: %s, 错误: %v", lockKey, err)
		} else if released {
			g.Log().Infof(ctx, "成功释放订单锁: %s", lockKey)
		} else {
			g.Log().Warningf(ctx, "订单锁已不存在或被其他实例持有: %s", lockKey)
		}
	}()

	// 5. 执行业务逻辑
	g.Log().Infof(ctx, "开始处理订单: %s", orderId)
	// ... 订单处理逻辑 ...
	time.Sleep(2 * time.Second) // 模拟业务处理
	g.Log().Infof(ctx, "订单处理完成: %s", orderId)

	return nil
}

// ExamplePaymentLock 示例：支付回调中使用分布式锁
func ExamplePaymentLock(ctx context.Context, paymentNo string, amount int64) error {
	lockValue := fmt.Sprintf("payment-%d-%s", time.Now().UnixNano(), uuid.New().String())
	lockKey := fmt.Sprintf("payment:lock:%s", paymentNo)
	lockExp := int64(20) // 20秒

	// 简单封装：获取锁
	getLock := func() (bool, error) {
		return GetLockWithValue(ctx, lockKey, lockValue, lockExp)
	}

	// 简单封装：释放锁
	releaseLock := func() {
		if released, err := DelLockWithValue(ctx, lockKey, lockValue); err != nil {
			g.Log().Errorf(ctx, "释放支付锁失败: %v", err)
		} else if !released {
			g.Log().Warningf(ctx, "支付锁已被释放或不属于当前实例")
		}
	}

	// 获取锁
	hasLock, err := getLock()
	if err != nil {
		return fmt.Errorf("获取支付锁失败: %w", err)
	}
	if !hasLock {
		return fmt.Errorf("支付正在处理中，请勿重复提交")
	}

	defer releaseLock()

	// 执行支付业务逻辑
	g.Log().Infof(ctx, "处理支付回调: %s, 金额: %d", paymentNo, amount)
	// ... 支付处理逻辑 ...

	return nil
}

// ExampleLockWithHelper 示例：使用辅助函数简化锁的使用
func ExampleLockWithHelper(ctx context.Context, orderId string) error {
	// 使用辅助函数获取锁
	unlock, err := AcquireLockWithRetry(ctx, fmt.Sprintf("order:process:%s", orderId), 30, 3)
	if err != nil {
		return err
	}
	defer unlock() // 自动释放锁

	// 执行业务逻辑
	g.Log().Infof(ctx, "处理订单业务: %s", orderId)
	return nil
}

// AcquireLockWithRetry 辅助函数：获取锁并返回释放函数
// key: 锁的key
// expSeconds: 锁过期时间（秒）
// maxRetries: 最大重试次数
func AcquireLockWithRetry(ctx context.Context, key string, expSeconds int64, maxRetries int) (unlock func(), err error) {
	// 生成唯一锁值
	lockValue := fmt.Sprintf("%s-%d-%s",
		gconv.String(g.Cfg().MustGet(ctx, "server.instance_id", "unknown")),
		gtime.Now().UnixNano(),
		uuid.New().String(),
	)

	retryInterval := 500 * time.Millisecond
	hasLock := false

	// 尝试获取锁
	for i := 0; i < maxRetries; i++ {
		hasLock, err = GetLockWithValue(ctx, key, lockValue, expSeconds)
		if err != nil {
			g.Log().Errorf(ctx, "获取锁失败 [%s], 第 %d 次尝试: %v", key, i+1, err)
			if i == maxRetries-1 {
				return nil, fmt.Errorf("获取锁失败: %w", err)
			}
			time.Sleep(retryInterval)
			continue
		}

		if hasLock {
			g.Log().Infof(ctx, "成功获取锁 [%s]", key)
			break
		}

		if i < maxRetries-1 {
			g.Log().Infof(ctx, "锁被占用 [%s], 第 %d 次重试", key, i+1)
			time.Sleep(retryInterval)
		}
	}

	if !hasLock {
		return nil, fmt.Errorf("锁被占用，请稍后重试 [%s]", key)
	}

	// 返回释放锁的函数
	unlock = func() {
		// 使用 Background context 避免父 context 已取消导致释放失败
		releaseCtx := context.Background()
		if released, err := DelLockWithValue(releaseCtx, key, lockValue); err != nil {
			g.Log().Errorf(releaseCtx, "释放锁失败 [%s]: %v", key, err)
		} else if released {
			g.Log().Infof(releaseCtx, "成功释放锁 [%s]", key)
		} else {
			g.Log().Warningf(releaseCtx, "锁已不存在或不属于当前实例 [%s]", key)
		}
	}

	return unlock, nil
}

// ExampleBatchOperation 示例：批量操作中使用锁
func ExampleBatchOperation(ctx context.Context, orderIds []string) error {
	// 为整个批次获取锁
	batchId := uuid.New().String()
	unlock, err := AcquireLockWithRetry(ctx, fmt.Sprintf("batch:lock:%s", batchId), 60, 3)
	if err != nil {
		return err
	}
	defer unlock()

	// 批量处理订单
	for _, orderId := range orderIds {
		g.Log().Infof(ctx, "批量处理订单: %s", orderId)
		// ... 处理逻辑 ...
	}

	return nil
}

// ExampleConcurrentSafe 示例：并发安全的库存扣减
func ExampleConcurrentSafe(ctx context.Context, productId string, quantity int) error {
	unlock, err := AcquireLockWithRetry(ctx, fmt.Sprintf("product:stock:%s", productId), 10, 5)
	if err != nil {
		return fmt.Errorf("无法获取库存锁: %w", err)
	}
	defer unlock()

	// 读取库存
	g.Log().Infof(ctx, "检查产品 %s 的库存", productId)
	// currentStock := getStock(productId)

	// 扣减库存
	// if currentStock >= quantity {
	//     updateStock(productId, currentStock - quantity)
	// }

	return nil
}
