package dongle

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

var hmacSha1Test = []struct {
	input    string
	key      string
	toHex    string
	toBase64 string
}{
	{"", "", "", ""},
	{"hello world", "dongle", "91c103ef93ba7420902b0d1bf0903251c94b4a62", "kcED75O6dCCQKw0b8JAyUclLSmI="},
}

func TestEncrypt_ByHmacSha1_ToString(t *testing.T) {
	for index, test := range hmacSha1Test {
		e := Encrypt.FromString(test.input).ByHmacSha1(test.key)
		assert.Nil(t, e.Error)

		assert.Equal(t, test.toHex, e.ToHexString(), "Hex test index is "+strconv.Itoa(index))
		assert.Equal(t, test.toBase64, e.ToBase64String(), "Base64 test index is "+strconv.Itoa(index))
	}
}

func TestEncrypt_ByHmacSha1_ToBytes(t *testing.T) {
	for index, test := range hmacSha1Test {
		e := Encrypt.FromBytes([]byte(test.input)).ByHmacSha1([]byte(test.key))
		assert.Nil(t, e.Error)

		assert.Equal(t, []byte(test.toHex), e.ToHexBytes(), "Hex test test index is "+strconv.Itoa(index))
		assert.Equal(t, []byte(test.toBase64), e.ToBase64Bytes(), "Base64 test index is "+strconv.Itoa(index))
	}
}
