package xslice

// RemoveDuplicates 删除切片中的重复元素
func RemoveDuplicates(slice []string) []string {
	allKeys := make(map[string]bool)
	var list = make([]string, 0, len(slice))
	for _, item := range slice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

// InArray 检查整数 s 是否在切片 a 中
func InArray(a []int, s int) bool {
	for _, item := range a {
		if item == s {
			return true
		}
	}
	return false
}

// InArrayInt64 检查 int64 类型的 s 是否在切片 a 中
func InArrayInt64(a []int64, s int64) bool {
	for _, item := range a {
		if item == s {
			return true
		}
	}
	return false
}
