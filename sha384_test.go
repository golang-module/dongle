package dongle

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncrypt_BySha384_FromString_ToString(t *testing.T) {
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
		{"hello world", "hex", "fdbd8e75a67f29f701a4e040385e2e23986303ea10239211af907fcbb83578b3e417cb71ce646efd0819dd8c088de1bd"},
		{"hello world", "base32", "7W6Y45NGP4U7OANE4BADQXROEOMGGA7KCARZEENPSB74XOBVPCZ6IF6LOHHGI3X5BAM53DAIRXQ32==="},
		{"hello world", "base64", "/b2OdaZ/KfcBpOBAOF4uI5hjA+oQI5IRr5B/y7g1eLPkF8txzmRu/QgZ3YwIjeG9"},
	}

	for index, test := range tests {
		e := Encrypt.FromString(test.input).BySha384()
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToString(test.encodeMode), "Current test index is "+strconv.Itoa(index))
	}
}

func TestEncrypt_BySha384_FromString_ToBytes(t *testing.T) {
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
		{"hello world", "hex", []byte("fdbd8e75a67f29f701a4e040385e2e23986303ea10239211af907fcbb83578b3e417cb71ce646efd0819dd8c088de1bd")},
		{"hello world", "base32", []byte("7W6Y45NGP4U7OANE4BADQXROEOMGGA7KCARZEENPSB74XOBVPCZ6IF6LOHHGI3X5BAM53DAIRXQ32===")},
		{"hello world", "base64", []byte("/b2OdaZ/KfcBpOBAOF4uI5hjA+oQI5IRr5B/y7g1eLPkF8txzmRu/QgZ3YwIjeG9")},
	}

	for index, test := range tests {
		e := Encrypt.FromString(test.input).BySha384()
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToBytes(test.encodeMode), "Current test index is "+strconv.Itoa(index))
	}
}

func TestEncrypt_BySha384_FromBytes_ToString(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input      []byte // 输入值
		encodeMode string // 编码模式
		expected   string // 期望值
	}{
		{[]byte(""), "hex", ""},
		{[]byte(""), "base32", ""},
		{[]byte(""), "base64", ""},
		{[]byte("hello world"), "xxx", ""},
		{[]byte("hello world"), "hex", "fdbd8e75a67f29f701a4e040385e2e23986303ea10239211af907fcbb83578b3e417cb71ce646efd0819dd8c088de1bd"},
		{[]byte("hello world"), "base32", "7W6Y45NGP4U7OANE4BADQXROEOMGGA7KCARZEENPSB74XOBVPCZ6IF6LOHHGI3X5BAM53DAIRXQ32==="},
		{[]byte("hello world"), "base64", "/b2OdaZ/KfcBpOBAOF4uI5hjA+oQI5IRr5B/y7g1eLPkF8txzmRu/QgZ3YwIjeG9"},
	}

	for index, test := range tests {
		e := Encrypt.FromBytes(test.input).BySha384()
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToString(test.encodeMode), "Current test index is "+strconv.Itoa(index))
	}
}

func TestEncrypt_BySha384_FromBytes_ToBytes(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input      []byte // 输入值
		encodeMode string // 编码模式
		expected   []byte // 期望值
	}{
		{[]byte(""), "hex", []byte{}},
		{[]byte(""), "base32", []byte{}},
		{[]byte(""), "base64", []byte{}},
		{[]byte("hello world"), "xxx", []byte("")},
		{[]byte("hello world"), "hex", []byte("fdbd8e75a67f29f701a4e040385e2e23986303ea10239211af907fcbb83578b3e417cb71ce646efd0819dd8c088de1bd")},
		{[]byte("hello world"), "base32", []byte("7W6Y45NGP4U7OANE4BADQXROEOMGGA7KCARZEENPSB74XOBVPCZ6IF6LOHHGI3X5BAM53DAIRXQ32===")},
		{[]byte("hello world"), "base64", []byte("/b2OdaZ/KfcBpOBAOF4uI5hjA+oQI5IRr5B/y7g1eLPkF8txzmRu/QgZ3YwIjeG9")},
	}

	for index, test := range tests {
		e := Encrypt.FromBytes(test.input).BySha384()
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToBytes(test.encodeMode), "Current test index is "+strconv.Itoa(index))
	}
}

func TestError_BySha384_FromFile(t *testing.T) {
	file := "./demo.txt"
	e := Encrypt.FromFile(file).BySha384()
	assert.Equal(t, invalidFileError(file), e.Error, "Should catch an exception")
}
