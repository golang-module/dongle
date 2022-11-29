package dongle

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	tripleDesInput = "hello world"
	tripleDesKey   = "0123456789abcdefghijklmn"
	tripleDesIV    = "12345678"
)

var tripleDesCipher = func(mode cipherMode, padding cipherPadding, key, iv interface{}) *Cipher {
	cipher := NewCipher()
	cipher.SetMode(mode)
	cipher.SetPadding(padding)
	cipher.SetKey(key)
	cipher.SetIV(iv)
	return cipher
}

func Test3Des_Key_Error(t *testing.T) {
	e := Encrypt.FromString(tripleDesInput).By3Des(tripleDesCipher(CBC, PKCS7, "xxxx", tripleDesIV))
	assert.Equal(t, invalid3DesKeyError(), e.Error)

	d := Decrypt.FromHexString("68656c6c6f20776f726c64").By3Des(tripleDesCipher(CBC, PKCS7, "xxxx", tripleDesIV))
	assert.Equal(t, invalid3DesKeyError(), d.Error)
}

func Test3Des_IV_Error(t *testing.T) {
	e := Encrypt.FromString(tripleDesInput).By3Des(tripleDesCipher(OFB, PKCS7, tripleDesKey, "xxxx"))
	assert.Equal(t, invalid3DesIVError(), e.Error)

	d := Decrypt.FromHexString("68656c6c6f20776f726c64").By3Des(tripleDesCipher(CBC, PKCS7, tripleDesKey, "xxxx"))
	assert.Equal(t, invalid3DesIVError(), d.Error)
}

func Test3Des_Mode_Error(t *testing.T) {
	e := Encrypt.FromString(tripleDesInput).By3Des(tripleDesCipher("xxxx", PKCS7, tripleDesKey, tripleDesIV))
	assert.Equal(t, invalidModeError("xxxx"), e.Error)

	d := Decrypt.FromHexString("68656c6c6f20776f726c64").By3Des(tripleDesCipher("xxxx", PKCS7, tripleDesKey, tripleDesIV))
	assert.Equal(t, invalidModeError("xxxx"), d.Error)
}

func Test3Des_Padding_Error(t *testing.T) {
	e := Encrypt.FromString(tripleDesInput).By3Des(tripleDesCipher(CFB, "xxxx", tripleDesKey, tripleDesIV))
	assert.Equal(t, invalidPaddingError("xxxx"), e.Error)

	d := Decrypt.FromHexString("68656c6c6f20776f726c64").By3Des(tripleDesCipher(CBC, "xxxx", tripleDesKey, tripleDesIV))
	assert.Equal(t, invalidPaddingError("xxxx"), d.Error)
}

func Test3Des_Ciphertext_Error(t *testing.T) {
	d1 := Decrypt.FromHexBytes([]byte("xxxx")).By3Des(tripleDesCipher(CTR, Zero, []byte(aesKey), []byte(aesIV)))
	assert.Equal(t, invalidCiphertextError("hex"), d1.Error)

	d2 := Decrypt.FromBase32Bytes([]byte("xxxx")).By3Des(tripleDesCipher(CBC, PKCS5, []byte(aesKey), []byte(aesIV)))
	assert.Equal(t, invalidCiphertextError("base32"), d2.Error)

	d3 := Decrypt.FromBase64Bytes([]byte("xxxxxx")).By3Des(tripleDesCipher(CFB, PKCS7, []byte(aesKey), []byte(aesIV)))
	assert.Equal(t, invalidCiphertextError("base64"), d3.Error)
}
