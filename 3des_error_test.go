package dongle

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	tripleDesRightInput   = "hello world"
	tripleDesRightMode    = CBC
	tripleDesRightPadding = PKCS7
	tripleDesRightKey     = "123456781234567812345678"
	tripleDesRightIV      = "12345678"

	tripleDesErrorInput   = "xxx"
	tripleDesErrorMode    = "xxx"
	tripleDesErrorPadding = "xxx"
	tripleDesErrorKey     = "xxx"
	tripleDesErrorIV      = "xxx"
)

func TestEncryptModeError_By3Des(t *testing.T) {
	cipher := NewCipher()
	cipher.SetMode(tripleDesErrorMode)
	cipher.SetPadding(tripleDesRightPadding)
	cipher.SetKey(tripleDesRightKey)
	cipher.SetIV(tripleDesRightIV)

	expected := invalidModeOrPaddingError(tripleDesErrorMode, tripleDesRightPadding)

	e1 := Encrypt.FromString(tripleDesRightInput).By3Des(cipher)
	assert.Equal(t, expected, e1.Error, "Should catch an exception")

	e2 := Encrypt.FromBytes([]byte(tripleDesRightInput)).By3Des(cipher)
	assert.Equal(t, expected, e2.Error, "Should catch an exception")
}

func TestDecryptModeError_By3Des(t *testing.T) {
	cipher := NewCipher()
	cipher.SetMode(tripleDesErrorMode)
	cipher.SetPadding(tripleDesRightPadding)
	cipher.SetKey([]byte(tripleDesRightKey))
	cipher.SetIV([]byte(tripleDesRightIV))

	expected := invalidModeOrPaddingError(tripleDesErrorMode, tripleDesRightPadding)

	d1 := Decrypt.FromString(tripleDesRightInput).By3Des(cipher)
	assert.Equal(t, expected, d1.Error, "Should catch an exception")

	d2 := Decrypt.FromBytes([]byte(tripleDesRightInput)).By3Des(cipher)
	assert.Equal(t, expected, d2.Error, "Should catch an exception")
}

func TestEncryptPaddingError_By3Des(t *testing.T) {
	cipher := NewCipher()
	cipher.SetMode(tripleDesRightMode)
	cipher.SetPadding(tripleDesErrorPadding)
	cipher.SetKey(tripleDesRightKey)
	cipher.SetIV(tripleDesRightIV)

	expected := invalidModeOrPaddingError(tripleDesRightMode, tripleDesErrorPadding)

	e1 := Encrypt.FromString(tripleDesRightInput).By3Des(cipher)
	assert.Equal(t, expected, e1.Error, "Should catch an exception")

	e2 := Encrypt.FromBytes([]byte(tripleDesRightInput)).By3Des(cipher)
	assert.Equal(t, expected, e2.Error, "Should catch an exception")
}

func TestDecryptPaddingError_By3Des(t *testing.T) {
	cipher := NewCipher()
	cipher.SetMode(tripleDesRightMode)
	cipher.SetPadding(tripleDesErrorPadding)
	cipher.SetKey([]byte(tripleDesRightKey))
	cipher.SetIV([]byte(tripleDesRightIV))

	expected := invalidModeOrPaddingError(tripleDesRightMode, tripleDesErrorPadding)

	d1 := Decrypt.FromString(tripleDesRightInput).By3Des(cipher)
	assert.Equal(t, expected, d1.Error, "Should catch an exception")

	d2 := Decrypt.FromBytes([]byte(tripleDesRightInput)).By3Des(cipher)
	assert.Equal(t, expected, d2.Error, "Should catch an exception")
}

func TestEncryptKeyError_By3Des(t *testing.T) {
	cipher := NewCipher()
	cipher.SetMode(CBC)
	cipher.SetPadding(tripleDesRightPadding)
	cipher.SetKey(tripleDesErrorKey)
	cipher.SetIV(tripleDesRightIV)

	expected := invalid3DesKeyError(3)

	e1 := Encrypt.FromString(tripleDesRightInput).By3Des(cipher)
	assert.Equal(t, expected, e1.Error, "Should catch an exception")

	e2 := Encrypt.FromBytes([]byte(tripleDesRightInput)).By3Des(cipher)
	assert.Equal(t, expected, e2.Error, "Should catch an exception")
}

func TestDecryptKeyError_By3Des(t *testing.T) {
	cipher := NewCipher()
	cipher.SetMode(CBC)
	cipher.SetPadding(tripleDesRightPadding)
	cipher.SetKey([]byte(tripleDesErrorKey))
	cipher.SetIV([]byte(tripleDesRightIV))

	expected := invalid3DesKeyError(len(tripleDesErrorKey))

	d1 := Decrypt.FromString(tripleDesRightInput).By3Des(cipher)
	assert.Equal(t, expected, d1.Error, "Should catch an exception")

	d2 := Decrypt.FromBytes([]byte(tripleDesRightInput)).By3Des(cipher)
	assert.Equal(t, expected, d2.Error, "Should catch an exception")
}

func TestEncryptIVError_By3Des(t *testing.T) {
	cipher := NewCipher()
	cipher.SetMode(CBC)
	cipher.SetPadding(tripleDesRightPadding)
	cipher.SetKey(tripleDesRightKey)
	cipher.SetIV(tripleDesErrorIV)

	expected := invalidIVError(len(tripleDesErrorIV), 8)

	e1 := Encrypt.FromString(tripleDesRightInput).By3Des(cipher)
	assert.Equal(t, expected, e1.Error, "Should catch an exception")

	e2 := Encrypt.FromBytes([]byte(tripleDesRightInput)).By3Des(cipher)
	assert.Equal(t, expected, e2.Error, "Should catch an exception")
}

func TestDecryptIVError_By3Des(t *testing.T) {
	cipher := NewCipher()
	cipher.SetMode(CBC)
	cipher.SetPadding(tripleDesRightPadding)
	cipher.SetKey([]byte(tripleDesRightKey))
	cipher.SetIV([]byte(tripleDesErrorIV))

	expected := invalidIVError(len(tripleDesErrorIV), 8)

	d1 := Decrypt.FromString(tripleDesRightInput).By3Des(cipher)
	assert.Equal(t, expected, d1.Error, "Should catch an exception")

	d2 := Decrypt.FromBytes([]byte(tripleDesRightInput)).By3Des(cipher)
	assert.Equal(t, expected, d2.Error, "Should catch an exception")
}

func TestEncryptSrcError_By3Des(t *testing.T) {
	cipher := NewCipher()
	cipher.SetMode(CBC)
	cipher.SetPadding(No)
	cipher.SetKey(tripleDesRightKey)
	cipher.SetIV(tripleDesRightIV)

	expected := invalidSrcError(len(tripleDesErrorInput))

	e1 := Encrypt.FromString(tripleDesErrorInput).By3Des(cipher)
	assert.Equal(t, expected, e1.Error, "Should catch an exception")

	e2 := Encrypt.FromBytes([]byte(tripleDesErrorInput)).By3Des(cipher)
	assert.Equal(t, expected, e2.Error, "Should catch an exception")
}

func TestDecryptSrcError_By3Des(t *testing.T) {
	cipher := NewCipher()
	cipher.SetMode(CBC)
	cipher.SetPadding(No)
	cipher.SetKey([]byte(tripleDesRightKey))
	cipher.SetIV([]byte(tripleDesRightIV))

	expected := invalidSrcError(len(tripleDesErrorInput))

	d1 := Decrypt.FromString(tripleDesErrorInput).By3Des(cipher)
	assert.Equal(t, expected, d1.Error, "Should catch an exception")

	d2 := Decrypt.FromBytes([]byte(tripleDesErrorInput)).By3Des(cipher)
	assert.Equal(t, expected, d2.Error, "Should catch an exception")
}

func TestEncryptDecodeError_By3Des(t *testing.T) {
	cipher := NewCipher()
	cipher.SetMode(CBC)
	cipher.SetPadding(PKCS7)
	cipher.SetKey(tripleDesRightKey)
	cipher.SetIV(tripleDesRightIV)

	e1 := Decrypt.FromHexString(tripleDesErrorInput).By3Des(cipher)
	assert.Equal(t, decodeSrcError("hex"), e1.Error, "Should catch an exception")

	e2 := Decrypt.FromBase32String(tripleDesErrorInput).By3Des(cipher)
	assert.Equal(t, decodeSrcError("base32"), e2.Error, "Should catch an exception")

	e3 := Decrypt.FromBase64String(tripleDesErrorInput).By3Des(cipher)
	assert.Equal(t, decodeSrcError("base64"), e3.Error, "Should catch an exception")
}

func TestDecryptDecodeError_By3Des(t *testing.T) {
	cipher := NewCipher()
	cipher.SetMode(CBC)
	cipher.SetPadding(PKCS7)
	cipher.SetKey([]byte(tripleDesRightKey))
	cipher.SetIV([]byte(tripleDesRightIV))

	d1 := Decrypt.FromHexBytes([]byte(tripleDesErrorInput)).By3Des(cipher)
	assert.Equal(t, decodeSrcError("hex"), d1.Error, "Should catch an exception")

	d2 := Decrypt.FromBase32Bytes([]byte(tripleDesErrorInput)).By3Des(cipher)
	assert.Equal(t, decodeSrcError("base32"), d2.Error, "Should catch an exception")

	d3 := Decrypt.FromBase64Bytes([]byte(tripleDesErrorInput)).By3Des(cipher)
	assert.Equal(t, decodeSrcError("base64"), d3.Error, "Should catch an exception")
}
