package dongle

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	hmacMd4Input, hmacMd4Key = "hello world", "dongle"
	hmacMd4HexExpected       = "7a9df5247cbf76a8bc17c9c4f5a75b6b"
	hmacMd4Base32Expected    = "PKO7KJD4X53KRPAXZHCPLJ23NM======"
	hmacMd4Base64Expected    = "ep31JHy/dqi8F8nE9adbaw=="
)

func TestEncrypt_ByHmacMd4_FromStringToString(t *testing.T) {
	assert := assert.New(t)

	hexTests := []struct {
		input1   string // 输入值1
		input2   string // 输入值2
		expected string // 期望值
	}{
		{"", "", ""},
		{"", hmacMd4Key, ""},
		{hmacMd4Input, hmacMd4Key, hmacMd4HexExpected},
	}

	for index, test := range hexTests {
		e := Encrypt.FromString(test.input1).ByHmacMd4(test.input2)
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToHexString(), "Current test index is "+strconv.Itoa(index))
	}

	base32Tests := []struct {
		input1   string // 输入值1
		input2   string // 输入值2
		expected string // 期望值
	}{
		{"", "", ""},
		{"", hmacMd4Key, ""},
		{hmacMd4Input, hmacMd4Key, hmacMd4Base32Expected},
	}

	for index, test := range base32Tests {
		e := Encrypt.FromString(test.input1).ByHmacMd4(test.input2)
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToBase32String(), "Current test index is "+strconv.Itoa(index))
	}

	base64Tests := []struct {
		input1   string // 输入值1
		input2   string // 输入值2
		expected string // 期望值
	}{
		{"", "", ""},
		{"", hmacMd4Key, ""},
		{hmacMd4Input, hmacMd4Key, hmacMd4Base64Expected},
	}

	for index, test := range base64Tests {
		e := Encrypt.FromString(test.input1).ByHmacMd4(test.input2)
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToBase64String(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestEncrypt_ByHmacMd4_FromBytesToBytes(t *testing.T) {
	assert := assert.New(t)

	hexTests := []struct {
		input1   []byte // 输入值1
		input2   []byte // 输入值2
		expected []byte // 期望值
	}{
		{[]byte(""), []byte(""), []byte("")},
		{[]byte(""), []byte(hmacMd4Key), []byte("")},
		{[]byte(hmacMd4Input), []byte(hmacMd4Key), []byte(hmacMd4HexExpected)},
	}

	for index, test := range hexTests {
		e := Encrypt.FromBytes(test.input1).ByHmacMd4(test.input2)
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToHexBytes(), "Current test index is "+strconv.Itoa(index))
	}

	base32Tests := []struct {
		input1   []byte // 输入值1
		input2   []byte // 输入值2
		expected []byte // 期望值
	}{
		{[]byte(""), []byte(""), []byte("")},
		{[]byte(""), []byte(hmacMd4Key), []byte("")},
		{[]byte(hmacMd4Input), []byte(hmacMd4Key), []byte(hmacMd4Base32Expected)},
	}

	for index, test := range base32Tests {
		e := Encrypt.FromBytes(test.input1).ByHmacMd4(test.input2)
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToBase32Bytes(), "Current test index is "+strconv.Itoa(index))
	}

	base64Tests := []struct {
		input1   []byte // 输入值1
		input2   []byte // 输入值2
		expected []byte // 期望值
	}{
		{[]byte(""), []byte(""), []byte("")},
		{[]byte(""), []byte(hmacMd4Key), []byte("")},
		{[]byte(hmacMd4Input), []byte(hmacMd4Key), []byte(hmacMd4Base64Expected)},
	}

	for index, test := range base64Tests {
		e := Encrypt.FromBytes(test.input1).ByHmacMd4(test.input2)
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToBase64Bytes(), "Current test index is "+strconv.Itoa(index))
	}
}
