package dongle

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

var (
	desCfbInput = "hello world"
	desCfbKey   = "12345678"
	desCfbIV    = "12345678"

	// DES-CFB-NoPadding
	desCfbNoHexExpected = "a7e231bc4de3bbb1b64cda645f2ae161"
	// desCfbNoHexExpected    = "a7e231bc4de3bbb1b64cda645f2ae161" // https://oktools.net/des
	// desCfbNoHexExpected    = "a747a55a077088a8eb29a927317f4d7e" // http://tool.chacuo.net/cryptdes
	desCfbNoBase32Expected = "U7RDDPCN4O53DNSM3JSF6KXBME======"
	desCfbNoBase64Expected = "p+IxvE3ju7G2TNpkXyrhYQ=="

	// DES-CFB-ZeroPadding
	desCfbZeroHexExpected = "feb56ee417f5fbe61a4e572c995b5b2a"
	// desCfbZeroHexExpected    = "feb56ee417f5fbe61a4e572c995b5b2a" // https://oktools.net/des
	// desCfbZeroHexExpected    = "feb56ee417f5fbe61a4e572c995b5b2a" // https://tool.lmeee.com/jiami/des
	// desCfbZeroHexExpected    = "fe0c815405fd275284e6f6ce79114dfa" // http://tool.chacuo.net/cryptdes
	desCfbZeroBase32Expected = "722W5ZAX6X56MGSOK4WJSW23FI======"
	desCfbZeroBase64Expected = "/rVu5Bf1++YaTlcsmVtbKg=="

	// DES-CFB-PKCS5Padding
	desCfbPkcs5HexExpected = "feb56ee417f5fbe61a4e57299c5e5e2f"
	// desCfbPkcs5HexExpected    = "feb56ee417f5fbe61a4e57299c5e5e2f" // https://oktools.net/des
	// desCfbPkcs5HexExpected    = "feb56ee417f5fbe61a4e57299c5e5e2f" // https://tool.lmeee.com/jiami/des
	// desCfbPkcs5HexExpected    = "fe0c815405fd275284e6f6cb08599485" // http://tool.chacuo.net/cryptdes
	desCfbPkcs5Base32Expected = "722W5ZAX6X56MGSOK4UZYXS6F4======"
	desCfbPkcs5Base64Expected = "/rVu5Bf1++YaTlcpnF5eLw=="

	// DES-CFB-PKCS7Padding
	desCfbPkcs7HexExpected = "feb56ee417f5fbe61a4e57299c5e5e2f"
	// desCfbPkcs7HexExpected    = "feb56ee417f5fbe61a4e57299c5e5e2f" // https://oktools.net/des
	// desCfbPkcs7HexExpected    = "feb56ee417f5fbe61a4e57299c5e5e2f" // https://tool.lmeee.com/jiami/des
	// desCfbPkcs7HexExpected    = "fe0c815405fd275284e6f6cb08599485" // http://tool.chacuo.net/cryptdes
	desCfbPkcs7Base32Expected = "722W5ZAX6X56MGSOK4UZYXS6F4======"
	desCfbPkcs7Base64Expected = "/rVu5Bf1++YaTlcpnF5eLw=="
)

func desCfbCipher(padding string) *Cipher {
	cipher := NewCipher()
	cipher.SetMode(CFB)
	cipher.SetPadding(padding)
	cipher.SetKey(desCfbKey)
	cipher.SetIV(desCfbIV)
	return cipher
}

func TestEncrypt_ByDes_CFB_NoPadding(t *testing.T) {
	input := "12345678asdfghjk"

	hexTests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{input, desCfbNoHexExpected},
	}

	for index, test := range hexTests {
		e := Encrypt.FromString(test.input).ByDes(desCfbCipher(No))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToHexString(), "Current test index is "+strconv.Itoa(index))
	}

	base32Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{input, desCfbNoBase32Expected},
	}

	for index, test := range base32Tests {
		e := Encrypt.FromString(test.input).ByDes(desCfbCipher(No))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToBase32String(), "Current test index is "+strconv.Itoa(index))
	}

	base64Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{input, desCfbNoBase64Expected},
	}

	for index, test := range base64Tests {
		e := Encrypt.FromString(test.input).ByDes(desCfbCipher(No))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToBase64String(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestDecrypt_ByDes_CFB_NoPadding(t *testing.T) {
	input := "12345678asdfghjk"

	hexTests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{desCfbNoHexExpected, input},
	}

	for index, test := range hexTests {
		e := Decrypt.FromHexString(test.input).ByDes(desCfbCipher(No))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}

	base32Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{desCfbNoBase32Expected, input},
	}

	for index, test := range base32Tests {
		e := Decrypt.FromBase32String(test.input).ByDes(desCfbCipher(No))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}

	base64Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{desCfbNoBase64Expected, input},
	}

	for index, test := range base64Tests {
		e := Decrypt.FromBase64String(test.input).ByDes(desCfbCipher(No))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestEncrypt_ByDes_CFB_ZeroPadding(t *testing.T) {
	hexTests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{desCfbInput, desCfbZeroHexExpected},
	}

	for index, test := range hexTests {
		e := Encrypt.FromString(test.input).ByDes(desCfbCipher(Zero))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToHexString(), "Current test index is "+strconv.Itoa(index))
	}

	base32Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{desCfbInput, desCfbZeroBase32Expected},
	}

	for index, test := range base32Tests {
		e := Encrypt.FromString(test.input).ByDes(desCfbCipher(Zero))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToBase32String(), "Current test index is "+strconv.Itoa(index))
	}

	base64Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{desCfbInput, desCfbZeroBase64Expected},
	}

	for index, test := range base64Tests {
		e := Encrypt.FromString(test.input).ByDes(desCfbCipher(Zero))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToBase64String(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestDecrypt_ByDes_CFB_ZeroPadding(t *testing.T) {
	hexTests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{desCfbZeroHexExpected, desCfbInput},
	}

	for index, test := range hexTests {
		e := Decrypt.FromHexString(test.input).ByDes(desCfbCipher(Zero))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}

	base32Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{desCfbZeroBase32Expected, desCfbInput},
	}

	for index, test := range base32Tests {
		e := Decrypt.FromBase32String(test.input).ByDes(desCfbCipher(Zero))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}

	base64Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{desCfbZeroBase64Expected, desCfbInput},
	}

	for index, test := range base64Tests {
		e := Decrypt.FromBase64String(test.input).ByDes(desCfbCipher(Zero))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestEncrypt_ByDes_CFB_PKCS5Padding(t *testing.T) {
	hexTests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{desCfbInput, desCfbPkcs5HexExpected},
	}

	for index, test := range hexTests {
		e := Encrypt.FromString(test.input).ByDes(desCfbCipher(PKCS5))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToHexString(), "Current test index is "+strconv.Itoa(index))
	}

	base32Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{desCfbInput, desCfbPkcs5Base32Expected},
	}

	for index, test := range base32Tests {
		e := Encrypt.FromString(test.input).ByDes(desCfbCipher(PKCS5))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToBase32String(), "Current test index is "+strconv.Itoa(index))
	}

	base64Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{desCfbInput, desCfbPkcs5Base64Expected},
	}

	for index, test := range base64Tests {
		e := Encrypt.FromString(test.input).ByDes(desCfbCipher(PKCS5))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToBase64String(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestDecrypt_ByDes_CFB_PKCS5Padding(t *testing.T) {
	hexTests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{desCfbPkcs5HexExpected, desCfbInput},
	}

	for index, test := range hexTests {
		e := Decrypt.FromHexString(test.input).ByDes(desCfbCipher(PKCS5))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}

	base32Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{desCfbPkcs5Base32Expected, desCfbInput},
	}

	for index, test := range base32Tests {
		e := Decrypt.FromBase32String(test.input).ByDes(desCfbCipher(PKCS5))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}

	base64Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{desCfbPkcs5Base64Expected, desCfbInput},
	}

	for index, test := range base64Tests {
		e := Decrypt.FromBase64String(test.input).ByDes(desCfbCipher(PKCS5))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestEncrypt_ByDes_CFB_PKCS7Padding(t *testing.T) {
	hexTests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{desCfbInput, desCfbPkcs7HexExpected},
	}

	for index, test := range hexTests {
		e := Encrypt.FromString(test.input).ByDes(desCfbCipher(PKCS7))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToHexString(), "Current test index is "+strconv.Itoa(index))
	}

	base32Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{desCfbInput, desCfbPkcs7Base32Expected},
	}

	for index, test := range base32Tests {
		e := Encrypt.FromString(test.input).ByDes(desCfbCipher(PKCS7))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToBase32String(), "Current test index is "+strconv.Itoa(index))
	}

	base64Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{desCfbInput, desCfbPkcs7Base64Expected},
	}

	for index, test := range base64Tests {
		e := Encrypt.FromString(test.input).ByDes(desCfbCipher(PKCS7))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToBase64String(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestDecrypt_ByDes_CFB_PKCS7Padding(t *testing.T) {
	hexTests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{desCfbPkcs7HexExpected, desCfbInput},
	}

	for index, test := range hexTests {
		e := Decrypt.FromHexString(test.input).ByDes(desCfbCipher(PKCS7))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}

	base32Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{desCfbPkcs7Base32Expected, desCfbInput},
	}

	for index, test := range base32Tests {
		e := Decrypt.FromBase32String(test.input).ByDes(desCfbCipher(PKCS7))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}

	base64Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{desCfbPkcs7Base64Expected, desCfbInput},
	}

	for index, test := range base64Tests {
		e := Decrypt.FromBase64String(test.input).ByDes(desCfbCipher(PKCS7))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}
}
