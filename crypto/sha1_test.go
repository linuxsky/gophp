package crypto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSHA1(t *testing.T) {
	plain := "This page intentionally left blank."
	cipher := "af064923bbf2301596aac4c273ba32178ebc4a96"

	assert.Equal(t, cipher, SHA1([]byte(plain)))
}
