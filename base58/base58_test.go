package base58

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	input  string // 输入值
	output string // 期望值
}{
	{"", ""},
	{"hello world", "StV1DL6CwTryKyV"},
}

func TestEncode(t *testing.T) {
	for index, test := range tests {
		assert.Equal(t, test.output, string(Encode([]byte(test.input))), "Current test index is "+strconv.Itoa(index))
		assert.Equal(t, []byte(test.output), Encode([]byte(test.input)), "Current test index is "+strconv.Itoa(index))
	}
}

func TestDecode(t *testing.T) {
	for index, test := range tests {
		assert.Equal(t, test.output, string(Encode([]byte(test.input))), "Current test index is "+strconv.Itoa(index))
		assert.Equal(t, []byte(test.input), Decode([]byte(test.output)), "Current test index is "+strconv.Itoa(index))
	}
}
