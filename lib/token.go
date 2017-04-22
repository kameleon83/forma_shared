package lib

import (
	"crypto/rand"
)

func TokenCreate() string {
	b := make([]byte, 20)
	rand.Read(b)
	return string(b)
}
