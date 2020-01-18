package utils

import (
	"crypto/md5"
	"crypto/rand"
	"fmt"
)

func Salt() string {
	b := make([]byte, 32)
	rand.Read(b)
	return fmt.Sprintf("%x", md5.Sum(b))
}

// md5(password + salt)
func EncryptPassword(password, salt string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(password+salt)))
}
