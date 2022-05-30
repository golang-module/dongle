package dongle

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	HexInput    = "hello world"
	HexExpected = "68656c6c6f20776f726c64"
)

func TestEncode_ByHex_FromStringToString(t *testing.T) {
	tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{HexInput, HexExpected},
	}

	for index, test := range tests {
		e := Encode.FromString(test.input).ByHex()
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToString(), "Current test id is "+strconv.Itoa(index))
	}
}

func TestDecode_ByHex_FromStringToString(t *testing.T) {
	tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{HexExpected, HexInput},
	}

	for index, test := range tests {
		d := Decode.FromString(test.input).ByHex()
		assert.Nil(t, d.Error)
		assert.Equal(t, test.expected, d.ToString(), "Current test id is "+strconv.Itoa(index))
	}
}

func TestEncode_ByHex_FromBytesToBytes(t *testing.T) {
	tests := []struct {
		input    []byte // 输入值
		expected []byte // 期望值
	}{
		{[]byte(""), []byte("")},
		{[]byte(HexInput), []byte(HexExpected)},
	}

	for index, test := range tests {
		e := Encode.FromBytes(test.input).ByHex()
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToBytes(), "Current test id is "+strconv.Itoa(index))
	}
}

func TestDecode_ByHex_FromBytesToBytes(t *testing.T) {
	tests := []struct {
		input    []byte // 输入值
		expected []byte // 期望值
	}{
		{[]byte(""), []byte("")},
		{[]byte(HexExpected), []byte(HexInput)},
	}

	for index, test := range tests {
		d := Decode.FromBytes(test.input).ByHex()
		assert.Nil(t, d.Error)
		assert.Equal(t, test.expected, d.ToBytes(), "Current test id is "+strconv.Itoa(index))
	}
}
