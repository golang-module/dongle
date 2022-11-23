package dongle

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

var desCfbTest = []struct {
	padding  cipherPadding
	input    string
	key      string
	iv       string
	toHex    string
	toBase64 string
}{
	{"", "", "12345678", "12345678", "", ""},
	{No, "hello world, go!", "12345678", "12345678", "feb56ee417f5fbe61a4e5700b93c340b", "/rVu5Bf1++YaTlcAuTw0Cw=="},
	{Zero, "hello world", "12345678", "12345678", "feb56ee417f5fbe61a4e572c995b5b2a", "/rVu5Bf1++YaTlcsmVtbKg=="},
	{PKCS5, "hello world", "12345678", "12345678", "feb56ee417f5fbe61a4e57299c5e5e2f", "/rVu5Bf1++YaTlcpnF5eLw=="},
	{PKCS7, "hello world", "12345678", "12345678", "feb56ee417f5fbe61a4e57299c5e5e2f", "/rVu5Bf1++YaTlcpnF5eLw=="},
}

func TestEncrypt_ByDes_CFB_ToString(t *testing.T) {
	for index, test := range desCfbTest {
		e := Encrypt.FromString(test.input).ByDes(desCipher(CFB, test.padding, test.key, test.iv))
		assert.Nil(t, e.Error)

		assert.Equal(t, test.toHex, e.ToHexString(), "Hex test index is "+strconv.Itoa(index))
		assert.Equal(t, test.toBase64, e.ToBase64String(), "Base64 test index is "+strconv.Itoa(index))
	}
}

func TestDecrypt_ByDes_CFB_ToString(t *testing.T) {
	for index, test := range desCfbTest {
		e := Decrypt.FromHexString(test.toHex).ByDes(desCipher(CFB, test.padding, test.key, test.iv))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.input, e.ToString(), "Hex test index is "+strconv.Itoa(index))
	}

	for index, test := range desCfbTest {
		e := Decrypt.FromBase64String(test.toBase64).ByDes(desCipher(CFB, test.padding, test.key, test.iv))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.input, e.ToString(), "Base64 test index is "+strconv.Itoa(index))
	}
}

func TestEncrypt_ByDes_CFB_ToBytes(t *testing.T) {
	for index, test := range desCfbTest {
		e := Encrypt.FromString(test.input).ByDes(desCipher(CFB, test.padding, test.key, test.iv))
		assert.Nil(t, e.Error)

		assert.Equal(t, []byte(test.toHex), e.ToHexBytes(), "Hex test index is "+strconv.Itoa(index))
		assert.Equal(t, []byte(test.toBase64), e.ToBase64Bytes(), "Base64 test index is "+strconv.Itoa(index))
	}
}

func TestDecrypt_ByDes_CFB_ToBytes(t *testing.T) {
	for index, test := range desCfbTest {
		e := Decrypt.FromHexBytes([]byte(test.toHex)).ByDes(desCipher(CFB, test.padding, test.key, test.iv))
		assert.Nil(t, e.Error)
		assert.Equal(t, []byte(test.input), e.ToBytes(), "Hex test index is "+strconv.Itoa(index))
	}

	for index, test := range desCfbTest {
		e := Decrypt.FromBase64Bytes([]byte(test.toBase64)).ByDes(desCipher(CFB, test.padding, test.key, test.iv))
		assert.Nil(t, e.Error)
		assert.Equal(t, []byte(test.input), e.ToBytes(), "Base64 test index is "+strconv.Itoa(index))
	}
}
