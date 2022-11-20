package dongle

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

var sha224Test = []struct {
	input    string
	toHex    string
	toBase64 string
}{
	{"", "", ""},
	{"hello world", "2f05477fc24bb4faefd86517156dafdecec45b8ad3cf2522a563582b", "LwVHf8JLtPrv2GUXFW2v3s7EW4rTzyUipWNYKw=="},
}

func TestEncrypt_BySha224_FromStringToString(t *testing.T) {
	for index, test := range sha224Test {
		e := Encrypt.FromString(test.input).BySha224()
		assert.Nil(t, e.Error)

		assert.Equal(t, test.toHex, e.ToHexString(), "Current test index is "+strconv.Itoa(index))
		assert.Equal(t, test.toBase64, e.ToBase64String(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestEncrypt_BySha224_FromBytesToBytes(t *testing.T) {
	for index, test := range sha224Test {
		e := Encrypt.FromBytes([]byte(test.input)).BySha224()
		assert.Nil(t, e.Error)

		assert.Equal(t, []byte(test.toHex), e.ToHexBytes(), "Current test index is "+strconv.Itoa(index))
		assert.Equal(t, []byte(test.toBase64), e.ToBase64Bytes(), "Current test index is "+strconv.Itoa(index))
	}
}
