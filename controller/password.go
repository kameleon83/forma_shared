package controller

import (
	"crypto/sha256"
	"encoding/base64"
)

func CheckEqualsPasswordString(pass1, pass2 string) string {
	if pass1 == pass2 {
		return encryptionPassword(pass1)
	}
	return ""
}

func CheckEqualsPasswordEncrypt(passInput, passBDD string) bool {
	if encryptionPassword(passInput) == passBDD {
		return true
	}
	return false
}

func encryptionPassword(pass string) string {
	hasher := sha256.New()
	hasher.Write([]byte(pass))
	sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	return sha
}
