package dongle

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

var sm3Test = []struct {
	input    string
	toHex    string
	toBase64 string
}{
	{"", "", ""},
	{"hello world", "44f0061e69fa6fdfc290c494654a05dc0c053da7e5c52b84ef93a9d67d3fff88", "RPAGHmn6b9/CkMSUZUoF3AwFPaflxSuE75Op1n0//4g="},
}

func TestEncrypt_BySm3_ToString(t *testing.T) {
	for index, test := range sm3Test {
		e := Encrypt.FromString(test.input).BySm3()
		assert.Nil(t, e.Error)

		assert.Equal(t, test.toHex, e.ToHexString(), "Current test index is "+strconv.Itoa(index))
		assert.Equal(t, test.toBase64, e.ToBase64String(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestEncrypt_BySm3_ToBytes(t *testing.T) {
	for index, test := range sm3Test {
		e := Encrypt.FromBytes([]byte(test.input)).BySm3()
		assert.Nil(t, e.Error)

		assert.Equal(t, []byte(test.toHex), e.ToHexBytes(), "Current test index is "+strconv.Itoa(index))
		assert.Equal(t, []byte(test.toBase64), e.ToBase64Bytes(), "Current test index is "+strconv.Itoa(index))
	}
}
