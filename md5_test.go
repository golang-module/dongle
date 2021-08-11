package dongle

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncrypt_ByMd5_FromString_ToString(t *testing.T) {
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
		{"hello world", "hex", "5eb63bbbe01eeed093cb22bb8f5acdc3"},
		{"hello world", "base32", "L23DXO7AD3XNBE6LEK5Y6WWNYM======"},
		{"hello world", "base64", "XrY7u+Ae7tCTyyK7j1rNww=="},
	}

	for index, test := range tests {
		e := Encrypt.FromString(test.input).ByMd5()
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToString(test.encodeMode), "Current test index is "+strconv.Itoa(index))
	}
}

func TestEncrypt_ByMd5_FromString_ToBytes(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input      string // 输入值
		encodeMode string // 编码模式
		expected   []byte // 期望值
	}{
		{"", "hex", []byte("")},
		{"", "base32", []byte("")},
		{"", "base64", []byte("")},
		{"", "xxx", []byte("")},
		{"hello world", "xxx", []byte("")},
		{"hello world", "hex", []byte("5eb63bbbe01eeed093cb22bb8f5acdc3")},
		{"hello world", "base32", []byte("L23DXO7AD3XNBE6LEK5Y6WWNYM======")},
		{"hello world", "base64", []byte("XrY7u+Ae7tCTyyK7j1rNww==")},
	}

	for index, test := range tests {
		e := Encrypt.FromString(test.input).ByMd5()
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToBytes(test.encodeMode), "Current test index is "+strconv.Itoa(index))
	}
}

func TestEncrypt_ByMd5_FromBytes_ToString(t *testing.T) {
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
		{[]byte("hello world"), "hex", "5eb63bbbe01eeed093cb22bb8f5acdc3"},
		{[]byte("hello world"), "base32", "L23DXO7AD3XNBE6LEK5Y6WWNYM======"},
		{[]byte("hello world"), "base64", "XrY7u+Ae7tCTyyK7j1rNww=="},
	}

	for index, test := range tests {
		e := Encrypt.FromBytes(test.input).ByMd5()
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToString(test.encodeMode), "Current test index is "+strconv.Itoa(index))
	}
}

func TestEncrypt_ByMd5_FromBytes_ToBytes(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input      []byte // 输入值
		encodeMode string // 编码模式
		expected   []byte // 期望值
	}{
		{[]byte(""), "hex", []byte("")},
		{[]byte(""), "base32", []byte("")},
		{[]byte(""), "base64", []byte("")},
		{[]byte(""), "xxx", []byte("")},
		{[]byte("hello world"), "xxx", []byte("")},
		{[]byte("hello world"), "hex", []byte("5eb63bbbe01eeed093cb22bb8f5acdc3")},
		{[]byte("hello world"), "base32", []byte("L23DXO7AD3XNBE6LEK5Y6WWNYM======")},
		{[]byte("hello world"), "base64", []byte("XrY7u+Ae7tCTyyK7j1rNww==")},
	}

	for index, test := range tests {
		e := Encrypt.FromBytes(test.input).ByMd5()
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToBytes(test.encodeMode), "Current test index is "+strconv.Itoa(index))
	}
}

func TestEncrypt_ByMd5_FromFile(t *testing.T) {
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
		{"./LICENSE", "xxx", ""},
		{"./LICENSE", "hex", "014f03f9025ea81a8a0e9734be540c53"},
		{"./LICENSE", "base32", "AFHQH6ICL2UBVCQOS42L4VAMKM======"},
		{"./LICENSE", "base64", "AU8D+QJeqBqKDpc0vlQMUw=="},
	}

	for index, test := range tests {
		e := Encrypt.FromFile(test.input).ByMd5()
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToString(test.encodeMode), "Current test index is "+strconv.Itoa(index))
	}
}

func TestError_ByMd5_FromFile(t *testing.T) {
	file := "./demo.txt"
	e := Encrypt.FromFile(file).ByMd5()
	assert.Equal(t, invalidFileError(file), e.Error, "Should catch an exception")
}
