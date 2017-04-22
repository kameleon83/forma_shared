package lib

import (
	"crypto/rand"
	"fmt"
)

func TokenCreate() string {
	b := make([]byte, 20)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}
