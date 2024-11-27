package dongle

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var rc4Tests = []struct {
	input    string
	key      string
	toHex    string
	toBase64 string
}{
	{"", "dongle", "", ""},
	{"hello world", "dongle", "eba154b4cb5a9038dbbf9d", "66FUtMtakDjbv50="},
}

func TestRc4_Encrypt_String(t *testing.T) {
	for index, test := range rc4Tests {
		e := Encrypt.FromString(test.input).ByRc4([]byte(test.key))

		t.Run(fmt.Sprintf("test_%d", index), func(t *testing.T) {
			assert.Nil(t, e.Error)
			assert.Equal(t, test.toHex, e.ToHexString())
			assert.Equal(t, test.toBase64, e.ToBase64String())
		})
	}
}

func TestRc4_Decrypt_String(t *testing.T) {
	for index, test := range rc4Tests {
		d1 := Decrypt.FromHexString(test.toHex).ByRc4([]byte(test.key))

		t.Run(fmt.Sprintf("test_hex_%d", index), func(t *testing.T) {
			assert.Nil(t, d1.Error)
			assert.Equal(t, test.input, d1.ToString())
		})

		t.Run(fmt.Sprintf("test_base64_%d", index), func(t *testing.T) {
			d := Decrypt.FromBase64String(test.toBase64).ByRc4([]byte(test.key))
			assert.Nil(t, d.Error)
			assert.Equal(t, test.input, d.ToString())
		})
	}
}

func TestRc4_Encrypt_Bytes(t *testing.T) {
	for index, test := range rc4Tests {
		e := Encrypt.FromBytes([]byte(test.input)).ByRc4([]byte(test.key))

		t.Run(fmt.Sprintf("test_%d", index), func(t *testing.T) {
			assert.Nil(t, e.Error)
			assert.Equal(t, []byte(test.toHex), e.ToHexBytes())
			assert.Equal(t, []byte(test.toBase64), e.ToBase64Bytes())
		})
	}
}

func TestRc4_Decrypt_Bytes(t *testing.T) {
	for index, test := range rc4Tests {
		d1 := Decrypt.FromHexBytes([]byte(test.toHex)).ByRc4([]byte(test.key))
		t.Run(fmt.Sprintf("test_hex_%d", index), func(t *testing.T) {
			assert.Nil(t, d1.Error)
			assert.Equal(t, []byte(test.input), d1.ToBytes())
		})

		d2 := Decrypt.FromBase64Bytes([]byte(test.toBase64)).ByRc4([]byte(test.key))
		t.Run(fmt.Sprintf("test_base64_%d", index), func(t *testing.T) {
			assert.Nil(t, d2.Error)
			assert.Equal(t, []byte(test.input), d2.ToBytes())
		})
	}
}

func TestRc4_Key_Error(t *testing.T) {
	key := []byte("")

	e := Encrypt.FromString("hello go").ByRc4(key)
	assert.Equal(t, rc4Error.KeyError(), e.Error)

	d := Decrypt.FromRawString("hello go").ByRc4(key)
	assert.Equal(t, rc4Error.KeyError(), d.Error)
}
