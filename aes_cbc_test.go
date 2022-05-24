package dongle

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

var (
	aesCbcInput = "hello world"
	aesCbcKey   = "1234567887654321"
	aesCbcIV    = "1234567887654321"

	// AES-CBC-NonePadding
	aesCbcNoneHexExpected    = "fb5baf6e799b4da405d63b1ed3e14d62"
	aesCbcNoneBase32Expected = "7NN263TZTNG2IBOWHMPNHYKNMI======"
	aesCbcNoneBase64Expected = "+1uvbnmbTaQF1jse0+FNYg=="

	// AES-CBC-ZeroPadding
	aesCbcZeroHexExpected    = "5eb15d89da69f77372ee8d9bd02b0252"
	aesCbcZeroBase32Expected = "L2YV3CO2NH3XG4XORWN5AKYCKI======"
	aesCbcZeroBase64Expected = "XrFdidpp93Ny7o2b0CsCUg=="

	// AES-CBC-PKCS5Padding
	aesCbcPkcs5HexExpected    = "65d823bdf1c581a1ded1cba42e03ca52"
	aesCbcPkcs5Base32Expected = "MXMCHPPRYWA2DXWRZOSC4A6KKI======"
	aesCbcPkcs5Base64Expected = "ZdgjvfHFgaHe0cukLgPKUg=="

	// AES-CBC-PKCS7Padding
	aesCbcPkcs7HexExpected    = "65d823bdf1c581a1ded1cba42e03ca52"
	aesCbcPkcs7Base32Expected = "MXMCHPPRYWA2DXWRZOSC4A6KKI======"
	aesCbcPkcs7Base64Expected = "ZdgjvfHFgaHe0cukLgPKUg=="
)

func aesCbcCipher(padding string) *Cipher {
	cipher := NewCipher()
	cipher.SetMode(CBC)
	cipher.SetPadding(padding)
	cipher.SetKey(aesCbcKey)
	cipher.SetIV(aesCbcIV)
	return cipher
}

func TestEncrypt_ByAes_CBC_None(t *testing.T) {
	assert := assert.New(t)

	input := "12345678asdfghjk"

	hexTests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{input, aesCbcNoneHexExpected},
	}

	for index, test := range hexTests {
		e := Encrypt.FromString(test.input).ByAes(aesCbcCipher(None))
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToHexString(), "Current test index is "+strconv.Itoa(index))
	}

	base32Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{input, aesCbcNoneBase32Expected},
	}

	for index, test := range base32Tests {
		e := Encrypt.FromString(test.input).ByAes(aesCbcCipher(None))
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToBase32String(), "Current test index is "+strconv.Itoa(index))
	}

	base64Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{input, aesCbcNoneBase64Expected},
	}

	for index, test := range base64Tests {
		e := Encrypt.FromString(test.input).ByAes(aesCbcCipher(None))
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToBase64String(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestDecrypt_ByAes_CBC_None(t *testing.T) {
	assert := assert.New(t)

	input := "12345678asdfghjk"

	hexTests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{aesCbcNoneHexExpected, input},
	}

	for index, test := range hexTests {
		e := Decrypt.FromHexString(test.input).ByAes(aesCbcCipher(None))
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}

	base32Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{aesCbcNoneBase32Expected, input},
	}

	for index, test := range base32Tests {
		e := Decrypt.FromBase32String(test.input).ByAes(aesCbcCipher(None))
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}

	base64Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{aesCbcNoneBase64Expected, input},
	}

	for index, test := range base64Tests {
		e := Decrypt.FromBase64String(test.input).ByAes(aesCbcCipher(None))
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestEncrypt_ByAes_CBC_Zero(t *testing.T) {
	assert := assert.New(t)

	hexTests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{aesCbcInput, aesCbcZeroHexExpected},
	}

	for index, test := range hexTests {
		e := Encrypt.FromString(test.input).ByAes(aesCbcCipher(Zero))
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToHexString(), "Current test index is "+strconv.Itoa(index))
	}

	base32Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{aesCbcInput, aesCbcZeroBase32Expected},
	}

	for index, test := range base32Tests {
		e := Encrypt.FromString(test.input).ByAes(aesCbcCipher(Zero))
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToBase32String(), "Current test index is "+strconv.Itoa(index))
	}

	base64Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{aesCbcInput, aesCbcZeroBase64Expected},
	}

	for index, test := range base64Tests {
		e := Encrypt.FromString(test.input).ByAes(aesCbcCipher(Zero))
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToBase64String(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestDecrypt_ByAes_CBC_Zero(t *testing.T) {
	assert := assert.New(t)

	hexTests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{aesCbcZeroHexExpected, aesCbcInput},
	}

	for index, test := range hexTests {
		e := Decrypt.FromHexString(test.input).ByAes(aesCbcCipher(Zero))
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}

	base32Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{aesCbcZeroBase32Expected, aesCbcInput},
	}

	for index, test := range base32Tests {
		e := Decrypt.FromBase32String(test.input).ByAes(aesCbcCipher(Zero))
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}

	base64Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{aesCbcZeroBase64Expected, aesCbcInput},
	}

	for index, test := range base64Tests {
		e := Decrypt.FromBase64String(test.input).ByAes(aesCbcCipher(Zero))
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
		{aesCbcInput, aesCbcPkcs5HexExpected},
	}

	for index, test := range hexTests {
		e := Encrypt.FromString(test.input).ByAes(aesCbcCipher(PKCS5))
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToHexString(), "Current test index is "+strconv.Itoa(index))
	}

	base32Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{aesCbcInput, aesCbcPkcs5Base32Expected},
	}

	for index, test := range base32Tests {
		e := Encrypt.FromString(test.input).ByAes(aesCbcCipher(PKCS5))
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToBase32String(), "Current test index is "+strconv.Itoa(index))
	}

	base64Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{aesCbcInput, aesCbcPkcs5Base64Expected},
	}

	for index, test := range base64Tests {
		e := Encrypt.FromString(test.input).ByAes(aesCbcCipher(PKCS5))
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
		{aesCbcPkcs5HexExpected, aesCbcInput},
	}

	for index, test := range hexTests {
		e := Decrypt.FromHexString(test.input).ByAes(aesCbcCipher(PKCS5))
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}

	base32Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{aesCbcPkcs5Base32Expected, aesCbcInput},
	}

	for index, test := range base32Tests {
		e := Decrypt.FromBase32String(test.input).ByAes(aesCbcCipher(PKCS5))
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}

	base64Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{aesCbcPkcs5Base64Expected, aesCbcInput},
	}

	for index, test := range base64Tests {
		e := Decrypt.FromBase64String(test.input).ByAes(aesCbcCipher(PKCS5))
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
		{aesCbcInput, aesCbcPkcs7HexExpected},
	}

	for index, test := range hexTests {
		e := Encrypt.FromString(test.input).ByAes(aesCbcCipher(PKCS7))
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToHexString(), "Current test index is "+strconv.Itoa(index))
	}

	base32Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{aesCbcInput, aesCbcPkcs7Base32Expected},
	}

	for index, test := range base32Tests {
		e := Encrypt.FromString(test.input).ByAes(aesCbcCipher(PKCS7))
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToBase32String(), "Current test index is "+strconv.Itoa(index))
	}

	base64Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{aesCbcInput, aesCbcPkcs7Base64Expected},
	}

	for index, test := range base64Tests {
		e := Encrypt.FromString(test.input).ByAes(aesCbcCipher(PKCS7))
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
		{aesCbcPkcs7HexExpected, aesCbcInput},
	}

	for index, test := range hexTests {
		e := Decrypt.FromHexString(test.input).ByAes(aesCbcCipher(PKCS7))
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}

	base32Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{aesCbcPkcs7Base32Expected, aesCbcInput},
	}

	for index, test := range base32Tests {
		e := Decrypt.FromBase32String(test.input).ByAes(aesCbcCipher(PKCS7))
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}

	base64Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{aesCbcPkcs7Base64Expected, aesCbcInput},
	}

	for index, test := range base64Tests {
		e := Decrypt.FromBase64String(test.input).ByAes(aesCbcCipher(PKCS7))
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}
}
