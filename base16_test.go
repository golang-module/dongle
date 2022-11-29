package dongle

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

var base16Test = []struct {
	input  string // 输入值
	output string // 期望值
}{
	{"", ""},
	{"hello world", "68656c6c6f20776f726c64"},
}

func TestEncode_ByBase16_ToString(t *testing.T) {
	for index, test := range base16Test {
		e := Encode.FromString(test.input).ByBase16()
		assert.Nil(t, e.Error)
		assert.Equal(t, test.output, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestDecode_ByBase16_ToString(t *testing.T) {
	for index, test := range base16Test {
		d := Decode.FromString(test.output).ByBase16()
		assert.Nil(t, d.Error)
		assert.Equal(t, test.input, d.ToString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestEncode_ByBase16_ToBytes(t *testing.T) {
	for index, test := range base16Test {
		e := Encode.FromBytes([]byte(test.input)).ByBase16()
		assert.Nil(t, e.Error)
		assert.Equal(t, []byte(test.output), e.ToBytes(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestDecode_ByBase16_ToBytes(t *testing.T) {
	for index, test := range base16Test {
		d := Decode.FromBytes([]byte(test.output)).ByBase16()
		assert.Nil(t, d.Error)
		assert.Equal(t, []byte(test.input), d.ToBytes(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestBase16_Ciphertext_Error(t *testing.T) {
	d1 := Decode.FromString("xxxxxx").ByBase16()
	assert.Equal(t, invalidCiphertextError("base16"), d1.Error)

	d2 := Decode.FromBytes([]byte("xxxxxx")).ByBase16()
	assert.Equal(t, invalidCiphertextError("base16"), d2.Error)
}
