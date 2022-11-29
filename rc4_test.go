package dongle

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

var rc4Test = []struct {
	input    string
	key      string
	toHex    string
	toBase64 string
}{
	{"", "", "", ""},
	{"hello world", "dongle", "eba154b4cb5a9038dbbf9d", "66FUtMtakDjbv50="},
}

func TestEncrypt_ByRc4_ToString(t *testing.T) {
	for index, test := range rc4Test {
		e := Encrypt.FromString(test.input).ByRc4(test.key)
		assert.Nil(t, e.Error)

		assert.Equal(t, test.toHex, e.ToHexString(), "Hex test index is "+strconv.Itoa(index))
		assert.Equal(t, test.toBase64, e.ToBase64String(), "Base64 test index is "+strconv.Itoa(index))
	}
}

func TestEncrypt_ByRc4_ToBytes(t *testing.T) {
	for index, test := range rc4Test {
		e := Encrypt.FromBytes([]byte(test.input)).ByRc4([]byte(test.key))
		assert.Nil(t, e.Error)

		assert.Equal(t, []byte(test.toHex), e.ToHexBytes(), "Hex test test index is "+strconv.Itoa(index))
		assert.Equal(t, []byte(test.toBase64), e.ToBase64Bytes(), "Base64 test index is "+strconv.Itoa(index))
	}
}

func TestRc4_Ciphertext_Error(t *testing.T) {
	e1 := Encrypt.FromString("hello world").ByRc4("")
	assert.Equal(t, invalidRc4KeyError(0), e1.Error, "Should catch an exception")

	e2 := Encrypt.FromBytes([]byte("hello world")).ByRc4([]byte(""))
	assert.Equal(t, invalidRc4KeyError(0), e2.Error, "Should catch an exception")
}
