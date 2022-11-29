package dongle

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

var hmacMd4Test = []struct {
	input    string
	key      string
	toHex    string
	toBase64 string
}{
	{"", "", "", ""},
	{"hello world", "dongle", "7a9df5247cbf76a8bc17c9c4f5a75b6b", "ep31JHy/dqi8F8nE9adbaw=="},
}

func TestEncrypt_ByHmacMd4_ToString(t *testing.T) {
	for index, test := range hmacMd4Test {
		e := Encrypt.FromString(test.input).ByHmacMd4(test.key)
		assert.Nil(t, e.Error)

		assert.Equal(t, test.toHex, e.ToHexString(), "Hex test index is "+strconv.Itoa(index))
		assert.Equal(t, test.toBase64, e.ToBase64String(), "Base64 test index is "+strconv.Itoa(index))
	}
}

func TestEncrypt_ByHmacMd4_ToBytes(t *testing.T) {
	for index, test := range hmacMd4Test {
		e := Encrypt.FromBytes([]byte(test.input)).ByHmacMd4([]byte(test.key))
		assert.Nil(t, e.Error)

		assert.Equal(t, []byte(test.toHex), e.ToHexBytes(), "Hex test test index is "+strconv.Itoa(index))
		assert.Equal(t, []byte(test.toBase64), e.ToBase64Bytes(), "Base64 test index is "+strconv.Itoa(index))
	}
}
