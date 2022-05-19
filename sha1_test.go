package dongle

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	sha1Input          = "hello world"
	sha1HexExpected    = "2aae6c35c94fcfb415dbe95f408b9ce91ee846ed"
	sha1Base32Expected = "FKXGYNOJJ7H3IFO35FPUBC445EPOQRXN"
	sha1Base64Expected = "Kq5sNclPz7QV2+lfQIuc6R7oRu0="
)

func TestEncrypt_BySha1_FromStringToString(t *testing.T) {
	assert := assert.New(t)

	hexTests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{sha1Input, sha1HexExpected},
	}

	for index, test := range hexTests {
		e := Encrypt.FromString(test.input).BySha1()
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToHexString(), "Current test index is "+strconv.Itoa(index))
	}

	base32Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{sha1Input, sha1Base32Expected},
	}

	for index, test := range base32Tests {
		e := Encrypt.FromString(test.input).BySha1()
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToBase32String(), "Current test index is "+strconv.Itoa(index))
	}

	base64Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{sha1Input, sha1Base64Expected},
	}

	for index, test := range base64Tests {
		e := Encrypt.FromString(test.input).BySha1()
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToBase64String(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestEncrypt_BySha1_FromBytesToBytes(t *testing.T) {
	assert := assert.New(t)

	hexTests := []struct {
		input    []byte // 输入值
		expected []byte // 期望值
	}{
		{[]byte(""), []byte("")},
		{[]byte(sha1Input), []byte(sha1HexExpected)},
	}

	for index, test := range hexTests {
		e := Encrypt.FromBytes(test.input).BySha1()
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToHexBytes(), "Current test index is "+strconv.Itoa(index))
	}

	base32Tests := []struct {
		input    []byte // 输入值
		expected []byte // 期望值
	}{
		{[]byte(""), []byte("")},
		{[]byte(sha1Input), []byte(sha1Base32Expected)},
	}

	for index, test := range base32Tests {
		e := Encrypt.FromBytes(test.input).BySha1()
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToBase32Bytes(), "Current test index is "+strconv.Itoa(index))
	}

	base64Tests := []struct {
		input    []byte // 输入值
		expected []byte // 期望值
	}{
		{[]byte(""), []byte("")},
		{[]byte(sha1Input), []byte(sha1Base64Expected)},
	}

	for index, test := range base64Tests {
		e := Encrypt.FromBytes(test.input).BySha1()
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToBase64Bytes(), "Current test index is "+strconv.Itoa(index))
	}

}
