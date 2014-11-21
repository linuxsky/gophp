package crypto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMD5(t *testing.T) {
	plain := "These pretzels are making me thirsty."
	cipher := "b0804ec967f48520697662a204f5fe72"

	assert.Equal(t, cipher, MD5([]byte(plain)))
}
