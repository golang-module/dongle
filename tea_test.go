package dongle

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

var teaTests = []struct {
	input    string
	key      string
	rounds   int
	toHex    string
	toBase64 string
}{
	{"", "", 64, "", ""},
	{"", "0123456789abcdef", 64, "", ""},
	{"hello go", "0123456789abcdef", 0, "3f9a9d3f2f58277f", "P5qdPy9YJ38="},
	{"hello world", "0123456789abcdef", 64, "bfa8d956247c0dcecfe2b0ef44b08aab", "v6jZViR8Dc7P4rDvRLCKqw=="},
}

func TestTea_Encrypt_String(t *testing.T) {
	for index, test := range teaTests {
		e1 := Encrypt.FromString(test.input).ByTea(test.key, test.rounds)
		if test.rounds == 0 {
			e1 = Encrypt.FromString(test.input).ByTea(test.key)
		}
		assert.Nil(t, e1.Error)
		assert.Equal(t, test.toHex, e1.ToHexString(), "Hex test index is "+strconv.Itoa(index))
		assert.Equal(t, test.toBase64, e1.ToBase64String(), "Base64 test index is "+strconv.Itoa(index))

		e2 := Encrypt.FromString(test.input).ByTea(test.key, test.rounds)
		if test.rounds == 0 {
			e2 = Encrypt.FromString(test.input).ByTea(test.key)
		}
		assert.Nil(t, e2.Error)
		assert.Equal(t, test.toHex, e2.ToHexString(), "Hex test index is "+strconv.Itoa(index))
		assert.Equal(t, test.toBase64, e2.ToBase64String(), "Base64 test index is "+strconv.Itoa(index))
	}
}

func TestTea_Decrypt_String(t *testing.T) {
	for index, test := range teaTests {
		d1 := Decrypt.FromHexString(test.toHex).ByTea(test.key, test.rounds)
		if test.rounds == 0 {
			d1 = Decrypt.FromHexString(test.toHex).ByTea(test.key)
		}
		assert.Nil(t, d1.Error)
		assert.Equal(t, test.input, d1.ToString(), "Hex test index is "+strconv.Itoa(index))

		d2 := Decrypt.FromBase64String(test.toBase64).ByTea(test.key, test.rounds)
		if test.rounds == 0 {
			d2 = Decrypt.FromBase64String(test.toBase64).ByTea(test.key)
		}
		assert.Nil(t, d2.Error)
		assert.Equal(t, test.input, d2.ToString(), "base64 test index is "+strconv.Itoa(index))
	}
}

func TestTea_Encrypt_Bytes(t *testing.T) {
	for index, test := range teaTests {
		e := Encrypt.FromBytes([]byte(test.input)).ByTea([]byte(test.key), test.rounds)
		if test.rounds == 0 {
			e = Encrypt.FromBytes([]byte(test.input)).ByTea([]byte(test.key))
		}
		assert.Nil(t, e.Error)
		assert.Equal(t, []byte(test.toHex), e.ToHexBytes(), "Hex test test index is "+strconv.Itoa(index))
		assert.Equal(t, []byte(test.toBase64), e.ToBase64Bytes(), "Base64 test index is "+strconv.Itoa(index))
	}
}

func TestTea_Decrypt_Bytes(t *testing.T) {
	for index, test := range teaTests {
		e1 := Decrypt.FromHexBytes([]byte(test.toHex)).ByTea([]byte(test.key), test.rounds)
		if test.rounds == 0 {
			e1 = Decrypt.FromHexBytes([]byte(test.toHex)).ByTea([]byte(test.key))
		}
		assert.Nil(t, e1.Error)
		assert.Equal(t, []byte(test.input), e1.ToBytes(), "Hex test index is "+strconv.Itoa(index))

		e2 := Decrypt.FromBase64Bytes([]byte(test.toBase64)).ByTea([]byte(test.key), test.rounds)
		if test.rounds == 0 {
			e2 = Decrypt.FromBase64Bytes([]byte(test.toBase64)).ByTea([]byte(test.key))
		}
		assert.Nil(t, e2.Error)
		assert.Equal(t, []byte(test.input), e2.ToBytes(), "base64 test index is "+strconv.Itoa(index))
	}
}

func TestTea_Key_Error(t *testing.T) {
	e := Encrypt.FromString("hello go").ByTea("xxxx", 8)
	assert.Equal(t, invalidTeaKeyError(), e.Error)

	d := Decrypt.FromRawString("hello go").ByTea("xxxx", 8)
	assert.Equal(t, invalidTeaKeyError(), d.Error)
}

func TestTea_Rounds_Error(t *testing.T) {
	e := Encrypt.FromString("hello go").ByTea("0123456789abcdefghijklmn", 1)
	assert.Equal(t, invalidTeaRoundsError(), e.Error)

	d := Decrypt.FromRawString("hello go").ByTea("0123456789abcdefghijklmn", 1)
	assert.Equal(t, invalidTeaRoundsError(), d.Error)
}
