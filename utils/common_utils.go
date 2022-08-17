package utils

import "strconv"

// StringToInt64 : conver string to int
func StringToInt64(text string) (int, error) {
	return strconv.Atoi(text)
}
