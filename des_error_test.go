package dongle

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

var desCipher = func(mode, padding string, key, iv interface{}) *Cipher {
	cipher := NewCipher()
	cipher.SetMode(mode)
	cipher.SetPadding(padding)
	cipher.SetKey(key)
	cipher.SetIV(iv)
	return cipher
}

var desErrorTest = []struct {
	mode    string
	padding string
	input   string
	key     string
	iv      string
	error   error
}{
	{"xxxx", Zero, "hello world", "12345678", "12345678", invalidModeError("xxxx")},
	{CBC, "xxxx", "hello world", "12345678", "12345678", invalidPaddingError("xxxx")},
	{CFB, PKCS5, "hello world", "xxxx", "12345678", invalidDesKeyError(4)},
	{OFB, PKCS7, "hello world", "12345678", "xxxx", invalidDesIVError(4)},
	{ECB, No, "xxxx", "12345678", "12345678", invalidPlaintextError()},
}

func TestEncrypt_ByDes_Error(t *testing.T) {
	for index, test := range desErrorTest {
		e := Encrypt.FromString(test.input).ByDes(desCipher(test.mode, test.padding, test.key, test.iv))
		assert.Equal(t, test.error, e.Error, "Should catch an exception, current test index is "+strconv.Itoa(index))
	}
}

func TestDecrypt_ByDes_Error(t *testing.T) {
	for index, test := range desErrorTest {
		e := Decrypt.FromString(test.input).ByDes(desCipher(test.mode, test.padding, test.key, test.iv))
		assert.Equal(t, test.error, e.Error, "Should catch an exception, current test index is "+strconv.Itoa(index))
	}
}
