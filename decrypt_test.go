package dongle

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecrypt_String(t *testing.T) {
	tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{"hello world", "hello world"},
	}

	for index, test := range tests {
		d := newDecrypt().FromString(test.input)
		d.dst = []byte(test.input)
		assert.Nil(t, d.Error)
		assert.Equal(t, test.expected, fmt.Sprintf("%s", d), "Current test index is "+strconv.Itoa(index))
	}
}

func TestDecrypt_ToString(t *testing.T) {
	tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{"hello world", "hello world"},
	}

	for index, test := range tests {
		d := newDecrypt().FromBytes([]byte(test.input))
		d.dst = []byte(test.input)
		assert.Nil(t, d.Error)
		assert.Equal(t, test.expected, d.ToString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestDecrypt_ToBytes(t *testing.T) {
	tests := []struct {
		input    string // 输入值
		expected []byte // 期望值
	}{
		{"", []byte("")},
		{"hello world", []byte("hello world")},
	}

	for index, test := range tests {
		d := newDecrypt()
		d.dst = []byte(test.input)
		assert.Nil(t, d.Error)
		assert.Equal(t, test.expected, d.ToBytes(), "Current test index is "+strconv.Itoa(index))
	}
}
