package helper

import (
	"regexp"
	"strings"
)

func ValidateEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	if len(email) < 10 || len(email) > 50 {
		return false
	}
	return emailRegex.MatchString(email)
}

func ValidatePassword(password string) bool {
	if len(password) < 12 || len(password) > 15 {
		return false
	}
	return true
}

// Check Empty Form
func IsEmpty(str string) bool {
	// Trimming Whitespace
	return len(strings.Trim(str, " ")) == 0
}