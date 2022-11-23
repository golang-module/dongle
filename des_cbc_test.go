package dongle

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

var desCbcTest = []struct {
	padding  cipherPadding
	input    string
	key      string
	iv       string
	toHex    string
	toBase64 string
}{
	{"", "", "12345678", "12345678", "", ""},
	{No, "hello world, go!", "12345678", "12345678", "0b2a92e81fb49ce1b161ab1a2e35d7ce", "CyqS6B+0nOGxYasaLjXXzg=="},
	{Zero, "hello world", "12345678", "12345678", "0b2a92e81fb49ce1e02ecfa986df2ec8", "CyqS6B+0nOHgLs+pht8uyA=="},
	{PKCS5, "hello world", "12345678", "12345678", "0b2a92e81fb49ce1a43266aacaea7b81", "CyqS6B+0nOGkMmaqyup7gQ=="},
	{PKCS7, "hello world", "12345678", "12345678", "0b2a92e81fb49ce1a43266aacaea7b81", "CyqS6B+0nOGkMmaqyup7gQ=="},
}

func TestEncrypt_ByDes_CBC_ToString(t *testing.T) {
	for index, test := range desCbcTest {
		e := Encrypt.FromString(test.input).ByDes(desCipher(CBC, test.padding, test.key, test.iv))
		assert.Nil(t, e.Error)

		assert.Equal(t, test.toHex, e.ToHexString(), "Hex test index is "+strconv.Itoa(index))
		assert.Equal(t, test.toBase64, e.ToBase64String(), "Base64 test index is "+strconv.Itoa(index))
	}
}

func TestDecrypt_ByDes_CBC_ToString(t *testing.T) {
	for index, test := range desCbcTest {
		e := Decrypt.FromHexString(test.toHex).ByDes(desCipher(CBC, test.padding, test.key, test.iv))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.input, e.ToString(), "Hex test index is "+strconv.Itoa(index))
	}

	for index, test := range desCbcTest {
		e := Decrypt.FromBase64String(test.toBase64).ByDes(desCipher(CBC, test.padding, test.key, test.iv))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.input, e.ToString(), "Base64 test index is "+strconv.Itoa(index))
	}
}

func TestEncrypt_ByDes_CBC_ToBytes(t *testing.T) {
	for index, test := range desCbcTest {
		e := Encrypt.FromString(test.input).ByDes(desCipher(CBC, test.padding, test.key, test.iv))
		assert.Nil(t, e.Error)

		assert.Equal(t, []byte(test.toHex), e.ToHexBytes(), "Hex test index is "+strconv.Itoa(index))
		assert.Equal(t, []byte(test.toBase64), e.ToBase64Bytes(), "Base64 test index is "+strconv.Itoa(index))
	}
}

func TestDecrypt_ByDes_CBC_ToBytes(t *testing.T) {
	for index, test := range desCbcTest {
		e := Decrypt.FromHexBytes([]byte(test.toHex)).ByDes(desCipher(CBC, test.padding, test.key, test.iv))
		assert.Nil(t, e.Error)
		assert.Equal(t, []byte(test.input), e.ToBytes(), "Hex test index is "+strconv.Itoa(index))
	}

	for index, test := range desCbcTest {
		e := Decrypt.FromBase64Bytes([]byte(test.toBase64)).ByDes(desCipher(CBC, test.padding, test.key, test.iv))
		assert.Nil(t, e.Error)
		assert.Equal(t, []byte(test.input), e.ToBytes(), "Base64 test index is "+strconv.Itoa(index))
	}
}
