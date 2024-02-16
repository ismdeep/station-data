package sha256util

import (
	"crypto/sha256"
	"fmt"
)

func SHA256(content []byte) string {
	h := sha256.New()
	h.Write(content)
	return fmt.Sprintf("%x", h.Sum(nil))
}
