package dongle

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

var tripleDesCtrTest = []struct {
	padding  string
	input    string
	key      string
	iv       string
	toHex    string
	toBase32 string
	toBase64 string
}{
	{"", "", "0123456789abcdefghijklmn", "12345678", "", "", ""},
	{No, "hello world, go!", "0123456789abcdefghijklmn", "12345678", "f52dc896ceff0ecc366b2281038f6f7f", "6UW4RFWO74HMYNTLEKAQHD3PP4======", "9S3Ils7/Dsw2ayKBA49vfw=="},
	{Zero, "hello world", "0123456789abcdefghijklmn", "12345678", "f52dc896ceff0ecc366b22ad23e8005e", "6UW4RFWO74HMYNTLEKWSH2AALY======", "9S3Ils7/Dsw2ayKtI+gAXg=="},
	{PKCS5, "hello world", "0123456789abcdefghijklmn", "12345678", "f52dc896ceff0ecc366b22a826ed055b", "6UW4RFWO74HMYNTLEKUCN3IFLM======", "9S3Ils7/Dsw2ayKoJu0FWw=="},
	{PKCS7, "hello world", "0123456789abcdefghijklmn", "12345678", "f52dc896ceff0ecc366b22a826ed055b", "6UW4RFWO74HMYNTLEKUCN3IFLM======", "9S3Ils7/Dsw2ayKoJu0FWw=="},
}

func TestEncrypt_By3Des_CTR_ToString(t *testing.T) {
	for index, test := range tripleDesCtrTest {
		e := Encrypt.FromString(test.input).By3Des(tripleDesCipher(CTR, test.padding, test.key, test.iv))
		assert.Nil(t, e.Error)

		assert.Equal(t, test.toHex, e.ToHexString(), "Hex test index is "+strconv.Itoa(index))
		assert.Equal(t, test.toBase32, e.ToBase32String(), "Base32 test index is "+strconv.Itoa(index))
		assert.Equal(t, test.toBase64, e.ToBase64String(), "Base64 test index is "+strconv.Itoa(index))
	}
}

func TestDecrypt_By3Des_CTR_ToString(t *testing.T) {
	for index, test := range tripleDesCtrTest {
		e := Decrypt.FromHexString(test.toHex).By3Des(tripleDesCipher(CTR, test.padding, test.key, test.iv))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.input, e.ToString(), "Hex test index is "+strconv.Itoa(index))
	}

	for index, test := range tripleDesCtrTest {
		e := Decrypt.FromBase32String(test.toBase32).By3Des(tripleDesCipher(CTR, test.padding, test.key, test.iv))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.input, e.ToString(), "Base32 test index is "+strconv.Itoa(index))
	}

	for index, test := range tripleDesCtrTest {
		e := Decrypt.FromBase64String(test.toBase64).By3Des(tripleDesCipher(CTR, test.padding, test.key, test.iv))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.input, e.ToString(), "Base64 test index is "+strconv.Itoa(index))
	}
}

func TestEncrypt_By3Des_CTR_ToBytes(t *testing.T) {
	for index, test := range tripleDesCtrTest {
		e := Encrypt.FromString(test.input).By3Des(tripleDesCipher(CTR, test.padding, []byte(test.key), []byte(test.iv)))
		assert.Nil(t, e.Error)

		assert.Equal(t, []byte(test.toHex), e.ToHexBytes(), "Hex test index is "+strconv.Itoa(index))
		assert.Equal(t, []byte(test.toBase32), e.ToBase32Bytes(), "Base32 test index is "+strconv.Itoa(index))
		assert.Equal(t, []byte(test.toBase64), e.ToBase64Bytes(), "Base64 test index is "+strconv.Itoa(index))
	}
}

func TestDecrypt_By3Des_CTR_ToBytes(t *testing.T) {
	for index, test := range tripleDesCtrTest {
		e := Decrypt.FromHexBytes([]byte(test.toHex)).By3Des(tripleDesCipher(CTR, test.padding, []byte(test.key), []byte(test.iv)))
		assert.Nil(t, e.Error)
		assert.Equal(t, []byte(test.input), e.ToBytes(), "Hex test index is "+strconv.Itoa(index))
	}

	for index, test := range tripleDesCtrTest {
		e := Decrypt.FromBase32Bytes([]byte(test.toBase32)).By3Des(tripleDesCipher(CTR, test.padding, []byte(test.key), []byte(test.iv)))
		assert.Nil(t, e.Error)
		assert.Equal(t, []byte(test.input), e.ToBytes(), "Base32 test index is "+strconv.Itoa(index))
	}

	for index, test := range tripleDesCtrTest {
		e := Decrypt.FromBase64Bytes([]byte(test.toBase64)).By3Des(tripleDesCipher(CTR, test.padding, []byte(test.key), []byte(test.iv)))
		assert.Nil(t, e.Error)
		assert.Equal(t, []byte(test.input), e.ToBytes(), "Base64 test index is "+strconv.Itoa(index))
	}
}
