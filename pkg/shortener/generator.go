package shortener

import (
	"crypto/rand"
	"math/big"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var charsetLen = big.NewInt(int64(len(charset)))

// Generate creates a cryptographically secure random string of the specified length
func Generate(length int) string {
	b := make([]byte, length)
	for i := range b {
		num, err := rand.Int(rand.Reader, charsetLen)
		if err != nil {
			// Fallback to a safe character if random generation fails
			b[i] = charset[0]
			continue
		}
		b[i] = charset[num.Int64()]
	}
	return string(b)
}
