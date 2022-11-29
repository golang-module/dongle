package dongle

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

var aesEcbTest = []struct {
	padding  cipherPadding
	input    string
	key      string
	iv       string
	toHex    string
	toBase32 string
	toBase64 string
}{
	{"", "", "0123456789abcdef", "0123456789abcdef", "", "", ""},
	{No, "hello world, go!", "0123456789abcdef", "0123456789abcdef", "f82a4c0db7a82f70c7fa84c39fa7627b", "7AVEYDNXVAXXBR72QTBZ7J3CPM======", "+CpMDbeoL3DH+oTDn6diew=="},
	{Zero, "hello world", "0123456789abcdef", "0123456789abcdef", "769c326290511c93bd59bba9c24d8904", "O2ODEYUQKEOJHPKZXOU4ETMJAQ======", "dpwyYpBRHJO9Wbupwk2JBA=="},
	{PKCS5, "hello world", "0123456789abcdef", "0123456789abcdef", "8169bed4ef49a8874559c5b200daade7", "QFU35VHPJGUIORKZYWZABWVN44======", "gWm+1O9JqIdFWcWyANqt5w=="},
	{PKCS7, "hello world", "0123456789abcdef", "0123456789abcdef", "8169bed4ef49a8874559c5b200daade7", "QFU35VHPJGUIORKZYWZABWVN44======", "gWm+1O9JqIdFWcWyANqt5w=="},
}

func TestEncrypt_ByAes_ECB_ToString(t *testing.T) {
	for index, test := range aesEcbTest {
		e := Encrypt.FromString(test.input).ByAes(aesCipher(ECB, test.padding, test.key, test.iv))
		assert.Nil(t, e.Error)

		assert.Equal(t, test.toHex, e.ToHexString(), "Hex test index is "+strconv.Itoa(index))
		assert.Equal(t, test.toBase32, e.ToBase32String(), "Base32 test index is "+strconv.Itoa(index))
		assert.Equal(t, test.toBase64, e.ToBase64String(), "Base64 test index is "+strconv.Itoa(index))
	}
}

func TestDecrypt_ByAes_ECB_ToString(t *testing.T) {
	for index, test := range aesEcbTest {
		e := Decrypt.FromHexString(test.toHex).ByAes(aesCipher(ECB, test.padding, test.key, test.iv))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.input, e.ToString(), "Hex test index is "+strconv.Itoa(index))
	}

	for index, test := range aesEcbTest {
		e := Decrypt.FromBase32String(test.toBase32).ByAes(aesCipher(ECB, test.padding, test.key, test.iv))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.input, e.ToString(), "Base32 test index is "+strconv.Itoa(index))
	}

	for index, test := range aesEcbTest {
		e := Decrypt.FromBase64String(test.toBase64).ByAes(aesCipher(ECB, test.padding, test.key, test.iv))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.input, e.ToString(), "Base64 test index is "+strconv.Itoa(index))
	}
}

func TestEncrypt_ByAes_ECB_ToBytes(t *testing.T) {
	for index, test := range aesEcbTest {
		e := Encrypt.FromString(test.input).ByAes(aesCipher(ECB, test.padding, test.key, test.iv))
		assert.Nil(t, e.Error)

		assert.Equal(t, []byte(test.toHex), e.ToHexBytes(), "Hex test index is "+strconv.Itoa(index))
		assert.Equal(t, []byte(test.toBase32), e.ToBase32Bytes(), "Base32 test index is "+strconv.Itoa(index))
		assert.Equal(t, []byte(test.toBase64), e.ToBase64Bytes(), "Base64 test index is "+strconv.Itoa(index))
	}
}

func TestDecrypt_ByAes_ECB_ToBytes(t *testing.T) {
	for index, test := range aesEcbTest {
		e := Decrypt.FromHexBytes([]byte(test.toHex)).ByAes(aesCipher(ECB, test.padding, test.key, test.iv))
		assert.Nil(t, e.Error)
		assert.Equal(t, []byte(test.input), e.ToBytes(), "Hex test index is "+strconv.Itoa(index))
	}

	for index, test := range aesEcbTest {
		e := Decrypt.FromBase32Bytes([]byte(test.toBase32)).ByAes(aesCipher(ECB, test.padding, test.key, test.iv))
		assert.Nil(t, e.Error)
		assert.Equal(t, []byte(test.input), e.ToBytes(), "Base32 test index is "+strconv.Itoa(index))
	}

	for index, test := range aesEcbTest {
		e := Decrypt.FromBase64Bytes([]byte(test.toBase64)).ByAes(aesCipher(ECB, test.padding, test.key, test.iv))
		assert.Nil(t, e.Error)
		assert.Equal(t, []byte(test.input), e.ToBytes(), "Base64 test index is "+strconv.Itoa(index))
	}
}
