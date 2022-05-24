package dongle

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

var (
	aesOfbInput = "hello world"
	aesOfbKey   = "1234567887654321"
	aesOfbIV    = "1234567887654321"

	// AES-OFB-NonePadding
	aesOfbNoneHexExpected    = "66ec296f79f7f2f6015df0c9ea601e6e"
	aesOfbNoneBase32Expected = "M3WCS33Z67ZPMAK56DE6UYA6NY======"
	aesOfbNoneBase64Expected = "Zuwpb3n38vYBXfDJ6mAebg=="

	// AES-OFB-ZeroPadding
	aesOfbZeroHexExpected    = "3fbb763723e1b2a11242f0af8d087405"
	aesOfbZeroBase32Expected = "H65XMNZD4GZKCESC6CXY2CDUAU======"
	aesOfbZeroBase64Expected = "P7t2NyPhsqESQvCvjQh0BQ=="

	// AES-OFB-PKCS5Padding
	aesOfbPkcs5HexExpected    = "3fbb763723e1b2a11242f0aa880d7100"
	aesOfbPkcs5Base32Expected = "H65XMNZD4GZKCESC6CVIQDLRAA======"
	aesOfbPkcs5Base64Expected = "P7t2NyPhsqESQvCqiA1xAA=="

	// AES-OFB-PKCS7Padding
	aesOfbPkcs7HexExpected    = "3fbb763723e1b2a11242f0aa880d7100"
	aesOfbPkcs7Base32Expected = "H65XMNZD4GZKCESC6CVIQDLRAA======"
	aesOfbPkcs7Base64Expected = "P7t2NyPhsqESQvCqiA1xAA=="
)

func aesOfbCipher(padding string) *Cipher {
	cipher := NewCipher()
	cipher.SetMode(OFB)
	cipher.SetPadding(padding)
	cipher.SetKey(aesOfbKey)
	cipher.SetIV(aesOfbIV)
	return cipher
}

func TestEncrypt_ByAes_OFB_None(t *testing.T) {
	assert := assert.New(t)

	input := "12345678asdfghjk"

	hexTests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{input, aesOfbNoneHexExpected},
	}

	for index, test := range hexTests {
		e := Encrypt.FromString(test.input).ByAes(aesOfbCipher(None))
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToHexString(), "Current test index is "+strconv.Itoa(index))
	}

	base32Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{input, aesOfbNoneBase32Expected},
	}

	for index, test := range base32Tests {
		e := Encrypt.FromString(test.input).ByAes(aesOfbCipher(None))
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToBase32String(), "Current test index is "+strconv.Itoa(index))
	}

	base64Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{input, aesOfbNoneBase64Expected},
	}

	for index, test := range base64Tests {
		e := Encrypt.FromString(test.input).ByAes(aesOfbCipher(None))
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToBase64String(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestDecrypt_ByAes_OFB_None(t *testing.T) {
	assert := assert.New(t)

	input := "12345678asdfghjk"

	hexTests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{aesOfbNoneHexExpected, input},
	}

	for index, test := range hexTests {
		e := Decrypt.FromHexString(test.input).ByAes(aesOfbCipher(None))
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}

	base32Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{aesOfbNoneBase32Expected, input},
	}

	for index, test := range base32Tests {
		e := Decrypt.FromBase32String(test.input).ByAes(aesOfbCipher(None))
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}

	base64Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{aesOfbNoneBase64Expected, input},
	}

	for index, test := range base64Tests {
		e := Decrypt.FromBase64String(test.input).ByAes(aesOfbCipher(None))
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestEncrypt_ByAes_OFB_Zero(t *testing.T) {
	assert := assert.New(t)

	hexTests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{aesOfbInput, aesOfbZeroHexExpected},
	}

	for index, test := range hexTests {
		e := Encrypt.FromString(test.input).ByAes(aesOfbCipher(Zero))
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToHexString(), "Current test index is "+strconv.Itoa(index))
	}

	base32Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{aesOfbInput, aesOfbZeroBase32Expected},
	}

	for index, test := range base32Tests {
		e := Encrypt.FromString(test.input).ByAes(aesOfbCipher(Zero))
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToBase32String(), "Current test index is "+strconv.Itoa(index))
	}

	base64Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{aesOfbInput, aesOfbZeroBase64Expected},
	}

	for index, test := range base64Tests {
		e := Encrypt.FromString(test.input).ByAes(aesOfbCipher(Zero))
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToBase64String(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestDecrypt_ByAes_OFB_Zero(t *testing.T) {
	assert := assert.New(t)

	hexTests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{aesOfbZeroHexExpected, aesOfbInput},
	}

	for index, test := range hexTests {
		e := Decrypt.FromHexString(test.input).ByAes(aesOfbCipher(Zero))
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}

	base32Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{aesOfbZeroBase32Expected, aesOfbInput},
	}

	for index, test := range base32Tests {
		e := Decrypt.FromBase32String(test.input).ByAes(aesOfbCipher(Zero))
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}

	base64Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{aesOfbZeroBase64Expected, aesOfbInput},
	}

	for index, test := range base64Tests {
		e := Decrypt.FromBase64String(test.input).ByAes(aesOfbCipher(Zero))
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestEncrypt_ByAes_OFB_PKCS5(t *testing.T) {
	assert := assert.New(t)

	hexTests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{aesOfbInput, aesOfbPkcs5HexExpected},
	}

	for index, test := range hexTests {
		e := Encrypt.FromString(test.input).ByAes(aesOfbCipher(PKCS5))
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToHexString(), "Current test index is "+strconv.Itoa(index))
	}

	base32Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{aesOfbInput, aesOfbPkcs5Base32Expected},
	}

	for index, test := range base32Tests {
		e := Encrypt.FromString(test.input).ByAes(aesOfbCipher(PKCS5))
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToBase32String(), "Current test index is "+strconv.Itoa(index))
	}

	base64Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{aesOfbInput, aesOfbPkcs5Base64Expected},
	}

	for index, test := range base64Tests {
		e := Encrypt.FromString(test.input).ByAes(aesOfbCipher(PKCS5))
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToBase64String(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestDecrypt_ByAes_OFB_PKCS5(t *testing.T) {
	assert := assert.New(t)

	hexTests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{aesOfbPkcs5HexExpected, aesOfbInput},
	}

	for index, test := range hexTests {
		e := Decrypt.FromHexString(test.input).ByAes(aesOfbCipher(PKCS5))
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}

	base32Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{aesOfbPkcs5Base32Expected, aesOfbInput},
	}

	for index, test := range base32Tests {
		e := Decrypt.FromBase32String(test.input).ByAes(aesOfbCipher(PKCS5))
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}

	base64Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{aesOfbPkcs5Base64Expected, aesOfbInput},
	}

	for index, test := range base64Tests {
		e := Decrypt.FromBase64String(test.input).ByAes(aesOfbCipher(PKCS5))
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestEncrypt_ByAes_OFB_PKCS7(t *testing.T) {
	assert := assert.New(t)

	hexTests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{aesOfbInput, aesOfbPkcs7HexExpected},
	}

	for index, test := range hexTests {
		e := Encrypt.FromString(test.input).ByAes(aesOfbCipher(PKCS7))
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToHexString(), "Current test index is "+strconv.Itoa(index))
	}

	base32Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{aesOfbInput, aesOfbPkcs7Base32Expected},
	}

	for index, test := range base32Tests {
		e := Encrypt.FromString(test.input).ByAes(aesOfbCipher(PKCS7))
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToBase32String(), "Current test index is "+strconv.Itoa(index))
	}

	base64Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{aesOfbInput, aesOfbPkcs7Base64Expected},
	}

	for index, test := range base64Tests {
		e := Encrypt.FromString(test.input).ByAes(aesOfbCipher(PKCS7))
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToBase64String(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestDecrypt_ByAes_OFB_PKCS7(t *testing.T) {
	assert := assert.New(t)

	hexTests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{aesOfbPkcs7HexExpected, aesOfbInput},
	}

	for index, test := range hexTests {
		e := Decrypt.FromHexString(test.input).ByAes(aesOfbCipher(PKCS7))
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}

	base32Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{aesOfbPkcs7Base32Expected, aesOfbInput},
	}

	for index, test := range base32Tests {
		e := Decrypt.FromBase32String(test.input).ByAes(aesOfbCipher(PKCS7))
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}

	base64Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{aesOfbPkcs7Base64Expected, aesOfbInput},
	}

	for index, test := range base64Tests {
		e := Decrypt.FromBase64String(test.input).ByAes(aesOfbCipher(PKCS7))
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}
}
