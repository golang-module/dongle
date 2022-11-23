package dongle

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

var aesCbcTest = []struct {
	padding  cipherPadding
	input    string
	key      string
	iv       string
	toHex    string
	toBase32 string
	toBase64 string
}{
	{"", "", "0123456789abcdef", "0123456789abcdef", "", "", ""},
	{No, "hello world, go!", "0123456789abcdef", "0123456789abcdef", "77aa39926f9b2f3f22254bfd422fa75d", "O6VDTETPTMXT6IRFJP6UEL5HLU======", "d6o5km+bLz8iJUv9Qi+nXQ=="},
	{Zero, "hello world", "0123456789abcdef", "0123456789abcdef", "889935b7a0c64b0333d713cafaee08fe", "RCMTLN5AYZFQGM6XCPFPV3QI7Y======", "iJk1t6DGSwMz1xPK+u4I/g=="},
	{PKCS5, "hello world", "0123456789abcdef", "0123456789abcdef", "c1e9b4529aac9793010f4677f6358efe", "YHU3IUU2VSLZGAIPIZ37MNMO7Y======", "wem0Upqsl5MBD0Z39jWO/g=="},
	{PKCS7, "hello world", "0123456789abcdef", "0123456789abcdef", "c1e9b4529aac9793010f4677f6358efe", "YHU3IUU2VSLZGAIPIZ37MNMO7Y======", "wem0Upqsl5MBD0Z39jWO/g=="},
}

func TestEncrypt_ByAes_CBC_ToString(t *testing.T) {
	for index, test := range aesCbcTest {
		e := Encrypt.FromString(test.input).ByAes(aesCipher(CBC, test.padding, test.key, test.iv))
		assert.Nil(t, e.Error)

		assert.Equal(t, test.toHex, e.ToHexString(), "Hex test index is "+strconv.Itoa(index))
		assert.Equal(t, test.toBase32, e.ToBase32String(), "Base32 test index is "+strconv.Itoa(index))
		assert.Equal(t, test.toBase64, e.ToBase64String(), "Base64 test index is "+strconv.Itoa(index))
	}
}

func TestDecrypt_ByAes_CBC_ToString(t *testing.T) {
	for index, test := range aesCbcTest {
		e := Decrypt.FromHexString(test.toHex).ByAes(aesCipher(CBC, test.padding, test.key, test.iv))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.input, e.ToString(), "Hex test index is "+strconv.Itoa(index))
	}

	for index, test := range aesCbcTest {
		e := Decrypt.FromBase32String(test.toBase32).ByAes(aesCipher(CBC, test.padding, test.key, test.iv))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.input, e.ToString(), "Base32 test index is "+strconv.Itoa(index))
	}

	for index, test := range aesCbcTest {
		e := Decrypt.FromBase64String(test.toBase64).ByAes(aesCipher(CBC, test.padding, test.key, test.iv))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.input, e.ToString(), "Base64 test index is "+strconv.Itoa(index))
	}
}

func TestEncrypt_ByAes_CBC_ToBytes(t *testing.T) {
	for index, test := range aesCbcTest {
		e := Encrypt.FromString(test.input).ByAes(aesCipher(CBC, test.padding, []byte(test.key), []byte(test.iv)))
		assert.Nil(t, e.Error)

		assert.Equal(t, []byte(test.toHex), e.ToHexBytes(), "Hex test index is "+strconv.Itoa(index))
		assert.Equal(t, []byte(test.toBase32), e.ToBase32Bytes(), "Base32 test index is "+strconv.Itoa(index))
		assert.Equal(t, []byte(test.toBase64), e.ToBase64Bytes(), "Base64 test index is "+strconv.Itoa(index))
	}
}

func TestDecrypt_ByAes_CBC_ToBytes(t *testing.T) {
	for index, test := range aesCbcTest {
		e := Decrypt.FromHexBytes([]byte(test.toHex)).ByAes(aesCipher(CBC, test.padding, test.key, test.iv))
		assert.Nil(t, e.Error)
		assert.Equal(t, []byte(test.input), e.ToBytes(), "Hex test index is "+strconv.Itoa(index))
	}

	for index, test := range aesCbcTest {
		e := Decrypt.FromBase32Bytes([]byte(test.toBase32)).ByAes(aesCipher(CBC, test.padding, test.key, test.iv))
		assert.Nil(t, e.Error)
		assert.Equal(t, []byte(test.input), e.ToBytes(), "Base32 test index is "+strconv.Itoa(index))
	}

	for index, test := range aesCbcTest {
		e := Decrypt.FromBase64Bytes([]byte(test.toBase64)).ByAes(aesCipher(CBC, test.padding, test.key, test.iv))
		assert.Nil(t, e.Error)
		assert.Equal(t, []byte(test.input), e.ToBytes(), "Base64 test index is "+strconv.Itoa(index))
	}
}
