package dongle

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	base58Input    = "hello world"
	base58Expected = "StV1DL6CwTryKyV"
)

func TestEncode_ByBase58_FromStringToString(t *testing.T) {
	tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{base58Input, base58Expected},
	}

	for index, test := range tests {
		e := Encode.FromString(test.input).ByBase58()
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToString(), "Current test id is "+strconv.Itoa(index))
	}
}

func TestDecode_ByBase58_FromStringToString(t *testing.T) {
	tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{base58Expected, base58Input},
	}

	for index, test := range tests {
		d := Decode.FromString(test.input).ByBase58()
		assert.Nil(t, d.Error)
		assert.Equal(t, test.expected, d.ToString(), "Current test id is "+strconv.Itoa(index))
	}
}

func TestEncode_ByBase58_FromBytesToBytes(t *testing.T) {
	tests := []struct {
		input    []byte // 输入值
		expected []byte // 期望值
	}{
		{[]byte(""), []byte("")},
		{[]byte(base58Input), []byte(base58Expected)},
	}

	for index, test := range tests {
		e := Encode.FromBytes(test.input).ByBase58()
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToBytes(), "Current test id is "+strconv.Itoa(index))
	}
}

func TestDecode_ByBase58_FromBytesToBytes(t *testing.T) {
	tests := []struct {
		input    []byte // 输入值
		expected []byte // 期望值
	}{
		{[]byte(""), []byte("")},
		{[]byte(base58Expected), []byte(base58Input)},
	}

	for index, test := range tests {
		d := Decode.FromBytes(test.input).ByBase58()
		assert.Nil(t, d.Error)
		assert.Equal(t, test.expected, d.ToBytes(), "Current test id is "+strconv.Itoa(index))
	}
}
