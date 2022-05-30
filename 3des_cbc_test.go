package dongle

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

var (
	tripleDesCbcInput = "hello world"
	tripleDesCbcKey   = "123456781234567812345678"
	tripleDesCbcIV    = "12345678"

	// 3DES-CBC-NoPadding
	tripleDesCbcNoHexExpected    = "3d7595a98bff809d3ea6f8aa28affd7f"
	tripleDesCbcNoBase32Expected = "HV2ZLKML76AJ2PVG7CVCRL75P4======"
	tripleDesCbcNoBase64Expected = "PXWVqYv/gJ0+pviqKK/9fw=="

	// 3DES-CBC-ZeroPadding
	tripleDesCbcZeroHexExpected    = "0b2a92e81fb49ce1e02ecfa986df2ec8"
	tripleDesCbcZeroBase32Expected = "BMVJF2A7WSOODYBOZ6UYNXZOZA======"
	tripleDesCbcZeroBase64Expected = "CyqS6B+0nOHgLs+pht8uyA=="

	// 3DES-CBC-PKCS5Padding
	tripleDesCbcPkcs5HexExpected    = "0b2a92e81fb49ce1a43266aacaea7b81"
	tripleDesCbcPkcs5Base32Expected = "BMVJF2A7WSOODJBSM2VMV2T3QE======"
	tripleDesCbcPkcs5Base64Expected = "CyqS6B+0nOGkMmaqyup7gQ=="

	// 3DES-CBC-PKCS7Padding
	tripleDesCbcPkcs7HexExpected    = "0b2a92e81fb49ce1a43266aacaea7b81"
	tripleDesCbcPkcs7Base32Expected = "BMVJF2A7WSOODJBSM2VMV2T3QE======"
	tripleDesCbcPkcs7Base64Expected = "CyqS6B+0nOGkMmaqyup7gQ=="
)

func tripleDesCbcCipher(padding string) *Cipher {
	cipher := NewCipher()
	cipher.SetMode(CBC)
	cipher.SetPadding(padding)
	cipher.SetKey(tripleDesCbcKey)
	cipher.SetIV(tripleDesCbcIV)
	return cipher
}

func TestEncrypt_By3Des_CBC_NoPadding(t *testing.T) {
	input := "12345678asdfghjk"

	hexTests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{input, tripleDesCbcNoHexExpected},
	}

	for index, test := range hexTests {
		e := Encrypt.FromString(test.input).By3Des(tripleDesCbcCipher(No))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToHexString(), "Current test index is "+strconv.Itoa(index))
	}

	base32Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{input, tripleDesCbcNoBase32Expected},
	}

	for index, test := range base32Tests {
		e := Encrypt.FromString(test.input).By3Des(tripleDesCbcCipher(No))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToBase32String(), "Current test index is "+strconv.Itoa(index))
	}

	base64Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{input, tripleDesCbcNoBase64Expected},
	}

	for index, test := range base64Tests {
		e := Encrypt.FromString(test.input).By3Des(tripleDesCbcCipher(No))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToBase64String(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestDecrypt_By3Des_CBC_NoPadding(t *testing.T) {
	input := "12345678asdfghjk"

	hexTests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{tripleDesCbcNoHexExpected, input},
	}

	for index, test := range hexTests {
		e := Decrypt.FromHexString(test.input).By3Des(tripleDesCbcCipher(No))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}

	base32Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{tripleDesCbcNoBase32Expected, input},
	}

	for index, test := range base32Tests {
		e := Decrypt.FromBase32String(test.input).By3Des(tripleDesCbcCipher(No))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}

	base64Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{tripleDesCbcNoBase64Expected, input},
	}

	for index, test := range base64Tests {
		e := Decrypt.FromBase64String(test.input).By3Des(tripleDesCbcCipher(No))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestEncrypt_By3Des_CBC_ZeroPadding(t *testing.T) {
	hexTests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{tripleDesCbcInput, tripleDesCbcZeroHexExpected},
	}

	for index, test := range hexTests {
		e := Encrypt.FromString(test.input).By3Des(tripleDesCbcCipher(Zero))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToHexString(), "Current test index is "+strconv.Itoa(index))
	}

	base32Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{tripleDesCbcInput, tripleDesCbcZeroBase32Expected},
	}

	for index, test := range base32Tests {
		e := Encrypt.FromString(test.input).By3Des(tripleDesCbcCipher(Zero))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToBase32String(), "Current test index is "+strconv.Itoa(index))
	}

	base64Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{tripleDesCbcInput, tripleDesCbcZeroBase64Expected},
	}

	for index, test := range base64Tests {
		e := Encrypt.FromString(test.input).By3Des(tripleDesCbcCipher(Zero))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToBase64String(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestDecrypt_By3Des_CBC_ZeroPadding(t *testing.T) {
	hexTests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{tripleDesCbcZeroHexExpected, tripleDesCbcInput},
	}

	for index, test := range hexTests {
		e := Decrypt.FromHexString(test.input).By3Des(tripleDesCbcCipher(Zero))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}

	base32Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{tripleDesCbcZeroBase32Expected, tripleDesCbcInput},
	}

	for index, test := range base32Tests {
		e := Decrypt.FromBase32String(test.input).By3Des(tripleDesCbcCipher(Zero))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}

	base64Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{tripleDesCbcZeroBase64Expected, tripleDesCbcInput},
	}

	for index, test := range base64Tests {
		e := Decrypt.FromBase64String(test.input).By3Des(tripleDesCbcCipher(Zero))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestEncrypt_By3Des_CBC_PKCS5Padding(t *testing.T) {
	hexTests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{tripleDesCbcInput, tripleDesCbcPkcs5HexExpected},
	}

	for index, test := range hexTests {
		e := Encrypt.FromString(test.input).By3Des(tripleDesCbcCipher(PKCS5))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToHexString(), "Current test index is "+strconv.Itoa(index))
	}

	base32Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{tripleDesCbcInput, tripleDesCbcPkcs5Base32Expected},
	}

	for index, test := range base32Tests {
		e := Encrypt.FromString(test.input).By3Des(tripleDesCbcCipher(PKCS5))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToBase32String(), "Current test index is "+strconv.Itoa(index))
	}

	base64Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{tripleDesCbcInput, tripleDesCbcPkcs5Base64Expected},
	}

	for index, test := range base64Tests {
		e := Encrypt.FromString(test.input).By3Des(tripleDesCbcCipher(PKCS5))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToBase64String(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestDecrypt_By3Des_CBC_PKCS5Padding(t *testing.T) {
	hexTests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{tripleDesCbcPkcs5HexExpected, tripleDesCbcInput},
	}

	for index, test := range hexTests {
		e := Decrypt.FromHexString(test.input).By3Des(tripleDesCbcCipher(PKCS5))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}

	base32Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{tripleDesCbcPkcs5Base32Expected, tripleDesCbcInput},
	}

	for index, test := range base32Tests {
		e := Decrypt.FromBase32String(test.input).By3Des(tripleDesCbcCipher(PKCS5))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}

	base64Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{tripleDesCbcPkcs5Base64Expected, tripleDesCbcInput},
	}

	for index, test := range base64Tests {
		e := Decrypt.FromBase64String(test.input).By3Des(tripleDesCbcCipher(PKCS5))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestEncrypt_By3Des_CBC_PKCS7Padding(t *testing.T) {
	hexTests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{tripleDesCbcInput, tripleDesCbcPkcs7HexExpected},
	}

	for index, test := range hexTests {
		e := Encrypt.FromString(test.input).By3Des(tripleDesCbcCipher(PKCS7))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToHexString(), "Current test index is "+strconv.Itoa(index))
	}

	base32Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{tripleDesCbcInput, tripleDesCbcPkcs7Base32Expected},
	}

	for index, test := range base32Tests {
		e := Encrypt.FromString(test.input).By3Des(tripleDesCbcCipher(PKCS7))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToBase32String(), "Current test index is "+strconv.Itoa(index))
	}

	base64Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{tripleDesCbcInput, tripleDesCbcPkcs7Base64Expected},
	}

	for index, test := range base64Tests {
		e := Encrypt.FromString(test.input).By3Des(tripleDesCbcCipher(PKCS7))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToBase64String(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestDecrypt_By3Des_CBC_PKCS7Padding(t *testing.T) {
	hexTests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{tripleDesCbcPkcs7HexExpected, tripleDesCbcInput},
	}

	for index, test := range hexTests {
		e := Decrypt.FromHexString(test.input).By3Des(tripleDesCbcCipher(PKCS7))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}

	base32Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{tripleDesCbcPkcs7Base32Expected, tripleDesCbcInput},
	}

	for index, test := range base32Tests {
		e := Decrypt.FromBase32String(test.input).By3Des(tripleDesCbcCipher(PKCS7))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}

	base64Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{tripleDesCbcPkcs7Base64Expected, tripleDesCbcInput},
	}

	for index, test := range base64Tests {
		e := Decrypt.FromBase64String(test.input).By3Des(tripleDesCbcCipher(PKCS7))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}
}
