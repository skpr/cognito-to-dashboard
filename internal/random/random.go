package random

import (
	"math/rand"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seed = rand.New(rand.NewSource(time.Now().UnixNano()))

// StringWithCharset creates a random string with characters from the charset.
func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)

	for i := range b {
		b[i] = charset[seed.Intn(len(charset))]
	}

	return string(b)
}

// String creates a random alphanumeric string.
func String(length int) string {
	return StringWithCharset(length, charset)
}
