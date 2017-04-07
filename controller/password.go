package controller

import (
	"crypto/sha256"
	"encoding/base64"
	"time"
)

// CheckEqualsPasswordString e
func CheckEqualsPasswordString(pass1, pass2 string) string {
	if pass1 == pass2 {
		return EncryptionPassword(pass1)
	}
	return ""
}

// CheckEqualsPasswordEncrypt e
func CheckEqualsPasswordEncrypt(passInput, passBDD string) bool {
	if EncryptionPassword(passInput) == passBDD {
		return true
	}
	return false
}

// EncryptionPassword e
func EncryptionPassword(pass string) string {
	hasher := sha256.New()
	hasher.Write([]byte(pass))
	sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	return sha
}

// EncryptionEmail e
func EncryptionEmail(email string) string {
	hasher := sha256.New()
	hasher.Write([]byte(email + "_" + formatDateEncryption()))
	sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	return sha
}

// EncryptionEmailRescue e
func EncryptionEmailRescue(email string) string {
	hasher := sha256.New()
	hasher.Write([]byte("rescue" + email + "_" + formatDateEncryption()))
	sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	return sha
}

func formatDateEncryption() string {
	now := time.Now()
	time := "2006-01-02"

	return now.Format(time)
}
