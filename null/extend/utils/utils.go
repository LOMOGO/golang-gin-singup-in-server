package utils

import (
	"crypto/sha1"
	"encoding/hex"
)

func MakeSha1(source string) string {
	sha1Hash := sha1.New()
	sha1Hash.Write([]byte(source))
	return hex.EncodeToString(sha1Hash.Sum(nil))
}
