package util

import (
	"crypto/md5"
	"encoding/hex"
)

// ToHexStr md5加密
func ToHexStr(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	cipherStr := h.Sum(nil)
	hexStr := hex.EncodeToString(cipherStr)
	return hexStr
}
