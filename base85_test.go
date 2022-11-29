package dongle

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

var base85Test = []struct {
	input  string // 输入值
	output string // 期望值
}{
	{"", ""},
	{"hello world", "BOu!rD]j7BEbo7"},
}

func TestEncode_ByBase85_ToString(t *testing.T) {
	for index, test := range base85Test {
		e := Encode.FromString(test.input).ByBase85()
		assert.Nil(t, e.Error)
		assert.Equal(t, test.output, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestDecode_ByBase85_ToString(t *testing.T) {
	for index, test := range base85Test {
		d := Decode.FromString(test.output).ByBase85()
		assert.Nil(t, d.Error)
		assert.Equal(t, test.input, d.ToString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestEncode_ByBase85_ToBytes(t *testing.T) {
	for index, test := range base85Test {
		e := Encode.FromBytes([]byte(test.input)).ByBase85()
		assert.Nil(t, e.Error)
		assert.Equal(t, []byte(test.output), e.ToBytes(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestDecode_ByBase85_ToBytes(t *testing.T) {
	for index, test := range base85Test {
		d := Decode.FromBytes([]byte(test.output)).ByBase85()
		assert.Nil(t, d.Error)
		assert.Equal(t, []byte(test.input), d.ToBytes(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestBase85_Ciphertext_Error(t *testing.T) {
	d1 := Decode.FromString("xxxxxx").ByBase85()
	assert.Equal(t, invalidCiphertextError("base85"), d1.Error)

	d2 := Decode.FromBytes([]byte("xxxxxx")).ByBase85()
	assert.Equal(t, invalidCiphertextError("base85"), d2.Error)
}
