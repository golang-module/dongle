package dongle

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

var base100Test = []struct {
	input  string // 输入值
	output string // 期望值
}{
	{"", ""},
	{"1234567890", "🐨🐩🐪🐫🐬🐭🐮🐯🐰🐧"},
	{"hello world", "👟👜👣👣👦🐗👮👦👩👣👛"},
}

func TestEncode_ByBase100_FromStringToString(t *testing.T) {
	for index, test := range base100Test {
		e := Encode.FromString(test.input).ByBase100()
		assert.Nil(t, e.Error)
		assert.Equal(t, test.output, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestDecode_ByBase100_FromStringToString(t *testing.T) {
	for index, test := range base100Test {
		d := Decode.FromString(test.output).ByBase100()
		assert.Nil(t, d.Error)
		assert.Equal(t, test.input, d.ToString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestEncode_ByBase100_FromBytesToBytes(t *testing.T) {
	for index, test := range base100Test {
		e := Encode.FromBytes([]byte(test.input)).ByBase100()
		assert.Nil(t, e.Error)
		assert.Equal(t, []byte(test.output), e.ToBytes(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestDecode_ByBase100_FromBytesToBytes(t *testing.T) {
	for index, test := range base100Test {
		d := Decode.FromBytes([]byte(test.output)).ByBase100()
		assert.Nil(t, d.Error)
		assert.Equal(t, []byte(test.input), d.ToBytes(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestDecode_ByBase100_Error(t *testing.T) {
	d1 := Decode.FromString("-").ByBase100()
	assert.Equal(t, invalidCiphertextError("base100"), d1.Error)

	d2 := Decode.FromString("'").ByBase100()
	assert.Equal(t, invalidCiphertextError("base100"), d2.Error)

	d3 := Decode.FromBytes([]byte("\\")).ByBase100()
	assert.Equal(t, invalidCiphertextError("base100"), d3.Error)

	d4 := Decode.FromBytes([]byte("~_1H=x_t{ |$AjJX(nMFdjL~:?1b3HgM")).ByBase100()
	assert.Equal(t, invalidCiphertextError("base100"), d4.Error)
}
