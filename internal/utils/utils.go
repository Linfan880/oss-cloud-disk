package utils

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/google/uuid"
)

func GenerateUuid() string {
	return uuid.NewString()
}

func Md5(str string) string {
	hash := md5.New()
	hash.Write([]byte(str))
	passwordHash := hex.EncodeToString(hash.Sum(nil))

	return passwordHash
}
