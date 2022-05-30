package dongle

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	sha512Input          = "hello world"
	sha512HexExpected    = "309ecc489c12d6eb4cc40f50c902f2b4d0ed77ee511a7c7a9bcd3ca86d4cd86f989dd35bc5ff499670da34255b45b0cfd830e81f605dcf7dc5542e93ae9cd76f"
	sha512Base32Expected = "GCPMYSE4CLLOWTGEB5IMSAXSWTIO257OKENHY6U3ZU6KQ3KM3BXZRHOTLPC76SMWODNDIJK3IWYM7WBQ5APWAXOPPXCVILUTV2ONO3Y="
	sha512Base64Expected = "MJ7MSJwS1utMxA9QyQLytNDtd+5RGnx6m808qG1M2G+YndNbxf9JlnDaNCVbRbDP2DDoH2Bdz33FVC6TrpzXbw=="
)

func TestEncrypt_BySha512_FromStringToString(t *testing.T) {
	hexTests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{sha512Input, sha512HexExpected},
	}

	for index, test := range hexTests {
		e := Encrypt.FromString(test.input).BySha512()
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToHexString(), "Current test index is "+strconv.Itoa(index))
	}

	base32Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{sha512Input, sha512Base32Expected},
	}

	for index, test := range base32Tests {
		e := Encrypt.FromString(test.input).BySha512()
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToBase32String(), "Current test index is "+strconv.Itoa(index))
	}

	base64Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{sha512Input, sha512Base64Expected},
	}

	for index, test := range base64Tests {
		e := Encrypt.FromString(test.input).BySha512()
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToBase64String(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestEncrypt_BySha512_FromBytesToString(t *testing.T) {
	hexTests := []struct {
		input    []byte // 输入值
		expected string // 期望值
	}{
		{[]byte(""), ""},
		{[]byte(sha512Input), sha512HexExpected},
	}

	for index, test := range hexTests {
		e := Encrypt.FromBytes(test.input).BySha512()
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToHexString(), "Current test index is "+strconv.Itoa(index))
	}

	base32Tests := []struct {
		input    []byte // 输入值
		expected string // 期望值
	}{
		{[]byte(""), ""},
		{[]byte(sha512Input), sha512Base32Expected},
	}

	for index, test := range base32Tests {
		e := Encrypt.FromBytes(test.input).BySha512()
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToBase32String(), "Current test index is "+strconv.Itoa(index))
	}

	base64Tests := []struct {
		input    []byte // 输入值
		expected string // 期望值
	}{
		{[]byte(""), ""},
		{[]byte(sha512Input), sha512Base64Expected},
	}

	for index, test := range base64Tests {
		e := Encrypt.FromBytes(test.input).BySha512()
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToBase64String(), "Current test index is "+strconv.Itoa(index))
	}
}
