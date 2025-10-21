package xstr

// SubStr 返回字符串 s 从 start 开始长度为 length 的子串
func SubStr(s string, start, length int) string {
	// 提前处理空字符串或length为0的情况
	if s == "" || length == 0 {
		return ""
	}

	// 将字符串转换为 rune 切片，以便正确处理 Unicode 字符
	runes := []rune(s)
	runeCount := len(runes)

	// 处理负数start（支持从末尾偏移）
	if start < 0 {
		start += runeCount
	}

	// 确保start在有效范围内
	if start < 0 {
		start = 0
	} else if start >= runeCount {
		return ""
	}

	// 处理负数length（支持从末尾截断）
	if length < 0 {
		length += runeCount - start
		// 调整后仍无效则返回空
		if length <= 0 {
			return ""
		}
	}

	end := start + length
	if end > runeCount {
		end = runeCount
	}

	return string(runes[start:end])
}
