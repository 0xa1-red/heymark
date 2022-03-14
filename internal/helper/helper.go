package helper

import (
	"crypto/sha256"
	"fmt"
)

func Sha256(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	return fmt.Sprintf("%x", h.Sum(nil))
}
