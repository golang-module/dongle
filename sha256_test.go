package dongle

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

var sha256Test = []struct {
	input    string
	toHex    string
	toBase64 string
}{
	{"", "", ""},
	{"hello world", "b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9", "uU0nuZNNPgilLlLX2n2r+sSE7+N6U4DukIj3rOLvzek="},
}

func TestEncrypt_BySha256_ToString(t *testing.T) {
	for index, test := range sha256Test {
		e := Encrypt.FromString(test.input).BySha256()
		assert.Nil(t, e.Error)

		assert.Equal(t, test.toHex, e.ToHexString(), "Current test index is "+strconv.Itoa(index))
		assert.Equal(t, test.toBase64, e.ToBase64String(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestEncrypt_BySha256_ToBytes(t *testing.T) {
	for index, test := range sha256Test {
		e := Encrypt.FromBytes([]byte(test.input)).BySha256()
		assert.Nil(t, e.Error)

		assert.Equal(t, []byte(test.toHex), e.ToHexBytes(), "Current test index is "+strconv.Itoa(index))
		assert.Equal(t, []byte(test.toBase64), e.ToBase64Bytes(), "Current test index is "+strconv.Itoa(index))
	}
}
