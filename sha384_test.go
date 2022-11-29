package dongle

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

var sha384Test = []struct {
	input    string
	toHex    string
	toBase64 string
}{
	{"", "", ""},
	{"hello world", "fdbd8e75a67f29f701a4e040385e2e23986303ea10239211af907fcbb83578b3e417cb71ce646efd0819dd8c088de1bd", "/b2OdaZ/KfcBpOBAOF4uI5hjA+oQI5IRr5B/y7g1eLPkF8txzmRu/QgZ3YwIjeG9"},
}

func TestEncrypt_BySha384_ToString(t *testing.T) {
	for index, test := range sha384Test {
		e := Encrypt.FromString(test.input).BySha384()
		assert.Nil(t, e.Error)

		assert.Equal(t, test.toHex, e.ToHexString(), "Current test index is "+strconv.Itoa(index))
		assert.Equal(t, test.toBase64, e.ToBase64String(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestEncrypt_BySha384_ToBytes(t *testing.T) {
	for index, test := range sha384Test {
		e := Encrypt.FromBytes([]byte(test.input)).BySha384()
		assert.Nil(t, e.Error)

		assert.Equal(t, []byte(test.toHex), e.ToHexBytes(), "Current test index is "+strconv.Itoa(index))
		assert.Equal(t, []byte(test.toBase64), e.ToBase64Bytes(), "Current test index is "+strconv.Itoa(index))
	}
}
