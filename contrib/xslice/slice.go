package xslice

func RemoveDuplicates(slice []string) []string {
	allKeys := make(map[string]bool)
	var list = make([]string, 0)
	for _, item := range slice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}
