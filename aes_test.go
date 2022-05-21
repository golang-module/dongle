package dongle

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

var (
	aesInput = "hello world"
	aesKey   = "1234567887654321"
	aesIV    = "1234567887654321"
)

var (
	// AES-CBC-ZeroPadding
	aesCbcZeroHexExpected    = "5eb15d89da69f77372ee8d9bd02b0252"
	aesCbcZeroBase32Expected = "L2YV3CO2NH3XG4XORWN5AKYCKI======"
	aesCbcZeroBase64Expected = "XrFdidpp93Ny7o2b0CsCUg=="
	aesCbcZeroCipher         = func() *Cipher {
		cipher := NewCipher()
		cipher.SetMode(CBC)
		cipher.SetPadding(ZERO)
		cipher.SetKey(aesKey)
		cipher.SetIV(aesIV)
		return cipher
	}

	// AES-CBC-PKCS5Padding
	aesCbcPkcs5HexExpected    = "65d823bdf1c581a1ded1cba42e03ca52"
	aesCbcPkcs5Base32Expected = "MXMCHPPRYWA2DXWRZOSC4A6KKI======"
	aesCbcPkcs5Base64Expected = "ZdgjvfHFgaHe0cukLgPKUg=="
	aesCbcPkcs5Cipher         = func() *Cipher {
		cipher := NewCipher()
		cipher.SetMode(CBC)
		cipher.SetPadding(PKCS5)
		cipher.SetKey(aesKey)
		cipher.SetIV(aesIV)
		return cipher
	}

	// AES-CBC-PKCS7Padding
	aesCbcPkcs7HexExpected    = "65d823bdf1c581a1ded1cba42e03ca52"
	aesCbcPkcs7Base32Expected = "MXMCHPPRYWA2DXWRZOSC4A6KKI======"
	aesCbcPkcs7Base64Expected = "ZdgjvfHFgaHe0cukLgPKUg=="
	aesCbcPkcs7Cipher         = func() *Cipher {
		cipher := NewCipher()
		cipher.SetMode(CBC)
		cipher.SetPadding(PKCS7)
		cipher.SetKey([]byte(aesKey))
		cipher.SetIV([]byte(aesIV))
		return cipher
	}

	// AES-CFB-PKCS5Padding
	aesCfbPkcs5HexExpected    = "3fbb763723e1b2a11242f0aa880d7100"
	aesCfbPkcs5Base32Expected = "H65XMNZD4GZKCESC6CVIQDLRAA======"
	aesCfbPkcs5Base64Expected = "P7t2NyPhsqESQvCqiA1xAA=="
	aesCfbPkcs5Cipher         = func() *Cipher {
		cipher := NewCipher()
		cipher.SetMode(CFB)
		cipher.SetPadding(PKCS5)
		cipher.SetKey(aesKey)
		cipher.SetIV(aesIV)
		return cipher
	}

	// AES-CFB-PKCS7Padding
	aesCfbPkcs7HexExpected    = "3fbb763723e1b2a11242f0aa880d7100"
	aesCfbPkcs7Base32Expected = "H65XMNZD4GZKCESC6CVIQDLRAA======"
	aesCfbPkcs7Base64Expected = "P7t2NyPhsqESQvCqiA1xAA=="
	aesCfbPkcs7Cipher         = func() *Cipher {
		cipher := NewCipher()
		cipher.SetMode(CFB)
		cipher.SetPadding(PKCS7)
		cipher.SetKey(aesKey)
		cipher.SetIV(aesIV)
		return cipher
	}
)

func TestEncrypt_ByAes_CBC_ZERO(t *testing.T) {
	assert := assert.New(t)

	hexTests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{aesInput, aesCbcZeroHexExpected},
	}

	for index, test := range hexTests {
		e := Encrypt.FromString(test.input).ByAes(aesCbcZeroCipher())
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToHexString(), "Current test index is "+strconv.Itoa(index))
	}

	base32Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{aesInput, aesCbcZeroBase32Expected},
	}

	for index, test := range base32Tests {
		e := Encrypt.FromString(test.input).ByAes(aesCbcZeroCipher())
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToBase32String(), "Current test index is "+strconv.Itoa(index))
	}

	base64Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{aesInput, aesCbcZeroBase64Expected},
	}

	for index, test := range base64Tests {
		e := Encrypt.FromString(test.input).ByAes(aesCbcZeroCipher())
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToBase64String(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestDecrypt_ByAes_CBC_ZERO(t *testing.T) {
	assert := assert.New(t)

	hexTests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{aesCbcZeroHexExpected, aesInput},
	}

	for index, test := range hexTests {
		e := Decrypt.FromHexString(test.input).ByAes(aesCbcZeroCipher())
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}

	base32Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{aesCbcZeroBase32Expected, aesInput},
	}

	for index, test := range base32Tests {
		e := Decrypt.FromBase32String(test.input).ByAes(aesCbcZeroCipher())
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}

	base64Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{aesCbcZeroBase64Expected, aesInput},
	}

	for index, test := range base64Tests {
		e := Decrypt.FromBase64String(test.input).ByAes(aesCbcZeroCipher())
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestEncrypt_ByAes_CBC_PKCS5(t *testing.T) {
	assert := assert.New(t)

	hexTests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{aesInput, aesCbcPkcs5HexExpected},
	}

	for index, test := range hexTests {
		e := Encrypt.FromString(test.input).ByAes(aesCbcPkcs5Cipher())
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToHexString(), "Current test index is "+strconv.Itoa(index))
	}

	base32Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{aesInput, aesCbcPkcs5Base32Expected},
	}

	for index, test := range base32Tests {
		e := Encrypt.FromString(test.input).ByAes(aesCbcPkcs5Cipher())
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToBase32String(), "Current test index is "+strconv.Itoa(index))
	}

	base64Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{aesInput, aesCbcPkcs5Base64Expected},
	}

	for index, test := range base64Tests {
		e := Encrypt.FromString(test.input).ByAes(aesCbcPkcs5Cipher())
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToBase64String(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestDecrypt_ByAes_CBC_PKCS5(t *testing.T) {
	assert := assert.New(t)

	hexTests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{aesCbcPkcs5HexExpected, aesInput},
	}

	for index, test := range hexTests {
		e := Decrypt.FromHexString(test.input).ByAes(aesCbcPkcs5Cipher())
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}

	base32Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{aesCbcPkcs5Base32Expected, aesInput},
	}

	for index, test := range base32Tests {
		e := Decrypt.FromBase32String(test.input).ByAes(aesCbcPkcs5Cipher())
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}

	base64Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{aesCbcPkcs5Base64Expected, aesInput},
	}

	for index, test := range base64Tests {
		e := Decrypt.FromBase64String(test.input).ByAes(aesCbcPkcs5Cipher())
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestEncrypt_ByAes_CBC_PKCS7(t *testing.T) {
	assert := assert.New(t)

	hexTests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{aesInput, aesCbcPkcs7HexExpected},
	}

	for index, test := range hexTests {
		e := Encrypt.FromString(test.input).ByAes(aesCbcPkcs7Cipher())
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToHexString(), "Current test index is "+strconv.Itoa(index))
	}

	base32Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{aesInput, aesCbcPkcs7Base32Expected},
	}

	for index, test := range base32Tests {
		e := Encrypt.FromString(test.input).ByAes(aesCbcPkcs7Cipher())
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToBase32String(), "Current test index is "+strconv.Itoa(index))
	}

	base64Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{aesInput, aesCbcPkcs7Base64Expected},
	}

	for index, test := range base64Tests {
		e := Encrypt.FromString(test.input).ByAes(aesCbcPkcs7Cipher())
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToBase64String(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestDecrypt_ByAes_CBC_PKCS7(t *testing.T) {
	assert := assert.New(t)

	hexTests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{aesCbcPkcs7HexExpected, aesInput},
	}

	for index, test := range hexTests {
		e := Decrypt.FromHexString(test.input).ByAes(aesCbcPkcs7Cipher())
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}

	base32Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{aesCbcPkcs7Base32Expected, aesInput},
	}

	for index, test := range base32Tests {
		e := Decrypt.FromBase32String(test.input).ByAes(aesCbcPkcs7Cipher())
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}

	base64Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{aesCbcPkcs7Base64Expected, aesInput},
	}

	for index, test := range base64Tests {
		e := Decrypt.FromBase64String(test.input).ByAes(aesCbcPkcs7Cipher())
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}
}

// func TestEncrypt_ByAes_CFB_PKCS7(t *testing.T) {
// 	assert := assert.New(t)
//
// 	hexTests := []struct {
// 		input    string // 输入值
// 		expected string // 期望值
// 	}{
// 		{"", ""},
// 		{aesInput, aesCfbPkcs7HexExpected},
// 	}
//
// 	for index, test := range hexTests {
// 		e := Encrypt.FromString(test.input).ByAes(aesCfbPkcs7Cipher())
// 		assert.Nil(e.Error)
// 		assert.Equal(test.expected, e.ToHexString(), "Current test index is "+strconv.Itoa(index))
// 	}
//
// 	base32Tests := []struct {
// 		input    string // 输入值
// 		expected string // 期望值
// 	}{
// 		{"", ""},
// 		{aesInput, aesCfbPkcs7Base32Expected},
// 	}
//
// 	for index, test := range base32Tests {
// 		e := Encrypt.FromString(test.input).ByAes(aesCfbPkcs7Cipher())
// 		assert.Nil(e.Error)
// 		assert.Equal(test.expected, e.ToBase32String(), "Current test index is "+strconv.Itoa(index))
// 	}
//
// 	base64Tests := []struct {
// 		input    string // 输入值
// 		expected string // 期望值
// 	}{
// 		{"", ""},
// 		{aesInput, aesCfbPkcs7Base64Expected},
// 	}
//
// 	for index, test := range base64Tests {
// 		e := Encrypt.FromString(test.input).ByAes(aesCfbPkcs7Cipher())
// 		assert.Nil(e.Error)
// 		assert.Equal(test.expected, e.ToBase64String(), "Current test index is "+strconv.Itoa(index))
// 	}
// }
//
// func TestDecrypt_ByAes_CFB_PKCS5(t *testing.T) {
// 	assert := assert.New(t)
//
// 	hexTests := []struct {
// 		input    string // 输入值
// 		expected string // 期望值
// 	}{
// 		{"", ""},
// 		{aesInput, ""},
// 	}
//
// 	for index, test := range hexTests {
// 		e := Decrypt.FromHexString(test.input).ByAes(aesCfbPkcs5Cipher())
// 		assert.Nil(e.Error)
// 		assert.Equal(test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
// 	}
//
// 	base32Tests := []struct {
// 		input    string // 输入值
// 		expected string // 期望值
// 	}{
// 		{"", ""},
// 		{aesCbcPkcs5Base32Expected, aesInput},
// 	}
//
// 	for index, test := range base32Tests {
// 		e := Decrypt.FromBase32String(test.input).ByAes(aesCfbPkcs5Cipher())
// 		assert.Nil(e.Error)
// 		assert.Equal(test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
// 	}
//
// 	base64Tests := []struct {
// 		input    string // 输入值
// 		expected string // 期望值
// 	}{
// 		{"", ""},
// 		{aesCbcPkcs5Base64Expected, aesInput},
// 	}
//
// 	for index, test := range base64Tests {
// 		e := Decrypt.FromBase64String(test.input).ByAes(aesCfbPkcs5Cipher())
// 		assert.Nil(e.Error)
// 		assert.Equal(test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
// 	}
// }
//
// func TestEncrypt_ByAes_CFB_PKCS5(t *testing.T) {
// 	assert := assert.New(t)
//
// 	hexTests := []struct {
// 		input    string // 输入值
// 		expected string // 期望值
// 	}{
// 		{"", ""},
// 		{aesInput, aesCfbPkcs5HexExpected},
// 	}
//
// 	for index, test := range hexTests {
// 		e := Encrypt.FromString(test.input).ByAes(aesCfbPkcs5Cipher())
// 		assert.Nil(e.Error)
// 		assert.Equal(test.expected, e.ToHexString(), "Current test index is "+strconv.Itoa(index))
// 	}
//
// 	base32Tests := []struct {
// 		input    string // 输入值
// 		expected string // 期望值
// 	}{
// 		{"", ""},
// 		{aesInput, aesCfbPkcs5Base32Expected},
// 	}
//
// 	for index, test := range base32Tests {
// 		e := Encrypt.FromString(test.input).ByAes(aesCfbPkcs5Cipher())
// 		assert.Nil(e.Error)
// 		assert.Equal(test.expected, e.ToBase32String(), "Current test index is "+strconv.Itoa(index))
// 	}
//
// 	base64Tests := []struct {
// 		input    string // 输入值
// 		expected string // 期望值
// 	}{
// 		{"", ""},
// 		{aesInput, aesCfbPkcs5Base64Expected},
// 	}
//
// 	for index, test := range base64Tests {
// 		e := Encrypt.FromString(test.input).ByAes(aesCfbPkcs5Cipher())
// 		assert.Nil(e.Error)
// 		assert.Equal(test.expected, e.ToBase64String(), "Current test index is "+strconv.Itoa(index))
// 	}
// }
//
// func TestDecrypt_ByAes_CFB_PKCS7(t *testing.T) {
// 	assert := assert.New(t)
//
// 	hexTests := []struct {
// 		input    string // 输入值
// 		expected string // 期望值
// 	}{
// 		{"", ""},
// 		{aesInput, ""},
// 	}
//
// 	for index, test := range hexTests {
// 		e := Decrypt.FromHexString(test.input).ByAes(aesCfbPkcs7Cipher())
// 		assert.Nil(e.Error)
// 		assert.Equal(test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
// 	}
//
// 	base32Tests := []struct {
// 		input    string // 输入值
// 		expected string // 期望值
// 	}{
// 		{"", ""},
// 		{aesCbcPkcs7Base32Expected, aesInput},
// 	}
//
// 	for index, test := range base32Tests {
// 		e := Decrypt.FromBase32String(test.input).ByAes(aesCfbPkcs7Cipher())
// 		assert.Nil(e.Error)
// 		assert.Equal(test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
// 	}
//
// 	base64Tests := []struct {
// 		input    string // 输入值
// 		expected string // 期望值
// 	}{
// 		{"", ""},
// 		{aesCbcPkcs7Base64Expected, aesInput},
// 	}
//
// 	for index, test := range base64Tests {
// 		e := Decrypt.FromBase64String(test.input).ByAes(aesCfbPkcs7Cipher())
// 		assert.Nil(e.Error)
// 		assert.Equal(test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
// 	}
// }

func TestEncryptModeError_ByAes(t *testing.T) {
	cipher := NewCipher()
	cipher.SetMode("xxx")
	cipher.SetPadding(PKCS7)
	cipher.SetKey(aesKey)
	cipher.SetIV(aesIV)

	expected := invalidModeOrPaddingError("xxx", PKCS7)

	e1 := Encrypt.FromString(aesInput).ByAes(cipher)
	assert.Equal(t, expected, e1.Error, "Should catch an exception")

	e2 := Encrypt.FromBytes([]byte(aesInput)).ByAes(cipher)
	assert.Equal(t, expected, e2.Error, "Should catch an exception")
}

func TestDecryptModeError_ByAes(t *testing.T) {
	cipher := NewCipher()
	cipher.SetMode("xxx")
	cipher.SetPadding(PKCS7)
	cipher.SetKey(aesKey)
	cipher.SetIV(aesIV)

	expected := invalidModeOrPaddingError("xxx", PKCS7)

	d1 := Decrypt.FromHexString(aesCbcPkcs7HexExpected).ByAes(cipher)
	assert.Equal(t, expected, d1.Error, "Should catch an exception")

	d2 := Decrypt.FromHexBytes([]byte(aesCbcPkcs7HexExpected)).ByAes(cipher)
	assert.Equal(t, expected, d2.Error, "Should catch an exception")
}

func TestEncryptKeySizeError_ByAes(t *testing.T) {
	cipher := NewCipher()
	cipher.SetMode(CBC)
	cipher.SetPadding(PKCS7)
	cipher.SetKey("xxx")
	cipher.SetIV(aesIV)

	expected := invalidKeyError(3)

	e1 := Encrypt.FromString(aesInput).ByAes(cipher)
	assert.Equal(t, expected, e1.Error, "Should catch an exception")

	e2 := Encrypt.FromBytes([]byte(aesInput)).ByAes(cipher)
	assert.Equal(t, expected, e2.Error, "Should catch an exception")
}

func TestDecryptKeySizeError_ByAes(t *testing.T) {
	cipher := NewCipher()
	cipher.SetMode(CBC)
	cipher.SetPadding(PKCS7)
	cipher.SetKey("xxx")
	cipher.SetIV(aesIV)

	expected := invalidKeyError(3)

	d1 := Decrypt.FromBase32String(aesCbcPkcs7HexExpected).ByAes(cipher)
	assert.Equal(t, expected, d1.Error, "Should catch an exception")

	d2 := Decrypt.FromBase32Bytes([]byte(aesCbcPkcs7HexExpected)).ByAes(cipher)
	assert.Equal(t, expected, d2.Error, "Should catch an exception")
}

func TestEncryptIVSizeError_ByAes(t *testing.T) {
	cipher := NewCipher()
	cipher.SetMode(CBC)
	cipher.SetPadding(PKCS7)
	cipher.SetKey(aesKey)
	cipher.SetIV("xxx")

	expected := invalidIVError(3, 16)

	e1 := Encrypt.FromString(aesInput).ByAes(cipher)
	assert.Equal(t, expected, e1.Error, "Should catch an exception")

	e2 := Encrypt.FromBytes([]byte(aesInput)).ByAes(cipher)
	assert.Equal(t, expected, e2.Error, "Should catch an exception")
}

func TestDecryptIVSizeError_ByAes(t *testing.T) {
	cipher := NewCipher()
	cipher.SetMode(CBC)
	cipher.SetPadding(PKCS7)
	cipher.SetKey(aesKey)
	cipher.SetIV("xxx")

	expected := invalidIVError(3, 16)

	d1 := Decrypt.FromBase64String(aesCbcPkcs7HexExpected).ByAes(cipher)
	assert.Equal(t, expected, d1.Error, "Should catch an exception")

	d2 := Decrypt.FromBase64Bytes([]byte(aesCbcPkcs7HexExpected)).ByAes(cipher)
	assert.Equal(t, expected, d2.Error, "Should catch an exception")
}
