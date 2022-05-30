package dongle

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

var (
	tripleDesCfbInput = "hello world"
	tripleDesCfbKey   = "123456781234567812345678"
	tripleDesCfbIV    = "12345678"

	// 3DES-CFB-NoPadding
	tripleDesCfbNoHexExpected = "a7e231bc4de3bbb1b64cda645f2ae161"
	// tripleDesCfbNoHexExpected    = "a7e231bc4de3bbb1b64cda645f2ae161" // https://oktools.net/des
	// tripleDesCfbNoHexExpected    = "a747a55a077088a8eb29a927317f4d7e" // http://tool.chacuo.net/cryptdes
	tripleDesCfbNoBase32Expected = "U7RDDPCN4O53DNSM3JSF6KXBME======"
	tripleDesCfbNoBase64Expected = "p+IxvE3ju7G2TNpkXyrhYQ=="

	// 3DES-CFB-ZeroPadding
	tripleDesCfbZeroHexExpected = "feb56ee417f5fbe61a4e572c995b5b2a"
	// tripleDesCfbZeroHexExpected    = "feb56ee417f5fbe61a4e572c995b5b2a" // https://oktools.net/des
	// tripleDesCfbZeroHexExpected    = "feb56ee417f5fbe61a4e572c995b5b2a" // https://tool.lmeee.com/jiami/des
	// tripleDesCfbZeroHexExpected    = "fe0c815405fd275284e6f6ce79114dfa" // http://tool.chacuo.net/cryptdes
	tripleDesCfbZeroBase32Expected = "722W5ZAX6X56MGSOK4WJSW23FI======"
	tripleDesCfbZeroBase64Expected = "/rVu5Bf1++YaTlcsmVtbKg=="

	// 3DES-CFB-PKCS5Padding
	tripleDesCfbPkcs5HexExpected = "feb56ee417f5fbe61a4e57299c5e5e2f"
	// tripleDesCfbPkcs5HexExpected    = "feb56ee417f5fbe61a4e57299c5e5e2f" // https://oktools.net/des
	// tripleDesCfbPkcs5HexExpected    = "feb56ee417f5fbe61a4e57299c5e5e2f" // https://tool.lmeee.com/jiami/des
	// tripleDesCfbPkcs5HexExpected    = "fe0c815405fd275284e6f6cb08599485" // http://tool.chacuo.net/cryptdes
	tripleDesCfbPkcs5Base32Expected = "722W5ZAX6X56MGSOK4UZYXS6F4======"
	tripleDesCfbPkcs5Base64Expected = "/rVu5Bf1++YaTlcpnF5eLw=="

	// 3DES-CFB-PKCS7Padding
	tripleDesCfbPkcs7HexExpected = "feb56ee417f5fbe61a4e57299c5e5e2f"
	// tripleDesCfbPkcs7HexExpected    = "feb56ee417f5fbe61a4e57299c5e5e2f" // https://oktools.net/des
	// tripleDesCfbPkcs7HexExpected    = "feb56ee417f5fbe61a4e57299c5e5e2f" // https://tool.lmeee.com/jiami/des
	// tripleDesCfbPkcs7HexExpected    = "fe0c815405fd275284e6f6cb08599485" // http://tool.chacuo.net/cryptdes
	tripleDesCfbPkcs7Base32Expected = "722W5ZAX6X56MGSOK4UZYXS6F4======"
	tripleDesCfbPkcs7Base64Expected = "/rVu5Bf1++YaTlcpnF5eLw=="
)

func tripleDesCfbCipher(padding string) *Cipher {
	cipher := NewCipher()
	cipher.SetMode(CFB)
	cipher.SetPadding(padding)
	cipher.SetKey(tripleDesCfbKey)
	cipher.SetIV(tripleDesCfbIV)
	return cipher
}

func TestEncrypt_By3Des_CFB_NoPadding(t *testing.T) {
	input := "12345678asdfghjk"

	hexTests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{input, tripleDesCfbNoHexExpected},
	}

	for index, test := range hexTests {
		e := Encrypt.FromString(test.input).By3Des(tripleDesCfbCipher(No))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToHexString(), "Current test index is "+strconv.Itoa(index))
	}

	base32Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{input, tripleDesCfbNoBase32Expected},
	}

	for index, test := range base32Tests {
		e := Encrypt.FromString(test.input).By3Des(tripleDesCfbCipher(No))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToBase32String(), "Current test index is "+strconv.Itoa(index))
	}

	base64Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{input, tripleDesCfbNoBase64Expected},
	}

	for index, test := range base64Tests {
		e := Encrypt.FromString(test.input).By3Des(tripleDesCfbCipher(No))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToBase64String(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestDecrypt_By3Des_CFB_NoPadding(t *testing.T) {
	input := "12345678asdfghjk"

	hexTests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{tripleDesCfbNoHexExpected, input},
	}

	for index, test := range hexTests {
		e := Decrypt.FromHexString(test.input).By3Des(tripleDesCfbCipher(No))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}

	base32Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{tripleDesCfbNoBase32Expected, input},
	}

	for index, test := range base32Tests {
		e := Decrypt.FromBase32String(test.input).By3Des(tripleDesCfbCipher(No))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}

	base64Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{tripleDesCfbNoBase64Expected, input},
	}

	for index, test := range base64Tests {
		e := Decrypt.FromBase64String(test.input).By3Des(tripleDesCfbCipher(No))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestEncrypt_By3Des_CFB_ZeroPadding(t *testing.T) {
	hexTests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{tripleDesCfbInput, tripleDesCfbZeroHexExpected},
	}

	for index, test := range hexTests {
		e := Encrypt.FromString(test.input).By3Des(tripleDesCfbCipher(Zero))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToHexString(), "Current test index is "+strconv.Itoa(index))
	}

	base32Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{tripleDesCfbInput, tripleDesCfbZeroBase32Expected},
	}

	for index, test := range base32Tests {
		e := Encrypt.FromString(test.input).By3Des(tripleDesCfbCipher(Zero))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToBase32String(), "Current test index is "+strconv.Itoa(index))
	}

	base64Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{tripleDesCfbInput, tripleDesCfbZeroBase64Expected},
	}

	for index, test := range base64Tests {
		e := Encrypt.FromString(test.input).By3Des(tripleDesCfbCipher(Zero))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToBase64String(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestDecrypt_By3Des_CFB_ZeroPadding(t *testing.T) {
	hexTests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{tripleDesCfbZeroHexExpected, tripleDesCfbInput},
	}

	for index, test := range hexTests {
		e := Decrypt.FromHexString(test.input).By3Des(tripleDesCfbCipher(Zero))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}

	base32Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{tripleDesCfbZeroBase32Expected, tripleDesCfbInput},
	}

	for index, test := range base32Tests {
		e := Decrypt.FromBase32String(test.input).By3Des(tripleDesCfbCipher(Zero))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}

	base64Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{tripleDesCfbZeroBase64Expected, tripleDesCfbInput},
	}

	for index, test := range base64Tests {
		e := Decrypt.FromBase64String(test.input).By3Des(tripleDesCfbCipher(Zero))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestEncrypt_By3Des_CFB_PKCS5Padding(t *testing.T) {
	hexTests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{tripleDesCfbInput, tripleDesCfbPkcs5HexExpected},
	}

	for index, test := range hexTests {
		e := Encrypt.FromString(test.input).By3Des(tripleDesCfbCipher(PKCS5))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToHexString(), "Current test index is "+strconv.Itoa(index))
	}

	base32Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{tripleDesCfbInput, tripleDesCfbPkcs5Base32Expected},
	}

	for index, test := range base32Tests {
		e := Encrypt.FromString(test.input).By3Des(tripleDesCfbCipher(PKCS5))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToBase32String(), "Current test index is "+strconv.Itoa(index))
	}

	base64Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{tripleDesCfbInput, tripleDesCfbPkcs5Base64Expected},
	}

	for index, test := range base64Tests {
		e := Encrypt.FromString(test.input).By3Des(tripleDesCfbCipher(PKCS5))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToBase64String(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestDecrypt_By3Des_CFB_PKCS5Padding(t *testing.T) {
	hexTests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{tripleDesCfbPkcs5HexExpected, tripleDesCfbInput},
	}

	for index, test := range hexTests {
		e := Decrypt.FromHexString(test.input).By3Des(tripleDesCfbCipher(PKCS5))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}

	base32Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{tripleDesCfbPkcs5Base32Expected, tripleDesCfbInput},
	}

	for index, test := range base32Tests {
		e := Decrypt.FromBase32String(test.input).By3Des(tripleDesCfbCipher(PKCS5))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}

	base64Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{tripleDesCfbPkcs5Base64Expected, tripleDesCfbInput},
	}

	for index, test := range base64Tests {
		e := Decrypt.FromBase64String(test.input).By3Des(tripleDesCfbCipher(PKCS5))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestEncrypt_By3Des_CFB_PKCS7Padding(t *testing.T) {
	hexTests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{tripleDesCfbInput, tripleDesCfbPkcs7HexExpected},
	}

	for index, test := range hexTests {
		e := Encrypt.FromString(test.input).By3Des(tripleDesCfbCipher(PKCS7))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToHexString(), "Current test index is "+strconv.Itoa(index))
	}

	base32Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{tripleDesCfbInput, tripleDesCfbPkcs7Base32Expected},
	}

	for index, test := range base32Tests {
		e := Encrypt.FromString(test.input).By3Des(tripleDesCfbCipher(PKCS7))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToBase32String(), "Current test index is "+strconv.Itoa(index))
	}

	base64Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{tripleDesCfbInput, tripleDesCfbPkcs7Base64Expected},
	}

	for index, test := range base64Tests {
		e := Encrypt.FromString(test.input).By3Des(tripleDesCfbCipher(PKCS7))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToBase64String(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestDecrypt_By3Des_CFB_PKCS7Padding(t *testing.T) {
	hexTests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{tripleDesCfbPkcs7HexExpected, tripleDesCfbInput},
	}

	for index, test := range hexTests {
		e := Decrypt.FromHexString(test.input).By3Des(tripleDesCfbCipher(PKCS7))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}

	base32Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{tripleDesCfbPkcs7Base32Expected, tripleDesCfbInput},
	}

	for index, test := range base32Tests {
		e := Decrypt.FromBase32String(test.input).By3Des(tripleDesCfbCipher(PKCS7))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}

	base64Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{tripleDesCfbPkcs7Base64Expected, tripleDesCfbInput},
	}

	for index, test := range base64Tests {
		e := Decrypt.FromBase64String(test.input).By3Des(tripleDesCfbCipher(PKCS7))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}
}
