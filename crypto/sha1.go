package crypto

import (
	"crypto/sha1"
	"encoding/hex"
)

func SHA1(data []byte) string {
	a := sha1.Sum(data)
	return hex.EncodeToString(a[:])
}
