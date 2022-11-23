package dongle

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

var base91Test = []struct {
	input  string // 输入值
	output string // 期望值
}{
	{"", ""},
	{"1", "xA"},
	{"1234567890", "QztEml0o[2;(A"},
	{"hello world", "TPwJh>Io2Tv!lE"},
	{"\x14\xfb\x9c\x03", "Q<c[A"},
	{"\xf2\x8e\x88\x31\x1a\xf0\x68\xce\x7a\x3f", "EquimSayaka~A"},
	{"\x35\x5e\x56\xe0\xc6\x29\x38\xf4\x81\x00\xab\x81\x7e\xd7\x08\x95\x62\x20\xa7\xda\x64\xa2\xce\xb3\xc5", "~_1H=x_t{|$AjJX(nMFdjL~:?1b3HgM"},
}

func TestEncode_ByBase91_FromStringToString(t *testing.T) {
	for index, test := range base91Test {
		e := Encode.FromString(test.input).ByBase91()
		assert.Nil(t, e.Error)
		assert.Equal(t, test.output, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestDecode_ByBase91_FromStringToString(t *testing.T) {
	for index, test := range base91Test {
		d := Decode.FromString(test.output).ByBase91()
		assert.Nil(t, d.Error)
		assert.Equal(t, test.input, d.ToString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestEncode_ByBase91_FromBytesToBytes(t *testing.T) {
	for index, test := range base91Test {
		e := Encode.FromBytes([]byte(test.input)).ByBase91()
		assert.Nil(t, e.Error)
		assert.Equal(t, []byte(test.output), e.ToBytes(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestDecode_ByBase91_FromBytesToBytes(t *testing.T) {
	for index, test := range base91Test {
		d := Decode.FromBytes([]byte(test.output)).ByBase91()
		assert.Nil(t, d.Error)
		assert.Equal(t, []byte(test.input), d.ToBytes(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestDecode_ByBase91_Error(t *testing.T) {
	d1 := Decode.FromString("-").ByBase91()
	assert.Equal(t, invalidCiphertextError("base91"), d1.Error)

	d2 := Decode.FromString("'").ByBase91()
	assert.Equal(t, invalidCiphertextError("base91"), d2.Error)

	d3 := Decode.FromBytes([]byte("\\")).ByBase91()
	assert.Equal(t, invalidCiphertextError("base91"), d3.Error)

	d4 := Decode.FromBytes([]byte("~_1H=x_t{ |$AjJX(nMFdjL~:?1b3HgM")).ByBase91()
	assert.Equal(t, invalidCiphertextError("base91"), d4.Error)
}
