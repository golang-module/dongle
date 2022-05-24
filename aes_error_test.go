package dongle

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	aesRightInput   = "hello world"
	aesRightMode    = CBC
	aesRightPadding = PKCS7
	aesRightKey     = "1234567887654321"
	aesRightIV      = "1234567887654321"

	aesErrorInput   = "xxx"
	aesErrorMode    = "xxx"
	aesErrorPadding = "xxx"
	aesErrorKey     = "xxx"
	aesErrorIV      = "xxx"
)

func TestEncryptModeError_ByAes(t *testing.T) {
	cipher := NewCipher()
	cipher.SetMode(aesErrorMode)
	cipher.SetPadding(aesRightPadding)
	cipher.SetKey(aesRightKey)
	cipher.SetIV(aesRightIV)

	expected := invalidModeOrPaddingError(aesErrorMode, aesRightPadding)

	e1 := Encrypt.FromString(aesRightInput).ByAes(cipher)
	assert.Equal(t, expected, e1.Error, "Should catch an exception")

	e2 := Encrypt.FromBytes([]byte(aesRightInput)).ByAes(cipher)
	assert.Equal(t, expected, e2.Error, "Should catch an exception")
}

func TestDecryptModeError_ByAes(t *testing.T) {
	cipher := NewCipher()
	cipher.SetMode(aesErrorMode)
	cipher.SetPadding(aesRightPadding)
	cipher.SetKey([]byte(aesRightKey))
	cipher.SetIV([]byte(aesRightIV))

	expected := invalidModeOrPaddingError(aesErrorMode, aesRightPadding)

	d1 := Decrypt.FromString(aesRightInput).ByAes(cipher)
	assert.Equal(t, expected, d1.Error, "Should catch an exception")

	d2 := Decrypt.FromBytes([]byte(aesRightInput)).ByAes(cipher)
	assert.Equal(t, expected, d2.Error, "Should catch an exception")
}

func TestEncryptPaddingError_ByAes(t *testing.T) {
	cipher := NewCipher()
	cipher.SetMode(aesRightMode)
	cipher.SetPadding(aesErrorPadding)
	cipher.SetKey(aesRightKey)
	cipher.SetIV(aesRightIV)

	expected := invalidModeOrPaddingError(aesRightMode, aesErrorPadding)

	e1 := Encrypt.FromString(aesRightInput).ByAes(cipher)
	assert.Equal(t, expected, e1.Error, "Should catch an exception")

	e2 := Encrypt.FromBytes([]byte(aesRightInput)).ByAes(cipher)
	assert.Equal(t, expected, e2.Error, "Should catch an exception")
}

func TestDecryptPaddingError_ByAes(t *testing.T) {
	cipher := NewCipher()
	cipher.SetMode(aesRightMode)
	cipher.SetPadding(aesErrorPadding)
	cipher.SetKey([]byte(aesRightKey))
	cipher.SetIV([]byte(aesRightIV))

	expected := invalidModeOrPaddingError(aesRightMode, aesErrorPadding)

	d1 := Decrypt.FromString(aesRightInput).ByAes(cipher)
	assert.Equal(t, expected, d1.Error, "Should catch an exception")

	d2 := Decrypt.FromBytes([]byte(aesRightInput)).ByAes(cipher)
	assert.Equal(t, expected, d2.Error, "Should catch an exception")
}

func TestEncryptKeyError_ByAes(t *testing.T) {
	cipher := NewCipher()
	cipher.SetMode(CBC)
	cipher.SetPadding(aesRightPadding)
	cipher.SetKey(aesErrorKey)
	cipher.SetIV(aesRightIV)

	expected := invalidKeyError(3)

	e1 := Encrypt.FromString(aesRightInput).ByAes(cipher)
	assert.Equal(t, expected, e1.Error, "Should catch an exception")

	e2 := Encrypt.FromBytes([]byte(aesRightInput)).ByAes(cipher)
	assert.Equal(t, expected, e2.Error, "Should catch an exception")
}

func TestDecryptKeyError_ByAes(t *testing.T) {
	cipher := NewCipher()
	cipher.SetMode(CBC)
	cipher.SetPadding(aesRightPadding)
	cipher.SetKey([]byte(aesErrorKey))
	cipher.SetIV([]byte(aesRightIV))

	expected := invalidKeyError(len(aesErrorKey))

	d1 := Decrypt.FromString(aesRightInput).ByAes(cipher)
	assert.Equal(t, expected, d1.Error, "Should catch an exception")

	d2 := Decrypt.FromBytes([]byte(aesRightInput)).ByAes(cipher)
	assert.Equal(t, expected, d2.Error, "Should catch an exception")
}

func TestEncryptIVError_ByAes(t *testing.T) {
	cipher := NewCipher()
	cipher.SetMode(CBC)
	cipher.SetPadding(aesRightPadding)
	cipher.SetKey(aesRightKey)
	cipher.SetIV(aesErrorIV)

	expected := invalidIVError(len(aesErrorIV), 16)

	e1 := Encrypt.FromString(aesRightInput).ByAes(cipher)
	assert.Equal(t, expected, e1.Error, "Should catch an exception")

	e2 := Encrypt.FromBytes([]byte(aesRightInput)).ByAes(cipher)
	assert.Equal(t, expected, e2.Error, "Should catch an exception")
}

func TestDecryptIVError_ByAes(t *testing.T) {
	cipher := NewCipher()
	cipher.SetMode(CBC)
	cipher.SetPadding(aesRightPadding)
	cipher.SetKey([]byte(aesRightKey))
	cipher.SetIV([]byte(aesErrorIV))

	expected := invalidIVError(len(aesErrorIV), 16)

	d1 := Decrypt.FromString(aesRightInput).ByAes(cipher)
	assert.Equal(t, expected, d1.Error, "Should catch an exception")

	d2 := Decrypt.FromBytes([]byte(aesRightInput)).ByAes(cipher)
	assert.Equal(t, expected, d2.Error, "Should catch an exception")
}

func TestEncryptSrcError_ByAes(t *testing.T) {
	cipher := NewCipher()
	cipher.SetMode(CBC)
	cipher.SetPadding(No)
	cipher.SetKey(aesRightKey)
	cipher.SetIV(aesRightIV)

	expected := invalidSrcError(len(aesErrorInput))

	e1 := Encrypt.FromString(aesErrorInput).ByAes(cipher)
	assert.Equal(t, expected, e1.Error, "Should catch an exception")

	e2 := Encrypt.FromBytes([]byte(aesErrorInput)).ByAes(cipher)
	assert.Equal(t, expected, e2.Error, "Should catch an exception")
}

func TestDecryptSrcError_ByAes(t *testing.T) {
	cipher := NewCipher()
	cipher.SetMode(CBC)
	cipher.SetPadding(No)
	cipher.SetKey([]byte(aesRightKey))
	cipher.SetIV([]byte(aesRightIV))

	expected := invalidSrcError(len(aesErrorInput))

	d1 := Decrypt.FromString(aesErrorInput).ByAes(cipher)
	assert.Equal(t, expected, d1.Error, "Should catch an exception")

	d2 := Decrypt.FromBytes([]byte(aesErrorInput)).ByAes(cipher)
	assert.Equal(t, expected, d2.Error, "Should catch an exception")
}

func TestEncryptDecodeError_ByAes(t *testing.T) {
	cipher := NewCipher()
	cipher.SetMode(CBC)
	cipher.SetPadding(PKCS7)
	cipher.SetKey(aesRightKey)
	cipher.SetIV(aesRightIV)

	e1 := Decrypt.FromHexString(aesErrorInput).ByAes(cipher)
	assert.Equal(t, decodeSrcError("hex"), e1.Error, "Should catch an exception")

	e2 := Decrypt.FromBase32String(aesErrorInput).ByAes(cipher)
	assert.Equal(t, decodeSrcError("base32"), e2.Error, "Should catch an exception")

	e3 := Decrypt.FromBase64String(aesErrorInput).ByAes(cipher)
	assert.Equal(t, decodeSrcError("base64"), e3.Error, "Should catch an exception")
}

func TestDecryptDecodeError_ByAes(t *testing.T) {
	cipher := NewCipher()
	cipher.SetMode(CBC)
	cipher.SetPadding(PKCS7)
	cipher.SetKey([]byte(aesRightKey))
	cipher.SetIV([]byte(aesRightIV))

	d1 := Decrypt.FromHexBytes([]byte(aesErrorInput)).ByAes(cipher)
	assert.Equal(t, decodeSrcError("hex"), d1.Error, "Should catch an exception")

	d2 := Decrypt.FromBase32Bytes([]byte(aesErrorInput)).ByAes(cipher)
	assert.Equal(t, decodeSrcError("base32"), d2.Error, "Should catch an exception")

	d3 := Decrypt.FromBase64Bytes([]byte(aesErrorInput)).ByAes(cipher)
	assert.Equal(t, decodeSrcError("base64"), d3.Error, "Should catch an exception")
}
