package dongle

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

var tripleDesCfbTest = []struct {
	padding  string
	input    string
	key      string
	iv       string
	toHex    string
	toBase32 string
	toBase64 string
}{
	{"", "", "0123456789abcdefghijklmn", "12345678", "", "", ""},
	{No, "hello world, go!", "0123456789abcdefghijklmn", "12345678", "f52dc896ceff0ecc9393c19d4e1e7591", "6UW4RFWO74HMZE4TYGOU4HTVSE======", "9S3Ils7/DsyTk8GdTh51kQ=="},
	{Zero, "hello world", "0123456789abcdefghijklmn", "12345678", "f52dc896ceff0ecc9393c1b16e791ab0", "6UW4RFWO74HMZE4TYGYW46I2WA======", "9S3Ils7/DsyTk8GxbnkasA=="},
	{PKCS5, "hello world", "0123456789abcdefghijklmn", "12345678", "f52dc896ceff0ecc9393c1b46b7c1fb5", "6UW4RFWO74HMZE4TYG2GW7A7WU======", "9S3Ils7/DsyTk8G0a3wftQ=="},
	{PKCS7, "hello world", "0123456789abcdefghijklmn", "12345678", "f52dc896ceff0ecc9393c1b46b7c1fb5", "6UW4RFWO74HMZE4TYG2GW7A7WU======", "9S3Ils7/DsyTk8G0a3wftQ=="},
}

func TestEncrypt_By3Des_CFB_ToString(t *testing.T) {
	for index, test := range tripleDesCfbTest {
		e := Encrypt.FromString(test.input).By3Des(tripleDesCipher(CFB, test.padding, test.key, test.iv))
		assert.Nil(t, e.Error)

		assert.Equal(t, test.toHex, e.ToHexString(), "Hex test index is "+strconv.Itoa(index))
		assert.Equal(t, test.toBase32, e.ToBase32String(), "Base32 test index is "+strconv.Itoa(index))
		assert.Equal(t, test.toBase64, e.ToBase64String(), "Base64 test index is "+strconv.Itoa(index))
	}
}

func TestDecrypt_By3Des_CFB_ToString(t *testing.T) {
	for index, test := range tripleDesCfbTest {
		e := Decrypt.FromHexString(test.toHex).By3Des(tripleDesCipher(CFB, test.padding, test.key, test.iv))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.input, e.ToString(), "Hex test index is "+strconv.Itoa(index))
	}

	for index, test := range tripleDesCfbTest {
		e := Decrypt.FromBase32String(test.toBase32).By3Des(tripleDesCipher(CFB, test.padding, test.key, test.iv))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.input, e.ToString(), "Base32 test index is "+strconv.Itoa(index))
	}

	for index, test := range tripleDesCfbTest {
		e := Decrypt.FromBase64String(test.toBase64).By3Des(tripleDesCipher(CFB, test.padding, test.key, test.iv))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.input, e.ToString(), "Base64 test index is "+strconv.Itoa(index))
	}
}

func TestEncrypt_By3Des_CFB_ToBytes(t *testing.T) {
	for index, test := range tripleDesCfbTest {
		e := Encrypt.FromString(test.input).By3Des(tripleDesCipher(CFB, test.padding, []byte(test.key), []byte(test.iv)))
		assert.Nil(t, e.Error)

		assert.Equal(t, []byte(test.toHex), e.ToHexBytes(), "Hex test index is "+strconv.Itoa(index))
		assert.Equal(t, []byte(test.toBase32), e.ToBase32Bytes(), "Base32 test index is "+strconv.Itoa(index))
		assert.Equal(t, []byte(test.toBase64), e.ToBase64Bytes(), "Base64 test index is "+strconv.Itoa(index))
	}
}

func TestDecrypt_By3Des_CFB_ToBytes(t *testing.T) {
	for index, test := range tripleDesCfbTest {
		e := Decrypt.FromHexBytes([]byte(test.toHex)).By3Des(tripleDesCipher(CFB, test.padding, []byte(test.key), []byte(test.iv)))
		assert.Nil(t, e.Error)
		assert.Equal(t, []byte(test.input), e.ToBytes(), "Hex test index is "+strconv.Itoa(index))
	}

	for index, test := range tripleDesCfbTest {
		e := Decrypt.FromBase32Bytes([]byte(test.toBase32)).By3Des(tripleDesCipher(CFB, test.padding, []byte(test.key), []byte(test.iv)))
		assert.Nil(t, e.Error)
		assert.Equal(t, []byte(test.input), e.ToBytes(), "Base32 test index is "+strconv.Itoa(index))
	}

	for index, test := range tripleDesCfbTest {
		e := Decrypt.FromBase64Bytes([]byte(test.toBase64)).By3Des(tripleDesCipher(CFB, test.padding, []byte(test.key), []byte(test.iv)))
		assert.Nil(t, e.Error)
		assert.Equal(t, []byte(test.input), e.ToBytes(), "Base64 test index is "+strconv.Itoa(index))
	}
}
