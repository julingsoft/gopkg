package xstr

import "hash/crc32"

func StringToID(s string) int {
	// 计算 CRC32 校验和
	checksum := crc32.ChecksumIEEE([]byte(s))
	// 转为 int（在 64 位系统 int 是 int64，32 位系统是 int32）
	return int(checksum)
}
