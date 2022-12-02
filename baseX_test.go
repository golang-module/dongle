package dongle

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

var baseXTest = []struct {
	baseX  string
	input  string // 输入值
	output string // 期望值
}{
	{"hex", "", ""},
	{"hex", "hello world", "68656c6c6f20776f726c64"},

	{"base16", "", ""},
	{"base16", "hello world", "68656c6c6f20776f726c64"},

	{"base32", "", ""},
	{"base32", "hello world", "NBSWY3DPEB3W64TMMQ======"},

	{"base58", "", ""},
	{"base58", "hello world", "StV1DL6CwTryKyV"},

	{"base62", "", ""},
	{"base62", "1", "n"},
	{"base62", "f", "1e"},
	{"base62", "1\n2", "DU4g"},
	{"base62", "1234567890", "1A0afZkibIAR2O"},
	{"base62", "hello world", "AAwf93rvy4aWQVw"},
	{"base62", "\x14\xfb\x9c\x03", "Np64R"},
	{"base62", "\xf2\x8e\x88\x31\x1a\xf0\x68\xce\x7a\x3f", "5j2FgcEB7vzAy7"},

	{"base64", "", ""},
	{"base64", "hello world", "aGVsbG8gd29ybGQ="},

	{"base64URL", "", ""},
	{"base64URL", "www.gouguoyin.cn", "d3d3LmdvdWd1b3lpbi5jbg=="},

	{"base85", "", ""},
	{"base85", "hello world", "BOu!rD]j7BEbo7"},

	{"base91", "", ""},
	{"base91", "1", "xA"},
	{"base91", "1234567890", "QztEml0o[2;(A"},
	{"base91", "hello world", "TPwJh>Io2Tv!lE"},
	{"base91", "\x14\xfb\x9c\x03", "Q<c[A"},
	{"base91", "\xf2\x8e\x88\x31\x1a\xf0\x68\xce\x7a\x3f", "EquimSayaka~A"},
	{"base91", "\x35\x5e\x56\xe0\xc6\x29\x38\xf4\x81\x00\xab\x81\x7e\xd7\x08\x95\x62\x20\xa7\xda\x64\xa2\xce\xb3\xc5", "~_1H=x_t{|$AjJX(nMFdjL~:?1b3HgM"},

	{"base100", "", ""},
	{"base100", "1234567890", "🐨🐩🐪🐫🐬🐭🐮🐯🐰🐧"},
	{"base100", "hello world", "👟👜👣👣👦🐗👮👦👩👣👛"},
}

func TestBaseX_Encode_ToString(t *testing.T) {
	for index, test := range baseXTest {
		e := Encode.FromString(test.input)

		switch test.baseX {
		case "hex":
			e = e.ByHex()
		case "base16":
			e = e.ByBase16()
		case "base32":
			e = e.ByBase32()
		case "base58":
			e = e.ByBase58()
		case "base62":
			e = e.ByBase62()
		case "base64":
			e = e.ByBase64()
		case "base64URL":
			e = e.ByBase64URL()
		case "base85":
			e = e.ByBase85()
		case "base91":
			e = e.ByBase91()
		case "base100":
			e = e.ByBase100()
		}

		assert.Nil(t, e.Error)
		assert.Equal(t, test.output, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestBaseX_Decode_ToString(t *testing.T) {
	for index, test := range baseXTest {
		d := Decode.FromString(test.output)

		switch test.baseX {
		case "hex":
			d = d.ByHex()
		case "base16":
			d = d.ByBase16()
		case "base32":
			d = d.ByBase32()
		case "base58":
			d = d.ByBase58()
		case "base62":
			d = d.ByBase62()
		case "base64":
			d = d.ByBase64()
		case "base64URL":
			d = d.ByBase64URL()
		case "base85":
			d = d.ByBase85()
		case "base91":
			d = d.ByBase91()
		case "base100":
			d = d.ByBase100()
		}

		assert.Nil(t, d.Error)
		assert.Equal(t, test.input, d.ToString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestBaseX_Encode_ToBytes(t *testing.T) {
	for index, test := range baseXTest {
		e := Encode.FromBytes([]byte(test.input))

		switch test.baseX {
		case "hex":
			e = e.ByHex()
		case "base16":
			e = e.ByBase16()
		case "base32":
			e = e.ByBase32()
		case "base58":
			e = e.ByBase58()
		case "base62":
			e = e.ByBase62()
		case "base64":
			e = e.ByBase64()
		case "base64URL":
			e = e.ByBase64URL()
		case "base85":
			e = e.ByBase85()
		case "base91":
			e = e.ByBase91()
		case "base100":
			e = e.ByBase100()
		}

		assert.Nil(t, e.Error)
		assert.Equal(t, []byte(test.output), e.ToBytes(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestBaseX_Decode_ToBytes(t *testing.T) {
	for index, test := range baseXTest {
		d := Decode.FromBytes([]byte(test.output))

		switch test.baseX {
		case "hex":
			d = d.ByHex()
		case "base16":
			d = d.ByBase16()
		case "base32":
			d = d.ByBase32()
		case "base58":
			d = d.ByBase58()
		case "base62":
			d = d.ByBase62()
		case "base64":
			d = d.ByBase64()
		case "base64URL":
			d = d.ByBase64URL()
		case "base85":
			d = d.ByBase85()
		case "base91":
			d = d.ByBase91()
		case "base100":
			d = d.ByBase100()
		}

		assert.Nil(t, d.Error)
		assert.Equal(t, []byte(test.input), d.ToBytes(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestBaseX_Ciphertext_Error(t *testing.T) {

	tests := []struct {
		baseX string
		input string // 输入值
		error error  // 期望值
	}{
		{"hex", "xxxxxx", invalidCiphertextError("hex")},
		{"base16", "xxxxxx", invalidCiphertextError("base16")},
		{"base32", "xxxxxx", invalidCiphertextError("base32")},
		{"base62", "~_1H=x_t{ |$AjJX(nMFdjL~:?1b3HgM", invalidCiphertextError("base62")},
		{"base64", "xxxxxx", invalidCiphertextError("base64")},
		{"base64URL", "xxxxxx", invalidCiphertextError("base64URL")},
		{"base85", "xxxxxx", invalidCiphertextError("base85")},
		{"base91", "'", invalidCiphertextError("base91")},
		{"base100", "\\", invalidCiphertextError("base100")},
		{"base100", "~_1H=x_t{ |$AjJX(nMFdjL~:?1b3HgM", invalidCiphertextError("base100")},
	}

	for index, test := range tests {
		d := Decode.FromString(test.input)
		switch test.baseX {
		case "hex":
			d = d.ByHex()
		case "base16":
			d = d.ByBase16()
		case "base32":
			d = d.ByBase32()
		case "base58":
			d = d.ByBase58()
		case "base62":
			d = d.ByBase62()
		case "base64":
			d = d.ByBase64()
		case "base64URL":
			d = d.ByBase64URL()
		case "base85":
			d = d.ByBase85()
		case "base91":
			d = d.ByBase91()
		case "base100":
			d = d.ByBase100()
		}
		assert.Equal(t, invalidCiphertextError(test.baseX), d.Error, "Current test index is "+strconv.Itoa(index))
	}
}
