package dongle

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

var tripleDesCbcTest = []struct {
	padding  string
	input    string
	key      string
	iv       string
	toHex    string
	toBase32 string
	toBase64 string
}{
	{"", "", "0123456789abcdefghijklmn", "12345678", "", "", ""},
	{No, "hello world, go!", "0123456789abcdefghijklmn", "12345678", "7e9194cc827a325d49111eaa503110fe", "P2IZJTECPIZF2SIRD2VFAMIQ7Y======", "fpGUzIJ6Ml1JER6qUDEQ/g=="},
	{Zero, "hello world", "0123456789abcdefghijklmn", "12345678", "7e9194cc827a325ddaee992a89c5cd8d", "P2IZJTECPIZF3WXOTEVITRONRU======", "fpGUzIJ6Ml3a7pkqicXNjQ=="},
	{PKCS5, "hello world", "0123456789abcdefghijklmn", "12345678", "7e9194cc827a325db9c765859716cc97", "P2IZJTECPIZF3OOHMWCZOFWMS4======", "fpGUzIJ6Ml25x2WFlxbMlw=="},
	{PKCS7, "hello world", "0123456789abcdefghijklmn", "12345678", "7e9194cc827a325db9c765859716cc97", "P2IZJTECPIZF3OOHMWCZOFWMS4======", "fpGUzIJ6Ml25x2WFlxbMlw=="},
}

func TestEncrypt_By3Des_CBC_ToString(t *testing.T) {
	for index, test := range tripleDesCbcTest {
		e := Encrypt.FromString(test.input).By3Des(tripleDesCipher(CBC, test.padding, test.key, test.iv))
		assert.Nil(t, e.Error)

		assert.Equal(t, test.toHex, e.ToHexString(), "Hex test index is "+strconv.Itoa(index))
		assert.Equal(t, test.toBase32, e.ToBase32String(), "Base32 test index is "+strconv.Itoa(index))
		assert.Equal(t, test.toBase64, e.ToBase64String(), "Base64 test index is "+strconv.Itoa(index))
	}
}

func TestDecrypt_By3Des_CBC_ToString(t *testing.T) {
	for index, test := range tripleDesCbcTest {
		e := Decrypt.FromHexString(test.toHex).By3Des(tripleDesCipher(CBC, test.padding, test.key, test.iv))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.input, e.ToString(), "Hex test index is "+strconv.Itoa(index))
	}

	for index, test := range tripleDesCbcTest {
		e := Decrypt.FromBase32String(test.toBase32).By3Des(tripleDesCipher(CBC, test.padding, test.key, test.iv))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.input, e.ToString(), "Base32 test index is "+strconv.Itoa(index))
	}

	for index, test := range tripleDesCbcTest {
		e := Decrypt.FromBase64String(test.toBase64).By3Des(tripleDesCipher(CBC, test.padding, test.key, test.iv))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.input, e.ToString(), "Base64 test index is "+strconv.Itoa(index))
	}
}

func TestEncrypt_By3Des_CBC_ToBytes(t *testing.T) {
	for index, test := range tripleDesCbcTest {
		e := Encrypt.FromString(test.input).By3Des(tripleDesCipher(CBC, test.padding, []byte(test.key), []byte(test.iv)))
		assert.Nil(t, e.Error)

		assert.Equal(t, []byte(test.toHex), e.ToHexBytes(), "Hex test index is "+strconv.Itoa(index))
		assert.Equal(t, []byte(test.toBase32), e.ToBase32Bytes(), "Base32 test index is "+strconv.Itoa(index))
		assert.Equal(t, []byte(test.toBase64), e.ToBase64Bytes(), "Base64 test index is "+strconv.Itoa(index))
	}
}

func TestDecrypt_By3Des_CBC_ToBytes(t *testing.T) {
	for index, test := range tripleDesCbcTest {
		e := Decrypt.FromHexBytes([]byte(test.toHex)).By3Des(tripleDesCipher(CBC, test.padding, []byte(test.key), []byte(test.iv)))
		assert.Nil(t, e.Error)
		assert.Equal(t, []byte(test.input), e.ToBytes(), "Hex test index is "+strconv.Itoa(index))
	}

	for index, test := range tripleDesCbcTest {
		e := Decrypt.FromBase32Bytes([]byte(test.toBase32)).By3Des(tripleDesCipher(CBC, test.padding, []byte(test.key), []byte(test.iv)))
		assert.Nil(t, e.Error)
		assert.Equal(t, []byte(test.input), e.ToBytes(), "Base32 test index is "+strconv.Itoa(index))
	}

	for index, test := range tripleDesCbcTest {
		e := Decrypt.FromBase64Bytes([]byte(test.toBase64)).By3Des(tripleDesCipher(CBC, test.padding, []byte(test.key), []byte(test.iv)))
		assert.Nil(t, e.Error)
		assert.Equal(t, []byte(test.input), e.ToBytes(), "Base64 test index is "+strconv.Itoa(index))
	}
}
