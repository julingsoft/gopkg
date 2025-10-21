package xguid

import (
	"fmt"
	"net"
)

func PrivateIPv4() (string, error) {
	// 获取所有网络接口
	interfaces, err := net.Interfaces()
	if err != nil {
		return "", fmt.Errorf("无法获取网络接口: %v", err)
	}

	// 遍历所有网络接口
	for _, iface := range interfaces {
		// 排除回环接口和未启用的接口
		if iface.Flags&net.FlagLoopback == 0 && iface.Flags&net.FlagUp != 0 {
			// 获取接口的地址
			addrs, err := iface.Addrs()
			if err != nil {
				return "", fmt.Errorf("无法获取接口地址: %v", err)
			}

			// 遍历接口的地址
			for _, addr := range addrs {
				// 检查地址是否为 IP 地址
				ipNet, ok := addr.(*net.IPNet)
				if !ok {
					continue
				}

				// 排除 IPv6 地址和本地回环地址
				ip := ipNet.IP
				if ip.IsLoopback() || ip.To4() == nil {
					continue
				}

				// 检查是否为私有 IP 地址
				if ip.IsPrivate() {
					return ip.String(), nil
				}
			}
		}
	}

	return "", fmt.Errorf("未找到私有 IP 地址")
}

func MachineID() (uint16, error) {
	ip, err := PrivateIPv4()
	if err != nil {
		return 0, err
	}

	return uint16(ip[2])<<8 + uint16(ip[3]), nil
}
