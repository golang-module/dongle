package dongle

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var hashTests = []struct {
	algo     string
	input    string
	toHex    string
	toBase64 string
}{
	{"md2", "", "", ""},
	{"md2", "hello world", "d9cce882ee690a5c1ce70beff3a78c77", "2czogu5pClwc5wvv86eMdw=="},

	{"md4", "", "", ""},
	{"md4", "hello world", "aa010fbc1d14c795d86ef98c95479d17", "qgEPvB0Ux5XYbvmMlUedFw=="},

	{"md5", "", "", ""},
	{"md5", "hello world", "5eb63bbbe01eeed093cb22bb8f5acdc3", "XrY7u+Ae7tCTyyK7j1rNww=="},

	{"sha1", "", "", ""},
	{"sha1", "hello world", "2aae6c35c94fcfb415dbe95f408b9ce91ee846ed", "Kq5sNclPz7QV2+lfQIuc6R7oRu0="},

	{"sha3-224", "", "", ""},
	{"sha3-224", "hello world", "dfb7f18c77e928bb56faeb2da27291bd790bc1045cde45f3210bb6c5", "37fxjHfpKLtW+ustonKRvXkLwQRc3kXzIQu2xQ=="},

	{"sha3-256", "", "", ""},
	{"sha3-256", "hello world", "644bcc7e564373040999aac89e7622f3ca71fba1d972fd94a31c3bfbf24e3938", "ZEvMflZDcwQJmarInnYi88px+6HZcv2Uoxw7+/JOOTg="},

	{"sha3-384", "", "", ""},
	{"sha3-384", "hello world", "83bff28dde1b1bf5810071c6643c08e5b05bdb836effd70b403ea8ea0a634dc4997eb1053aa3593f590f9c63630dd90b", "g7/yjd4bG/WBAHHGZDwI5bBb24Nu/9cLQD6o6gpjTcSZfrEFOqNZP1kPnGNjDdkL"},

	{"sha3-512", "", "", ""},
	{"sha3-512", "hello world", "840006653e9ac9e95117a15c915caab81662918e925de9e004f774ff82d7079a40d4d27b1b372657c61d46d470304c88c788b3a4527ad074d1dccbee5dbaa99a", "hAAGZT6ayelRF6FckVyquBZikY6SXengBPd0/4LXB5pA1NJ7GzcmV8YdRtRwMEyIx4izpFJ60HTR3MvuXbqpmg=="},

	{"sha224", "", "", ""},
	{"sha224", "hello world", "2f05477fc24bb4faefd86517156dafdecec45b8ad3cf2522a563582b", "LwVHf8JLtPrv2GUXFW2v3s7EW4rTzyUipWNYKw=="},

	{"sha256", "", "", ""},
	{"sha256", "hello world", "b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9", "uU0nuZNNPgilLlLX2n2r+sSE7+N6U4DukIj3rOLvzek="},

	{"sha384", "", "", ""},
	{"sha384", "hello world", "fdbd8e75a67f29f701a4e040385e2e23986303ea10239211af907fcbb83578b3e417cb71ce646efd0819dd8c088de1bd", "/b2OdaZ/KfcBpOBAOF4uI5hjA+oQI5IRr5B/y7g1eLPkF8txzmRu/QgZ3YwIjeG9"},

	{"sha512", "", "", ""},
	{"sha512", "hello world", "309ecc489c12d6eb4cc40f50c902f2b4d0ed77ee511a7c7a9bcd3ca86d4cd86f989dd35bc5ff499670da34255b45b0cfd830e81f605dcf7dc5542e93ae9cd76f", "MJ7MSJwS1utMxA9QyQLytNDtd+5RGnx6m808qG1M2G+YndNbxf9JlnDaNCVbRbDP2DDoH2Bdz33FVC6TrpzXbw=="},

	{"sha512-224", "", "", ""},
	{"sha512-224", "hello world", "22e0d52336f64a998085078b05a6e37b26f8120f43bf4db4c43a64ee", "IuDVIzb2SpmAhQeLBabjeyb4Eg9Dv020xDpk7g=="},

	{"sha512-256", "", "", ""},
	{"sha512-256", "hello world", "0ac561fac838104e3f2e4ad107b4bee3e938bf15f2b15f009ccccd61a913f017", "CsVh+sg4EE4/LkrRB7S+4+k4vxXysV8AnMzNYakT8Bc="},

	{"ripemd160", "", "", ""},
	{"ripemd160", "hello world", "98c615784ccb5fe5936fbc0cbe9dfdb408d92f0f", "mMYVeEzLX+WTb7wMvp39tAjZLw8="},
}

func TestHash_Encrypt_String(t *testing.T) {
	for index, test := range hashTests {
		e := Encrypt.FromString(test.input)

		switch test.algo {
		case "md2":
			e = e.ByMd2()
		case "md4":
			e = e.ByMd4()
		case "md5":
			e = e.ByMd5()
		case "sha1":
			e = e.BySha1()
		case "sha3-224":
			e = e.BySha3(224)
		case "sha3-256":
			e = e.BySha3(256)
		case "sha3-384":
			e = e.BySha3(384)
		case "sha3-512":
			e = e.BySha3(512)
		case "sha224":
			e = e.BySha224()
		case "sha256":
			e = e.BySha256()
		case "sha384":
			e = e.BySha384()
		case "sha512":
			e = e.BySha512()
		case "sha512-224":
			e = e.BySha512(224)
		case "sha512-256":
			e = e.BySha512(256)
		case "ripemd160":
			e = e.ByRipemd160()
		}

		t.Run(fmt.Sprintf(test.algo+"_test_%d", index), func(t *testing.T) {
			assert.Nil(t, e.Error)
			assert.Equal(t, fmt.Sprintf("%s", e), e.ToRawString())
			assert.Equal(t, test.toHex, e.ToHexString())
			assert.Equal(t, test.toBase64, e.ToBase64String())
		})
	}
}

func TestHash_Encrypt_Bytes(t *testing.T) {
	for index, test := range hashTests {
		e := Encrypt.FromBytes([]byte(test.input))

		switch test.algo {
		case "md2":
			e = e.ByMd2()
		case "md4":
			e = e.ByMd4()
		case "md5":
			e = e.ByMd5()
		case "sha1":
			e = e.BySha1()
		case "sha3-224":
			e = e.BySha3(224)
		case "sha3-256":
			e = e.BySha3(256)
		case "sha3-384":
			e = e.BySha3(384)
		case "sha3-512":
			e = e.BySha3(512)
		case "sha224":
			e = e.BySha224()
		case "sha256":
			e = e.BySha256()
		case "sha384":
			e = e.BySha384()
		case "sha512":
			e = e.BySha512()
		case "sha512-224":
			e = e.BySha512(224)
		case "sha512-256":
			e = e.BySha512(256)
		case "ripemd160":
			e = e.ByRipemd160()
		}

		t.Run(fmt.Sprintf(test.algo+"_test_%d", index), func(t *testing.T) {
			assert.Nil(t, e.Error)
			assert.Equal(t, []byte(test.toHex), e.ToHexBytes())
			assert.Equal(t, []byte(test.toBase64), e.ToBase64Bytes())
		})
	}
}

func TestHash_Size_Error(t *testing.T) {
	e1 := Encrypt.FromString("hello world").BySha3(100)
	assert.Equal(t, invalidHashSizeError(), e1.Error)
	e2 := Encrypt.FromBytes([]byte("hello world")).BySha3(100)
	assert.Equal(t, invalidHashSizeError(), e2.Error)

	e3 := Encrypt.FromString("hello world").BySha512(100)
	assert.Equal(t, invalidHashSizeError(), e3.Error)
	e4 := Encrypt.FromBytes([]byte("hello world")).BySha512(100)
	assert.Equal(t, invalidHashSizeError(), e4.Error)
}
