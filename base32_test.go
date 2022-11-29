package dongle

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

var base32Test = []struct {
	input  string // 输入值
	output string // 期望值
}{
	{"", ""},
	{"hello world", "NBSWY3DPEB3W64TMMQ======"},
}

func TestEncode_ByBase32_ToString(t *testing.T) {
	for index, test := range base32Test {
		e := Encode.FromString(test.input).ByBase32()
		assert.Nil(t, e.Error)
		assert.Equal(t, test.output, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestDecode_ByBase32_ToString(t *testing.T) {
	for index, test := range base32Test {
		d := Decode.FromString(test.output).ByBase32()
		assert.Nil(t, d.Error)
		assert.Equal(t, test.input, d.ToString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestEncode_ByBase32_ToBytes(t *testing.T) {
	for index, test := range base32Test {
		e := Encode.FromBytes([]byte(test.input)).ByBase32()
		assert.Nil(t, e.Error)
		assert.Equal(t, []byte(test.output), e.ToBytes(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestDecode_ByBase32_ToBytes(t *testing.T) {
	for index, test := range base32Test {
		d := Decode.FromBytes([]byte(test.output)).ByBase32()
		assert.Nil(t, d.Error)
		assert.Equal(t, []byte(test.input), d.ToBytes(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestBase32_Ciphertext_Error(t *testing.T) {
	d1 := Decode.FromString("xxxxxx").ByBase32()
	assert.Equal(t, invalidCiphertextError("base32"), d1.Error)

	d2 := Decode.FromBytes([]byte("xxxxxx")).ByBase32()
	assert.Equal(t, invalidCiphertextError("base32"), d2.Error)
}
