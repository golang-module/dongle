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

func TestAes_Key_Error(t *testing.T) {
	e := Encrypt.FromString(aesInput).ByAes(aesCipher(CBC, PKCS7, "xxxx", aesIV))
	assert.Equal(t, invalidAesKeyError(), e.Error)

	d := Decrypt.FromHexString("68656c6c6f20776f726c64").ByAes(aesCipher(CBC, PKCS7, "xxxx", aesIV))
	assert.Equal(t, invalidAesKeyError(), d.Error)
}

func TestAes_IV_Error(t *testing.T) {
	e := Encrypt.FromString(aesInput).ByAes(aesCipher(OFB, PKCS7, aesKey, "xxxx"))
	assert.Equal(t, invalidAesIVError(), e.Error)

	d := Decrypt.FromHexString("68656c6c6f20776f726c64").ByAes(aesCipher(CBC, PKCS7, aesKey, "xxxx"))
	assert.Equal(t, invalidAesIVError(), d.Error)
}

func TestAes_Mode_Error(t *testing.T) {
	e := Encrypt.FromString(aesInput).ByAes(aesCipher("xxxx", PKCS7, aesKey, aesIV))
	assert.Equal(t, invalidModeError("xxxx"), e.Error)

	d := Decrypt.FromHexString("68656c6c6f20776f726c64").ByAes(aesCipher("xxxx", PKCS7, aesKey, aesIV))
	assert.Equal(t, invalidModeError("xxxx"), d.Error)
}

func TestAes_Padding_Error(t *testing.T) {
	e := Encrypt.FromString(aesInput).ByAes(aesCipher(CFB, "xxxx", aesKey, aesIV))
	assert.Equal(t, invalidPaddingError("xxxx"), e.Error)

	d := Decrypt.FromHexString("68656c6c6f20776f726c64").ByAes(aesCipher(CBC, "xxxx", aesKey, aesIV))
	assert.Equal(t, invalidPaddingError("xxxx"), d.Error)
}

func TestAes_Plaintext_Error(t *testing.T) {
	e := Encrypt.FromString(aesInput).ByAes(aesCipher(CFB, No, aesKey, aesIV))
	assert.Equal(t, invalidPlaintextError(), e.Error)

	d := Decrypt.FromHexString("68656c6c6f20776f726c64").ByAes(aesCipher(CBC, No, aesKey, aesIV))
	assert.Equal(t, invalidPlaintextError(), d.Error)
}

func TestAes_Ciphertext_Error(t *testing.T) {
	d1 := Decrypt.FromHexString("xxxx").ByAes(aesCipher(CTR, Zero, aesKey, aesIV))
	assert.Equal(t, invalidCiphertextError("hex"), d1.Error)
	d2 := Decrypt.FromHexBytes([]byte("xxxx")).ByAes(aesCipher(CTR, Zero, []byte(aesKey), []byte(aesIV)))
	assert.Equal(t, invalidCiphertextError("hex"), d2.Error)

	d3 := Decrypt.FromBase32String("xxxx").ByAes(aesCipher(CBC, PKCS5, aesKey, aesIV))
	assert.Equal(t, invalidCiphertextError("base32"), d3.Error)
	d4 := Decrypt.FromBase32Bytes([]byte("xxxx")).ByAes(aesCipher(CBC, PKCS5, []byte(aesKey), []byte(aesIV)))
	assert.Equal(t, invalidCiphertextError("base32"), d4.Error)

	d5 := Decrypt.FromBase64String("xxxxxx").ByAes(aesCipher(CFB, PKCS7, aesKey, aesIV))
	assert.Equal(t, invalidCiphertextError("base64"), d5.Error)
	d6 := Decrypt.FromBase64Bytes([]byte("xxxxxx")).ByAes(aesCipher(CFB, PKCS7, []byte(aesKey), []byte(aesIV)))
	assert.Equal(t, invalidCiphertextError("base64"), d6.Error)
}
