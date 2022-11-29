package dongle

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

var hmacSha224Test = []struct {
	input    string
	key      string
	toHex    string
	toBase64 string
}{
	{"", "", "", ""},
	{"hello world", "dongle", "e15b9e5a7eccb1f17dc81dc07c909a891936dc3429dc0d940accbcec", "4VueWn7MsfF9yB3AfJCaiRk23DQp3A2UCsy87A=="},
}

func TestEncrypt_ByHmacSha224_ToString(t *testing.T) {
	for index, test := range hmacSha224Test {
		e := Encrypt.FromString(test.input).ByHmacSha224(test.key)
		assert.Nil(t, e.Error)

		assert.Equal(t, test.toHex, e.ToHexString(), "Hex test index is "+strconv.Itoa(index))
		assert.Equal(t, test.toBase64, e.ToBase64String(), "Base64 test index is "+strconv.Itoa(index))
	}
}

func TestEncrypt_ByHmacSha224_ToBytes(t *testing.T) {
	for index, test := range hmacSha224Test {
		e := Encrypt.FromBytes([]byte(test.input)).ByHmacSha224([]byte(test.key))
		assert.Nil(t, e.Error)

		assert.Equal(t, []byte(test.toHex), e.ToHexBytes(), "Hex test test index is "+strconv.Itoa(index))
		assert.Equal(t, []byte(test.toBase64), e.ToBase64Bytes(), "Base64 test index is "+strconv.Itoa(index))
	}
}
