package hash

import (
	"crypto/rand"
	"fmt"
)

type HashGenerator struct {
	Charset string
	Length  int
}

func (h *HashGenerator) Generate() (string, error) {
	b := make([]byte, h.Length)
	_, err := rand.Read(b)
	if err != nil {
		return "", fmt.Errorf("failed to generate random bytes")
	}
	for i := range b {
		b[i] = h.Charset[int(b[i])%len(h.Charset)]
	}
	return string(b), nil
}
