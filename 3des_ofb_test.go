package dongle

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

var tripleDesOfbTest = []struct {
	padding  string
	input    string
	key      string
	iv       string
	toHex    string
	toBase32 string
	toBase64 string
}{
	{"", "", "0123456789abcdefghijklmn", "12345678", "", "", ""},
	{No, "hello world, go!", "0123456789abcdefghijklmn", "12345678", "f52dc896ceff0eccf869cb59735c3766", "6UW4RFWO74HMZ6DJZNMXGXBXMY======", "9S3Ils7/Dsz4actZc1w3Zg=="},
	{Zero, "hello world", "0123456789abcdefghijklmn", "12345678", "f52dc896ceff0eccf869cb75533b5847", "6UW4RFWO74HMZ6DJZN2VGO2YI4======", "9S3Ils7/Dsz4act1UztYRw=="},
	{PKCS5, "hello world", "0123456789abcdefghijklmn", "12345678", "f52dc896ceff0eccf869cb70563e5d42", "6UW4RFWO74HMZ6DJZNYFMPS5II======", "9S3Ils7/Dsz4actwVj5dQg=="},
	{PKCS7, "hello world", "0123456789abcdefghijklmn", "12345678", "f52dc896ceff0eccf869cb70563e5d42", "6UW4RFWO74HMZ6DJZNYFMPS5II======", "9S3Ils7/Dsz4actwVj5dQg=="},
}

func TestEncrypt_By3Des_OFB_ToString(t *testing.T) {
	for index, test := range tripleDesOfbTest {
		e := Encrypt.FromString(test.input).By3Des(tripleDesCipher(OFB, test.padding, test.key, test.iv))
		assert.Nil(t, e.Error)

		assert.Equal(t, test.toHex, e.ToHexString(), "Hex test index is "+strconv.Itoa(index))
		assert.Equal(t, test.toBase32, e.ToBase32String(), "Base32 test index is "+strconv.Itoa(index))
		assert.Equal(t, test.toBase64, e.ToBase64String(), "Base64 test index is "+strconv.Itoa(index))
	}
}

func TestDecrypt_By3Des_OFB_ToString(t *testing.T) {
	for index, test := range tripleDesOfbTest {
		e := Decrypt.FromHexString(test.toHex).By3Des(tripleDesCipher(OFB, test.padding, test.key, test.iv))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.input, e.ToString(), "Hex test index is "+strconv.Itoa(index))
	}

	for index, test := range tripleDesOfbTest {
		e := Decrypt.FromBase32String(test.toBase32).By3Des(tripleDesCipher(OFB, test.padding, test.key, test.iv))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.input, e.ToString(), "Base32 test index is "+strconv.Itoa(index))
	}

	for index, test := range tripleDesOfbTest {
		e := Decrypt.FromBase64String(test.toBase64).By3Des(tripleDesCipher(OFB, test.padding, test.key, test.iv))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.input, e.ToString(), "Base64 test index is "+strconv.Itoa(index))
	}
}

func TestEncrypt_By3Des_OFB_ToBytes(t *testing.T) {
	for index, test := range tripleDesOfbTest {
		e := Encrypt.FromString(test.input).By3Des(tripleDesCipher(OFB, test.padding, test.key, test.iv))
		assert.Nil(t, e.Error)

		assert.Equal(t, []byte(test.toHex), e.ToHexBytes(), "Hex test index is "+strconv.Itoa(index))
		assert.Equal(t, []byte(test.toBase32), e.ToBase32Bytes(), "Base32 test index is "+strconv.Itoa(index))
		assert.Equal(t, []byte(test.toBase64), e.ToBase64Bytes(), "Base64 test index is "+strconv.Itoa(index))
	}
}

func TestDecrypt_By3Des_OFB_ToBytes(t *testing.T) {
	for index, test := range tripleDesOfbTest {
		e := Decrypt.FromHexBytes([]byte(test.toHex)).By3Des(tripleDesCipher(OFB, test.padding, test.key, test.iv))
		assert.Nil(t, e.Error)
		assert.Equal(t, []byte(test.input), e.ToBytes(), "Hex test index is "+strconv.Itoa(index))
	}

	for index, test := range tripleDesOfbTest {
		e := Decrypt.FromBase32Bytes([]byte(test.toBase32)).By3Des(tripleDesCipher(OFB, test.padding, test.key, test.iv))
		assert.Nil(t, e.Error)
		assert.Equal(t, []byte(test.input), e.ToBytes(), "Base32 test index is "+strconv.Itoa(index))
	}

	for index, test := range tripleDesOfbTest {
		e := Decrypt.FromBase64Bytes([]byte(test.toBase64)).By3Des(tripleDesCipher(OFB, test.padding, test.key, test.iv))
		assert.Nil(t, e.Error)
		assert.Equal(t, []byte(test.input), e.ToBytes(), "Base64 test index is "+strconv.Itoa(index))
	}
}
