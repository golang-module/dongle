package dongle

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

var hmacTest = []struct {
	hash     string
	input    string
	key      string
	toHex    string
	toBase64 string
}{
	{"md4", "", "", "", ""},
	{"md4", "hello world", "dongle", "7a9df5247cbf76a8bc17c9c4f5a75b6b", "ep31JHy/dqi8F8nE9adbaw=="},

	{"md5", "", "", "", ""},
	{"md5", "hello world", "dongle", "4790626a275f776956386e5a3ea7b726", "R5Biaidfd2lWOG5aPqe3Jg=="},

	{"sha1", "", "", "", ""},
	{"sha1", "hello world", "dongle", "91c103ef93ba7420902b0d1bf0903251c94b4a62", "kcED75O6dCCQKw0b8JAyUclLSmI="},

	{"sha224", "", "", "", ""},
	{"sha224", "hello world", "dongle", "e15b9e5a7eccb1f17dc81dc07c909a891936dc3429dc0d940accbcec", "4VueWn7MsfF9yB3AfJCaiRk23DQp3A2UCsy87A=="},

	{"sha256", "", "", "", ""},
	{"sha256", "hello world", "dongle", "77f5c8ce4147600543e70b12701e7b78b5b95172332ebbb06de65fcea7112179", "d/XIzkFHYAVD5wsScB57eLW5UXIzLruwbeZfzqcRIXk="},

	{"sha384", "", "", "", ""},
	{"sha384", "hello world", "dongle", "421fcaa740216a31bbcd1f86f2212e0c68aa4b156a8ebc2ae55b3e75c4ee0509ea0325a0570ae739006b61d91d817fe8", "Qh/Kp0AhajG7zR+G8iEuDGiqSxVqjrwq5Vs+dcTuBQnqAyWgVwrnOQBrYdkdgX/o"},

	{"sha512", "", "", "", ""},
	{"sha512", "hello world", "dongle", "d971b790bbc2a4ac81062bbffac693c9c234bae176c8faf5e304dbdb153032a826f12353964b4a4fb87abecd2dc237638a630cbad54a6b94b1f6ef5d5e2835d1", "2XG3kLvCpKyBBiu/+saTycI0uuF2yPr14wTb2xUwMqgm8SNTlktKT7h6vs0twjdjimMMutVKa5Sx9u9dXig10Q=="},

	{"ripemd160", "", "", "", ""},
	{"ripemd160", "hello world", "dongle", "3691ad040e80c43dc6e8ffe9bc6ef3d5bd8786b8", "NpGtBA6AxD3G6P/pvG7z1b2Hhrg="},

	{"sm3", "", "", "", ""},
	{"sm3", "hello world", "dongle", "8c733aae1d553c466a08c3e9e5daac3e99ae220181c7c1bc8c2564961de751b3", "jHM6rh1VPEZqCMPp5dqsPpmuIgGBx8G8jCVklh3nUbM="},
}

func TestHmac_Encrypt_ToString(t *testing.T) {
	for index, test := range hmacTest {
		e := Encrypt.FromString(test.input)

		switch test.hash {
		case "md4":
			e = e.ByHmacMd4(test.key)
		case "md5":
			e = e.ByHmacMd5(test.key)
		case "sha1":
			e = e.ByHmacSha1(test.key)
		case "sha224":
			e = e.ByHmacSha224(test.key)
		case "sha256":
			e = e.ByHmacSha256(test.key)
		case "sha384":
			e = e.ByHmacSha384(test.key)
		case "sha512":
			e = e.ByHmacSha512(test.key)
		case "ripemd160":
			e = e.ByHmacRipemd160(test.key)
		case "sm3":
			e = e.ByHmacSm3(test.key)
		}

		assert.Nil(t, e.Error)

		assert.Equal(t, test.toHex, e.ToHexString(), "Hex test index is "+strconv.Itoa(index))
		assert.Equal(t, test.toBase64, e.ToBase64String(), "Base64 test index is "+strconv.Itoa(index))
	}
}

func TestHmac_Encrypt_ToBytes(t *testing.T) {
	for index, test := range hmacTest {
		e := Encrypt.FromBytes([]byte(test.input)).ByHmacMd4([]byte(test.key))

		switch test.hash {
		case "md4":
			e = e.ByHmacMd4([]byte(test.key))
		case "md5":
			e = e.ByHmacMd5([]byte(test.key))
		case "sha1":
			e = e.ByHmacSha1([]byte(test.key))
		case "sha224":
			e = e.ByHmacSha224([]byte(test.key))
		case "sha256":
			e = e.ByHmacSha256([]byte(test.key))
		case "sha384":
			e = e.ByHmacSha384([]byte(test.key))
		case "sha512":
			e = e.ByHmacSha512([]byte(test.key))
		case "ripemd160":
			e = e.ByHmacRipemd160([]byte(test.key))
		case "sm3":
			e = e.ByHmacSm3([]byte(test.key))
		}

		assert.Nil(t, e.Error)

		assert.Equal(t, []byte(test.toHex), e.ToHexBytes(), "Hex test test index is "+strconv.Itoa(index))
		assert.Equal(t, []byte(test.toBase64), e.ToBase64Bytes(), "Base64 test index is "+strconv.Itoa(index))
	}
}
