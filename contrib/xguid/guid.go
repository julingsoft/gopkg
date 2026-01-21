package xguid

import (
	"fmt"
	"net"
	"time"

	"github.com/sony/sonyflake/v2"
)

// New 创建一个新的 Sonyflake ID，机器 ID 从 IP 地址自动生成
func New() (int64, error) {
	st := sonyflake.Settings{
		TimeUnit: time.Millisecond,
		// MachineID 默认为 nil，Sonyflake 将从 IP 地址自动生成
		// 它使用找到的第一个私有 IP 地址的低 16 位
	}

	s, err := sonyflake.New(st)
	if err != nil {
		return 0, err
	}

	return s.NextID()
}

// NewWithMachineID 使用自定义机器 ID 创建一个新的 Sonyflake ID
// machineID 应该在 0 到 65535 之间（16 位）
func NewWithMachineID(machineID int) (int64, error) {
	st := sonyflake.Settings{
		TimeUnit: time.Millisecond,
		MachineID: func() (int, error) {
			return machineID, nil
		},
	}

	s, err := sonyflake.New(st)
	if err != nil {
		return 0, err
	}

	return s.NextID()
}

// GetDefaultMachineID 返回 Sonyflake 自动生成的机器 ID
// 这对于调试或日志记录很有用
func GetDefaultMachineID() (uint16, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return 0, fmt.Errorf("failed to get network interfaces: %w", err)
	}

	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				// 使用 IP 地址的低 16 位作为机器 ID
				ip := ipnet.IP.To4()
				return uint16(ip[2])<<8 + uint16(ip[3]), nil
			}
		}
	}

	return 0, fmt.Errorf("no valid private IP address found")
}

func NextID() int64 {
	for {
		if id, err := New(); err == nil {
			return id
		}
	}
}
