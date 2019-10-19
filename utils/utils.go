package utils

import (
	"crypto/md5"
	"encoding/hex"
)

// HashToMD5 will change string data to a 16 bytes long hex string.
func HashToMD5(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}
