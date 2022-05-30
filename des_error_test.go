package dongle

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	desRightInput   = "hello world"
	desRightMode    = CBC
	desRightPadding = PKCS7
	desRightKey     = "12345678"
	desRightIV      = "12345678"

	desErrorInput   = "xxx"
	desErrorMode    = "xxx"
	desErrorPadding = "xxx"
	desErrorKey     = "xxx"
	desErrorIV      = "xxx"
)

func TestEncryptModeError_ByDes(t *testing.T) {
	cipher := NewCipher()
	cipher.SetMode(desErrorMode)
	cipher.SetPadding(desRightPadding)
	cipher.SetKey(desRightKey)
	cipher.SetIV(desRightIV)

	expected := invalidModeOrPaddingError(desErrorMode, desRightPadding)

	e1 := Encrypt.FromString(aesRightInput).ByDes(cipher)
	assert.Equal(t, expected, e1.Error, "Should catch an exception")

	e2 := Encrypt.FromBytes([]byte(aesRightInput)).ByDes(cipher)
	assert.Equal(t, expected, e2.Error, "Should catch an exception")
}

func TestDecryptModeError_ByDes(t *testing.T) {
	cipher := NewCipher()
	cipher.SetMode(desErrorMode)
	cipher.SetPadding(desRightPadding)
	cipher.SetKey([]byte(desRightKey))
	cipher.SetIV([]byte(desRightIV))

	expected := invalidModeOrPaddingError(desErrorMode, desRightPadding)

	d1 := Decrypt.FromString(aesRightInput).ByDes(cipher)
	assert.Equal(t, expected, d1.Error, "Should catch an exception")

	d2 := Decrypt.FromBytes([]byte(aesRightInput)).ByDes(cipher)
	assert.Equal(t, expected, d2.Error, "Should catch an exception")
}

func TestEncryptPaddingError_ByDes(t *testing.T) {
	cipher := NewCipher()
	cipher.SetMode(aesRightMode)
	cipher.SetPadding(aesErrorPadding)
	cipher.SetKey(desRightKey)
	cipher.SetIV(desRightIV)

	expected := invalidModeOrPaddingError(aesRightMode, aesErrorPadding)

	e1 := Encrypt.FromString(aesRightInput).ByDes(cipher)
	assert.Equal(t, expected, e1.Error, "Should catch an exception")

	e2 := Encrypt.FromBytes([]byte(aesRightInput)).ByDes(cipher)
	assert.Equal(t, expected, e2.Error, "Should catch an exception")
}

func TestDecryptPaddingError_ByDes(t *testing.T) {
	cipher := NewCipher()
	cipher.SetMode(aesRightMode)
	cipher.SetPadding(aesErrorPadding)
	cipher.SetKey([]byte(desRightKey))
	cipher.SetIV([]byte(desRightIV))

	expected := invalidModeOrPaddingError(aesRightMode, aesErrorPadding)

	d1 := Decrypt.FromString(aesRightInput).ByDes(cipher)
	assert.Equal(t, expected, d1.Error, "Should catch an exception")

	d2 := Decrypt.FromBytes([]byte(aesRightInput)).ByDes(cipher)
	assert.Equal(t, expected, d2.Error, "Should catch an exception")
}

func TestEncryptKeyError_ByDes(t *testing.T) {
	cipher := NewCipher()
	cipher.SetMode(CBC)
	cipher.SetPadding(desRightPadding)
	cipher.SetKey(aesErrorKey)
	cipher.SetIV(desRightIV)

	expected := invalidDesKeyError(3)

	e1 := Encrypt.FromString(aesRightInput).ByDes(cipher)
	assert.Equal(t, expected, e1.Error, "Should catch an exception")

	e2 := Encrypt.FromBytes([]byte(aesRightInput)).ByDes(cipher)
	assert.Equal(t, expected, e2.Error, "Should catch an exception")
}

func TestDecryptKeyError_ByDes(t *testing.T) {
	cipher := NewCipher()
	cipher.SetMode(CBC)
	cipher.SetPadding(desRightPadding)
	cipher.SetKey([]byte(aesErrorKey))
	cipher.SetIV([]byte(desRightIV))

	expected := invalidDesKeyError(len(aesErrorKey))

	d1 := Decrypt.FromString(aesRightInput).ByDes(cipher)
	assert.Equal(t, expected, d1.Error, "Should catch an exception")

	d2 := Decrypt.FromBytes([]byte(aesRightInput)).ByDes(cipher)
	assert.Equal(t, expected, d2.Error, "Should catch an exception")
}

func TestEncryptIVError_ByDes(t *testing.T) {
	cipher := NewCipher()
	cipher.SetMode(CBC)
	cipher.SetPadding(desRightPadding)
	cipher.SetKey(desRightKey)
	cipher.SetIV(aesErrorIV)

	expected := invalidIVError(len(aesErrorIV), 8)

	e1 := Encrypt.FromString(aesRightInput).ByDes(cipher)
	assert.Equal(t, expected, e1.Error, "Should catch an exception")

	e2 := Encrypt.FromBytes([]byte(aesRightInput)).ByDes(cipher)
	assert.Equal(t, expected, e2.Error, "Should catch an exception")
}

func TestDecryptIVError_ByDes(t *testing.T) {
	cipher := NewCipher()
	cipher.SetMode(CBC)
	cipher.SetPadding(desRightPadding)
	cipher.SetKey([]byte(desRightKey))
	cipher.SetIV([]byte(aesErrorIV))

	expected := invalidIVError(len(aesErrorIV), 8)

	d1 := Decrypt.FromString(aesRightInput).ByDes(cipher)
	assert.Equal(t, expected, d1.Error, "Should catch an exception")

	d2 := Decrypt.FromBytes([]byte(aesRightInput)).ByDes(cipher)
	assert.Equal(t, expected, d2.Error, "Should catch an exception")
}

func TestEncryptSrcError_ByDes(t *testing.T) {
	cipher := NewCipher()
	cipher.SetMode(CBC)
	cipher.SetPadding(No)
	cipher.SetKey(desRightKey)
	cipher.SetIV(desRightIV)

	expected := invalidSrcError(len(aesErrorInput))

	e1 := Encrypt.FromString(aesErrorInput).ByDes(cipher)
	assert.Equal(t, expected, e1.Error, "Should catch an exception")

	e2 := Encrypt.FromBytes([]byte(aesErrorInput)).ByDes(cipher)
	assert.Equal(t, expected, e2.Error, "Should catch an exception")
}

func TestDecryptSrcError_ByDes(t *testing.T) {
	cipher := NewCipher()
	cipher.SetMode(CBC)
	cipher.SetPadding(No)
	cipher.SetKey([]byte(desRightKey))
	cipher.SetIV([]byte(desRightIV))

	expected := invalidSrcError(len(aesErrorInput))

	d1 := Decrypt.FromString(aesErrorInput).ByDes(cipher)
	assert.Equal(t, expected, d1.Error, "Should catch an exception")

	d2 := Decrypt.FromBytes([]byte(aesErrorInput)).ByDes(cipher)
	assert.Equal(t, expected, d2.Error, "Should catch an exception")
}

func TestEncryptDecodeError_ByDes(t *testing.T) {
	cipher := NewCipher()
	cipher.SetMode(CBC)
	cipher.SetPadding(PKCS7)
	cipher.SetKey(desRightKey)
	cipher.SetIV(desRightIV)

	e1 := Decrypt.FromHexString(aesErrorInput).ByDes(cipher)
	assert.Equal(t, decodeSrcError("hex"), e1.Error, "Should catch an exception")

	e2 := Decrypt.FromBase32String(aesErrorInput).ByDes(cipher)
	assert.Equal(t, decodeSrcError("base32"), e2.Error, "Should catch an exception")

	e3 := Decrypt.FromBase64String(aesErrorInput).ByDes(cipher)
	assert.Equal(t, decodeSrcError("base64"), e3.Error, "Should catch an exception")
}

func TestDecryptDecodeError_ByDes(t *testing.T) {
	cipher := NewCipher()
	cipher.SetMode(CBC)
	cipher.SetPadding(PKCS7)
	cipher.SetKey([]byte(desRightKey))
	cipher.SetIV([]byte(desRightIV))

	d1 := Decrypt.FromHexBytes([]byte(aesErrorInput)).ByDes(cipher)
	assert.Equal(t, decodeSrcError("hex"), d1.Error, "Should catch an exception")

	d2 := Decrypt.FromBase32Bytes([]byte(aesErrorInput)).ByDes(cipher)
	assert.Equal(t, decodeSrcError("base32"), d2.Error, "Should catch an exception")

	d3 := Decrypt.FromBase64Bytes([]byte(aesErrorInput)).ByDes(cipher)
	assert.Equal(t, decodeSrcError("base64"), d3.Error, "Should catch an exception")
}
