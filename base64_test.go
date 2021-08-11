package dongle

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncode_ByBase64_FromString(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{"hello world", "aGVsbG8gd29ybGQ="},
	}

	for index, test := range tests {
		e := Encode.FromString(test.input).ByBase64()
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToString(), "Current test id is "+strconv.Itoa(index))
	}
}

func TestEncode_ByBase64_FromBytes(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    []byte // 输入值
		expected []byte // 期望值
	}{
		{[]byte(""), []byte("")},
		{[]byte("hello world"), []byte("aGVsbG8gd29ybGQ=")},
	}

	for index, test := range tests {
		e := Encode.FromBytes(test.input).ByBase64()
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToBytes(), "Current test id is "+strconv.Itoa(index))
	}
}

func TestEncode_ByBase64_Multiple(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{"hello world", "YUdWc2JHOGdkMjl5YkdRPQ=="},
	}

	for index, test := range tests {
		e := Encode.FromString(test.input).ByBase64().ByBase64()
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToString(), "Current test id is "+strconv.Itoa(index))
	}
}

func TestDecode_ByBase64_FromString(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{"aGVsbG8gd29ybGQ=", "hello world"},
	}

	for index, test := range tests {
		e := Decode.FromString(test.input).ByBase64()
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToString(), "Current test id is "+strconv.Itoa(index))
	}
}

func TestDecode_ByBase64_FromBytes(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    []byte // 输入值
		expected []byte // 期望值
	}{
		{[]byte(""), []byte("")},
		{[]byte("aGVsbG8gd29ybGQ="), []byte("hello world")},
	}

	for index, test := range tests {
		e := Decode.FromBytes(test.input).ByBase64()
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToBytes(), "Current test id is "+strconv.Itoa(index))
	}
}

func TestDecode_ByBase64_Multiple(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{"YUdWc2JHOGdkMjl5YkdRPQ==", "hello world"},
	}

	for index, test := range tests {
		e := Decode.FromString(test.input).ByBase64().ByBase64()
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToString(), "Current test id is "+strconv.Itoa(index))
	}
}
