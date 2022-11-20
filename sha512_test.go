package dongle

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

var sha512Test = []struct {
	input    string
	toHex    string
	toBase64 string
}{
	{"", "", ""},
	{"hello world", "309ecc489c12d6eb4cc40f50c902f2b4d0ed77ee511a7c7a9bcd3ca86d4cd86f989dd35bc5ff499670da34255b45b0cfd830e81f605dcf7dc5542e93ae9cd76f", "MJ7MSJwS1utMxA9QyQLytNDtd+5RGnx6m808qG1M2G+YndNbxf9JlnDaNCVbRbDP2DDoH2Bdz33FVC6TrpzXbw=="},
}

func TestEncrypt_BySha512_FromStringToString(t *testing.T) {
	for index, test := range sha512Test {
		e := Encrypt.FromString(test.input).BySha512()
		assert.Nil(t, e.Error)

		assert.Equal(t, test.toHex, e.ToHexString(), "Current test index is "+strconv.Itoa(index))
		assert.Equal(t, test.toBase64, e.ToBase64String(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestEncrypt_BySha512_FromBytesToBytes(t *testing.T) {
	for index, test := range sha512Test {
		e := Encrypt.FromBytes([]byte(test.input)).BySha512()
		assert.Nil(t, e.Error)

		assert.Equal(t, []byte(test.toHex), e.ToHexBytes(), "Current test index is "+strconv.Itoa(index))
		assert.Equal(t, []byte(test.toBase64), e.ToBase64Bytes(), "Current test index is "+strconv.Itoa(index))
	}
}
