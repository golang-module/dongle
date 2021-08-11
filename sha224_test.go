package dongle

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncrypt_BySha224_FromString_ToString(t *testing.T) {
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
		{"hello world", "hex", "2f05477fc24bb4faefd86517156dafdecec45b8ad3cf2522a563582b"},
		{"hello world", "base32", "F4CUO76CJO2PV36YMULRK3NP33HMIW4K2PHSKIVFMNMCW==="},
		{"hello world", "base64", "LwVHf8JLtPrv2GUXFW2v3s7EW4rTzyUipWNYKw=="},
	}

	for index, test := range tests {
		e := Encrypt.FromString(test.input).BySha224()
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToString(test.encodeMode), "Current test index is "+strconv.Itoa(index))
	}
}

func TestEncrypt_BySha224_FromString_ToBytes(t *testing.T) {
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
		{"hello world", "hex", []byte("2f05477fc24bb4faefd86517156dafdecec45b8ad3cf2522a563582b")},
		{"hello world", "base32", []byte("F4CUO76CJO2PV36YMULRK3NP33HMIW4K2PHSKIVFMNMCW===")},
		{"hello world", "base64", []byte("LwVHf8JLtPrv2GUXFW2v3s7EW4rTzyUipWNYKw==")},
	}

	for index, test := range tests {
		e := Encrypt.FromString(test.input).BySha224()
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToBytes(test.encodeMode), "Current test index is "+strconv.Itoa(index))
	}
}

func TestEncrypt_BySha224_FromBytes_ToString(t *testing.T) {
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
		{[]byte("hello world"), "hex", "2f05477fc24bb4faefd86517156dafdecec45b8ad3cf2522a563582b"},
		{[]byte("hello world"), "base32", "F4CUO76CJO2PV36YMULRK3NP33HMIW4K2PHSKIVFMNMCW==="},
		{[]byte("hello world"), "base64", "LwVHf8JLtPrv2GUXFW2v3s7EW4rTzyUipWNYKw=="},
	}

	for index, test := range tests {
		e := Encrypt.FromBytes(test.input).BySha224()
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToString(test.encodeMode), "Current test index is "+strconv.Itoa(index))
	}
}

func TestEncrypt_BySha224_FromBytes_ToBytes(t *testing.T) {
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
		{[]byte("hello world"), "hex", []byte("2f05477fc24bb4faefd86517156dafdecec45b8ad3cf2522a563582b")},
		{[]byte("hello world"), "base32", []byte("F4CUO76CJO2PV36YMULRK3NP33HMIW4K2PHSKIVFMNMCW===")},
		{[]byte("hello world"), "base64", []byte("LwVHf8JLtPrv2GUXFW2v3s7EW4rTzyUipWNYKw==")},
	}

	for index, test := range tests {
		e := Encrypt.FromBytes(test.input).BySha224()
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToBytes(test.encodeMode), "Current test index is "+strconv.Itoa(index))
	}
}

func TestError_BySha224_FromFile(t *testing.T) {
	file := "./demo.txt"
	e := Encrypt.FromFile(file).BySha224()
	assert.Equal(t, invalidFileError(file), e.Error, "Should catch an exception")
}
