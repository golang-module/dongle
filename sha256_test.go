package dongle

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncrypt_BySha256_FromString_ToString(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input      string // 输入值
		encodeMode string // 编码模式
		expected   string // 期望值
	}{
		{"", "hex", ""},
		{"", "base32", ""},
		{"", "base64", ""},
		{"", "xxx", ""},
		{"hello world", "xxx", ""},
		{"hello world", "hex", "b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9"},
		{"hello world", "base32", "XFGSPOMTJU7ARJJOKLL5U7NL7LCIJ37DPJJYB3UQRD32ZYXPZXUQ===="},
		{"hello world", "base64", "uU0nuZNNPgilLlLX2n2r+sSE7+N6U4DukIj3rOLvzek="},
	}

	for index, test := range tests {
		e := Encrypt.FromString(test.input).BySha256()
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToString(test.encodeMode), "Current test index is "+strconv.Itoa(index))
	}
}

func TestEncrypt_BySha256_FromString_ToBytes(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input      string // 输入值
		encodeMode string // 编码模式
		expected   []byte // 期望值
	}{
		{"", "hex", []byte{}},
		{"", "base32", []byte{}},
		{"", "base64", []byte{}},
		{"", "xxx", []byte{}},
		{"hello world", "xxx", []byte("")},
		{"hello world", "hex", []byte("b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9")},
		{"hello world", "base32", []byte("XFGSPOMTJU7ARJJOKLL5U7NL7LCIJ37DPJJYB3UQRD32ZYXPZXUQ====")},
		{"hello world", "base64", []byte("uU0nuZNNPgilLlLX2n2r+sSE7+N6U4DukIj3rOLvzek=")},
	}

	for index, test := range tests {
		e := Encrypt.FromString(test.input).BySha256()
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToBytes(test.encodeMode), "Current test index is "+strconv.Itoa(index))
	}
}

func TestEncrypt_BySha256_FromBytes_ToString(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input      []byte // 输入值
		encodeMode string // 编码模式
		expected   string // 期望值
	}{
		{[]byte(""), "hex", ""},
		{[]byte(""), "base32", ""},
		{[]byte(""), "base64", ""},
		{[]byte(""), "xxx", ""},
		{[]byte("hello world"), "xxx", ""},
		{[]byte("hello world"), "hex", "b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9"},
		{[]byte("hello world"), "base32", "XFGSPOMTJU7ARJJOKLL5U7NL7LCIJ37DPJJYB3UQRD32ZYXPZXUQ===="},
		{[]byte("hello world"), "base64", "uU0nuZNNPgilLlLX2n2r+sSE7+N6U4DukIj3rOLvzek="},
	}

	for index, test := range tests {
		e := Encrypt.FromBytes(test.input).BySha256()
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToString(test.encodeMode), "Current test index is "+strconv.Itoa(index))
	}
}

func TestEncrypt_BySha256_FromBytes_ToBytes(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input      []byte // 输入值
		encodeMode string // 编码模式
		expected   []byte // 期望值
	}{
		{[]byte(""), "hex", []byte{}},
		{[]byte(""), "base32", []byte{}},
		{[]byte(""), "base64", []byte{}},
		{[]byte(""), "xxx", []byte{}},
		{[]byte("hello world"), "xxx", []byte("")},
		{[]byte("hello world"), "hex", []byte("b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9")},
		{[]byte("hello world"), "base32", []byte("XFGSPOMTJU7ARJJOKLL5U7NL7LCIJ37DPJJYB3UQRD32ZYXPZXUQ====")},
		{[]byte("hello world"), "base64", []byte("uU0nuZNNPgilLlLX2n2r+sSE7+N6U4DukIj3rOLvzek=")},
	}

	for index, test := range tests {
		e := Encrypt.FromBytes(test.input).BySha256()
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToBytes(test.encodeMode), "Current test index is "+strconv.Itoa(index))
	}
}

func TestError_BySha256_FromFile(t *testing.T) {
	file := "./demo.txt"
	e := Encrypt.FromFile(file).BySha256()
	assert.Equal(t, invalidFileError(file), e.Error, "Should catch an exception")
}
