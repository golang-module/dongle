package dongle

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

var teaTest = []struct {
	input    string
	key      string
	rounds   int
	toHex    string
	toBase64 string
}{
	{"", "", 8, "", ""},
	{"", "0123456789abcdef", 8, "", ""},
	{"hello go", "0123456789abcdef", 8, "06f1e586e866a2b7", "BvHlhuhmorc="},
}

func TestTea_Encrypt_ToString(t *testing.T) {
	for index, test := range teaTest {
		e := Encrypt.FromString(test.input).ByTea(test.key, test.rounds)
		assert.Nil(t, e.Error)

		assert.Equal(t, test.toHex, e.ToHexString(), "Hex test index is "+strconv.Itoa(index))
		assert.Equal(t, test.toBase64, e.ToBase64String(), "Base64 test index is "+strconv.Itoa(index))
	}
}

func TestTea_Encrypt_ToBytes(t *testing.T) {
	for index, test := range teaTest {
		e := Encrypt.FromBytes([]byte(test.input)).ByTea([]byte(test.key), test.rounds)
		assert.Nil(t, e.Error)

		assert.Equal(t, []byte(test.toHex), e.ToHexBytes(), "Hex test test index is "+strconv.Itoa(index))
		assert.Equal(t, []byte(test.toBase64), e.ToBase64Bytes(), "Base64 test index is "+strconv.Itoa(index))
	}
}

func TestTea_Decrypt_ToString(t *testing.T) {
	for index, test := range teaTest {
		d1 := Decrypt.FromHexString(test.toHex).ByTea(test.key, test.rounds)
		assert.Nil(t, d1.Error)
		assert.Equal(t, test.input, d1.ToString(), "Hex test index is "+strconv.Itoa(index))

		d2 := Decrypt.FromBase64String(test.toBase64).ByTea(test.key, test.rounds)
		assert.Nil(t, d2.Error)
		assert.Equal(t, test.input, d2.ToString(), "base64 test index is "+strconv.Itoa(index))
	}
}

func TestTea_Decrypt_ToBytes(t *testing.T) {
	for index, test := range teaTest {
		e1 := Decrypt.FromHexBytes([]byte(test.toHex)).ByTea([]byte(test.key), test.rounds)
		assert.Nil(t, e1.Error)
		assert.Equal(t, []byte(test.input), e1.ToBytes(), "Hex test index is "+strconv.Itoa(index))

		e2 := Decrypt.FromBase64Bytes([]byte(test.toBase64)).ByTea([]byte(test.key), test.rounds)
		assert.Nil(t, e2.Error)
		assert.Equal(t, []byte(test.input), e2.ToBytes(), "base64 test index is "+strconv.Itoa(index))
	}
}

func TestTea_Src_Error(t *testing.T) {
	e := Encrypt.FromString("xxxx").ByTea("0123456789abcdefghijklmn", 8)
	assert.Equal(t, invalidTeaSrcError(), e.Error)

	d := Decrypt.FromString("xxxx").ByTea("0123456789abcdefghijklmn", 8)
	assert.Equal(t, invalidTeaSrcError(), d.Error)
}

func TestTea_Key_Error(t *testing.T) {
	e := Encrypt.FromString("hello go").ByTea("xxxx", 8)
	assert.Equal(t, invalidTeaKeyError(), e.Error)

	d := Decrypt.FromString("hello go").ByTea("xxxx", 8)
	assert.Equal(t, invalidTeaKeyError(), d.Error)
}

func TestTea_Rounds_Error(t *testing.T) {
	e := Encrypt.FromString("hello go").ByTea("0123456789abcdefghijklmn", 1)
	assert.Equal(t, invalidTeaRoundsError(), e.Error)

	d := Decrypt.FromString("hello go").ByTea("0123456789abcdefghijklmn", 1)
	assert.Equal(t, invalidTeaRoundsError(), d.Error)
}
