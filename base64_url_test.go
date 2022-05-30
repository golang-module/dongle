package dongle

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	base64UrlInput    = "www.gouguoyin.cn"
	base64UrlExpected = "d3d3LmdvdWd1b3lpbi5jbg=="
)

func TestEncode_ByBase64URL_FromStringToString(t *testing.T) {
	tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{base64UrlInput, base64UrlExpected},
	}

	for index, test := range tests {
		e := Encode.FromString(test.input).ByBase64URL()
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToString(), "Current test id is "+strconv.Itoa(index))
	}
}

func TestDecode_ByBase64URL_FromStringToString(t *testing.T) {
	tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{base64UrlExpected, base64UrlInput},
	}

	for index, test := range tests {
		d := Decode.FromString(test.input).ByBase64URL()
		assert.Nil(t, d.Error)
		assert.Equal(t, test.expected, d.ToString(), "Current test id is "+strconv.Itoa(index))
	}
}

func TestEncode_ByBase64URL_FromBytesToBytes(t *testing.T) {
	tests := []struct {
		input    []byte // 输入值
		expected []byte // 期望值
	}{
		{[]byte(""), []byte("")},
		{[]byte(base64UrlInput), []byte(base64UrlExpected)},
	}

	for index, test := range tests {
		e := Encode.FromBytes(test.input).ByBase64URL()
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToBytes(), "Current test id is "+strconv.Itoa(index))
	}
}

func TestDecode_ByBase64URL_FromBytesToBytes(t *testing.T) {
	tests := []struct {
		input    []byte // 输入值
		expected []byte // 期望值
	}{
		{[]byte(""), []byte("")},
		{[]byte(base64UrlExpected), []byte(base64UrlInput)},
	}

	for index, test := range tests {
		d := Decode.FromBytes(test.input).ByBase64URL()
		assert.Nil(t, d.Error)
		assert.Equal(t, test.expected, d.ToBytes(), "Current test id is "+strconv.Itoa(index))
	}
}
