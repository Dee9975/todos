package util

func IntToBool(i int) bool {
	return i > 0
}

func BoolToInt(i bool) int {
	if i {
		return 1
	}
	return 0
}
