package dongle

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	hmacMd5Input, hmacMd5Key = "hello world", "dongle"
	hmacMd5HexExpected       = "4790626a275f776956386e5a3ea7b726"
	hmacMd5Base32Expected    = "I6IGE2RHL53WSVRYNZND5J5XEY======"
	hmacMd5Base64Expected    = "R5Biaidfd2lWOG5aPqe3Jg=="
)

func TestEncrypt_ByHmacMd5_FromStringToString(t *testing.T) {
	hexTests := []struct {
		input1   string // 输入值1
		input2   string // 输入值2
		expected string // 期望值
	}{
		{"", "", ""},
		{hmacMd5Input, hmacMd5Key, hmacMd5HexExpected},
	}

	for index, test := range hexTests {
		e := Encrypt.FromString(test.input1).ByHmacMd5(test.input2)
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToHexString(), "Current test index is "+strconv.Itoa(index))
	}

	base32Tests := []struct {
		input1   string // 输入值1
		input2   string // 输入值2
		expected string // 期望值
	}{
		{"", "", ""},
		{hmacMd5Input, hmacMd5Key, hmacMd5Base32Expected},
	}

	for index, test := range base32Tests {
		e := Encrypt.FromString(test.input1).ByHmacMd5(test.input2)
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToBase32String(), "Current test index is "+strconv.Itoa(index))
	}

	base64Tests := []struct {
		input1   string // 输入值1
		input2   string // 输入值2
		expected string // 期望值
	}{
		{"", "", ""},
		{hmacMd5Input, hmacMd5Key, hmacMd5Base64Expected},
	}

	for index, test := range base64Tests {
		e := Encrypt.FromString(test.input1).ByHmacMd5(test.input2)
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToBase64String(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestEncrypt_ByHmacMd5_FromBytesToBytes(t *testing.T) {
	hexTests := []struct {
		input1   []byte // 输入值1
		input2   []byte // 输入值2
		expected []byte // 期望值
	}{
		{[]byte(""), []byte(""), []byte("")},
		{[]byte(hmacMd5Input), []byte(hmacMd5Key), []byte(hmacMd5HexExpected)},
	}

	for index, test := range hexTests {
		e := Encrypt.FromBytes(test.input1).ByHmacMd5(test.input2)
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToHexBytes(), "Current test index is "+strconv.Itoa(index))
	}

	base32Tests := []struct {
		input1   []byte // 输入值1
		input2   []byte // 输入值2
		expected []byte // 期望值
	}{
		{[]byte(""), []byte(""), []byte("")},
		{[]byte(hmacMd5Input), []byte(hmacMd5Key), []byte(hmacMd5Base32Expected)},
	}

	for index, test := range base32Tests {
		e := Encrypt.FromBytes(test.input1).ByHmacMd5(test.input2)
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToBase32Bytes(), "Current test index is "+strconv.Itoa(index))
	}

	base64Tests := []struct {
		input1   []byte // 输入值1
		input2   []byte // 输入值2
		expected []byte // 期望值
	}{
		{[]byte(""), []byte(""), []byte("")},
		{[]byte(hmacMd5Input), []byte(hmacMd5Key), []byte(hmacMd5Base64Expected)},
	}

	for index, test := range base64Tests {
		e := Encrypt.FromBytes(test.input1).ByHmacMd5(test.input2)
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToBase64Bytes(), "Current test index is "+strconv.Itoa(index))
	}
}
