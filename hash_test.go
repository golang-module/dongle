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

	{"shake128-256", "", "", ""},
	{"shake128-256", "hello world", "3a9159f071e4dd1c8c4f968607c30942e120d8156b8b1e72e0d376e8871cb8b8", "OpFZ8HHk3RyMT5aGB8MJQuEg2BVrix5y4NN26IccuLg="},

	{"shake128-512", "", "", ""},
	{"shake128-512", "hello world", "3a9159f071e4dd1c8c4f968607c30942e120d8156b8b1e72e0d376e8871cb8b899072665674f26cc494a4bcf027c58267e8ee2da60e942759de86d2670bba1aa", "OpFZ8HHk3RyMT5aGB8MJQuEg2BVrix5y4NN26IccuLiZByZlZ08mzElKS88CfFgmfo7i2mDpQnWd6G0mcLuhqg=="},

	{"shake256-256", "", "", ""},
	{"shake256-256", "hello world", "369771bb2cb9d2b04c1d54cca487e372d9f187f73f7ba3f65b95c8ee7798c527", "Npdxuyy50rBMHVTMpIfjctnxh/c/e6P2W5XI7neYxSc="},

	{"shake256-512", "", "", ""},
	{"shake256-512", "hello world", "369771bb2cb9d2b04c1d54cca487e372d9f187f73f7ba3f65b95c8ee7798c527f4f3c2d55c2d46a29f2e945d469c3df27853a8735271f5cc2d9e889544357116", "Npdxuyy50rBMHVTMpIfjctnxh/c/e6P2W5XI7neYxSf088LVXC1Gop8ulF1GnD3yeFOoc1Jx9cwtnoiVRDVxFg=="},

	{"blake2b-256", "", "", ""},
	{"blake2b-256", "hello world", "256c83b297114d201b30179f3f0ef0cace9783622da5974326b436178aeef610", "JWyDspcRTSAbMBefPw7wys6Xg2ItpZdDJrQ2F4ru9hA="},

	{"blake2b-384", "", "", ""},
	{"blake2b-384", "hello world", "8c653f8c9c9aa2177fb6f8cf5bb914828faa032d7b486c8150663d3f6524b086784f8e62693171ac51fc80b7d2cbb12b", "jGU/jJyaohd/tvjPW7kUgo+qAy17SGyBUGY9P2UksIZ4T45iaTFxrFH8gLfSy7Er"},

	{"blake2b-512", "", "", ""},
	{"blake2b-512", "hello world", "021ced8799296ceca557832ab941a50b4a11f83478cf141f51f933f653ab9fbcc05a037cddbed06e309bf334942c4e58cdf1a46e237911ccd7fcf9787cbc7fd0", "Ahzth5kpbOylV4MquUGlC0oR+DR4zxQfUfkz9lOrn7zAWgN83b7QbjCb8zSULE5YzfGkbiN5EczX/Pl4fLx/0A=="},

	{"blake2s-256", "", "", ""},
	{"blake2s-256", "hello world", "9aec6806794561107e594b1f6a8a6b0c92a0cba9acf5e5e93cca06f781813b0b", "muxoBnlFYRB+WUsfaoprDJKgy6ms9eXpPMoG94GBOws="},
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
		case "shake128-256":
			e = e.ByShake128(256)
		case "shake128-512":
			e = e.ByShake128(512)
		case "shake256-256":
			e = e.ByShake256(256)
		case "shake256-512":
			e = e.ByShake256(512)
		case "blake2b-256":
			e = e.ByBlake2b(256)
		case "blake2b-384":
			e = e.ByBlake2b(384)
		case "blake2b-512":
			e = e.ByBlake2b(512)
		case "blake2s-256":
			e = e.ByBlake2s(256)
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
		case "shake128-256":
			e = e.ByShake128(256)
		case "shake128-512":
			e = e.ByShake128(512)
		case "shake256-256":
			e = e.ByShake256(256)
		case "shake256-512":
			e = e.ByShake256(512)
		case "blake2b-256":
			e = e.ByBlake2b(256)
		case "blake2b-384":
			e = e.ByBlake2b(384)
		case "blake2b-512":
			e = e.ByBlake2b(512)
		case "blake2s-256":
			e = e.ByBlake2s(256)
		}

		t.Run(fmt.Sprintf(test.algo+"_test_%d", index), func(t *testing.T) {
			assert.Nil(t, e.Error)
			assert.Equal(t, []byte(test.toHex), e.ToHexBytes())
			assert.Equal(t, []byte(test.toBase64), e.ToBase64Bytes())
		})
	}
}

func TestHash_Size_Error(t *testing.T) {
	err := NewHashError()

	e1 := Encrypt.FromString("hello world").BySha3(100)
	assert.Equal(t, err.SizeError(), e1.Error)
	e2 := Encrypt.FromBytes([]byte("hello world")).BySha3(100)
	assert.Equal(t, err.SizeError(), e2.Error)

	e3 := Encrypt.FromString("hello world").BySha512(100)
	assert.Equal(t, err.SizeError(), e3.Error)
	e4 := Encrypt.FromBytes([]byte("hello world")).BySha512(100)
	assert.Equal(t, err.SizeError(), e4.Error)

	e5 := Encrypt.FromString("hello world").ByBlake2b(100)
	assert.Equal(t, err.SizeError(), e5.Error)
	e6 := Encrypt.FromBytes([]byte("hello world")).ByBlake2b(100)
	assert.Equal(t, err.SizeError(), e6.Error)

	e7 := Encrypt.FromString("hello world").ByBlake2s(100)
	assert.Equal(t, err.SizeError(), e7.Error)
	e8 := Encrypt.FromBytes([]byte("hello world")).ByBlake2s(100)
	assert.Equal(t, err.SizeError(), e8.Error)
}
