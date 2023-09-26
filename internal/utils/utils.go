package utils

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/google/uuid"
)

func GenerateUuid() string {
	return uuid.NewString()
}

func Md5(password string) string {
	hasher := md5.New()
	hasher.Write([]byte(password))
	passwordHash := hex.EncodeToString(hasher.Sum(nil))

	return passwordHash
}
