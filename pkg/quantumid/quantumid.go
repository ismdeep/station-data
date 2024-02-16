package quantumid

import (
	"crypto/rand"
	"fmt"
	"time"
)

// NewString new a quantum id
func NewString() string {
	raw := make([]byte, 16)
	s := time.Now().UnixNano()
	for i := 0; i < 8; i++ {
		raw[7-i] = byte(s % 256)
		s >>= 8
	}
	_, _ = rand.Read(raw[8:16])

	return fmt.Sprintf("%x-%x-%x-%x-%x", raw[0:4], raw[4:6], raw[6:8], raw[8:10], raw[10:16])
}
