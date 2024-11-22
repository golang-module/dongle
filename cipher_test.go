package dongle

import (
	"crypto/aes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCipher_Encrypt(t *testing.T) {
	cipher := NewCipher()
	cipher.SetMode(CBC)
	cipher.SetPadding(PKCS7)

	block, err1 := aes.NewCipher(aesKey)
	assert.Nil(t, err1)
	dst, err2 := cipher.Encrypt([]byte(""), block)
	assert.Nil(t, err2)
	assert.Equal(t, []byte(""), dst)
}

func TestCipher_Decrypt(t *testing.T) {
	cipher := NewCipher()
	cipher.SetMode(CBC)
	cipher.SetPadding(PKCS7)

	block, err1 := aes.NewCipher(aesKey)
	assert.Nil(t, err1)
	dst, err2 := cipher.Decrypt([]byte(""), block)
	assert.Nil(t, err2)
	assert.Equal(t, []byte(""), dst)
}

func TestCipher_Padding(t *testing.T) {
	cipher := NewCipher()
	cipher.SetMode(CBC)
	cipher.SetPadding(PKCS7)

	assert.Equal(t, []byte(""), cipher.AnsiX923UnPadding([]byte("")))
	assert.Equal(t, []byte(""), cipher.PKCS7UnPadding([]byte("")))
}
