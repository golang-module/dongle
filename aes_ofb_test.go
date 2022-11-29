package dongle

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

var aesOfbTest = []struct {
	padding  cipherPadding
	input    string
	key      string
	iv       string
	toHex    string
	toBase32 string
	toBase64 string
}{
	{"", "", "0123456789abcdef", "0123456789abcdef", "", "", ""},
	{No, "hello world, go!", "0123456789abcdef", "0123456789abcdef", "1a1712e471fc8a6e72cb7c44596eda44", "DILRFZDR7SFG44WLPRCFS3W2IQ======", "GhcS5HH8im5yy3xEWW7aRA=="},
	{Zero, "hello world", "0123456789abcdef", "0123456789abcdef", "1a1712e471fc8a6e72cb7c687909b565", "DILRFZDR7SFG44WLPRUHSCNVMU======", "GhcS5HH8im5yy3xoeQm1ZQ=="},
	{PKCS5, "hello world", "0123456789abcdef", "0123456789abcdef", "1a1712e471fc8a6e72cb7c6d7c0cb060", "DILRFZDR7SFG44WLPRWXYDFQMA======", "GhcS5HH8im5yy3xtfAywYA=="},
	{PKCS7, "hello world", "0123456789abcdef", "0123456789abcdef", "1a1712e471fc8a6e72cb7c6d7c0cb060", "DILRFZDR7SFG44WLPRWXYDFQMA======", "GhcS5HH8im5yy3xtfAywYA=="},
}

func TestEncrypt_ByAes_OFB_ToString(t *testing.T) {
	for index, test := range aesOfbTest {
		e := Encrypt.FromString(test.input).ByAes(aesCipher(OFB, test.padding, test.key, test.iv))
		assert.Nil(t, e.Error)

		assert.Equal(t, test.toHex, e.ToHexString(), "Hex test index is "+strconv.Itoa(index))
		assert.Equal(t, test.toBase32, e.ToBase32String(), "Base32 test index is "+strconv.Itoa(index))
		assert.Equal(t, test.toBase64, e.ToBase64String(), "Base64 test index is "+strconv.Itoa(index))
	}
}

func TestDecrypt_ByAes_OFB_ToString(t *testing.T) {
	for index, test := range aesOfbTest {
		e := Decrypt.FromHexString(test.toHex).ByAes(aesCipher(OFB, test.padding, test.key, test.iv))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.input, e.ToString(), "Hex test index is "+strconv.Itoa(index))
	}

	for index, test := range aesOfbTest {
		e := Decrypt.FromBase32String(test.toBase32).ByAes(aesCipher(OFB, test.padding, test.key, test.iv))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.input, e.ToString(), "Base32 test index is "+strconv.Itoa(index))
	}

	for index, test := range aesOfbTest {
		e := Decrypt.FromBase64String(test.toBase64).ByAes(aesCipher(OFB, test.padding, test.key, test.iv))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.input, e.ToString(), "Base64 test index is "+strconv.Itoa(index))
	}
}

func TestEncrypt_ByAes_OFB_ToBytes(t *testing.T) {
	for index, test := range aesOfbTest {
		e := Encrypt.FromString(test.input).ByAes(aesCipher(OFB, test.padding, test.key, test.iv))
		assert.Nil(t, e.Error)

		assert.Equal(t, []byte(test.toHex), e.ToHexBytes(), "Hex test index is "+strconv.Itoa(index))
		assert.Equal(t, []byte(test.toBase32), e.ToBase32Bytes(), "Base32 test index is "+strconv.Itoa(index))
		assert.Equal(t, []byte(test.toBase64), e.ToBase64Bytes(), "Base64 test index is "+strconv.Itoa(index))
	}
}

func TestDecrypt_ByAes_OFB_ToBytes(t *testing.T) {
	for index, test := range aesOfbTest {
		e := Decrypt.FromHexBytes([]byte(test.toHex)).ByAes(aesCipher(OFB, test.padding, test.key, test.iv))
		assert.Nil(t, e.Error)
		assert.Equal(t, []byte(test.input), e.ToBytes(), "Hex test index is "+strconv.Itoa(index))
	}

	for index, test := range aesOfbTest {
		e := Decrypt.FromBase32Bytes([]byte(test.toBase32)).ByAes(aesCipher(OFB, test.padding, test.key, test.iv))
		assert.Nil(t, e.Error)
		assert.Equal(t, []byte(test.input), e.ToBytes(), "Base32 test index is "+strconv.Itoa(index))
	}

	for index, test := range aesOfbTest {
		e := Decrypt.FromBase64Bytes([]byte(test.toBase64)).ByAes(aesCipher(OFB, test.padding, test.key, test.iv))
		assert.Nil(t, e.Error)
		assert.Equal(t, []byte(test.input), e.ToBytes(), "Base64 test index is "+strconv.Itoa(index))
	}
}
