package dongle

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

var hexTest = []struct {
	input  string // 输入值
	output string // 期望值
}{
	{"", ""},
	{"hello world", "68656c6c6f20776f726c64"},
}

func TestHex_Encode_ToString(t *testing.T) {
	for index, test := range hexTest {
		e := Encode.FromString(test.input).ByHex()
		assert.Nil(t, e.Error)
		assert.Equal(t, test.output, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestHex_Decode_ToString(t *testing.T) {
	for index, test := range hexTest {
		d := Decode.FromString(test.output).ByHex()
		assert.Nil(t, d.Error)
		assert.Equal(t, test.input, d.ToString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestHex_Encode_ToBytes(t *testing.T) {
	for index, test := range hexTest {
		e := Encode.FromBytes([]byte(test.input)).ByHex()
		assert.Nil(t, e.Error)
		assert.Equal(t, []byte(test.output), e.ToBytes(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestHex_Decode_ToBytes(t *testing.T) {
	for index, test := range hexTest {
		d := Decode.FromBytes([]byte(test.output)).ByHex()
		assert.Nil(t, d.Error)
		assert.Equal(t, []byte(test.input), d.ToBytes(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestHex_Ciphertext_Error(t *testing.T) {
	d1 := Decode.FromString("xxxxxx").ByHex()
	assert.Equal(t, invalidCiphertextError("hex"), d1.Error)

	d2 := Decode.FromBytes([]byte("xxxxxx")).ByHex()
	assert.Equal(t, invalidCiphertextError("hex"), d2.Error)
}
