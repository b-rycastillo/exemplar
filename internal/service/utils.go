package service

import (
	"regexp"
	"strconv"
	"time"
)

// IsValidEmail validates email format
func IsValidEmail(email string) bool {
	pattern := `^[^\s@]+@[^\s@]+\.[^\s@]+$`
	re := regexp.MustCompile(pattern)
	return re.MatchString(email)
}

// IsValidPassword validates password strength (minimum 8 characters)
func IsValidPassword(password string) bool {
	return len(password) >= 8
}

// GenerateID generates a unique ID
func GenerateID() string {
	return "user_" + strconv.FormatInt(time.Now().UnixNano(), 10) + "_" + strconv.FormatInt(int64(time.Now().Nanosecond()%1000), 10)
}
