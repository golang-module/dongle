package dongle

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

var hashTest = []struct {
	hash     string
	input    string
	toHex    string
	toBase64 string
}{
	{"md4", "", "", ""},
	{"md4", "hello world", "aa010fbc1d14c795d86ef98c95479d17", "qgEPvB0Ux5XYbvmMlUedFw=="},

	{"md5", "", "", ""},
	{"md5", "hello world", "5eb63bbbe01eeed093cb22bb8f5acdc3", "XrY7u+Ae7tCTyyK7j1rNww=="},

	{"sha1", "", "", ""},
	{"sha1", "hello world", "2aae6c35c94fcfb415dbe95f408b9ce91ee846ed", "Kq5sNclPz7QV2+lfQIuc6R7oRu0="},

	{"sha224", "", "", ""},
	{"sha224", "hello world", "2f05477fc24bb4faefd86517156dafdecec45b8ad3cf2522a563582b", "LwVHf8JLtPrv2GUXFW2v3s7EW4rTzyUipWNYKw=="},

	{"sha256", "", "", ""},
	{"sha256", "hello world", "b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9", "uU0nuZNNPgilLlLX2n2r+sSE7+N6U4DukIj3rOLvzek="},

	{"sha384", "", "", ""},
	{"sha384", "hello world", "fdbd8e75a67f29f701a4e040385e2e23986303ea10239211af907fcbb83578b3e417cb71ce646efd0819dd8c088de1bd", "/b2OdaZ/KfcBpOBAOF4uI5hjA+oQI5IRr5B/y7g1eLPkF8txzmRu/QgZ3YwIjeG9"},

	{"sha512", "", "", ""},
	{"sha512", "hello world", "309ecc489c12d6eb4cc40f50c902f2b4d0ed77ee511a7c7a9bcd3ca86d4cd86f989dd35bc5ff499670da34255b45b0cfd830e81f605dcf7dc5542e93ae9cd76f", "MJ7MSJwS1utMxA9QyQLytNDtd+5RGnx6m808qG1M2G+YndNbxf9JlnDaNCVbRbDP2DDoH2Bdz33FVC6TrpzXbw=="},

	{"ripemd160", "", "", ""},
	{"ripemd160", "hello world", "98c615784ccb5fe5936fbc0cbe9dfdb408d92f0f", "mMYVeEzLX+WTb7wMvp39tAjZLw8="},
}

func TestHash_Encrypt_ToString(t *testing.T) {
	for index, test := range hashTest {
		e := Encrypt.FromString(test.input)

		switch test.hash {
		case "md4":
			e = e.ByMd4()
		case "md5":
			e = e.ByMd5()
		case "sha1":
			e = e.BySha1()
		case "sha224":
			e = e.BySha224()
		case "sha256":
			e = e.BySha256()
		case "sha384":
			e = e.BySha384()
		case "sha512":
			e = e.BySha512()
		case "ripemd160":
			e = e.ByRipemd160()
		}

		assert.Nil(t, e.Error)

		assert.Equal(t, test.toHex, e.ToHexString(), "Current test index is "+strconv.Itoa(index))
		assert.Equal(t, test.toBase64, e.ToBase64String(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestHash_Encrypt_ToBytes(t *testing.T) {
	for index, test := range hashTest {
		e := Encrypt.FromBytes([]byte(test.input))

		switch test.hash {
		case "md4":
			e = e.ByMd4()
		case "md5":
			e = e.ByMd5()
		case "sha1":
			e = e.BySha1()
		case "sha224":
			e = e.BySha224()
		case "sha256":
			e = e.BySha256()
		case "sha384":
			e = e.BySha384()
		case "sha512":
			e = e.BySha512()
		case "ripemd160":
			e = e.ByRipemd160()
		}

		assert.Nil(t, e.Error)

		assert.Equal(t, []byte(test.toHex), e.ToHexBytes(), "Current test index is "+strconv.Itoa(index))
		assert.Equal(t, []byte(test.toBase64), e.ToBase64Bytes(), "Current test index is "+strconv.Itoa(index))
	}
}
