package crypto

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5(data []byte) string {
	a := md5.Sum(data)
	return hex.EncodeToString(a[:])
}
