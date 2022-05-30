package dongle

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

var (
	tripleDesCtrInput = "hello world"
	tripleDesCtrKey   = "123456781234567812345678"
	tripleDesCtrIV    = "12345678"

	// 3DES-CTR-NoPadding
	tripleDesCtrNoHexExpected    = "a7e231bc4de3bbb120618b70be119c63"
	tripleDesCtrNoBase32Expected = "U7RDDPCN4O53CIDBRNYL4EM4MM======"
	tripleDesCtrNoBase64Expected = "p+IxvE3ju7EgYYtwvhGcYw=="

	// 3DES-CTR-ZeroPadding
	tripleDesCtrZeroHexExpected    = "feb56ee417f5fbe6337e8b16d979f608"
	tripleDesCtrZeroBase32Expected = "722W5ZAX6X56MM36RMLNS6PWBA======"
	tripleDesCtrZeroBase64Expected = "/rVu5Bf1++YzfosW2Xn2CA=="

	// 3DES-CTR-PKCS5Padding
	tripleDesCtrPkcs5HexExpected    = "feb56ee417f5fbe6337e8b13dc7cf30d"
	tripleDesCtrPkcs5Base32Expected = "722W5ZAX6X56MM36RMJ5Y7HTBU======"
	tripleDesCtrPkcs5Base64Expected = "/rVu5Bf1++YzfosT3HzzDQ=="

	// 3DES-CTR-PKCS7Padding
	tripleDesCtrPkcs7HexExpected    = "feb56ee417f5fbe6337e8b13dc7cf30d"
	tripleDesCtrPkcs7Base32Expected = "722W5ZAX6X56MM36RMJ5Y7HTBU======"
	tripleDesCtrPkcs7Base64Expected = "/rVu5Bf1++YzfosT3HzzDQ=="
)

func tripleDesCtrCipher(padding string) *Cipher {
	cipher := NewCipher()
	cipher.SetMode(CTR)
	cipher.SetPadding(padding)
	cipher.SetKey(tripleDesCtrKey)
	cipher.SetIV(tripleDesCtrIV)
	return cipher
}

func TestEncrypt_By3Des_CTR_NoPadding(t *testing.T) {
	input := "12345678asdfghjk"

	hexTests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{input, tripleDesCtrNoHexExpected},
	}

	for index, test := range hexTests {
		e := Encrypt.FromString(test.input).By3Des(tripleDesCtrCipher(No))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToHexString(), "Current test index is "+strconv.Itoa(index))
	}

	base32Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{input, tripleDesCtrNoBase32Expected},
	}

	for index, test := range base32Tests {
		e := Encrypt.FromString(test.input).By3Des(tripleDesCtrCipher(No))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToBase32String(), "Current test index is "+strconv.Itoa(index))
	}

	base64Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{input, tripleDesCtrNoBase64Expected},
	}

	for index, test := range base64Tests {
		e := Encrypt.FromString(test.input).By3Des(tripleDesCtrCipher(No))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToBase64String(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestDecrypt_By3Des_CTR_NoPadding(t *testing.T) {
	input := "12345678asdfghjk"

	hexTests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{tripleDesCtrNoHexExpected, input},
	}

	for index, test := range hexTests {
		e := Decrypt.FromHexString(test.input).By3Des(tripleDesCtrCipher(No))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}

	base32Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{tripleDesCtrNoBase32Expected, input},
	}

	for index, test := range base32Tests {
		e := Decrypt.FromBase32String(test.input).By3Des(tripleDesCtrCipher(No))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}

	base64Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{tripleDesCtrNoBase64Expected, input},
	}

	for index, test := range base64Tests {
		e := Decrypt.FromBase64String(test.input).By3Des(tripleDesCtrCipher(No))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestEncrypt_By3Des_CTR_ZeroPadding(t *testing.T) {
	hexTests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{tripleDesCtrInput, tripleDesCtrZeroHexExpected},
	}

	for index, test := range hexTests {
		e := Encrypt.FromString(test.input).By3Des(tripleDesCtrCipher(Zero))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToHexString(), "Current test index is "+strconv.Itoa(index))
	}

	base32Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{tripleDesCtrInput, tripleDesCtrZeroBase32Expected},
	}

	for index, test := range base32Tests {
		e := Encrypt.FromString(test.input).By3Des(tripleDesCtrCipher(Zero))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToBase32String(), "Current test index is "+strconv.Itoa(index))
	}

	base64Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{tripleDesCtrInput, tripleDesCtrZeroBase64Expected},
	}

	for index, test := range base64Tests {
		e := Encrypt.FromString(test.input).By3Des(tripleDesCtrCipher(Zero))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToBase64String(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestDecrypt_By3Des_CTR_ZeroPadding(t *testing.T) {
	hexTests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{tripleDesCtrZeroHexExpected, tripleDesCtrInput},
	}

	for index, test := range hexTests {
		e := Decrypt.FromHexString(test.input).By3Des(tripleDesCtrCipher(Zero))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}

	base32Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{tripleDesCtrZeroBase32Expected, tripleDesCtrInput},
	}

	for index, test := range base32Tests {
		e := Decrypt.FromBase32String(test.input).By3Des(tripleDesCtrCipher(Zero))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}

	base64Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{tripleDesCtrZeroBase64Expected, tripleDesCtrInput},
	}

	for index, test := range base64Tests {
		e := Decrypt.FromBase64String(test.input).By3Des(tripleDesCtrCipher(Zero))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestEncrypt_By3Des_CTR_PKCS5Padding(t *testing.T) {
	hexTests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{tripleDesCtrInput, tripleDesCtrPkcs5HexExpected},
	}

	for index, test := range hexTests {
		e := Encrypt.FromString(test.input).By3Des(tripleDesCtrCipher(PKCS5))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToHexString(), "Current test index is "+strconv.Itoa(index))
	}

	base32Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{tripleDesCtrInput, tripleDesCtrPkcs5Base32Expected},
	}

	for index, test := range base32Tests {
		e := Encrypt.FromString(test.input).By3Des(tripleDesCtrCipher(PKCS5))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToBase32String(), "Current test index is "+strconv.Itoa(index))
	}

	base64Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{tripleDesCtrInput, tripleDesCtrPkcs5Base64Expected},
	}

	for index, test := range base64Tests {
		e := Encrypt.FromString(test.input).By3Des(tripleDesCtrCipher(PKCS5))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToBase64String(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestDecrypt_By3Des_CTR_PKCS5Padding(t *testing.T) {
	hexTests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{tripleDesCtrPkcs5HexExpected, tripleDesCtrInput},
	}

	for index, test := range hexTests {
		e := Decrypt.FromHexString(test.input).By3Des(tripleDesCtrCipher(PKCS5))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}

	base32Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{tripleDesCtrPkcs5Base32Expected, tripleDesCtrInput},
	}

	for index, test := range base32Tests {
		e := Decrypt.FromBase32String(test.input).By3Des(tripleDesCtrCipher(PKCS5))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}

	base64Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{tripleDesCtrPkcs5Base64Expected, tripleDesCtrInput},
	}

	for index, test := range base64Tests {
		e := Decrypt.FromBase64String(test.input).By3Des(tripleDesCtrCipher(PKCS5))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestEncrypt_By3Des_CTR_PKCS7Padding(t *testing.T) {
	hexTests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{tripleDesCtrInput, tripleDesCtrPkcs7HexExpected},
	}

	for index, test := range hexTests {
		e := Encrypt.FromString(test.input).By3Des(tripleDesCtrCipher(PKCS7))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToHexString(), "Current test index is "+strconv.Itoa(index))
	}

	base32Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{tripleDesCtrInput, tripleDesCtrPkcs7Base32Expected},
	}

	for index, test := range base32Tests {
		e := Encrypt.FromString(test.input).By3Des(tripleDesCtrCipher(PKCS7))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToBase32String(), "Current test index is "+strconv.Itoa(index))
	}

	base64Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{tripleDesCtrInput, tripleDesCtrPkcs7Base64Expected},
	}

	for index, test := range base64Tests {
		e := Encrypt.FromString(test.input).By3Des(tripleDesCtrCipher(PKCS7))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToBase64String(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestDecrypt_By3Des_CTR_PKCS7Padding(t *testing.T) {
	hexTests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{tripleDesCtrPkcs7HexExpected, tripleDesCtrInput},
	}

	for index, test := range hexTests {
		e := Decrypt.FromHexString(test.input).By3Des(tripleDesCtrCipher(PKCS7))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}

	base32Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{tripleDesCtrPkcs7Base32Expected, tripleDesCtrInput},
	}

	for index, test := range base32Tests {
		e := Decrypt.FromBase32String(test.input).By3Des(tripleDesCtrCipher(PKCS7))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}

	base64Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{tripleDesCtrPkcs7Base64Expected, tripleDesCtrInput},
	}

	for index, test := range base64Tests {
		e := Decrypt.FromBase64String(test.input).By3Des(tripleDesCtrCipher(PKCS7))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}
}
