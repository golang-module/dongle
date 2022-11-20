package dongle

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

var tripleDesEcbTest = []struct {
	padding  string
	input    string
	key      string
	iv       string
	toHex    string
	toBase32 string
	toBase64 string
}{
	{"", "", "0123456789abcdefghijklmn", "12345678", "", "", ""},
	{No, "hello world, go!", "0123456789abcdefghijklmn", "12345678", "b8097975c76319c623be7c7aa6e0f3fc", "XAEXS5OHMMM4MI56PR5KNYHT7Q======", "uAl5dcdjGcYjvnx6puDz/A=="},
	{Zero, "hello world", "0123456789abcdefghijklmn", "12345678", "b8097975c76319c61971a986e579cdf9", "XAEXS5OHMMM4MGLRVGDOK6ON7E======", "uAl5dcdjGcYZcamG5XnN+Q=="},
	{PKCS5, "hello world", "0123456789abcdefghijklmn", "12345678", "b8097975c76319c6172687e0d90fd4d1", "XAEXS5OHMMM4MFZGQ7QNSD6U2E======", "uAl5dcdjGcYXJofg2Q/U0Q=="},
	{PKCS7, "hello world", "0123456789abcdefghijklmn", "12345678", "b8097975c76319c6172687e0d90fd4d1", "XAEXS5OHMMM4MFZGQ7QNSD6U2E======", "uAl5dcdjGcYXJofg2Q/U0Q=="},
}

func TestEncrypt_By3Des_ECB_ToString(t *testing.T) {
	for index, test := range tripleDesEcbTest {
		e := Encrypt.FromString(test.input).By3Des(tripleDesCipher(ECB, test.padding, test.key, test.iv))
		assert.Nil(t, e.Error)

		assert.Equal(t, test.toHex, e.ToHexString(), "Hex test index is "+strconv.Itoa(index))
		assert.Equal(t, test.toBase32, e.ToBase32String(), "Base32 test index is "+strconv.Itoa(index))
		assert.Equal(t, test.toBase64, e.ToBase64String(), "Base64 test index is "+strconv.Itoa(index))
	}
}

func TestDecrypt_By3Des_ECB_ToString(t *testing.T) {
	for index, test := range tripleDesEcbTest {
		e := Decrypt.FromHexString(test.toHex).By3Des(tripleDesCipher(ECB, test.padding, test.key, test.iv))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.input, e.ToString(), "Hex test index is "+strconv.Itoa(index))
	}

	for index, test := range tripleDesEcbTest {
		e := Decrypt.FromBase32String(test.toBase32).By3Des(tripleDesCipher(ECB, test.padding, test.key, test.iv))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.input, e.ToString(), "Base32 test index is "+strconv.Itoa(index))
	}

	for index, test := range tripleDesEcbTest {
		e := Decrypt.FromBase64String(test.toBase64).By3Des(tripleDesCipher(ECB, test.padding, test.key, test.iv))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.input, e.ToString(), "Base64 test index is "+strconv.Itoa(index))
	}
}

func TestEncrypt_By3Des_ECB_ToBytes(t *testing.T) {
	for index, test := range tripleDesEcbTest {
		e := Encrypt.FromString(test.input).By3Des(tripleDesCipher(ECB, test.padding, []byte(test.key), []byte(test.iv)))
		assert.Nil(t, e.Error)

		assert.Equal(t, []byte(test.toHex), e.ToHexBytes(), "Hex test index is "+strconv.Itoa(index))
		assert.Equal(t, []byte(test.toBase32), e.ToBase32Bytes(), "Base32 test index is "+strconv.Itoa(index))
		assert.Equal(t, []byte(test.toBase64), e.ToBase64Bytes(), "Base64 test index is "+strconv.Itoa(index))
	}
}

func TestDecrypt_By3Des_ECB_ToBytes(t *testing.T) {
	for index, test := range tripleDesEcbTest {
		e := Decrypt.FromHexBytes([]byte(test.toHex)).By3Des(tripleDesCipher(ECB, test.padding, []byte(test.key), []byte(test.iv)))
		assert.Nil(t, e.Error)
		assert.Equal(t, []byte(test.input), e.ToBytes(), "Hex test index is "+strconv.Itoa(index))
	}

	for index, test := range tripleDesEcbTest {
		e := Decrypt.FromBase32Bytes([]byte(test.toBase32)).By3Des(tripleDesCipher(ECB, test.padding, []byte(test.key), []byte(test.iv)))
		assert.Nil(t, e.Error)
		assert.Equal(t, []byte(test.input), e.ToBytes(), "Base32 test index is "+strconv.Itoa(index))
	}

	for index, test := range tripleDesEcbTest {
		e := Decrypt.FromBase64Bytes([]byte(test.toBase64)).By3Des(tripleDesCipher(ECB, test.padding, []byte(test.key), []byte(test.iv)))
		assert.Nil(t, e.Error)
		assert.Equal(t, []byte(test.input), e.ToBytes(), "Base64 test index is "+strconv.Itoa(index))
	}
}
