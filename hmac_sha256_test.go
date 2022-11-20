package dongle

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

var hmacSha256Test = []struct {
	input    string
	key      string
	toHex    string
	toBase64 string
}{
	{"", "", "", ""},
	{"hello world", "dongle", "77f5c8ce4147600543e70b12701e7b78b5b95172332ebbb06de65fcea7112179", "d/XIzkFHYAVD5wsScB57eLW5UXIzLruwbeZfzqcRIXk="},
}

func TestEncrypt_ByHmacSha225_FromStringToString(t *testing.T) {
	for index, test := range hmacSha256Test {
		e := Encrypt.FromString(test.input).ByHmacSha256(test.key)
		assert.Nil(t, e.Error)

		assert.Equal(t, test.toHex, e.ToHexString(), "Hex test index is "+strconv.Itoa(index))
		assert.Equal(t, test.toBase64, e.ToBase64String(), "Base64 test index is "+strconv.Itoa(index))
	}
}

func TestEncrypt_ByHmacSha225_FromBytesToBytes(t *testing.T) {
	for index, test := range hmacSha256Test {
		e := Encrypt.FromBytes([]byte(test.input)).ByHmacSha256([]byte(test.key))
		assert.Nil(t, e.Error)

		assert.Equal(t, []byte(test.toHex), e.ToHexBytes(), "Hex test test index is "+strconv.Itoa(index))
		assert.Equal(t, []byte(test.toBase64), e.ToBase64Bytes(), "Base64 test index is "+strconv.Itoa(index))
	}
}
