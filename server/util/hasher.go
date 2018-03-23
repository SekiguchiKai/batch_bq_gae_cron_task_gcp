package util

import (
	"crypto/sha256"
	"encoding/hex"
)

// 引数からhashを生成する
func GetHash(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	sum := h.Sum(nil)
	return hex.EncodeToString(sum)
}