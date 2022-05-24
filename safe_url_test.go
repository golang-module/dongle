package dongle

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	safeUrlInput    = "www.gouguoyin.cn?sex=男&age=18"
	safeUrlExpected = "www.gouguoyin.cn%3Fsex%3D%E7%94%B7%26age%3D18"
)

func TestEncode_BySafeURL_FromStringToString(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{safeUrlInput, safeUrlExpected},
	}

	for index, test := range tests {
		e := Encode.FromString(test.input).BySafeURL()
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToString(), "Current test id is "+strconv.Itoa(index))
	}
}

func TestDecode_BySafeURL_FromStringToString(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{safeUrlExpected, safeUrlInput},
	}

	for index, test := range tests {
		d := Decode.FromString(test.input).BySafeURL()
		assert.Nil(d.Error)
		assert.Equal(test.expected, d.ToString(), "Current test id is "+strconv.Itoa(index))
	}
}

func TestEncode_ByBySafeURL_FromBytesToBytes(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    []byte // 输入值
		expected []byte // 期望值
	}{
		{[]byte(""), []byte("")},
		{[]byte(safeUrlInput), []byte(safeUrlExpected)},
	}

	for index, test := range tests {
		e := Encode.FromBytes(test.input).BySafeURL()
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToBytes(), "Current test id is "+strconv.Itoa(index))
	}
}

func TestDecode_ByBySafeURL_FromBytesToBytes(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    []byte // 输入值
		expected []byte // 期望值
	}{
		{[]byte(""), []byte("")},
		{[]byte(safeUrlExpected), []byte(safeUrlInput)},
	}

	for index, test := range tests {
		d := Decode.FromBytes(test.input).BySafeURL()
		assert.Nil(d.Error)
		assert.Equal(test.expected, d.ToBytes(), "Current test id is "+strconv.Itoa(index))
	}
}
