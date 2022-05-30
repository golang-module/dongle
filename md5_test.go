package dongle

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	md5Input          = "hello world"
	md5HexExpected    = "5eb63bbbe01eeed093cb22bb8f5acdc3"
	md5Base32Expected = "L23DXO7AD3XNBE6LEK5Y6WWNYM======"
	md5Base64Expected = "XrY7u+Ae7tCTyyK7j1rNww=="
)

func TestEncrypt_ByMd5_FromStringToString(t *testing.T) {
	hexTests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{md5Input, md5HexExpected},
	}

	for index, test := range hexTests {
		e := Encrypt.FromString(test.input).ByMd5()
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToHexString(), "Current test index is "+strconv.Itoa(index))
	}

	base32Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{md5Input, md5Base32Expected},
	}

	for index, test := range base32Tests {
		e := Encrypt.FromString(test.input).ByMd5()
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToBase32String(), "Current test index is "+strconv.Itoa(index))
	}

	base64Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{md5Input, md5Base64Expected},
	}

	for index, test := range base64Tests {
		e := Encrypt.FromString(test.input).ByMd5()
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToBase64String(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestEncrypt_ByMd5_FromBytesToBytes(t *testing.T) {
	hexTests := []struct {
		input    []byte // 输入值
		expected []byte // 期望值
	}{
		{[]byte(""), []byte("")},
		{[]byte(md5Input), []byte(md5HexExpected)},
	}

	for index, test := range hexTests {
		e := Encrypt.FromBytes(test.input).ByMd5()
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToHexBytes(), "Current test index is "+strconv.Itoa(index))
	}

	base32Tests := []struct {
		input    []byte // 输入值
		expected []byte // 期望值
	}{
		{[]byte(""), []byte("")},
		{[]byte(md5Input), []byte(md5Base32Expected)},
	}

	for index, test := range base32Tests {
		e := Encrypt.FromBytes(test.input).ByMd5()
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToBase32Bytes(), "Current test index is "+strconv.Itoa(index))
	}

	base64Tests := []struct {
		input    []byte // 输入值
		expected []byte // 期望值
	}{
		{[]byte(""), []byte("")},
		{[]byte(md5Input), []byte(md5Base64Expected)},
	}

	for index, test := range base64Tests {
		e := Encrypt.FromBytes(test.input).ByMd5()
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToBase64Bytes(), "Current test index is "+strconv.Itoa(index))
	}

}

func TestEncrypt_ByMd5_FromFileToString(t *testing.T) {
	hexTests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{"./LICENSE", "c7549a87626c65bd0970e32e02f130d7"},
	}

	for index, test := range hexTests {
		e := Encrypt.FromFile(test.input).ByMd5()
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToHexString(), "Current test index is "+strconv.Itoa(index))
	}

	base32Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{"./LICENSE", "Y5KJVB3CNRS32CLQ4MXAF4JQ24======"},
	}

	for index, test := range base32Tests {
		e := Encrypt.FromFile(test.input).ByMd5()
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToBase32String(), "Current test index is "+strconv.Itoa(index))
	}

	base64Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{"./LICENSE", "x1Sah2JsZb0JcOMuAvEw1w=="},
	}

	for index, test := range base64Tests {
		e := Encrypt.FromFile(test.input).ByMd5()
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToBase64String(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestEncrypt_ByMd5_FromFileToBytes(t *testing.T) {
	hexTests := []struct {
		input    []byte // 输入值
		expected []byte // 期望值
	}{
		{[]byte(""), []byte("")},
		{[]byte("./LICENSE"), []byte("c7549a87626c65bd0970e32e02f130d7")},
	}

	for index, test := range hexTests {
		e := Encrypt.FromFile(test.input).ByMd5()
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToHexBytes(), "Current test index is "+strconv.Itoa(index))
	}

	base32Tests := []struct {
		input    []byte // 输入值
		expected []byte // 期望值
	}{
		{[]byte(""), []byte("")},
		{[]byte("./LICENSE"), []byte("Y5KJVB3CNRS32CLQ4MXAF4JQ24======")},
	}

	for index, test := range base32Tests {
		e := Encrypt.FromFile(test.input).ByMd5()
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToBase32Bytes(), "Current test index is "+strconv.Itoa(index))
	}

	base64Tests := []struct {
		input    []byte // 输入值
		expected []byte // 期望值
	}{
		{[]byte(""), []byte("")},
		{[]byte("./LICENSE"), []byte("x1Sah2JsZb0JcOMuAvEw1w==")},
	}

	for index, test := range base64Tests {
		e := Encrypt.FromFile(test.input).ByMd5()
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToBase64Bytes(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestError_ByMd5(t *testing.T) {
	file := "./demo.txt"
	e := Encrypt.FromFile(file).ByMd5()
	assert.Equal(t, invalidFileError(file), e.Error, "Should catch an exception")
}
