package dongle

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

var base58Test = []struct {
	input  string // 输入值
	output string // 期望值
}{
	{"", ""},
	{"hello world", "StV1DL6CwTryKyV"},
}

func TestEncode_ByBase58_ToString(t *testing.T) {
	for index, test := range base58Test {
		e := Encode.FromString(test.input).ByBase58()
		assert.Nil(t, e.Error)
		assert.Equal(t, test.output, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestDecode_ByBase58_ToString(t *testing.T) {
	for index, test := range base58Test {
		d := Decode.FromString(test.output).ByBase58()
		assert.Nil(t, d.Error)
		assert.Equal(t, test.input, d.ToString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestEncode_ByBase58_ToBytes(t *testing.T) {
	for index, test := range base58Test {
		e := Encode.FromBytes([]byte(test.input)).ByBase58()
		assert.Nil(t, e.Error)
		assert.Equal(t, []byte(test.output), e.ToBytes(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestDecode_ByBase58_ToBytes(t *testing.T) {
	for index, test := range base58Test {
		d := Decode.FromBytes([]byte(test.output)).ByBase58()
		assert.Nil(t, d.Error)
		assert.Equal(t, []byte(test.input), d.ToBytes(), "Current test index is "+strconv.Itoa(index))
	}
}
