package dongle

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	md4Input          = "hello world"
	md4HexExpected    = "aa010fbc1d14c795d86ef98c95479d17"
	md4Base32Expected = "VIAQ7PA5CTDZLWDO7GGJKR45C4======"
	md4Base64Expected = "qgEPvB0Ux5XYbvmMlUedFw=="
)

func TestEncrypt_ByMd4_FromStringToString(t *testing.T) {
	assert := assert.New(t)

	hexTests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{md4Input, md4HexExpected},
	}

	for index, test := range hexTests {
		e := Encrypt.FromString(test.input).ByMd4()
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToHexString(), "Current test index is "+strconv.Itoa(index))
	}

	base32Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{md4Input, md4Base32Expected},
	}

	for index, test := range base32Tests {
		e := Encrypt.FromString(test.input).ByMd4()
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToBase32String(), "Current test index is "+strconv.Itoa(index))
	}

	base64Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{md4Input, md4Base64Expected},
	}

	for index, test := range base64Tests {
		e := Encrypt.FromString(test.input).ByMd4()
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToBase64String(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestEncrypt_ByMd4_FromBytesToBytes(t *testing.T) {
	assert := assert.New(t)

	hexTests := []struct {
		input    []byte // 输入值
		expected []byte // 期望值
	}{
		{[]byte(""), []byte("")},
		{[]byte(md4Input), []byte(md4HexExpected)},
	}

	for index, test := range hexTests {
		e := Encrypt.FromBytes(test.input).ByMd4()
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToHexBytes(), "Current test index is "+strconv.Itoa(index))
	}

	base32Tests := []struct {
		input    []byte // 输入值
		expected []byte // 期望值
	}{
		{[]byte(""), []byte("")},
		{[]byte(md4Input), []byte(md4Base32Expected)},
	}

	for index, test := range base32Tests {
		e := Encrypt.FromBytes(test.input).ByMd4()
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToBase32Bytes(), "Current test index is "+strconv.Itoa(index))
	}

	base64Tests := []struct {
		input    []byte // 输入值
		expected []byte // 期望值
	}{
		{[]byte(""), []byte("")},
		{[]byte(md4Input), []byte(md4Base64Expected)},
	}

	for index, test := range base64Tests {
		e := Encrypt.FromBytes(test.input).ByMd4()
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToBase64Bytes(), "Current test index is "+strconv.Itoa(index))
	}

}
