package helper

import (
	"crypto/sha256"
	"fmt"
)

func Generate4CharsPassword(input string) string {
	h := sha256.New()
	h.Write([]byte(input))
	output := fmt.Sprintf("%x", h.Sum(nil))

	return output[len(output)-4:]
}
