package dongle

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	rc4Input, rc4Key  = "hello world", "dongle"
	rc4HexExpected    = "eba154b4cb5a9038dbbf9d"
	rc4Base32Expected = "5OQVJNGLLKIDRW57TU======"
	rc4Base64Expected = "66FUtMtakDjbv50="
)

func TestEncrypt_ByRC4_FromStringToString(t *testing.T) {
	assert := assert.New(t)

	hexTests := []struct {
		input1   string // 输入值1
		input2   string // 输入值2
		expected string // 期望值
	}{
		{"", "", ""},
		{"", rc4Key, ""},
		{rc4Input, rc4Key, rc4HexExpected},
	}

	for index, test := range hexTests {
		e := Encrypt.FromString(test.input1).ByRc4(test.input2)
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToHexString(), "Current test index is "+strconv.Itoa(index))
	}

	base32Tests := []struct {
		input1   string // 输入值1
		input2   string // 输入值2
		expected string // 期望值
	}{
		{"", "", ""},
		{"", rc4Key, ""},
		{rc4Input, rc4Key, rc4Base32Expected},
	}

	for index, test := range base32Tests {
		e := Encrypt.FromString(test.input1).ByRc4(test.input2)
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToBase32String(), "Current test index is "+strconv.Itoa(index))
	}

	base64Tests := []struct {
		input1   string // 输入值1
		input2   string // 输入值2
		expected string // 期望值
	}{
		{"", "", ""},
		{"", rc4Key, ""},
		{rc4Input, rc4Key, rc4Base64Expected},
	}

	for index, test := range base64Tests {
		e := Encrypt.FromString(test.input1).ByRc4(test.input2)
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToBase64String(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestEncrypt_ByRC4_FromBytesToBytes(t *testing.T) {
	assert := assert.New(t)

	hexTests := []struct {
		input1   []byte // 输入值1
		input2   []byte // 输入值2
		expected []byte // 期望值
	}{
		{[]byte(""), []byte(""), []byte("")},
		{[]byte(""), []byte(rc4Key), []byte("")},
		{[]byte(rc4Input), []byte(rc4Key), []byte(rc4HexExpected)},
	}

	for index, test := range hexTests {
		e := Encrypt.FromBytes(test.input1).ByRc4(test.input2)
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToHexBytes(), "Current test index is "+strconv.Itoa(index))
	}

	base32Tests := []struct {
		input1   []byte // 输入值1
		input2   []byte // 输入值2
		expected []byte // 期望值
	}{
		{[]byte(""), []byte(""), []byte("")},
		{[]byte(""), []byte(rc4Key), []byte("")},
		{[]byte(rc4Input), []byte(rc4Key), []byte(rc4Base32Expected)},
	}

	for index, test := range base32Tests {
		e := Encrypt.FromBytes(test.input1).ByRc4(test.input2)
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToBase32Bytes(), "Current test index is "+strconv.Itoa(index))
	}

	base64Tests := []struct {
		input1   []byte // 输入值1
		input2   []byte // 输入值2
		expected []byte // 期望值
	}{
		{[]byte(""), []byte(""), []byte("")},
		{[]byte(""), []byte(rc4Key), []byte("")},
		{[]byte(rc4Input), []byte(rc4Key), []byte(rc4Base64Expected)},
	}

	for index, test := range base64Tests {
		e := Encrypt.FromBytes(test.input1).ByRc4(test.input2)
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToBase64Bytes(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestError_ByRc4_FromString(t *testing.T) {
	e := Encrypt.FromString("hello world").ByRc4("")
	expected := fmt.Errorf("invalid key size %d, key at least 1 byte and at most 256 bytes", 0)
	assert.Equal(t, expected, e.Error, "Should catch an exception")
}
