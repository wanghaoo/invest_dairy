package util

import (
	"crypto/hmac"
	"crypto/sha256"
	"strings"
	"encoding/hex"
)

type UserInRequest struct {
	UserID int
}

func GetSign(secret, data string) string {
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(data))
	expectedMAC := mac.Sum(nil)
	return strings.ToLower(hex.EncodeToString(expectedMAC))
}