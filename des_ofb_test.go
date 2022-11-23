package dongle

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

var desOfbTest = []struct {
	padding  cipherPadding
	input    string
	key      string
	iv       string
	toHex    string
	toBase64 string
}{
	{"", "", "12345678", "12345678", "", ""},
	{No, "hello world, go!", "12345678", "12345678", "feb56ee417f5fbe69b19fcf309087d8b", "/rVu5Bf1++abGfzzCQh9iw=="},
	{Zero, "hello world", "12345678", "12345678", "feb56ee417f5fbe69b19fcdf296f12aa", "/rVu5Bf1++abGfzfKW8Sqg=="},
	{PKCS5, "hello world", "12345678", "12345678", "feb56ee417f5fbe69b19fcda2c6a17af", "/rVu5Bf1++abGfzaLGoXrw=="},
	{PKCS7, "hello world", "12345678", "12345678", "feb56ee417f5fbe69b19fcda2c6a17af", "/rVu5Bf1++abGfzaLGoXrw=="},
}

func TestEncrypt_ByDes_OFB_ToString(t *testing.T) {
	for index, test := range desOfbTest {
		e := Encrypt.FromString(test.input).ByDes(desCipher(OFB, test.padding, test.key, test.iv))
		assert.Nil(t, e.Error)

		assert.Equal(t, test.toHex, e.ToHexString(), "Hex test index is "+strconv.Itoa(index))
		assert.Equal(t, test.toBase64, e.ToBase64String(), "Base64 test index is "+strconv.Itoa(index))
	}
}

func TestDecrypt_ByDes_OFB_ToString(t *testing.T) {
	for index, test := range desOfbTest {
		e := Decrypt.FromHexString(test.toHex).ByDes(desCipher(OFB, test.padding, test.key, test.iv))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.input, e.ToString(), "Hex test index is "+strconv.Itoa(index))
	}

	for index, test := range desOfbTest {
		e := Decrypt.FromBase64String(test.toBase64).ByDes(desCipher(OFB, test.padding, test.key, test.iv))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.input, e.ToString(), "Base64 test index is "+strconv.Itoa(index))
	}
}

func TestEncrypt_ByDes_OFB_ToBytes(t *testing.T) {
	for index, test := range desOfbTest {
		e := Encrypt.FromString(test.input).ByDes(desCipher(OFB, test.padding, test.key, test.iv))
		assert.Nil(t, e.Error)

		assert.Equal(t, []byte(test.toHex), e.ToHexBytes(), "Hex test index is "+strconv.Itoa(index))
		assert.Equal(t, []byte(test.toBase64), e.ToBase64Bytes(), "Base64 test index is "+strconv.Itoa(index))
	}
}

func TestDecrypt_ByDes_OFB_ToBytes(t *testing.T) {
	for index, test := range desOfbTest {
		e := Decrypt.FromHexBytes([]byte(test.toHex)).ByDes(desCipher(OFB, test.padding, test.key, test.iv))
		assert.Nil(t, e.Error)
		assert.Equal(t, []byte(test.input), e.ToBytes(), "Hex test index is "+strconv.Itoa(index))
	}

	for index, test := range desOfbTest {
		e := Decrypt.FromBase64Bytes([]byte(test.toBase64)).ByDes(desCipher(OFB, test.padding, test.key, test.iv))
		assert.Nil(t, e.Error)
		assert.Equal(t, []byte(test.input), e.ToBytes(), "Base64 test index is "+strconv.Itoa(index))
	}
}
