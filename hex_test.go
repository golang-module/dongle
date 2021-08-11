package dongle

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncode_ByHex_FromString(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{"hello world", "68656c6c6f20776f726c64"},
	}

	for index, test := range tests {
		e := Encode.FromString(test.input).ByHex()
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToString(), "Current test id is "+strconv.Itoa(index))
	}
}

func TestEncode_ByHex_FromBytes(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    []byte // 输入值
		expected []byte // 期望值
	}{
		{[]byte(""), []byte("")},
		{[]byte("hello world"), []byte("68656c6c6f20776f726c64")},
	}

	for index, test := range tests {
		e := Encode.FromBytes(test.input).ByHex()
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToBytes(), "Current test id is "+strconv.Itoa(index))
	}
}

func TestEncode_ByHex_Multiple(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{"hello world", "36383635366336633666323037373666373236633634"},
	}

	for index, test := range tests {
		e := Encode.FromString(test.input).ByHex().ByHex()
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToString(), "Current test id is "+strconv.Itoa(index))
	}
}

func TestDecode_ByHex_FromString(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{"68656c6c6f20776f726c64", "hello world"},
	}

	for index, test := range tests {
		e := Decode.FromString(test.input).ByHex()
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToString(), "Current test id is "+strconv.Itoa(index))
	}
}

func TestDecode_ByHex_FromBytes(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    []byte // 输入值
		expected []byte // 期望值
	}{
		{[]byte(""), []byte("")},
		{[]byte("68656c6c6f20776f726c64"), []byte("hello world")},
	}

	for index, test := range tests {
		e := Decode.FromBytes(test.input).ByHex()
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToBytes(), "Current test id is "+strconv.Itoa(index))
	}
}

func TestDecode_ByHex_Multiple(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{"36383635366336633666323037373666373236633634", "hello world"},
	}

	for index, test := range tests {
		e := Decode.FromString(test.input).ByHex().ByHex()
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToString(), "Current test id is "+strconv.Itoa(index))
	}
}
