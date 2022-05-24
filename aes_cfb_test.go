package dongle

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

var (
	aesCfbInput = "hello world"
	aesCfbKey   = "1234567887654321"
	aesCfbIV    = "1234567887654321"

	// AES-CFB-NonePadding
	aesCfbNoneHexExpected    = "66ec296f79f7f2f6015df0c9ea601e6e"
	aesCfbNoneBase32Expected = "M3WCS33Z67ZPMAK56DE6UYA6NY======"
	aesCfbNoneBase64Expected = "Zuwpb3n38vYBXfDJ6mAebg=="

	// AES-CFB-ZeroPadding
	aesCfbZeroHexExpected    = "3fbb763723e1b2a11242f0af8d087405"
	aesCfbZeroBase32Expected = "H65XMNZD4GZKCESC6CXY2CDUAU======"
	aesCfbZeroBase64Expected = "P7t2NyPhsqESQvCvjQh0BQ=="

	// AES-CFB-PKCS5Padding
	aesCfbPkcs5HexExpected    = "3fbb763723e1b2a11242f0aa880d7100"
	aesCfbPkcs5Base32Expected = "H65XMNZD4GZKCESC6CVIQDLRAA======"
	aesCfbPkcs5Base64Expected = "P7t2NyPhsqESQvCqiA1xAA=="

	// AES-CFB-PKCS7Padding
	aesCfbPkcs7HexExpected    = "3fbb763723e1b2a11242f0aa880d7100"
	aesCfbPkcs7Base32Expected = "H65XMNZD4GZKCESC6CVIQDLRAA======"
	aesCfbPkcs7Base64Expected = "P7t2NyPhsqESQvCqiA1xAA=="
)

func aesCfbCipher(padding string) *Cipher {
	cipher := NewCipher()
	cipher.SetMode(CFB)
	cipher.SetPadding(padding)
	cipher.SetKey(aesCfbKey)
	cipher.SetIV(aesCfbIV)
	return cipher
}

func TestEncrypt_ByAes_CFB_None(t *testing.T) {
	assert := assert.New(t)

	input := "12345678asdfghjk"

	hexTests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{input, aesCfbNoneHexExpected},
	}

	for index, test := range hexTests {
		e := Encrypt.FromString(test.input).ByAes(aesCfbCipher(None))
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToHexString(), "Current test index is "+strconv.Itoa(index))
	}

	base32Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{input, aesCfbNoneBase32Expected},
	}

	for index, test := range base32Tests {
		e := Encrypt.FromString(test.input).ByAes(aesCfbCipher(None))
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToBase32String(), "Current test index is "+strconv.Itoa(index))
	}

	base64Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{input, aesCfbNoneBase64Expected},
	}

	for index, test := range base64Tests {
		e := Encrypt.FromString(test.input).ByAes(aesCfbCipher(None))
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToBase64String(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestDecrypt_ByAes_CFB_None(t *testing.T) {
	assert := assert.New(t)

	input := "12345678asdfghjk"

	hexTests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{aesCfbNoneHexExpected, input},
	}

	for index, test := range hexTests {
		e := Decrypt.FromHexString(test.input).ByAes(aesCfbCipher(None))
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}

	base32Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{aesCfbNoneBase32Expected, input},
	}

	for index, test := range base32Tests {
		e := Decrypt.FromBase32String(test.input).ByAes(aesCfbCipher(None))
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}

	base64Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{aesCfbNoneBase64Expected, input},
	}

	for index, test := range base64Tests {
		e := Decrypt.FromBase64String(test.input).ByAes(aesCfbCipher(None))
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestEncrypt_ByAes_CFB_Zero(t *testing.T) {
	assert := assert.New(t)

	hexTests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{aesCfbInput, aesCfbZeroHexExpected},
	}

	for index, test := range hexTests {
		e := Encrypt.FromString(test.input).ByAes(aesCfbCipher(Zero))
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToHexString(), "Current test index is "+strconv.Itoa(index))
	}

	base32Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{aesCfbInput, aesCfbZeroBase32Expected},
	}

	for index, test := range base32Tests {
		e := Encrypt.FromString(test.input).ByAes(aesCfbCipher(Zero))
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToBase32String(), "Current test index is "+strconv.Itoa(index))
	}

	base64Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{aesCfbInput, aesCfbZeroBase64Expected},
	}

	for index, test := range base64Tests {
		e := Encrypt.FromString(test.input).ByAes(aesCfbCipher(Zero))
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToBase64String(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestDecrypt_ByAes_CFB_Zero(t *testing.T) {
	assert := assert.New(t)

	hexTests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{aesCfbZeroHexExpected, aesCfbInput},
	}

	for index, test := range hexTests {
		e := Decrypt.FromHexString(test.input).ByAes(aesCfbCipher(Zero))
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}

	base32Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{aesCfbZeroBase32Expected, aesCfbInput},
	}

	for index, test := range base32Tests {
		e := Decrypt.FromBase32String(test.input).ByAes(aesCfbCipher(Zero))
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}

	base64Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{aesCfbZeroBase64Expected, aesCfbInput},
	}

	for index, test := range base64Tests {
		e := Decrypt.FromBase64String(test.input).ByAes(aesCfbCipher(Zero))
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestEncrypt_ByAes_CFB_PKCS5(t *testing.T) {
	assert := assert.New(t)

	hexTests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{aesCfbInput, aesCfbPkcs5HexExpected},
	}

	for index, test := range hexTests {
		e := Encrypt.FromString(test.input).ByAes(aesCfbCipher(PKCS5))
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToHexString(), "Current test index is "+strconv.Itoa(index))
	}

	base32Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{aesCfbInput, aesCfbPkcs5Base32Expected},
	}

	for index, test := range base32Tests {
		e := Encrypt.FromString(test.input).ByAes(aesCfbCipher(PKCS5))
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToBase32String(), "Current test index is "+strconv.Itoa(index))
	}

	base64Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{aesCfbInput, aesCfbPkcs5Base64Expected},
	}

	for index, test := range base64Tests {
		e := Encrypt.FromString(test.input).ByAes(aesCfbCipher(PKCS5))
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToBase64String(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestDecrypt_ByAes_CFB_PKCS5(t *testing.T) {
	assert := assert.New(t)

	hexTests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{aesCfbPkcs5HexExpected, aesCfbInput},
	}

	for index, test := range hexTests {
		e := Decrypt.FromHexString(test.input).ByAes(aesCfbCipher(PKCS5))
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}

	base32Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{aesCfbPkcs5Base32Expected, aesCfbInput},
	}

	for index, test := range base32Tests {
		e := Decrypt.FromBase32String(test.input).ByAes(aesCfbCipher(PKCS5))
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}

	base64Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{aesCfbPkcs5Base64Expected, aesCfbInput},
	}

	for index, test := range base64Tests {
		e := Decrypt.FromBase64String(test.input).ByAes(aesCfbCipher(PKCS5))
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestEncrypt_ByAes_CFB_PKCS7(t *testing.T) {
	assert := assert.New(t)

	hexTests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{aesCfbInput, aesCfbPkcs7HexExpected},
	}

	for index, test := range hexTests {
		e := Encrypt.FromString(test.input).ByAes(aesCfbCipher(PKCS7))
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToHexString(), "Current test index is "+strconv.Itoa(index))
	}

	base32Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{aesCfbInput, aesCfbPkcs7Base32Expected},
	}

	for index, test := range base32Tests {
		e := Encrypt.FromString(test.input).ByAes(aesCfbCipher(PKCS7))
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToBase32String(), "Current test index is "+strconv.Itoa(index))
	}

	base64Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{aesCfbInput, aesCfbPkcs7Base64Expected},
	}

	for index, test := range base64Tests {
		e := Encrypt.FromString(test.input).ByAes(aesCfbCipher(PKCS7))
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToBase64String(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestDecrypt_ByAes_CFB_PKCS7(t *testing.T) {
	assert := assert.New(t)

	hexTests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{aesCfbPkcs7HexExpected, aesCfbInput},
	}

	for index, test := range hexTests {
		e := Decrypt.FromHexString(test.input).ByAes(aesCfbCipher(PKCS7))
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}

	base32Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{aesCfbPkcs7Base32Expected, aesCfbInput},
	}

	for index, test := range base32Tests {
		e := Decrypt.FromBase32String(test.input).ByAes(aesCfbCipher(PKCS7))
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}

	base64Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{aesCfbPkcs7Base64Expected, aesCfbInput},
	}

	for index, test := range base64Tests {
		e := Decrypt.FromBase64String(test.input).ByAes(aesCfbCipher(PKCS7))
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}
}
