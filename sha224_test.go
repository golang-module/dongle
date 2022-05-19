package dongle

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	sha224Input          = "hello world"
	sha224HexExpected    = "2f05477fc24bb4faefd86517156dafdecec45b8ad3cf2522a563582b"
	sha224Base32Expected = "F4CUO76CJO2PV36YMULRK3NP33HMIW4K2PHSKIVFMNMCW==="
	sha224Base64Expected = "LwVHf8JLtPrv2GUXFW2v3s7EW4rTzyUipWNYKw=="
)

func TestEncrypt_BySha224_FromStringToString(t *testing.T) {
	assert := assert.New(t)

	hexTests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{sha224Input, sha224HexExpected},
	}

	for index, test := range hexTests {
		e := Encrypt.FromString(test.input).BySha224()
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToHexString(), "Current test index is "+strconv.Itoa(index))
	}

	base32Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{sha224Input, sha224Base32Expected},
	}

	for index, test := range base32Tests {
		e := Encrypt.FromString(test.input).BySha224()
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToBase32String(), "Current test index is "+strconv.Itoa(index))
	}

	base64Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{sha224Input, sha224Base64Expected},
	}

	for index, test := range base64Tests {
		e := Encrypt.FromString(test.input).BySha224()
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToBase64String(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestEncrypt_BySha224_FromBytesToBytes(t *testing.T) {
	assert := assert.New(t)

	hexTests := []struct {
		input    []byte // 输入值
		expected []byte // 期望值
	}{
		{[]byte(""), []byte("")},
		{[]byte(sha224Input), []byte(sha224HexExpected)},
	}

	for index, test := range hexTests {
		e := Encrypt.FromBytes(test.input).BySha224()
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToHexBytes(), "Current test index is "+strconv.Itoa(index))
	}

	base32Tests := []struct {
		input    []byte // 输入值
		expected []byte // 期望值
	}{
		{[]byte(""), []byte("")},
		{[]byte(sha224Input), []byte(sha224Base32Expected)},
	}

	for index, test := range base32Tests {
		e := Encrypt.FromBytes(test.input).BySha224()
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToBase32Bytes(), "Current test index is "+strconv.Itoa(index))
	}

	base64Tests := []struct {
		input    []byte // 输入值
		expected []byte // 期望值
	}{
		{[]byte(""), []byte("")},
		{[]byte(sha224Input), []byte(sha224Base64Expected)},
	}

	for index, test := range base64Tests {
		e := Encrypt.FromBytes(test.input).BySha224()
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToBase64Bytes(), "Current test index is "+strconv.Itoa(index))
	}

}
