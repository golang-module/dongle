package dongle

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

var (
	desOfbInput = "hello world"
	desOfbKey   = "12345678"
	desOfbIV    = "12345678"

	// DES-OFB-NoPadding
	desOfbNoHexExpected = "a7e231bc4de3bbb18806fcb94e0778c1"
	// desOfbNoHexExpected = "a7e231bc4de3bbb18806fcb94e0778c1" // https://oktools.net/des
	// desOfbNoHexExpected = "a772b4cd10858f12be00fc3bcfc2b130" http://tool.chacuo.net/cryptdes
	desOfbNoBase32Expected = "U7RDDPCN4O53DCAG7S4U4B3YYE======"
	desOfbNoBase64Expected = "p+IxvE3ju7GIBvy5Tgd4wQ=="

	// DES-OFB-ZeroPadding
	desOfbZeroHexExpected = "feb56ee417f5fbe69b19fcdf296f12aa"
	// desOfbZeroHexExpected    = "feb56ee417f5fbe69b19fcdf296f12aa" // https://oktools.net/des
	// desOfbZeroHexExpected    = "feb56ee417f5fbe69b19fcdf296f12aa" // https://tool.lmeee.com/jiami/des
	// desOfbZeroHexExpected    = "fe25eb954a93cf45ad1ffc5da8aadb5b" // http://tool.chacuo.net/cryptdes
	desOfbZeroBase32Expected = "722W5ZAX6X56NGYZ7TPSS3YSVI======"
	desOfbZeroBase64Expected = "/rVu5Bf1++abGfzfKW8Sqg=="

	// DES-OFB-PKCS5Padding
	desOfbPkcs5HexExpected = "feb56ee417f5fbe69b19fcda2c6a17af"
	// desOfbPkcs5HexExpected    = "feb56ee417f5fbe69b19fcda2c6a17af" // https://oktools.net/des
	// desOfbPkcs5HexExpected    = "feb56ee417f5fbe69b19fcda2c6a17af" // https://tool.lmeee.com/jiami/des
	// desOfbPkcs5HexExpected    = "fe25eb954a93cf45ad1ffc58adafde5e" // http://tool.chacuo.net/cryptdes
	desOfbPkcs5Base32Expected = "722W5ZAX6X56NGYZ7TNCY2QXV4======"
	desOfbPkcs5Base64Expected = "/rVu5Bf1++abGfzaLGoXrw=="

	// DES-OFB-PKCS7Padding
	desOfbPkcs7HexExpected = "feb56ee417f5fbe69b19fcda2c6a17af"
	// desOfbPkcs7HexExpected    = "feb56ee417f5fbe69b19fcda2c6a17af" // https://oktools.net/des
	// desOfbPkcs7HexExpected    = "feb56ee417f5fbe69b19fcda2c6a17af" // https://tool.lmeee.com/jiami/des
	// desOfbPkcs7HexExpected    = "fe25eb954a93cf45ad1ffc58adafde5e" // http://tool.chacuo.net/cryptdes
	desOfbPkcs7Base32Expected = "722W5ZAX6X56NGYZ7TNCY2QXV4======"
	desOfbPkcs7Base64Expected = "/rVu5Bf1++abGfzaLGoXrw=="
)

func desOfbCipher(padding string) *Cipher {
	cipher := NewCipher()
	cipher.SetMode(OFB)
	cipher.SetPadding(padding)
	cipher.SetKey(desOfbKey)
	cipher.SetIV(desOfbIV)
	return cipher
}

func TestEncrypt_ByDes_OFB_NoPadding(t *testing.T) {
	input := "12345678asdfghjk"

	hexTests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{input, desOfbNoHexExpected},
	}

	for index, test := range hexTests {
		e := Encrypt.FromString(test.input).ByDes(desOfbCipher(No))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToHexString(), "Current test index is "+strconv.Itoa(index))
	}

	base32Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{input, desOfbNoBase32Expected},
	}

	for index, test := range base32Tests {
		e := Encrypt.FromString(test.input).ByDes(desOfbCipher(No))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToBase32String(), "Current test index is "+strconv.Itoa(index))
	}

	base64Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{input, desOfbNoBase64Expected},
	}

	for index, test := range base64Tests {
		e := Encrypt.FromString(test.input).ByDes(desOfbCipher(No))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToBase64String(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestDecrypt_ByDes_OFB_NoPadding(t *testing.T) {
	input := "12345678asdfghjk"

	hexTests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{desOfbNoHexExpected, input},
	}

	for index, test := range hexTests {
		e := Decrypt.FromHexString(test.input).ByDes(desOfbCipher(No))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}

	base32Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{desOfbNoBase32Expected, input},
	}

	for index, test := range base32Tests {
		e := Decrypt.FromBase32String(test.input).ByDes(desOfbCipher(No))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}

	base64Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{desOfbNoBase64Expected, input},
	}

	for index, test := range base64Tests {
		e := Decrypt.FromBase64String(test.input).ByDes(desOfbCipher(No))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestEncrypt_ByDes_OFB_ZeroPadding(t *testing.T) {
	hexTests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{desOfbInput, desOfbZeroHexExpected},
	}

	for index, test := range hexTests {
		e := Encrypt.FromString(test.input).ByDes(desOfbCipher(Zero))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToHexString(), "Current test index is "+strconv.Itoa(index))
	}

	base32Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{desOfbInput, desOfbZeroBase32Expected},
	}

	for index, test := range base32Tests {
		e := Encrypt.FromString(test.input).ByDes(desOfbCipher(Zero))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToBase32String(), "Current test index is "+strconv.Itoa(index))
	}

	base64Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{desOfbInput, desOfbZeroBase64Expected},
	}

	for index, test := range base64Tests {
		e := Encrypt.FromString(test.input).ByDes(desOfbCipher(Zero))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToBase64String(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestDecrypt_ByDes_OFB_ZeroPadding(t *testing.T) {
	hexTests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{desOfbZeroHexExpected, desOfbInput},
	}

	for index, test := range hexTests {
		e := Decrypt.FromHexString(test.input).ByDes(desOfbCipher(Zero))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}

	base32Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{desOfbZeroBase32Expected, desOfbInput},
	}

	for index, test := range base32Tests {
		e := Decrypt.FromBase32String(test.input).ByDes(desOfbCipher(Zero))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}

	base64Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{desOfbZeroBase64Expected, desOfbInput},
	}

	for index, test := range base64Tests {
		e := Decrypt.FromBase64String(test.input).ByDes(desOfbCipher(Zero))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestEncrypt_ByDes_OFB_PKCS5Padding(t *testing.T) {
	hexTests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{desOfbInput, desOfbPkcs5HexExpected},
	}

	for index, test := range hexTests {
		e := Encrypt.FromString(test.input).ByDes(desOfbCipher(PKCS5))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToHexString(), "Current test index is "+strconv.Itoa(index))
	}

	base32Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{desOfbInput, desOfbPkcs5Base32Expected},
	}

	for index, test := range base32Tests {
		e := Encrypt.FromString(test.input).ByDes(desOfbCipher(PKCS5))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToBase32String(), "Current test index is "+strconv.Itoa(index))
	}

	base64Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{desOfbInput, desOfbPkcs5Base64Expected},
	}

	for index, test := range base64Tests {
		e := Encrypt.FromString(test.input).ByDes(desOfbCipher(PKCS5))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToBase64String(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestDecrypt_ByDes_OFB_PKCS5Padding(t *testing.T) {
	hexTests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{desOfbPkcs5HexExpected, desOfbInput},
	}

	for index, test := range hexTests {
		e := Decrypt.FromHexString(test.input).ByDes(desOfbCipher(PKCS5))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}

	base32Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{desOfbPkcs5Base32Expected, desOfbInput},
	}

	for index, test := range base32Tests {
		e := Decrypt.FromBase32String(test.input).ByDes(desOfbCipher(PKCS5))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}

	base64Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{desOfbPkcs5Base64Expected, desOfbInput},
	}

	for index, test := range base64Tests {
		e := Decrypt.FromBase64String(test.input).ByDes(desOfbCipher(PKCS5))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestEncrypt_ByDes_OFB_PKCS7Padding(t *testing.T) {
	hexTests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{desOfbInput, desOfbPkcs7HexExpected},
	}

	for index, test := range hexTests {
		e := Encrypt.FromString(test.input).ByDes(desOfbCipher(PKCS7))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToHexString(), "Current test index is "+strconv.Itoa(index))
	}

	base32Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{desOfbInput, desOfbPkcs7Base32Expected},
	}

	for index, test := range base32Tests {
		e := Encrypt.FromString(test.input).ByDes(desOfbCipher(PKCS7))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToBase32String(), "Current test index is "+strconv.Itoa(index))
	}

	base64Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{desOfbInput, desOfbPkcs7Base64Expected},
	}

	for index, test := range base64Tests {
		e := Encrypt.FromString(test.input).ByDes(desOfbCipher(PKCS7))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToBase64String(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestDecrypt_ByDes_OFB_PKCS7Padding(t *testing.T) {
	hexTests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{desOfbPkcs7HexExpected, desOfbInput},
	}

	for index, test := range hexTests {
		e := Decrypt.FromHexString(test.input).ByDes(desOfbCipher(PKCS7))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}

	base32Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{desOfbPkcs7Base32Expected, desOfbInput},
	}

	for index, test := range base32Tests {
		e := Decrypt.FromBase32String(test.input).ByDes(desOfbCipher(PKCS7))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}

	base64Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{desOfbPkcs7Base64Expected, desOfbInput},
	}

	for index, test := range base64Tests {
		e := Decrypt.FromBase64String(test.input).ByDes(desOfbCipher(PKCS7))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}
}
