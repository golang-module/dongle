package dongle

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

var desCtrTest = []struct {
	padding  cipherPadding
	input    string
	key      string
	iv       string
	toHex    string
	toBase64 string
}{
	{"", "", "12345678", "12345678", "", ""},
	{No, "hello world, go!", "12345678", "12345678", "feb56ee417f5fbe6337e8b3af91e9929", "/rVu5Bf1++Yzfos6+R6ZKQ=="},
	{Zero, "hello world", "12345678", "12345678", "feb56ee417f5fbe6337e8b16d979f608", "/rVu5Bf1++YzfosW2Xn2CA=="},
	{PKCS5, "hello world", "12345678", "12345678", "feb56ee417f5fbe6337e8b13dc7cf30d", "/rVu5Bf1++YzfosT3HzzDQ=="},
	{PKCS7, "hello world", "12345678", "12345678", "feb56ee417f5fbe6337e8b13dc7cf30d", "/rVu5Bf1++YzfosT3HzzDQ=="},
}

func TestEncrypt_ByDes_CTR_ToString(t *testing.T) {
	for index, test := range desCtrTest {
		e := Encrypt.FromString(test.input).ByDes(desCipher(CTR, test.padding, test.key, test.iv))
		assert.Nil(t, e.Error)

		assert.Equal(t, test.toHex, e.ToHexString(), "Hex test index is "+strconv.Itoa(index))
		assert.Equal(t, test.toBase64, e.ToBase64String(), "Base64 test index is "+strconv.Itoa(index))
	}
}

func TestDecrypt_ByDes_CTR_ToString(t *testing.T) {
	for index, test := range desCtrTest {
		e := Decrypt.FromHexString(test.toHex).ByDes(desCipher(CTR, test.padding, test.key, test.iv))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.input, e.ToString(), "Hex test index is "+strconv.Itoa(index))
	}

	for index, test := range desCtrTest {
		e := Decrypt.FromBase64String(test.toBase64).ByDes(desCipher(CTR, test.padding, test.key, test.iv))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.input, e.ToString(), "Base64 test index is "+strconv.Itoa(index))
	}
}

func TestEncrypt_ByDes_CTR_ToBytes(t *testing.T) {
	for index, test := range desCtrTest {
		e := Encrypt.FromString(test.input).ByDes(desCipher(CTR, test.padding, test.key, test.iv))
		assert.Nil(t, e.Error)

		assert.Equal(t, []byte(test.toHex), e.ToHexBytes(), "Hex test index is "+strconv.Itoa(index))
		assert.Equal(t, []byte(test.toBase64), e.ToBase64Bytes(), "Base64 test index is "+strconv.Itoa(index))
	}
}

func TestDecrypt_ByDes_CTR_ToBytes(t *testing.T) {
	for index, test := range desCtrTest {
		e := Decrypt.FromHexBytes([]byte(test.toHex)).ByDes(desCipher(CTR, test.padding, test.key, test.iv))
		assert.Nil(t, e.Error)
		assert.Equal(t, []byte(test.input), e.ToBytes(), "Hex test index is "+strconv.Itoa(index))
	}

	for index, test := range desCtrTest {
		e := Decrypt.FromBase64Bytes([]byte(test.toBase64)).ByDes(desCipher(CTR, test.padding, test.key, test.iv))
		assert.Nil(t, e.Error)
		assert.Equal(t, []byte(test.input), e.ToBytes(), "Base64 test index is "+strconv.Itoa(index))
	}
}
