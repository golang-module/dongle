package dongle

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncode_ByBase58_FromString(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{"hello world", "StV1DL6CwTryKyV"},
		{"0123456789", "3i37NcgooY8f1S"},
	}

	for index, test := range tests {
		e := Encode.FromString(test.input).ByBase58()
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToString(), "Current test id is "+strconv.Itoa(index))
	}
}

func TestEncode_ByBase58_FromBytes(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    []byte // 输入值
		expected []byte // 期望值
	}{
		{[]byte(""), []byte("")},
		{[]byte("hello world"), []byte("StV1DL6CwTryKyV")},
	}

	for index, test := range tests {
		e := Encode.FromBytes(test.input).ByBase58()
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToBytes(), "Current test id is "+strconv.Itoa(index))
	}
}

func TestEncode_ByBase58_Multiple(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{"hello world", "3LRFsmWKLfsR7G5PqjytR"},
	}

	for index, test := range tests {
		e := Encode.FromString(test.input).ByBase58().ByBase58()
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToString(), "Current test id is "+strconv.Itoa(index))
	}
}

func TestDecode_ByBase58_FromString(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{"StV1DL6CwTryKyV", "hello world"},
	}

	for index, test := range tests {
		e := Decode.FromString(test.input).ByBase58()
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToString(), "Current test id is "+strconv.Itoa(index))
	}
}

func TestDecode_ByBase58_FromBytes(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    []byte // 输入值
		expected []byte // 期望值
	}{
		{[]byte(""), []byte("")},
		{[]byte("StV1DL6CwTryKyV"), []byte("hello world")},
	}

	for index, test := range tests {
		e := Decode.FromBytes(test.input).ByBase58()
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToBytes(), "Current test id is "+strconv.Itoa(index))
	}
}

func TestDecode_ByBase58_Multiple(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{"3LRFsmWKLfsR7G5PqjytR", "hello world"},
	}

	for index, test := range tests {
		e := Decode.FromString(test.input).ByBase58().ByBase58()
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToString(), "Current test id is "+strconv.Itoa(index))
	}
}
