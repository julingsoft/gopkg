# QA

## RocketMQ 5.x 的 Proxy 有 8080 和 8081 的端口，有什么区别？
  - 8080 端口（HTTP） ：主要用于运维管理，如 Topic 创建、队列管理等 RESTful 操作，不是客户端 SDK 用来收发消息的端口。
  - 8081 端口（gRPC） ：这是 RocketMQ 5.x 客户端 SDK 连接 Proxy 收发消息的端口 。你使用的 rocketmq-clients/golang/v5 就是通过 gRPC 协议连接的。

## RocketMQ Broker 的端口功能

| 端口 | 协议 | 功能 |
|------|------|------|
| **10909** | gRPC (Proxy) | Proxy gRPC 监听端口（5.x 新增） |
| **10911** | TCP (Remoting) | **主端口**，客户端连接 Broker 用于收发消息 |
| **10912** | TCP (Fast Remoting) | 快速通道，用于 Broker 间主从同步 |

## 重要区别

| 场景 | 端口 | 说明 |
|------|------|------|
| **RocketMQ 4.x 客户端** | 10911 | 直连 Broker 的 TCP 端口 |
| **RocketMQ 5.x 客户端** | 8081 (Proxy) | 通过 Proxy 的 gRPC 端口 |
| **RocketMQ 5.x + Proxy** | 10909 (Proxy) | Proxy 内嵌的 gRPC 端口 |

## 关键点

**RocketMQ 5.x 的客户端 SDK 不能直接连接 10911/10912**，这些是传统 TCP 协议端口，5.x 客户端使用 gRPC 协议。

所以如果你使用 RocketMQ 5.x 客户端：
- ✅ **必须通过 Proxy**（8081 或 10909）
- ❌ **不能直接连 Broker 的 10911/10912**

如果你有 Proxy 部署，应该配置：
```yaml
rocketmq:
  endpoint: 172.23.198.91:8081  # 或 10909
```