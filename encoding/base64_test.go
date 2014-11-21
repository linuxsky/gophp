package encoding

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	ori string = "any + old & data"
	enc string = "YW55ICsgb2xkICYgZGF0YQ=="
)

func TestBase64Encode(t *testing.T) {
	assert.Equal(t, enc, Base64Encode([]byte(ori)))
}

func TestBase64Decode(t *testing.T) {
	dec, err := Base64Decode(enc)

	assert.NoError(t, err)
	assert.Equal(t, []byte(ori), dec)
}
