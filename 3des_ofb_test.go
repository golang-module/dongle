package dongle

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

var (
	tripleDesOfbInput = "hello world"
	tripleDesOfbKey   = "123456781234567812345678"
	tripleDesOfbIV    = "12345678"

	// 3AES-OFB-NoPadding
	tripleDesOfbNoHexExpected    = "a7e231bc4de3bbb18806fcb94e0778c1" //
	tripleDesOfbNoBase32Expected = "U7RDDPCN4O53DCAG7S4U4B3YYE======"
	tripleDesOfbNoBase64Expected = "p+IxvE3ju7GIBvy5Tgd4wQ=="

	// 3AES-OFB-ZeroPadding
	tripleDesOfbZeroHexExpected    = "feb56ee417f5fbe69b19fcdf296f12aa"
	tripleDesOfbZeroBase32Expected = "722W5ZAX6X56NGYZ7TPSS3YSVI======"
	tripleDesOfbZeroBase64Expected = "/rVu5Bf1++abGfzfKW8Sqg=="

	// 3AES-OFB-PKCS5Padding
	tripleDesOfbPkcs5HexExpected    = "feb56ee417f5fbe69b19fcda2c6a17af"
	tripleDesOfbPkcs5Base32Expected = "722W5ZAX6X56NGYZ7TNCY2QXV4======"
	tripleDesOfbPkcs5Base64Expected = "/rVu5Bf1++abGfzaLGoXrw=="

	// 3AES-OFB-PKCS7Padding
	tripleDesOfbPkcs7HexExpected    = "feb56ee417f5fbe69b19fcda2c6a17af"
	tripleDesOfbPkcs7Base32Expected = "722W5ZAX6X56NGYZ7TNCY2QXV4======"
	tripleDesOfbPkcs7Base64Expected = "/rVu5Bf1++abGfzaLGoXrw=="
)

func tripleDesOfbCipher(padding string) *Cipher {
	cipher := NewCipher()
	cipher.SetMode(OFB)
	cipher.SetPadding(padding)
	cipher.SetKey(tripleDesOfbKey)
	cipher.SetIV(tripleDesOfbIV)
	return cipher
}

func TestEncrypt_By3Des_OFB_NoPadding(t *testing.T) {
	input := "12345678asdfghjk"

	hexTests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{input, tripleDesOfbNoHexExpected},
	}

	for index, test := range hexTests {
		e := Encrypt.FromString(test.input).By3Des(tripleDesOfbCipher(No))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToHexString(), "Current test index is "+strconv.Itoa(index))
	}

	base32Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{input, tripleDesOfbNoBase32Expected},
	}

	for index, test := range base32Tests {
		e := Encrypt.FromString(test.input).By3Des(tripleDesOfbCipher(No))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToBase32String(), "Current test index is "+strconv.Itoa(index))
	}

	base64Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{input, tripleDesOfbNoBase64Expected},
	}

	for index, test := range base64Tests {
		e := Encrypt.FromString(test.input).By3Des(tripleDesOfbCipher(No))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToBase64String(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestDecrypt_By3Des_OFB_NoPadding(t *testing.T) {
	input := "12345678asdfghjk"

	hexTests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{tripleDesOfbNoHexExpected, input},
	}

	for index, test := range hexTests {
		e := Decrypt.FromHexString(test.input).By3Des(tripleDesOfbCipher(No))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}

	base32Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{tripleDesOfbNoBase32Expected, input},
	}

	for index, test := range base32Tests {
		e := Decrypt.FromBase32String(test.input).By3Des(tripleDesOfbCipher(No))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}

	base64Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{tripleDesOfbNoBase64Expected, input},
	}

	for index, test := range base64Tests {
		e := Decrypt.FromBase64String(test.input).By3Des(tripleDesOfbCipher(No))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestEncrypt_By3Des_OFB_ZeroPadding(t *testing.T) {
	hexTests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{tripleDesOfbInput, tripleDesOfbZeroHexExpected},
	}

	for index, test := range hexTests {
		e := Encrypt.FromString(test.input).By3Des(tripleDesOfbCipher(Zero))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToHexString(), "Current test index is "+strconv.Itoa(index))
	}

	base32Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{tripleDesOfbInput, tripleDesOfbZeroBase32Expected},
	}

	for index, test := range base32Tests {
		e := Encrypt.FromString(test.input).By3Des(tripleDesOfbCipher(Zero))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToBase32String(), "Current test index is "+strconv.Itoa(index))
	}

	base64Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{tripleDesOfbInput, tripleDesOfbZeroBase64Expected},
	}

	for index, test := range base64Tests {
		e := Encrypt.FromString(test.input).By3Des(tripleDesOfbCipher(Zero))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToBase64String(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestDecrypt_By3Des_OFB_ZeroPadding(t *testing.T) {
	hexTests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{tripleDesOfbZeroHexExpected, tripleDesOfbInput},
	}

	for index, test := range hexTests {
		e := Decrypt.FromHexString(test.input).By3Des(tripleDesOfbCipher(Zero))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}

	base32Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{tripleDesOfbZeroBase32Expected, tripleDesOfbInput},
	}

	for index, test := range base32Tests {
		e := Decrypt.FromBase32String(test.input).By3Des(tripleDesOfbCipher(Zero))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}

	base64Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{tripleDesOfbZeroBase64Expected, tripleDesOfbInput},
	}

	for index, test := range base64Tests {
		e := Decrypt.FromBase64String(test.input).By3Des(tripleDesOfbCipher(Zero))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestEncrypt_By3Des_OFB_PKCS5Padding(t *testing.T) {
	hexTests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{tripleDesOfbInput, tripleDesOfbPkcs5HexExpected},
	}

	for index, test := range hexTests {
		e := Encrypt.FromString(test.input).By3Des(tripleDesOfbCipher(PKCS5))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToHexString(), "Current test index is "+strconv.Itoa(index))
	}

	base32Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{tripleDesOfbInput, tripleDesOfbPkcs5Base32Expected},
	}

	for index, test := range base32Tests {
		e := Encrypt.FromString(test.input).By3Des(tripleDesOfbCipher(PKCS5))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToBase32String(), "Current test index is "+strconv.Itoa(index))
	}

	base64Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{tripleDesOfbInput, tripleDesOfbPkcs5Base64Expected},
	}

	for index, test := range base64Tests {
		e := Encrypt.FromString(test.input).By3Des(tripleDesOfbCipher(PKCS5))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToBase64String(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestDecrypt_By3Des_OFB_PKCS5Padding(t *testing.T) {
	hexTests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{tripleDesOfbPkcs5HexExpected, tripleDesOfbInput},
	}

	for index, test := range hexTests {
		e := Decrypt.FromHexString(test.input).By3Des(tripleDesOfbCipher(PKCS5))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}

	base32Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{tripleDesOfbPkcs5Base32Expected, tripleDesOfbInput},
	}

	for index, test := range base32Tests {
		e := Decrypt.FromBase32String(test.input).By3Des(tripleDesOfbCipher(PKCS5))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}

	base64Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{tripleDesOfbPkcs5Base64Expected, tripleDesOfbInput},
	}

	for index, test := range base64Tests {
		e := Decrypt.FromBase64String(test.input).By3Des(tripleDesOfbCipher(PKCS5))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestEncrypt_By3Des_OFB_PKCS7Padding(t *testing.T) {
	hexTests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{tripleDesOfbInput, tripleDesOfbPkcs7HexExpected},
	}

	for index, test := range hexTests {
		e := Encrypt.FromString(test.input).By3Des(tripleDesOfbCipher(PKCS7))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToHexString(), "Current test index is "+strconv.Itoa(index))
	}

	base32Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{tripleDesOfbInput, tripleDesOfbPkcs7Base32Expected},
	}

	for index, test := range base32Tests {
		e := Encrypt.FromString(test.input).By3Des(tripleDesOfbCipher(PKCS7))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToBase32String(), "Current test index is "+strconv.Itoa(index))
	}

	base64Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{tripleDesOfbInput, tripleDesOfbPkcs7Base64Expected},
	}

	for index, test := range base64Tests {
		e := Encrypt.FromString(test.input).By3Des(tripleDesOfbCipher(PKCS7))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToBase64String(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestDecrypt_By3Des_OFB_PKCS7Padding(t *testing.T) {
	hexTests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{tripleDesOfbPkcs7HexExpected, tripleDesOfbInput},
	}

	for index, test := range hexTests {
		e := Decrypt.FromHexString(test.input).By3Des(tripleDesOfbCipher(PKCS7))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}

	base32Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{tripleDesOfbPkcs7Base32Expected, tripleDesOfbInput},
	}

	for index, test := range base32Tests {
		e := Decrypt.FromBase32String(test.input).By3Des(tripleDesOfbCipher(PKCS7))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}

	base64Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{tripleDesOfbPkcs7Base64Expected, tripleDesOfbInput},
	}

	for index, test := range base64Tests {
		e := Decrypt.FromBase64String(test.input).By3Des(tripleDesOfbCipher(PKCS7))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}
}
