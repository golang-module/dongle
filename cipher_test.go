package dongle

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCipher_SetMode(t *testing.T) {
	assert := assert.New(t)
	cipher := NewCipher()

	cipher.SetMode(CBC)
	assert.Equal(CBC, cipher.mode)
}

func TestCipher_SetPadding(t *testing.T) {
	assert := assert.New(t)
	cipher := NewCipher()

	cipher.SetPadding(PKCS7Padding)
	assert.Equal(PKCS7Padding, cipher.padding)
}

func TestCipher_SetIV(t *testing.T) {
	assert := assert.New(t)
	cipher := NewCipher()

	cipher.SetIV("0123456789abcdef")
	assert.Equal([]byte("0123456789abcdef"), cipher.iv)

	cipher.SetIV([]byte("0123456789abcdef"))
	assert.Equal([]byte("0123456789abcdef"), cipher.iv)
}

func TestCipher_SetKey(t *testing.T) {
	assert := assert.New(t)
	cipher := NewCipher()

	cipher.SetKey("0123456789abcdef")
	assert.Equal([]byte("0123456789abcdef"), cipher.key)

	cipher.SetKey([]byte("0123456789abcdef"))
	assert.Equal([]byte("0123456789abcdef"), cipher.key)
}
