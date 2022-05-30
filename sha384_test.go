package dongle

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	sha384Input          = "hello world"
	sha384HexExpected    = "fdbd8e75a67f29f701a4e040385e2e23986303ea10239211af907fcbb83578b3e417cb71ce646efd0819dd8c088de1bd"
	sha384Base32Expected = "7W6Y45NGP4U7OANE4BADQXROEOMGGA7KCARZEENPSB74XOBVPCZ6IF6LOHHGI3X5BAM53DAIRXQ32==="
	sha384Base64Expected = "/b2OdaZ/KfcBpOBAOF4uI5hjA+oQI5IRr5B/y7g1eLPkF8txzmRu/QgZ3YwIjeG9"
)

func TestEncrypt_BySha384_FromStringToString(t *testing.T) {
	hexTests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{sha384Input, sha384HexExpected},
	}

	for index, test := range hexTests {
		e := Encrypt.FromString(test.input).BySha384()
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToHexString(), "Current test index is "+strconv.Itoa(index))
	}

	base32Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{sha384Input, sha384Base32Expected},
	}

	for index, test := range base32Tests {
		e := Encrypt.FromString(test.input).BySha384()
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToBase32String(), "Current test index is "+strconv.Itoa(index))
	}

	base64Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{sha384Input, sha384Base64Expected},
	}

	for index, test := range base64Tests {
		e := Encrypt.FromString(test.input).BySha384()
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToBase64String(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestEncrypt_BySha384_FromBytesToBytes(t *testing.T) {
	hexTests := []struct {
		input    []byte // 输入值
		expected []byte // 期望值
	}{
		{[]byte(""), []byte("")},
		{[]byte(sha384Input), []byte(sha384HexExpected)},
	}

	for index, test := range hexTests {
		e := Encrypt.FromBytes(test.input).BySha384()
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToHexBytes(), "Current test index is "+strconv.Itoa(index))
	}

	base32Tests := []struct {
		input    []byte // 输入值
		expected []byte // 期望值
	}{
		{[]byte(""), []byte("")},
		{[]byte(sha384Input), []byte(sha384Base32Expected)},
	}

	for index, test := range base32Tests {
		e := Encrypt.FromBytes(test.input).BySha384()
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToBase32Bytes(), "Current test index is "+strconv.Itoa(index))
	}

	base64Tests := []struct {
		input    []byte // 输入值
		expected []byte // 期望值
	}{
		{[]byte(""), []byte("")},
		{[]byte(sha384Input), []byte(sha384Base64Expected)},
	}

	for index, test := range base64Tests {
		e := Encrypt.FromBytes(test.input).BySha384()
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToBase64Bytes(), "Current test index is "+strconv.Itoa(index))
	}
}
