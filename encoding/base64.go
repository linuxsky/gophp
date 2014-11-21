package encoding

import (
	"encoding/base64"
)

func Base64Encode(src []byte) string {
	return base64.URLEncoding.EncodeToString(src)
}

func Base64Decode(s string) ([]byte, error) {
	return base64.URLEncoding.DecodeString(s)
}
