package dongle

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncode_ByBase32_FromString(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{"hello world", "NBSWY3DPEB3W64TMMQ======"},
	}

	for index, test := range tests {
		e := Encode.FromString(test.input).ByBase32()
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToString(), "Current test id is "+strconv.Itoa(index))
	}
}

func TestEncode_ByBase32_FromBytes(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    []byte // 输入值
		expected []byte // 期望值
	}{
		{[]byte(""), []byte("")},
		{[]byte("hello world"), []byte("NBSWY3DPEB3W64TMMQ======")},
	}

	for index, test := range tests {
		e := Encode.FromBytes(test.input).ByBase32()
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToBytes(), "Current test id is "+strconv.Itoa(index))
	}
}

func TestEncode_ByBase32_Multiple(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{"hello world", "JZBFGV2ZGNCFARKCGNLTMNCUJVGVCPJ5HU6T2PI="},
	}

	for index, test := range tests {
		e := Encode.FromString(test.input).ByBase32().ByBase32()
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToString(), "Current test id is "+strconv.Itoa(index))
	}
}

func TestDecode_ByBase32_FromString(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{"NBSWY3DPEB3W64TMMQ======", "hello world"},
	}

	for index, test := range tests {
		e := Decode.FromString(test.input).ByBase32()
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToString(), "Current test id is "+strconv.Itoa(index))
	}
}

func TestDecode_ByBase32_FromBytes(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    []byte // 输入值
		expected []byte // 期望值
	}{
		{[]byte(""), []byte("")},
		{[]byte("NBSWY3DPEB3W64TMMQ======"), []byte("hello world")},
	}

	for index, test := range tests {
		e := Decode.FromBytes(test.input).ByBase32()
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToBytes(), "Current test id is "+strconv.Itoa(index))
	}
}

func TestDecode_ByBase32_Multiple(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{"JZBFGV2ZGNCFARKCGNLTMNCUJVGVCPJ5HU6T2PI=", "hello world"},
	}

	for index, test := range tests {
		e := Decode.FromString(test.input).ByBase32().ByBase32()
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToString(), "Current test id is "+strconv.Itoa(index))
	}
}
