package dongle

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

var md4Test = []struct {
	input    string
	toHex    string
	toBase64 string
}{
	{"", "", ""},
	{"hello world", "aa010fbc1d14c795d86ef98c95479d17", "qgEPvB0Ux5XYbvmMlUedFw=="},
}

func TestEncrypt_ByMd4_ToString(t *testing.T) {
	for index, test := range md4Test {
		e := Encrypt.FromString(test.input).ByMd4()
		assert.Nil(t, e.Error)

		assert.Equal(t, test.toHex, e.ToHexString(), "Current test index is "+strconv.Itoa(index))
		assert.Equal(t, test.toBase64, e.ToBase64String(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestEncrypt_ByMd4_ToBytes(t *testing.T) {
	for index, test := range md4Test {
		e := Encrypt.FromBytes([]byte(test.input)).ByMd4()
		assert.Nil(t, e.Error)

		assert.Equal(t, []byte(test.toHex), e.ToHexBytes(), "Current test index is "+strconv.Itoa(index))
		assert.Equal(t, []byte(test.toBase64), e.ToBase64Bytes(), "Current test index is "+strconv.Itoa(index))
	}
}
