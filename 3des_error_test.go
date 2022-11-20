package dongle

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

var tripleDesCipher = func(mode, padding string, key, iv interface{}) *Cipher {
	cipher := NewCipher()
	cipher.SetMode(mode)
	cipher.SetPadding(padding)
	cipher.SetKey(key)
	cipher.SetIV(iv)
	return cipher
}

var tripleDesErrorTest = []struct {
	mode    string
	padding string
	input   string
	key     string
	iv      string
	error   error
}{
	{"xxxx", Zero, "hello world", "0123456789abcdefghijklmn", "12345678", invalidModeError("xxxx")},
	{CBC, "xxxx", "hello world", "0123456789abcdefghijklmn", "12345678", invalidPaddingError("xxxx")},
	{CFB, PKCS5, "hello world", "xxxx", "12345678", invalid3DesKeyError(4)},
	{OFB, PKCS7, "hello world", "0123456789abcdefghijklmn", "xxxx", invalid3DesIVError(4)},
	{ECB, No, "xxxx", "0123456789abcdefghijklmn", "12345678", invalidPlaintextError()},
}

func TestEncrypt_By3Des_Error(t *testing.T) {
	for index, test := range tripleDesErrorTest {
		e := Encrypt.FromString(test.input).By3Des(tripleDesCipher(test.mode, test.padding, test.key, test.iv))
		assert.Equal(t, test.error, e.Error, "Should catch an exception, current test index is "+strconv.Itoa(index))
	}
}

func TestDecrypt_By3Des_Error(t *testing.T) {
	for index, test := range tripleDesErrorTest {
		e := Decrypt.FromString(test.input).By3Des(tripleDesCipher(test.mode, test.padding, test.key, test.iv))
		assert.Equal(t, test.error, e.Error, "Should catch an exception, current test index is "+strconv.Itoa(index))
	}
}
