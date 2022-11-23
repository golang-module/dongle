package dongle

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

var base62Test = []struct {
	input  string // 输入值
	output string // 期望值
}{
	{"", ""},
	{"1", "n"},
	{"f", "1e"},
	{"1\n2", "DU4g"},
	{"1234567890", "1A0afZkibIAR2O"},
	{"hello world", "AAwf93rvy4aWQVw"},
	{"\x14\xfb\x9c\x03", "Np64R"},
	{"\xf2\x8e\x88\x31\x1a\xf0\x68\xce\x7a\x3f", "5j2FgcEB7vzAy7"},
}

func TestEncode_ByBase62_FromStringToString(t *testing.T) {
	for index, test := range base62Test {
		e := Encode.FromString(test.input).ByBase62()
		assert.Nil(t, e.Error)
		assert.Equal(t, test.output, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestDecode_ByBase62_FromStringToString(t *testing.T) {
	for index, test := range base62Test {
		d := Decode.FromString(test.output).ByBase62()
		assert.Nil(t, d.Error)
		assert.Equal(t, test.input, d.ToString(), "Current test index is "+strconv.Itoa(index))
	}

	d := Decode.FromString("AAwf93rvy4aWQVw\n").ByBase62()
	assert.Nil(t, d.Error)
	assert.Equal(t, "hello world", d.ToString())
}

func TestEncode_ByBase62_FromBytesToBytes(t *testing.T) {
	for index, test := range base62Test {
		e := Encode.FromBytes([]byte(test.input)).ByBase62()
		assert.Nil(t, e.Error)
		assert.Equal(t, []byte(test.output), e.ToBytes(), "Current test index is "+strconv.Itoa(index))
	}

	d := Decode.FromBytes([]byte("AAwf93rvy4aWQVw\n")).ByBase62()
	assert.Nil(t, d.Error)
	assert.Equal(t, "hello world", d.ToString())
}

func TestDecode_ByBase62_FromBytesToBytes(t *testing.T) {
	for index, test := range base62Test {
		d := Decode.FromBytes([]byte(test.output)).ByBase62()
		assert.Nil(t, d.Error)
		assert.Equal(t, []byte(test.input), d.ToBytes(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestDecode_ByBase62_Error(t *testing.T) {
	d1 := Decode.FromString("-").ByBase62()
	assert.Equal(t, invalidCiphertextError("base62"), d1.Error)

	d2 := Decode.FromString("'").ByBase62()
	assert.Equal(t, invalidCiphertextError("base62"), d2.Error)

	d3 := Decode.FromBytes([]byte("\\")).ByBase62()
	assert.Equal(t, invalidCiphertextError("base62"), d3.Error)

	d4 := Decode.FromBytes([]byte("~_1H=x_t{ |$AjJX(nMFdjL~:?1b3HgM")).ByBase62()
	assert.Equal(t, invalidCiphertextError("base62"), d4.Error)
}
