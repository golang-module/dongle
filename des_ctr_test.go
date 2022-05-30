package dongle

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

var (
	desCtrInput = "hello world"
	desCtrKey   = "12345678"
	desCtrIV    = "12345678"

	// DES-CTR-NoPadding
	desCtrNoHexExpected = "a7e231bc4de3bbb120618b70be119c63"
	// desCtrNoHexExpected    = "a7e231bc4de3bbb120618b70be119c63" // https://oktools.net/des
	// desCtrNoHexExpected    = "a7e231bc4de3bbb120618b70be119c63" // http://tool.chacuo.net/cryptdes
	desCtrNoBase32Expected = "U7RDDPCN4O53CIDBRNYL4EM4MM======"
	desCtrNoBase64Expected = "p+IxvE3ju7EgYYtwvhGcYw=="

	// DES-CTR-ZeroPadding
	desCtrZeroHexExpected    = "feb56ee417f5fbe6337e8b16d979f608"
	desCtrZeroBase32Expected = "722W5ZAX6X56MM36RMLNS6PWBA======"
	desCtrZeroBase64Expected = "/rVu5Bf1++YzfosW2Xn2CA=="

	// DES-CTR-PKCS5Padding
	desCtrPkcs5HexExpected    = "feb56ee417f5fbe6337e8b13dc7cf30d"
	desCtrPkcs5Base32Expected = "722W5ZAX6X56MM36RMJ5Y7HTBU======"
	desCtrPkcs5Base64Expected = "/rVu5Bf1++YzfosT3HzzDQ=="

	// DES-CTR-PKCS7Padding
	desCtrPkcs7HexExpected    = "feb56ee417f5fbe6337e8b13dc7cf30d"
	desCtrPkcs7Base32Expected = "722W5ZAX6X56MM36RMJ5Y7HTBU======"
	desCtrPkcs7Base64Expected = "/rVu5Bf1++YzfosT3HzzDQ=="
)

func desCtrCipher(padding string) *Cipher {
	cipher := NewCipher()
	cipher.SetMode(CTR)
	cipher.SetPadding(padding)
	cipher.SetKey(desCtrKey)
	cipher.SetIV(desCtrIV)
	return cipher
}

func TestEncrypt_ByDes_CTR_NoPadding(t *testing.T) {
	input := "12345678asdfghjk"

	hexTests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{input, desCtrNoHexExpected},
	}

	for index, test := range hexTests {
		e := Encrypt.FromString(test.input).ByDes(desCtrCipher(No))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToHexString(), "Current test index is "+strconv.Itoa(index))
	}

	base32Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{input, desCtrNoBase32Expected},
	}

	for index, test := range base32Tests {
		e := Encrypt.FromString(test.input).ByDes(desCtrCipher(No))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToBase32String(), "Current test index is "+strconv.Itoa(index))
	}

	base64Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{input, desCtrNoBase64Expected},
	}

	for index, test := range base64Tests {
		e := Encrypt.FromString(test.input).ByDes(desCtrCipher(No))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToBase64String(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestDecrypt_ByDes_CTR_NoPadding(t *testing.T) {
	input := "12345678asdfghjk"

	hexTests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{desCtrNoHexExpected, input},
	}

	for index, test := range hexTests {
		e := Decrypt.FromHexString(test.input).ByDes(desCtrCipher(No))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}

	base32Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{desCtrNoBase32Expected, input},
	}

	for index, test := range base32Tests {
		e := Decrypt.FromBase32String(test.input).ByDes(desCtrCipher(No))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}

	base64Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{desCtrNoBase64Expected, input},
	}

	for index, test := range base64Tests {
		e := Decrypt.FromBase64String(test.input).ByDes(desCtrCipher(No))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestEncrypt_ByDes_CTR_ZeroPadding(t *testing.T) {
	hexTests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{desCtrInput, desCtrZeroHexExpected},
	}

	for index, test := range hexTests {
		e := Encrypt.FromString(test.input).ByDes(desCtrCipher(Zero))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToHexString(), "Current test index is "+strconv.Itoa(index))
	}

	base32Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{desCtrInput, desCtrZeroBase32Expected},
	}

	for index, test := range base32Tests {
		e := Encrypt.FromString(test.input).ByDes(desCtrCipher(Zero))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToBase32String(), "Current test index is "+strconv.Itoa(index))
	}

	base64Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{desCtrInput, desCtrZeroBase64Expected},
	}

	for index, test := range base64Tests {
		e := Encrypt.FromString(test.input).ByDes(desCtrCipher(Zero))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToBase64String(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestDecrypt_ByDes_CTR_ZeroPadding(t *testing.T) {
	hexTests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{desCtrZeroHexExpected, desCtrInput},
	}

	for index, test := range hexTests {
		e := Decrypt.FromHexString(test.input).ByDes(desCtrCipher(Zero))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}

	base32Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{desCtrZeroBase32Expected, desCtrInput},
	}

	for index, test := range base32Tests {
		e := Decrypt.FromBase32String(test.input).ByDes(desCtrCipher(Zero))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}

	base64Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{desCtrZeroBase64Expected, desCtrInput},
	}

	for index, test := range base64Tests {
		e := Decrypt.FromBase64String(test.input).ByDes(desCtrCipher(Zero))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestEncrypt_ByDes_CTR_PKCS5Padding(t *testing.T) {
	hexTests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{desCtrInput, desCtrPkcs5HexExpected},
	}

	for index, test := range hexTests {
		e := Encrypt.FromString(test.input).ByDes(desCtrCipher(PKCS5))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToHexString(), "Current test index is "+strconv.Itoa(index))
	}

	base32Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{desCtrInput, desCtrPkcs5Base32Expected},
	}

	for index, test := range base32Tests {
		e := Encrypt.FromString(test.input).ByDes(desCtrCipher(PKCS5))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToBase32String(), "Current test index is "+strconv.Itoa(index))
	}

	base64Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{desCtrInput, desCtrPkcs5Base64Expected},
	}

	for index, test := range base64Tests {
		e := Encrypt.FromString(test.input).ByDes(desCtrCipher(PKCS5))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToBase64String(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestDecrypt_ByDes_CTR_PKCS5Padding(t *testing.T) {
	hexTests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{desCtrPkcs5HexExpected, desCtrInput},
	}

	for index, test := range hexTests {
		e := Decrypt.FromHexString(test.input).ByDes(desCtrCipher(PKCS5))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}

	base32Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{desCtrPkcs5Base32Expected, desCtrInput},
	}

	for index, test := range base32Tests {
		e := Decrypt.FromBase32String(test.input).ByDes(desCtrCipher(PKCS5))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}

	base64Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{desCtrPkcs5Base64Expected, desCtrInput},
	}

	for index, test := range base64Tests {
		e := Decrypt.FromBase64String(test.input).ByDes(desCtrCipher(PKCS5))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestEncrypt_ByDes_CTR_PKCS7Padding(t *testing.T) {
	hexTests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{desCtrInput, desCtrPkcs7HexExpected},
	}

	for index, test := range hexTests {
		e := Encrypt.FromString(test.input).ByDes(desCtrCipher(PKCS7))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToHexString(), "Current test index is "+strconv.Itoa(index))
	}

	base32Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{desCtrInput, desCtrPkcs7Base32Expected},
	}

	for index, test := range base32Tests {
		e := Encrypt.FromString(test.input).ByDes(desCtrCipher(PKCS7))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToBase32String(), "Current test index is "+strconv.Itoa(index))
	}

	base64Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{desCtrInput, desCtrPkcs7Base64Expected},
	}

	for index, test := range base64Tests {
		e := Encrypt.FromString(test.input).ByDes(desCtrCipher(PKCS7))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToBase64String(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestDecrypt_ByDes_CTR_PKCS7Padding(t *testing.T) {
	hexTests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{desCtrPkcs7HexExpected, desCtrInput},
	}

	for index, test := range hexTests {
		e := Decrypt.FromHexString(test.input).ByDes(desCtrCipher(PKCS7))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}

	base32Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{desCtrPkcs7Base32Expected, desCtrInput},
	}

	for index, test := range base32Tests {
		e := Decrypt.FromBase32String(test.input).ByDes(desCtrCipher(PKCS7))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}

	base64Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{desCtrPkcs7Base64Expected, desCtrInput},
	}

	for index, test := range base64Tests {
		e := Decrypt.FromBase64String(test.input).ByDes(desCtrCipher(PKCS7))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}
}
