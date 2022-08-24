package utils

import "regexp"

func EmailValidation(text string) bool {
	res := regexp.MustCompile(`[a-z0-9._%+\-]+@[a-z0-9._%+\-]+\.[a-z0-9._%+\-]{2,4}$`)
	return res.MatchString(text)
}

func PasswordValidation(text string) bool {
	if len(text) > 16 {
		return true
	}
	return false
}
