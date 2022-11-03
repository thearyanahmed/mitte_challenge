package utils

import (
	"crypto/rand"
	"fmt"
)

func CreateRandomString(n int8) (string, error) {
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return fmt.Sprintf("%X", b), nil
}
