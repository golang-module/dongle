package dongle

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	desInput = "hello world"
	desKey   = "12345678"
	desIV    = "12345678"
)

var desCipher = func(mode cipherMode, padding cipherPadding, key, iv interface{}) *Cipher {
	cipher := NewCipher()
	cipher.SetMode(mode)
	cipher.SetPadding(padding)
	cipher.SetKey(key)
	cipher.SetIV(iv)
	return cipher
}

func TestDes_Key_Error(t *testing.T) {
	e := Encrypt.FromString(desInput).ByDes(desCipher(CBC, PKCS7, "xxxx", desIV))
	assert.Equal(t, invalidDesKeyError(), e.Error)

	d := Decrypt.FromHexString("68656c6c6f20776f726c64").ByDes(desCipher(CBC, PKCS7, "xxxx", desIV))
	assert.Equal(t, invalidDesKeyError(), d.Error)
}

func TestDes_IV_Error(t *testing.T) {
	e := Encrypt.FromString(desInput).ByDes(desCipher(OFB, PKCS7, desKey, "xxxx"))
	assert.Equal(t, invalidDesIVError(), e.Error)

	d := Decrypt.FromHexString("68656c6c6f20776f726c64").ByDes(desCipher(CBC, PKCS7, desKey, "xxxx"))
	assert.Equal(t, invalidDesIVError(), d.Error)
}

func TestDes_Mode_Error(t *testing.T) {
	e := Encrypt.FromString(desInput).ByDes(desCipher("xxxx", PKCS7, desKey, desIV))
	assert.Equal(t, invalidModeError("xxxx"), e.Error)

	d := Decrypt.FromHexString("68656c6c6f20776f726c64").ByDes(desCipher("xxxx", PKCS7, desKey, desIV))
	assert.Equal(t, invalidModeError("xxxx"), d.Error)
}

func TestDes_Padding_Error(t *testing.T) {
	e := Encrypt.FromString(desInput).ByDes(desCipher(CFB, "xxxx", desKey, desIV))
	assert.Equal(t, invalidPaddingError("xxxx"), e.Error)

	d := Decrypt.FromHexString("68656c6c6f20776f726c64").ByDes(desCipher(CBC, "xxxx", desKey, desIV))
	assert.Equal(t, invalidPaddingError("xxxx"), d.Error)
}

func TestDes_Ciphertext_Error(t *testing.T) {
	d1 := Decrypt.FromHexBytes([]byte("xxxx")).ByDes(desCipher(CTR, Zero, []byte(aesKey), []byte(aesIV)))
	assert.Equal(t, invalidCiphertextError("hex"), d1.Error)

	d2 := Decrypt.FromBase32Bytes([]byte("xxxx")).ByDes(desCipher(CBC, PKCS5, []byte(aesKey), []byte(aesIV)))
	assert.Equal(t, invalidCiphertextError("base32"), d2.Error)

	d3 := Decrypt.FromBase64Bytes([]byte("xxxxxx")).ByDes(desCipher(CFB, PKCS7, []byte(aesKey), []byte(aesIV)))
	assert.Equal(t, invalidCiphertextError("base64"), d3.Error)
}
