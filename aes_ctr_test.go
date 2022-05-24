package dongle

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

var (
	aesCtrInput = "hello world"
	aesCtrKey   = "1234567887654321"
	aesCtrIV    = "1234567887654321"

	// AES-CTR-NoPadding
	aesCtrNoHexExpected    = "66ec296f79f7f2f6015df0c9ea601e6e"
	aesCtrNoBase32Expected = "M3WCS33Z67ZPMAK56DE6UYA6NY======"
	aesCtrNoBase64Expected = "Zuwpb3n38vYBXfDJ6mAebg=="

	// AES-CTR-ZeroPadding
	aesCtrZeroHexExpected    = "3fbb763723e1b2a11242f0af8d087405"
	aesCtrZeroBase32Expected = "H65XMNZD4GZKCESC6CXY2CDUAU======"
	aesCtrZeroBase64Expected = "P7t2NyPhsqESQvCvjQh0BQ=="

	// AES-CTR-PKCS5Padding
	aesCtrPkcs5HexExpected    = "3fbb763723e1b2a11242f0aa880d7100"
	aesCtrPkcs5Base32Expected = "H65XMNZD4GZKCESC6CVIQDLRAA======"
	aesCtrPkcs5Base64Expected = "P7t2NyPhsqESQvCqiA1xAA=="

	// AES-CTR-PKCS7Padding
	aesCtrPkcs7HexExpected    = "3fbb763723e1b2a11242f0aa880d7100"
	aesCtrPkcs7Base32Expected = "H65XMNZD4GZKCESC6CVIQDLRAA======"
	aesCtrPkcs7Base64Expected = "P7t2NyPhsqESQvCqiA1xAA=="
)

func aesCtrCipher(padding string) *Cipher {
	cipher := NewCipher()
	cipher.SetMode(CTR)
	cipher.SetPadding(padding)
	cipher.SetKey(aesCtrKey)
	cipher.SetIV(aesCtrIV)
	return cipher
}

func TestEncrypt_ByAes_CTR_NoPadding(t *testing.T) {
	assert := assert.New(t)

	input := "12345678asdfghjk"

	hexTests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{input, aesCtrNoHexExpected},
	}

	for index, test := range hexTests {
		e := Encrypt.FromString(test.input).ByAes(aesCtrCipher(No))
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToHexString(), "Current test index is "+strconv.Itoa(index))
	}

	base32Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{input, aesCtrNoBase32Expected},
	}

	for index, test := range base32Tests {
		e := Encrypt.FromString(test.input).ByAes(aesCtrCipher(No))
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToBase32String(), "Current test index is "+strconv.Itoa(index))
	}

	base64Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{input, aesCtrNoBase64Expected},
	}

	for index, test := range base64Tests {
		e := Encrypt.FromString(test.input).ByAes(aesCtrCipher(No))
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToBase64String(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestDecrypt_ByAes_CTR_NoPadding(t *testing.T) {
	assert := assert.New(t)

	input := "12345678asdfghjk"

	hexTests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{aesCtrNoHexExpected, input},
	}

	for index, test := range hexTests {
		e := Decrypt.FromHexString(test.input).ByAes(aesCtrCipher(No))
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}

	base32Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{aesCtrNoBase32Expected, input},
	}

	for index, test := range base32Tests {
		e := Decrypt.FromBase32String(test.input).ByAes(aesCtrCipher(No))
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}

	base64Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{aesCtrNoBase64Expected, input},
	}

	for index, test := range base64Tests {
		e := Decrypt.FromBase64String(test.input).ByAes(aesCtrCipher(No))
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestEncrypt_ByAes_CTR_ZeroPadding(t *testing.T) {
	assert := assert.New(t)

	hexTests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{aesCtrInput, aesCtrZeroHexExpected},
	}

	for index, test := range hexTests {
		e := Encrypt.FromString(test.input).ByAes(aesCtrCipher(Zero))
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToHexString(), "Current test index is "+strconv.Itoa(index))
	}

	base32Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{aesCtrInput, aesCtrZeroBase32Expected},
	}

	for index, test := range base32Tests {
		e := Encrypt.FromString(test.input).ByAes(aesCtrCipher(Zero))
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToBase32String(), "Current test index is "+strconv.Itoa(index))
	}

	base64Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{aesCtrInput, aesCtrZeroBase64Expected},
	}

	for index, test := range base64Tests {
		e := Encrypt.FromString(test.input).ByAes(aesCtrCipher(Zero))
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToBase64String(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestDecrypt_ByAes_CTR_ZeroPadding(t *testing.T) {
	assert := assert.New(t)

	hexTests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{aesCtrZeroHexExpected, aesCtrInput},
	}

	for index, test := range hexTests {
		e := Decrypt.FromHexString(test.input).ByAes(aesCtrCipher(Zero))
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}

	base32Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{aesCtrZeroBase32Expected, aesCtrInput},
	}

	for index, test := range base32Tests {
		e := Decrypt.FromBase32String(test.input).ByAes(aesCtrCipher(Zero))
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}

	base64Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{aesCtrZeroBase64Expected, aesCtrInput},
	}

	for index, test := range base64Tests {
		e := Decrypt.FromBase64String(test.input).ByAes(aesCtrCipher(Zero))
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestEncrypt_ByAes_CTR_PKCS5Padding(t *testing.T) {
	assert := assert.New(t)

	hexTests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{aesCtrInput, aesCtrPkcs5HexExpected},
	}

	for index, test := range hexTests {
		e := Encrypt.FromString(test.input).ByAes(aesCtrCipher(PKCS5))
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToHexString(), "Current test index is "+strconv.Itoa(index))
	}

	base32Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{aesCtrInput, aesCtrPkcs5Base32Expected},
	}

	for index, test := range base32Tests {
		e := Encrypt.FromString(test.input).ByAes(aesCtrCipher(PKCS5))
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToBase32String(), "Current test index is "+strconv.Itoa(index))
	}

	base64Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{aesCtrInput, aesCtrPkcs5Base64Expected},
	}

	for index, test := range base64Tests {
		e := Encrypt.FromString(test.input).ByAes(aesCtrCipher(PKCS5))
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToBase64String(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestDecrypt_ByAes_CTR_PKCS5Padding(t *testing.T) {
	assert := assert.New(t)

	hexTests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{aesCtrPkcs5HexExpected, aesCtrInput},
	}

	for index, test := range hexTests {
		e := Decrypt.FromHexString(test.input).ByAes(aesCtrCipher(PKCS5))
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}

	base32Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{aesCtrPkcs5Base32Expected, aesCtrInput},
	}

	for index, test := range base32Tests {
		e := Decrypt.FromBase32String(test.input).ByAes(aesCtrCipher(PKCS5))
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}

	base64Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{aesCtrPkcs5Base64Expected, aesCtrInput},
	}

	for index, test := range base64Tests {
		e := Decrypt.FromBase64String(test.input).ByAes(aesCtrCipher(PKCS5))
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestEncrypt_ByAes_CTR_PKCS7Padding(t *testing.T) {
	assert := assert.New(t)

	hexTests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{aesCtrInput, aesCtrPkcs7HexExpected},
	}

	for index, test := range hexTests {
		e := Encrypt.FromString(test.input).ByAes(aesCtrCipher(PKCS7))
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToHexString(), "Current test index is "+strconv.Itoa(index))
	}

	base32Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{aesCtrInput, aesCtrPkcs7Base32Expected},
	}

	for index, test := range base32Tests {
		e := Encrypt.FromString(test.input).ByAes(aesCtrCipher(PKCS7))
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToBase32String(), "Current test index is "+strconv.Itoa(index))
	}

	base64Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{aesCtrInput, aesCtrPkcs7Base64Expected},
	}

	for index, test := range base64Tests {
		e := Encrypt.FromString(test.input).ByAes(aesCtrCipher(PKCS7))
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToBase64String(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestDecrypt_ByAes_CTR_PKCS7Padding(t *testing.T) {
	assert := assert.New(t)

	hexTests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{aesCtrPkcs7HexExpected, aesCtrInput},
	}

	for index, test := range hexTests {
		e := Decrypt.FromHexString(test.input).ByAes(aesCtrCipher(PKCS7))
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}

	base32Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{aesCtrPkcs7Base32Expected, aesCtrInput},
	}

	for index, test := range base32Tests {
		e := Decrypt.FromBase32String(test.input).ByAes(aesCtrCipher(PKCS7))
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}

	base64Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{aesCtrPkcs7Base64Expected, aesCtrInput},
	}

	for index, test := range base64Tests {
		e := Decrypt.FromBase64String(test.input).ByAes(aesCtrCipher(PKCS7))
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}
}
