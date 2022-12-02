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
	{"", "dongle", "", ""},
	{"hello world", "dongle", "eba154b4cb5a9038dbbf9d", "66FUtMtakDjbv50="},
}

func TestRc4_Encrypt_ToString(t *testing.T) {
	for index, test := range rc4Test {
		e := Encrypt.FromString(test.input).ByRc4(test.key)
		assert.Nil(t, e.Error)

		assert.Equal(t, test.toHex, e.ToHexString(), "Hex test index is "+strconv.Itoa(index))
		assert.Equal(t, test.toBase64, e.ToBase64String(), "Base64 test index is "+strconv.Itoa(index))
	}
}

func TestRc4_Encrypt_ToBytes(t *testing.T) {
	for index, test := range rc4Test {
		e := Encrypt.FromBytes([]byte(test.input)).ByRc4([]byte(test.key))
		assert.Nil(t, e.Error)

		assert.Equal(t, []byte(test.toHex), e.ToHexBytes(), "Hex test test index is "+strconv.Itoa(index))
		assert.Equal(t, []byte(test.toBase64), e.ToBase64Bytes(), "Base64 test index is "+strconv.Itoa(index))
	}
}

func TestRc4_Decrypt_ToString(t *testing.T) {
	for index, test := range rc4Test {
		d1 := Decrypt.FromHexString(test.toHex).ByRc4(test.key)
		assert.Nil(t, d1.Error)
		assert.Equal(t, test.input, d1.ToString(), "Hex test index is "+strconv.Itoa(index))

		d := Decrypt.FromBase64String(test.toBase64).ByRc4(test.key)
		assert.Nil(t, d.Error)
		assert.Equal(t, test.input, d.ToString(), "Base64 test index is "+strconv.Itoa(index))
	}
}

func TestRc4_Decrypt_ToBytes(t *testing.T) {
	for index, test := range rc4Test {
		d1 := Decrypt.FromHexBytes([]byte(test.toHex)).ByRc4([]byte(test.key))
		assert.Nil(t, d1.Error)
		assert.Equal(t, []byte(test.input), d1.ToBytes(), "Hex test index is "+strconv.Itoa(index))

		d2 := Decrypt.FromBase64Bytes([]byte(test.toBase64)).ByRc4([]byte(test.key))
		assert.Nil(t, d2.Error)
		assert.Equal(t, []byte(test.input), d2.ToBytes(), "Base64 test index is "+strconv.Itoa(index))
	}
}

func TestRc4_Key_Error(t *testing.T) {
	e := Encrypt.FromString("hello go").ByRc4("")
	assert.Equal(t, invalidRc4KeyError(), e.Error)

	d := Decrypt.FromString("hello go").ByRc4("")
	assert.Equal(t, invalidRc4KeyError(), d.Error)
}
