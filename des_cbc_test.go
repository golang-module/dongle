package dongle

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

var (
	desCbcInput = "hello world"
	desCbcKey   = "12345678"
	desCbcIV    = "12345678"

	// DES-CBC-NoPadding
	desCbcNoHexExpected = "3d7595a98bff809d3ea6f8aa28affd7f"
	// desCbcNoHexExpected    = "3d7595a98bff809d3ea6f8aa28affd7f" // https://oktools.net/des
	// desCbcNoHexExpected    = "3d7595a98bff809d3ea6f8aa28affd7f" // http://tool.chacuo.net/cryptdes
	desCbcNoBase32Expected = "HV2ZLKML76AJ2PVG7CVCRL75P4======"
	desCbcNoBase64Expected = "PXWVqYv/gJ0+pviqKK/9fw=="

	// DES-CBC-ZeroPadding
	desCbcZeroHexExpected    = "0b2a92e81fb49ce1e02ecfa986df2ec8"
	desCbcZeroBase32Expected = "BMVJF2A7WSOODYBOZ6UYNXZOZA======"
	desCbcZeroBase64Expected = "CyqS6B+0nOHgLs+pht8uyA=="

	// DES-CBC-PKCS5Padding
	desCbcPkcs5HexExpected    = "0b2a92e81fb49ce1a43266aacaea7b81"
	desCbcPkcs5Base32Expected = "BMVJF2A7WSOODJBSM2VMV2T3QE======"
	desCbcPkcs5Base64Expected = "CyqS6B+0nOGkMmaqyup7gQ=="

	// DES-CBC-PKCS7Padding
	desCbcPkcs7HexExpected    = "0b2a92e81fb49ce1a43266aacaea7b81"
	desCbcPkcs7Base32Expected = "BMVJF2A7WSOODJBSM2VMV2T3QE======"
	desCbcPkcs7Base64Expected = "CyqS6B+0nOGkMmaqyup7gQ=="
)

func desCbcCipher(padding string) *Cipher {
	cipher := NewCipher()
	cipher.SetMode(CBC)
	cipher.SetPadding(padding)
	cipher.SetKey(desCbcKey)
	cipher.SetIV(desCbcIV)
	return cipher
}

func TestEncrypt_ByDes_CBC_NoPadding(t *testing.T) {
	input := "12345678asdfghjk"

	hexTests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{input, desCbcNoHexExpected},
	}

	for index, test := range hexTests {
		e := Encrypt.FromString(test.input).ByDes(desCbcCipher(No))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToHexString(), "Current test index is "+strconv.Itoa(index))
	}

	base32Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{input, desCbcNoBase32Expected},
	}

	for index, test := range base32Tests {
		e := Encrypt.FromString(test.input).ByDes(desCbcCipher(No))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToBase32String(), "Current test index is "+strconv.Itoa(index))
	}

	base64Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{input, desCbcNoBase64Expected},
	}

	for index, test := range base64Tests {
		e := Encrypt.FromString(test.input).ByDes(desCbcCipher(No))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToBase64String(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestDecrypt_ByDes_CBC_NoPadding(t *testing.T) {
	input := "12345678asdfghjk"

	hexTests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{desCbcNoHexExpected, input},
	}

	for index, test := range hexTests {
		e := Decrypt.FromHexString(test.input).ByDes(desCbcCipher(No))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}

	base32Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{desCbcNoBase32Expected, input},
	}

	for index, test := range base32Tests {
		e := Decrypt.FromBase32String(test.input).ByDes(desCbcCipher(No))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}

	base64Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{desCbcNoBase64Expected, input},
	}

	for index, test := range base64Tests {
		e := Decrypt.FromBase64String(test.input).ByDes(desCbcCipher(No))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestEncrypt_ByDes_CBC_ZeroPadding(t *testing.T) {
	hexTests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{desCbcInput, desCbcZeroHexExpected},
	}

	for index, test := range hexTests {
		e := Encrypt.FromString(test.input).ByDes(desCbcCipher(Zero))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToHexString(), "Current test index is "+strconv.Itoa(index))
	}

	base32Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{desCbcInput, desCbcZeroBase32Expected},
	}

	for index, test := range base32Tests {
		e := Encrypt.FromString(test.input).ByDes(desCbcCipher(Zero))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToBase32String(), "Current test index is "+strconv.Itoa(index))
	}

	base64Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{desCbcInput, desCbcZeroBase64Expected},
	}

	for index, test := range base64Tests {
		e := Encrypt.FromString(test.input).ByDes(desCbcCipher(Zero))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToBase64String(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestDecrypt_ByDes_CBC_ZeroPadding(t *testing.T) {
	hexTests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{desCbcZeroHexExpected, desCbcInput},
	}

	for index, test := range hexTests {
		e := Decrypt.FromHexString(test.input).ByDes(desCbcCipher(Zero))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}

	base32Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{desCbcZeroBase32Expected, desCbcInput},
	}

	for index, test := range base32Tests {
		e := Decrypt.FromBase32String(test.input).ByDes(desCbcCipher(Zero))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}

	base64Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{desCbcZeroBase64Expected, desCbcInput},
	}

	for index, test := range base64Tests {
		e := Decrypt.FromBase64String(test.input).ByDes(desCbcCipher(Zero))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestEncrypt_ByDes_CBC_PKCS5Padding(t *testing.T) {
	hexTests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{desCbcInput, desCbcPkcs5HexExpected},
	}

	for index, test := range hexTests {
		e := Encrypt.FromString(test.input).ByDes(desCbcCipher(PKCS5))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToHexString(), "Current test index is "+strconv.Itoa(index))
	}

	base32Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{desCbcInput, desCbcPkcs5Base32Expected},
	}

	for index, test := range base32Tests {
		e := Encrypt.FromString(test.input).ByDes(desCbcCipher(PKCS5))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToBase32String(), "Current test index is "+strconv.Itoa(index))
	}

	base64Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{desCbcInput, desCbcPkcs5Base64Expected},
	}

	for index, test := range base64Tests {
		e := Encrypt.FromString(test.input).ByDes(desCbcCipher(PKCS5))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToBase64String(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestDecrypt_ByDes_CBC_PKCS5Padding(t *testing.T) {
	hexTests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{desCbcPkcs5HexExpected, desCbcInput},
	}

	for index, test := range hexTests {
		e := Decrypt.FromHexString(test.input).ByDes(desCbcCipher(PKCS5))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}

	base32Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{desCbcPkcs5Base32Expected, desCbcInput},
	}

	for index, test := range base32Tests {
		e := Decrypt.FromBase32String(test.input).ByDes(desCbcCipher(PKCS5))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}

	base64Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{desCbcPkcs5Base64Expected, desCbcInput},
	}

	for index, test := range base64Tests {
		e := Decrypt.FromBase64String(test.input).ByDes(desCbcCipher(PKCS5))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestEncrypt_ByDes_CBC_PKCS7Padding(t *testing.T) {
	hexTests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{desCbcInput, desCbcPkcs7HexExpected},
	}

	for index, test := range hexTests {
		e := Encrypt.FromString(test.input).ByDes(desCbcCipher(PKCS7))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToHexString(), "Current test index is "+strconv.Itoa(index))
	}

	base32Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{desCbcInput, desCbcPkcs7Base32Expected},
	}

	for index, test := range base32Tests {
		e := Encrypt.FromString(test.input).ByDes(desCbcCipher(PKCS7))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToBase32String(), "Current test index is "+strconv.Itoa(index))
	}

	base64Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{desCbcInput, desCbcPkcs7Base64Expected},
	}

	for index, test := range base64Tests {
		e := Encrypt.FromString(test.input).ByDes(desCbcCipher(PKCS7))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToBase64String(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestDecrypt_ByDes_CBC_PKCS7Padding(t *testing.T) {
	hexTests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{desCbcPkcs7HexExpected, desCbcInput},
	}

	for index, test := range hexTests {
		e := Decrypt.FromHexString(test.input).ByDes(desCbcCipher(PKCS7))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}

	base32Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{desCbcPkcs7Base32Expected, desCbcInput},
	}

	for index, test := range base32Tests {
		e := Decrypt.FromBase32String(test.input).ByDes(desCbcCipher(PKCS7))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}

	base64Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{desCbcPkcs7Base64Expected, desCbcInput},
	}

	for index, test := range base64Tests {
		e := Decrypt.FromBase64String(test.input).ByDes(desCbcCipher(PKCS7))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}
}
