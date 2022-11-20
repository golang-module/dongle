package dongle

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

var sha1Test = []struct {
	input    string
	toHex    string
	toBase64 string
}{
	{"", "", ""},
	{"hello world", "2aae6c35c94fcfb415dbe95f408b9ce91ee846ed", "Kq5sNclPz7QV2+lfQIuc6R7oRu0="},
}

func TestEncrypt_BySha1_FromStringToString(t *testing.T) {
	for index, test := range sha1Test {
		e := Encrypt.FromString(test.input).BySha1()
		assert.Nil(t, e.Error)

		assert.Equal(t, test.toHex, e.ToHexString(), "Current test index is "+strconv.Itoa(index))
		assert.Equal(t, test.toBase64, e.ToBase64String(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestEncrypt_BySha1_FromBytesToBytes(t *testing.T) {
	for index, test := range sha1Test {
		e := Encrypt.FromBytes([]byte(test.input)).BySha1()
		assert.Nil(t, e.Error)

		assert.Equal(t, []byte(test.toHex), e.ToHexBytes(), "Current test index is "+strconv.Itoa(index))
		assert.Equal(t, []byte(test.toBase64), e.ToBase64Bytes(), "Current test index is "+strconv.Itoa(index))
	}
}
