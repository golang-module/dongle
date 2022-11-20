package dongle

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

var md5Test = []struct {
	input    string
	toHex    string
	toBase64 string
}{
	{"", "", ""},
	{"hello world", "5eb63bbbe01eeed093cb22bb8f5acdc3", "XrY7u+Ae7tCTyyK7j1rNww=="},
}

func TestEncrypt_ByMd5_ToString(t *testing.T) {
	for index, test := range md5Test {
		e := Encrypt.FromString(test.input).ByMd5()
		assert.Nil(t, e.Error)

		assert.Equal(t, test.toHex, e.ToHexString(), "Current test index is "+strconv.Itoa(index))
		assert.Equal(t, test.toBase64, e.ToBase64String(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestEncrypt_ByMd5_ToBytes(t *testing.T) {
	for index, test := range md5Test {
		e := Encrypt.FromBytes([]byte(test.input)).ByMd5()
		assert.Nil(t, e.Error)

		assert.Equal(t, []byte(test.toHex), e.ToHexBytes(), "Current test index is "+strconv.Itoa(index))
		assert.Equal(t, []byte(test.toBase64), e.ToBase64Bytes(), "Current test index is "+strconv.Itoa(index))
	}
}
