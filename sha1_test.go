package dongle

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncrypt_BySha1_FromString_ToString(t *testing.T) {
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
		{"hello world", "hex", "2aae6c35c94fcfb415dbe95f408b9ce91ee846ed"},
		{"hello world", "base32", "FKXGYNOJJ7H3IFO35FPUBC445EPOQRXN"},
		{"hello world", "base64", "Kq5sNclPz7QV2+lfQIuc6R7oRu0="},
	}

	for index, test := range tests {
		e := Encrypt.FromString(test.input).BySha1()
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToString(test.encodeMode), "Current test index is "+strconv.Itoa(index))
	}
}

func TestEncrypt_BySha1_FromString_ToBytes(t *testing.T) {
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
		{"hello world", "hex", []byte("2aae6c35c94fcfb415dbe95f408b9ce91ee846ed")},
		{"hello world", "base32", []byte("FKXGYNOJJ7H3IFO35FPUBC445EPOQRXN")},
		{"hello world", "base64", []byte("Kq5sNclPz7QV2+lfQIuc6R7oRu0=")},
	}

	for index, test := range tests {
		e := Encrypt.FromString(test.input).BySha1()
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToBytes(test.encodeMode), "Current test index is "+strconv.Itoa(index))
	}
}

func TestEncrypt_BySha1_FromBytes_ToString(t *testing.T) {
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
		{[]byte("hello world"), "hex", "2aae6c35c94fcfb415dbe95f408b9ce91ee846ed"},
		{[]byte("hello world"), "base32", "FKXGYNOJJ7H3IFO35FPUBC445EPOQRXN"},
		{[]byte("hello world"), "base64", "Kq5sNclPz7QV2+lfQIuc6R7oRu0="},
	}

	for index, test := range tests {
		e := Encrypt.FromBytes(test.input).BySha1()
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToString(test.encodeMode), "Current test index is "+strconv.Itoa(index))
	}
}

func TestEncrypt_BySha1_FromBytes_ToBytes(t *testing.T) {
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
		{[]byte("hello world"), "hex", []byte("2aae6c35c94fcfb415dbe95f408b9ce91ee846ed")},
		{[]byte("hello world"), "base32", []byte("FKXGYNOJJ7H3IFO35FPUBC445EPOQRXN")},
		{[]byte("hello world"), "base64", []byte("Kq5sNclPz7QV2+lfQIuc6R7oRu0=")},
	}

	for index, test := range tests {
		e := Encrypt.FromBytes(test.input).BySha1()
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToBytes(test.encodeMode), "Current test index is "+strconv.Itoa(index))
	}
}

func TestError_BySha1_FromFile(t *testing.T) {
	file := "./demo.txt"
	e := Encrypt.FromFile(file).BySha1()
	assert.Equal(t, invalidFileError(file), e.Error, "Should catch an exception")
}
