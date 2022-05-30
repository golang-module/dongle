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
	hexTests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{md4Input, md4HexExpected},
	}

	for index, test := range hexTests {
		e := Encrypt.FromString(test.input).ByMd4()
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToHexString(), "Current test index is "+strconv.Itoa(index))
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
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToBase32String(), "Current test index is "+strconv.Itoa(index))
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
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToBase64String(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestEncrypt_ByMd4_FromBytesToBytes(t *testing.T) {
	hexTests := []struct {
		input    []byte // 输入值
		expected []byte // 期望值
	}{
		{[]byte(""), []byte("")},
		{[]byte(md4Input), []byte(md4HexExpected)},
	}

	for index, test := range hexTests {
		e := Encrypt.FromBytes(test.input).ByMd4()
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToHexBytes(), "Current test index is "+strconv.Itoa(index))
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
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToBase32Bytes(), "Current test index is "+strconv.Itoa(index))
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
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToBase64Bytes(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestEncrypt_ByMd4_FromFileToString(t *testing.T) {
	hexTests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{"./LICENSE", "1240c5c0fb26b585999357915c56b511"},
	}

	for index, test := range hexTests {
		e := Encrypt.FromFile(test.input).ByMd4()
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToHexString(), "Current test index is "+strconv.Itoa(index))
	}

	base32Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{"./LICENSE", "CJAMLQH3E22YLGMTK6IVYVVVCE======"},
	}

	for index, test := range base32Tests {
		e := Encrypt.FromFile(test.input).ByMd4()
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToBase32String(), "Current test index is "+strconv.Itoa(index))
	}

	base64Tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{"./LICENSE", "EkDFwPsmtYWZk1eRXFa1EQ=="},
	}

	for index, test := range base64Tests {
		e := Encrypt.FromFile(test.input).ByMd4()
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToBase64String(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestEncrypt_ByMd4_FromFileToBytes(t *testing.T) {
	hexTests := []struct {
		input    []byte // 输入值
		expected []byte // 期望值
	}{
		{[]byte(""), []byte("")},
		{[]byte("./LICENSE"), []byte("1240c5c0fb26b585999357915c56b511")},
	}

	for index, test := range hexTests {
		e := Encrypt.FromFile(test.input).ByMd4()
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToHexBytes(), "Current test index is "+strconv.Itoa(index))
	}

	base32Tests := []struct {
		input    []byte // 输入值
		expected []byte // 期望值
	}{
		{[]byte(""), []byte("")},
		{[]byte("./LICENSE"), []byte("CJAMLQH3E22YLGMTK6IVYVVVCE======")},
	}

	for index, test := range base32Tests {
		e := Encrypt.FromFile(test.input).ByMd4()
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToBase32Bytes(), "Current test index is "+strconv.Itoa(index))
	}

	base64Tests := []struct {
		input    []byte // 输入值
		expected []byte // 期望值
	}{
		{[]byte(""), []byte("")},
		{[]byte("./LICENSE"), []byte("EkDFwPsmtYWZk1eRXFa1EQ==")},
	}

	for index, test := range base64Tests {
		e := Encrypt.FromFile(test.input).ByMd4()
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToBase64Bytes(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestError_ByMd4(t *testing.T) {
	file := "./demo.txt"
	e := Encrypt.FromFile(file).ByMd4()
	assert.Equal(t, invalidFileError(file), e.Error, "Should catch an exception")
}
