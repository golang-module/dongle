package dongle

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

var desEcbTest = []struct {
	padding  string
	input    string
	key      string
	iv       string
	toHex    string
	toBase64 string
}{
	{"", "", "12345678", "12345678", "", ""},
	{No, "hello world, go!", "12345678", "12345678", "28dba02eb5f6dd47e5a66298af31d4ca", "KNugLrX23UflpmKYrzHUyg=="},
	{Zero, "hello world", "12345678", "12345678", "28dba02eb5f6dd476042daebfa59687a", "KNugLrX23UdgQtrr+lloeg=="},
	{PKCS5, "hello world", "12345678", "12345678", "28dba02eb5f6dd475d82e3681c83bb77", "KNugLrX23UddguNoHIO7dw=="},
	{PKCS7, "hello world", "12345678", "12345678", "28dba02eb5f6dd475d82e3681c83bb77", "KNugLrX23UddguNoHIO7dw=="},
}

func TestEncrypt_ByDes_ECB_ToString(t *testing.T) {
	for index, test := range desEcbTest {
		e := Encrypt.FromString(test.input).ByDes(desCipher(ECB, test.padding, test.key, test.iv))
		assert.Nil(t, e.Error)

		assert.Equal(t, test.toHex, e.ToHexString(), "Hex test index is "+strconv.Itoa(index))
		assert.Equal(t, test.toBase64, e.ToBase64String(), "Base64 test index is "+strconv.Itoa(index))
	}
}

func TestDecrypt_ByDes_ECB_ToString(t *testing.T) {
	for index, test := range desEcbTest {
		e := Decrypt.FromHexString(test.toHex).ByDes(desCipher(ECB, test.padding, test.key, test.iv))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.input, e.ToString(), "Hex test index is "+strconv.Itoa(index))
	}

	for index, test := range desEcbTest {
		e := Decrypt.FromBase64String(test.toBase64).ByDes(desCipher(ECB, test.padding, test.key, test.iv))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.input, e.ToString(), "Base64 test index is "+strconv.Itoa(index))
	}
}

func TestEncrypt_ByDes_ECB_ToBytes(t *testing.T) {
	for index, test := range desEcbTest {
		e := Encrypt.FromString(test.input).ByDes(desCipher(ECB, test.padding, test.key, test.iv))
		assert.Nil(t, e.Error)

		assert.Equal(t, []byte(test.toHex), e.ToHexBytes(), "Hex test index is "+strconv.Itoa(index))
		assert.Equal(t, []byte(test.toBase64), e.ToBase64Bytes(), "Base64 test index is "+strconv.Itoa(index))
	}
}

func TestDecrypt_ByDes_ECB_ToBytes(t *testing.T) {
	for index, test := range desEcbTest {
		e := Decrypt.FromHexBytes([]byte(test.toHex)).ByDes(desCipher(ECB, test.padding, test.key, test.iv))
		assert.Nil(t, e.Error)
		assert.Equal(t, []byte(test.input), e.ToBytes(), "Hex test index is "+strconv.Itoa(index))
	}

	for index, test := range desEcbTest {
		e := Decrypt.FromBase64Bytes([]byte(test.toBase64)).ByDes(desCipher(ECB, test.padding, test.key, test.iv))
		assert.Nil(t, e.Error)
		assert.Equal(t, []byte(test.input), e.ToBytes(), "Base64 test index is "+strconv.Itoa(index))
	}
}
