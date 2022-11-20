package dongle

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

var base64UrlTest = []struct {
	input  string // 输入值
	output string // 期望值
}{
	{"", ""},
	{"www.gouguoyin.cn", "d3d3LmdvdWd1b3lpbi5jbg=="},
}

func TestEncode_ByBase64URL_FromStringToString(t *testing.T) {
	for index, test := range base64UrlTest {
		e := Encode.FromString(test.input).ByBase64URL()
		assert.Nil(t, e.Error)
		assert.Equal(t, test.output, e.ToString(), "Current test id is "+strconv.Itoa(index))
	}
}

func TestDecode_ByBase64URL_FromStringToString(t *testing.T) {

	for index, test := range base64UrlTest {
		d := Decode.FromString(test.output).ByBase64URL()
		assert.Nil(t, d.Error)
		assert.Equal(t, test.input, d.ToString(), "Current test id is "+strconv.Itoa(index))
	}
}

func TestEncode_ByBase64URL_FromBytesToBytes(t *testing.T) {

	for index, test := range base64UrlTest {
		e := Encode.FromBytes([]byte(test.input)).ByBase64URL()
		assert.Nil(t, e.Error)
		assert.Equal(t, []byte(test.output), e.ToBytes(), "Current test id is "+strconv.Itoa(index))
	}
}

func TestDecode_ByBase64URL_FromBytesToBytes(t *testing.T) {
	for index, test := range base64UrlTest {
		d := Decode.FromBytes([]byte(test.output)).ByBase64URL()
		assert.Nil(t, d.Error)
		assert.Equal(t, []byte(test.input), d.ToBytes(), "Current test id is "+strconv.Itoa(index))
	}
}

func TestDecode_ByBase64URL_Error(t *testing.T) {
	d1 := Decode.FromString("xxxxxx").ByBase64URL()
	assert.Equal(t, invalidCiphertextError("base64URL"), d1.Error)

	d2 := Decode.FromBytes([]byte("xxxxxx")).ByBase64URL()
	assert.Equal(t, invalidCiphertextError("base64URL"), d2.Error)
}
