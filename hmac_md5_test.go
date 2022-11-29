package dongle

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

var hmacMd5Test = []struct {
	input    string
	key      string
	toHex    string
	toBase64 string
}{
	{"", "", "", ""},
	{"hello world", "dongle", "4790626a275f776956386e5a3ea7b726", "R5Biaidfd2lWOG5aPqe3Jg=="},
}

func TestEncrypt_ByHmacMd5_ToString(t *testing.T) {
	for index, test := range hmacMd5Test {
		e := Encrypt.FromString(test.input).ByHmacMd5(test.key)
		assert.Nil(t, e.Error)

		assert.Equal(t, test.toHex, e.ToHexString(), "Hex test index is "+strconv.Itoa(index))
		assert.Equal(t, test.toBase64, e.ToBase64String(), "Base64 test index is "+strconv.Itoa(index))
	}
}

func TestEncrypt_ByHmacMd5_ToBytes(t *testing.T) {
	for index, test := range hmacMd5Test {
		e := Encrypt.FromBytes([]byte(test.input)).ByHmacMd5([]byte(test.key))
		assert.Nil(t, e.Error)

		assert.Equal(t, []byte(test.toHex), e.ToHexBytes(), "Hex test test index is "+strconv.Itoa(index))
		assert.Equal(t, []byte(test.toBase64), e.ToBase64Bytes(), "Base64 test index is "+strconv.Itoa(index))
	}
}
