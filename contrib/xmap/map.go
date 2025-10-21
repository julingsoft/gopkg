package xmap

// GetUniqueKeyCount 获取map中key的唯一数量
func GetUniqueKeyCount(m map[string]interface{}) int {
	uniqueKeys := make(map[string]bool)
	for key := range m {
		uniqueKeys[key] = true
	}
	return len(uniqueKeys)
}
