package dongle

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	base32Input    = "hello world"
	base32Expected = "NBSWY3DPEB3W64TMMQ======"
)

func TestEncode_ByBase32_FromStringToString(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{base32Input, base32Expected},
	}

	for index, test := range tests {
		e := Encode.FromString(test.input).ByBase32()
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToString(), "Current test id is "+strconv.Itoa(index))
	}
}

func TestDecode_ByBase32_FromStringToString(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{base32Expected, base32Input},
	}

	for index, test := range tests {
		d := Decode.FromString(test.input).ByBase32()
		assert.Nil(d.Error)
		assert.Equal(test.expected, d.ToString(), "Current test id is "+strconv.Itoa(index))
	}
}

func TestEncode_ByBase32_FromBytesToBytes(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    []byte // 输入值
		expected []byte // 期望值
	}{
		{[]byte(""), []byte("")},
		{[]byte(base32Input), []byte(base32Expected)},
	}

	for index, test := range tests {
		e := Encode.FromBytes(test.input).ByBase32()
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToBytes(), "Current test id is "+strconv.Itoa(index))
	}
}

func TestDecode_ByBase32_FromBytesToBytes(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    []byte // 输入值
		expected []byte // 期望值
	}{
		{[]byte(""), []byte("")},
		{[]byte(base32Expected), []byte(base32Input)},
	}

	for index, test := range tests {
		d := Decode.FromBytes(test.input).ByBase32()
		assert.Nil(d.Error)
		assert.Equal(test.expected, d.ToBytes(), "Current test id is "+strconv.Itoa(index))
	}
}
