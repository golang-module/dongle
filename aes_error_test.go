package dongle

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	aesInput = "hello world"
	aesKey   = "0123456789abcdef"
	aesIV    = "0123456789abcdef"
)

var aesCipher = func(mode cipherMode, padding cipherPadding, key, iv interface{}) *Cipher {
	cipher := NewCipher()
	cipher.SetMode(mode)
	cipher.SetPadding(padding)
	cipher.SetKey(key)
	cipher.SetIV(iv)
	return cipher
}

func TestEncrypt_ByAes_FromString_Error(t *testing.T) {
	e1 := Encrypt.FromString(aesInput).ByAes(aesCipher("xxxx", PKCS7, aesKey, aesIV))
	assert.Equal(t, invalidModeError("xxxx"), e1.Error)

	e2 := Encrypt.FromString(aesInput).ByAes(aesCipher(CBC, "xxxx", aesKey, aesIV))
	assert.Equal(t, invalidPaddingError("xxxx"), e2.Error)

	e3 := Encrypt.FromString(aesInput).ByAes(aesCipher(CBC, PKCS7, "xxxx", aesIV))
	assert.Equal(t, invalidAesKeyError(4), e3.Error)

	e4 := Encrypt.FromString(aesInput).ByAes(aesCipher(CBC, PKCS7, aesKey, "xxxx"))
	assert.Equal(t, invalidAesIVError(4), e4.Error)
}

func TestEncrypt_ByAes_FromBytes_Error(t *testing.T) {
	e1 := Encrypt.FromBytes([]byte(aesInput)).ByAes(aesCipher("xxxx", PKCS7, []byte(aesKey), []byte(aesIV)))
	assert.Equal(t, invalidModeError("xxxx"), e1.Error)

	e2 := Encrypt.FromBytes([]byte(aesInput)).ByAes(aesCipher(CBC, "xxxx", []byte(aesKey), []byte(aesIV)))
	assert.Equal(t, invalidPaddingError("xxxx"), e2.Error)

	e3 := Encrypt.FromBytes([]byte(aesInput)).ByAes(aesCipher(CBC, PKCS7, []byte("xxxx"), []byte(aesIV)))
	assert.Equal(t, invalidAesKeyError(4), e3.Error)

	e4 := Encrypt.FromBytes([]byte(aesInput)).ByAes(aesCipher(CBC, PKCS7, []byte(aesKey), []byte("xxxx")))
	assert.Equal(t, invalidAesIVError(4), e4.Error)
}

func TestDecrypt_ByAes_FromString_Error(t *testing.T) {
	d1 := Decrypt.FromHexString("xxxx").ByAes(aesCipher(CBC, PKCS7, aesKey, aesIV))
	assert.Equal(t, invalidCiphertextError("hex"), d1.Error)
	d2 := Decrypt.FromBase32String("xxxx").ByAes(aesCipher(CBC, PKCS7, aesKey, aesIV))
	assert.Equal(t, invalidCiphertextError("base32"), d2.Error)
	d3 := Decrypt.FromBase64String("xxxxxx").ByAes(aesCipher(CBC, PKCS7, aesKey, aesIV))
	assert.Equal(t, invalidCiphertextError("base64"), d3.Error)

	d4 := Decrypt.FromHexString("68656c6c6f20776f726c64").ByAes(aesCipher("xxxx", PKCS7, aesKey, aesIV))
	assert.Equal(t, invalidModeError("xxxx"), d4.Error)
	d5 := Decrypt.FromHexString("68656c6c6f20776f726c64").ByAes(aesCipher(CBC, "xxxx", aesKey, aesIV))
	assert.Equal(t, invalidPaddingError("xxxx"), d5.Error)
	d6 := Decrypt.FromHexString("68656c6c6f20776f726c64").ByAes(aesCipher(CBC, PKCS7, "xxxx", aesIV))
	assert.Equal(t, invalidAesKeyError(4), d6.Error)
	d7 := Decrypt.FromHexString("68656c6c6f20776f726c64").ByAes(aesCipher(CBC, PKCS7, aesKey, "xxxx"))
	assert.Equal(t, invalidAesIVError(4), d7.Error)
}

func TestDecrypt_ByAes_FromBytes_Error(t *testing.T) {
	d1 := Decrypt.FromHexBytes([]byte("xxxx")).ByAes(aesCipher(CBC, PKCS7, aesKey, aesIV))
	assert.Equal(t, invalidCiphertextError("hex"), d1.Error)
	d2 := Decrypt.FromBase32Bytes([]byte("xxxx")).ByAes(aesCipher(CBC, PKCS7, aesKey, aesIV))
	assert.Equal(t, invalidCiphertextError("base32"), d2.Error)
	d3 := Decrypt.FromBase64Bytes([]byte("xxxxxx")).ByAes(aesCipher(CBC, PKCS7, aesKey, aesIV))
	assert.Equal(t, invalidCiphertextError("base64"), d3.Error)

	d4 := Decrypt.FromHexBytes([]byte("68656c6c6f20776f726c64")).ByAes(aesCipher("xxxx", PKCS7, []byte(aesKey), []byte(aesIV)))
	assert.Equal(t, invalidModeError("xxxx"), d4.Error, "mode error")
	d5 := Decrypt.FromHexBytes([]byte("68656c6c6f20776f726c64")).ByAes(aesCipher(CBC, "xxxx", []byte(aesKey), []byte(aesIV)))
	assert.Equal(t, invalidPaddingError("xxxx"), d5.Error, "padding error")
	d6 := Decrypt.FromHexBytes([]byte("68656c6c6f20776f726c64")).ByAes(aesCipher(CBC, PKCS7, []byte("xxxx"), []byte(aesIV)))
	assert.Equal(t, invalidAesKeyError(4), d6.Error, "key error")
	d7 := Decrypt.FromHexBytes([]byte("68656c6c6f20776f726c64")).ByAes(aesCipher(CBC, PKCS7, []byte(aesKey), []byte("xxxx")))
	assert.Equal(t, invalidAesIVError(4), d7.Error, "iv error")
}
