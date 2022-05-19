package dongle

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncode_String(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{"hello world", "hello world"},
	}

	for index, test := range tests {
		e := newEncode()
		e.dst = []byte(test.input)
		assert.Nil(e.Error)
		assert.Equal(test.expected, fmt.Sprintf("%s", e), "Current test index is "+strconv.Itoa(index))
	}
}

func TestEncode_ToString(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{"hello world", "hello world"},
	}

	for index, test := range tests {
		e := newEncode()
		e.dst = []byte(test.input)
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestEncode_ToBytes(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected []byte // 期望值
	}{
		{"", []byte("")},
		{"hello world", []byte("hello world")},
	}

	for index, test := range tests {
		e := newEncode()
		e.dst = []byte(test.input)
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToBytes(), "Current test index is "+strconv.Itoa(index))
	}
}
