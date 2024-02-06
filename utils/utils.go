package utils

// InList 判断key是否存在于列表中
func InList(key string, list []string) bool {
	for _, v := range list {
		if v == key {
			return true
		}
	}
	return false
}

// Reverse 反转切片
func Reverse[T any](slice []T) {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}
