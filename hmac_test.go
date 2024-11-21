package dongle

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var hmacTests = []struct {
	algo     string
	input    string
	key      string
	toHex    string
	toBase64 string
}{
	{"md2", "", "", "", ""},
	{"md2", "hello world", "dongle", "88ed6ef9ab699d03a702f2a6fb1c0673", "iO1u+atpnQOnAvKm+xwGcw=="},

	{"md4", "", "", "", ""},
	{"md4", "hello world", "dongle", "7a9df5247cbf76a8bc17c9c4f5a75b6b", "ep31JHy/dqi8F8nE9adbaw=="},

	{"md5", "", "", "", ""},
	{"md5", "hello world", "dongle", "4790626a275f776956386e5a3ea7b726", "R5Biaidfd2lWOG5aPqe3Jg=="},

	{"sha1", "", "", "", ""},
	{"sha1", "hello world", "dongle", "91c103ef93ba7420902b0d1bf0903251c94b4a62", "kcED75O6dCCQKw0b8JAyUclLSmI="},

	{"sha3-224", "", "", "", ""},
	{"sha3-224", "hello world", "dongle", "fb8f061d9d1dddd2f5d3b9064a5e98e3e4b6df27ea93ce67627583ce", "+48GHZ0d3dL107kGSl6Y4+S23yfqk85nYnWDzg=="},

	{"sha3-256", "", "", "", ""},
	{"sha3-256", "hello world", "dongle", "8193367fde28cf5c460adb449a04b3dd9c184f488bdccbabf0526c54f90c4460", "gZM2f94oz1xGCttEmgSz3ZwYT0iL3Mur8FJsVPkMRGA="},

	{"sha3-384", "", "", "", ""},
	{"sha3-384", "hello world", "dongle", "3f76f5cda69cada3ee6b33f8458cd498b063075db263dd8b33f2a3992a8804f9569a7c86ffa2b8f0748babeb7a6fc0e7", "P3b1zaacraPuazP4RYzUmLBjB12yY92LM/KjmSqIBPlWmnyG/6K48HSLq+t6b8Dn"},

	{"sha3-512", "", "", "", ""},
	{"sha3-512", "hello world", "dongle", "a99653d0407d659eccdeed43bb7cccd2e2b05a2c34fd3467c4198cf2ad26a466738513e88839fb55e64eb49df65bc52ed0fec2775bd9e086edd4fb4024add4a2", "qZZT0EB9ZZ7M3u1Du3zM0uKwWiw0/TRnxBmM8q0mpGZzhRPoiDn7VeZOtJ32W8Uu0P7Cd1vZ4Ibt1PtAJK3Uog=="},

	{"sha224", "", "", "", ""},
	{"sha224", "hello world", "dongle", "e15b9e5a7eccb1f17dc81dc07c909a891936dc3429dc0d940accbcec", "4VueWn7MsfF9yB3AfJCaiRk23DQp3A2UCsy87A=="},

	{"sha256", "", "", "", ""},
	{"sha256", "hello world", "dongle", "77f5c8ce4147600543e70b12701e7b78b5b95172332ebbb06de65fcea7112179", "d/XIzkFHYAVD5wsScB57eLW5UXIzLruwbeZfzqcRIXk="},

	{"sha384", "", "", "", ""},
	{"sha384", "hello world", "dongle", "421fcaa740216a31bbcd1f86f2212e0c68aa4b156a8ebc2ae55b3e75c4ee0509ea0325a0570ae739006b61d91d817fe8", "Qh/Kp0AhajG7zR+G8iEuDGiqSxVqjrwq5Vs+dcTuBQnqAyWgVwrnOQBrYdkdgX/o"},

	{"sha512", "", "", "", ""},
	{"sha512", "hello world", "dongle", "d971b790bbc2a4ac81062bbffac693c9c234bae176c8faf5e304dbdb153032a826f12353964b4a4fb87abecd2dc237638a630cbad54a6b94b1f6ef5d5e2835d1", "2XG3kLvCpKyBBiu/+saTycI0uuF2yPr14wTb2xUwMqgm8SNTlktKT7h6vs0twjdjimMMutVKa5Sx9u9dXig10Q=="},

	{"sha512-224", "", "", "", ""},
	{"sha512-224", "hello world", "dongle", "c4145bcc385c29f0e5683cd5450be9deb522d556de3b046a7ffa1eb3", "xBRbzDhcKfDlaDzVRQvp3rUi1VbeOwRqf/oesw=="},

	{"sha512-256", "", "", "", ""},
	{"sha512-256", "hello world", "dongle", "e99fae71bcb43651ae10e952989eadf897faccb43966ee5122bb1b1d82f7a7c2", "6Z+ucby0NlGuEOlSmJ6t+Jf6zLQ5Zu5RIrsbHYL3p8I="},

	{"ripemd160", "", "", "", ""},
	{"ripemd160", "hello world", "dongle", "3691ad040e80c43dc6e8ffe9bc6ef3d5bd8786b8", "NpGtBA6AxD3G6P/pvG7z1b2Hhrg="},

	{"sm3", "", "", "", ""},
	{"sm3", "hello world", "dongle", "8c733aae1d553c466a08c3e9e5daac3e99ae220181c7c1bc8c2564961de751b3", "jHM6rh1VPEZqCMPp5dqsPpmuIgGBx8G8jCVklh3nUbM="},
}

func TestHmac_Encrypt_String(t *testing.T) {
	for index, test := range hmacTests {
		key := []byte(test.key)
		e := Encrypt.FromString(test.input)

		switch test.algo {
		case "md2":
			e = e.ByHmacMd2(key)
		case "md4":
			e = e.ByHmacMd4(key)
		case "md5":
			e = e.ByHmacMd5(key)
		case "sha1":
			e = e.ByHmacSha1(key)
		case "sha3-224":
			e = e.ByHmacSha3(key, 224)
		case "sha3-256":
			e = e.ByHmacSha3(key, 256)
		case "sha3-384":
			e = e.ByHmacSha3(key, 384)
		case "sha3-512":
			e = e.ByHmacSha3(key, 512)
		case "sha224":
			e = e.ByHmacSha224(key)
		case "sha256":
			e = e.ByHmacSha256(key)
		case "sha384":
			e = e.ByHmacSha384(key)
		case "sha512":
			e = e.ByHmacSha512(key)
		case "sha512-224":
			e = e.ByHmacSha512(key, 224)
		case "sha512-256":
			e = e.ByHmacSha512(key, 256)
		case "ripemd160":
			e = e.ByHmacRipemd160(key)
		case "sm3":
			e = e.ByHmacSm3(key)
		}

		t.Run(fmt.Sprintf(test.algo+"_test_%d", index), func(t *testing.T) {
			assert.Nil(t, e.Error)
			assert.Equal(t, fmt.Sprintf("%s", e), e.ToRawString())
			assert.Equal(t, test.toHex, e.ToHexString())
			assert.Equal(t, test.toBase64, e.ToBase64String())
		})
	}
}

func TestHmac_Encrypt_Bytes(t *testing.T) {
	for index, test := range hmacTests {
		key := []byte(test.key)

		e := Encrypt.FromBytes([]byte(test.input)).ByHmacMd4(key)

		switch test.algo {
		case "md2":
			e = e.ByHmacMd2(key)
		case "md4":
			e = e.ByHmacMd4(key)
		case "md5":
			e = e.ByHmacMd5(key)
		case "sha1":
			e = e.ByHmacSha1(key)
		case "sha3-224":
			e = e.ByHmacSha3(key, 224)
		case "sha3-256":
			e = e.ByHmacSha3(key, 256)
		case "sha3-384":
			e = e.ByHmacSha3(key, 384)
		case "sha3-512":
			e = e.ByHmacSha3(key, 512)
		case "sha224":
			e = e.ByHmacSha224(key)
		case "sha256":
			e = e.ByHmacSha256(key)
		case "sha384":
			e = e.ByHmacSha384(key)
		case "sha512":
			e = e.ByHmacSha512(key)
		case "sha512-224":
			e = e.ByHmacSha512(key, 224)
		case "sha512-256":
			e = e.ByHmacSha512(key, 256)
		case "ripemd160":
			e = e.ByHmacRipemd160(key)
		case "sm3":
			e = e.ByHmacSm3(key)
		}

		t.Run(fmt.Sprintf(test.algo+"_test_%d", index), func(t *testing.T) {
			assert.Nil(t, e.Error)
			assert.Equal(t, []byte(test.toHex), e.ToHexBytes())
			assert.Equal(t, []byte(test.toBase64), e.ToBase64Bytes())
		})
	}
}

func TestHmac_Size_Error(t *testing.T) {
	err := NewHmacError()
	key := []byte("dongle")

	e1 := Encrypt.FromString("hello world").ByHmacSha3(key, 100)
	assert.Equal(t, err.SizeError(), e1.Error)
	e2 := Encrypt.FromBytes([]byte("hello world")).ByHmacSha3(key, 100)
	assert.Equal(t, err.SizeError(), e2.Error)

	e3 := Encrypt.FromString("hello world").ByHmacSha512(key, 100)
	assert.Equal(t, err.SizeError(), e3.Error)
	e4 := Encrypt.FromBytes([]byte("hello world")).ByHmacSha512(key, 100)
	assert.Equal(t, err.SizeError(), e4.Error)
}
