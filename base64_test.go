package dongle

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

var base64Test = []struct {
	input  string // 输入值
	output string // 期望值
}{
	{"", ""},
	{"hello world", "aGVsbG8gd29ybGQ="},
}

func TestEncode_ByBase64_FromStringToString(t *testing.T) {
	for index, test := range base64Test {
		e := Encode.FromString(test.input).ByBase64()
		assert.Nil(t, e.Error)
		assert.Equal(t, test.output, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestDecode_ByBase64_FromStringToString(t *testing.T) {
	for index, test := range base64Test {
		d := Decode.FromString(test.output).ByBase64()
		assert.Nil(t, d.Error)
		assert.Equal(t, test.input, d.ToString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestEncode_ByBase64_FromBytesToBytes(t *testing.T) {
	for index, test := range base64Test {
		e := Encode.FromBytes([]byte(test.input)).ByBase64()
		assert.Nil(t, e.Error)
		assert.Equal(t, []byte(test.output), e.ToBytes(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestDecode_ByBase64_FromBytesToBytes(t *testing.T) {
	for index, test := range base64Test {
		d := Decode.FromBytes([]byte(test.output)).ByBase64()
		assert.Nil(t, d.Error)
		assert.Equal(t, []byte(test.input), d.ToBytes(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestDecode_ByBase64_Error(t *testing.T) {
	d1 := Decode.FromString("xxxxxx").ByBase64()
	assert.Equal(t, invalidCiphertextError("base64"), d1.Error)

	d2 := Decode.FromBytes([]byte("xxxxxx")).ByBase64()
	assert.Equal(t, invalidCiphertextError("base64"), d2.Error)
}
