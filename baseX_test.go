package dongle

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

var baseXTests = []struct {
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

func TestBaseX_Encode_String(t *testing.T) {
	for index, test := range baseXTests {
		e1 := Encode.FromString(test.input)
		e2 := Encode.FromBytes([]byte(test.input))

		switch test.baseX {
		case "hex":
			e1 = e1.ByHex()
			e2 = e2.ByHex()
		case "base16":
			e1 = e1.ByBase16()
			e2 = e2.ByBase16()
		case "base32":
			e1 = e1.ByBase32()
			e2 = e2.ByBase32()
		case "base58":
			e1 = e1.ByBase58()
			e2 = e2.ByBase58()
		case "base62":
			e1 = e1.ByBase62()
			e2 = e2.ByBase62()
		case "base64":
			e1 = e1.ByBase64()
			e2 = e2.ByBase64()
		case "base64URL":
			e1 = e1.ByBase64URL()
			e2 = e2.ByBase64URL()
		case "base85":
			e1 = e1.ByBase85()
			e2 = e2.ByBase85()
		case "base91":
			e1 = e1.ByBase91()
			e2 = e2.ByBase91()
		case "base100":
			e1 = e1.ByBase100()
			e2 = e2.ByBase100()
		}

		t.Run(fmt.Sprintf(test.baseX+"_e1_test_%d", index), func(t *testing.T) {
			assert.Nil(t, e1.Error)
			assert.Equal(t, test.output, e1.ToString())
			assert.Equal(t, test.output, fmt.Sprintf("%s", e1))
		})

		t.Run(fmt.Sprintf(test.baseX+"_e2_%d", index), func(t *testing.T) {
			assert.Nil(t, e2.Error)
			assert.Equal(t, []byte(test.output), e2.ToBytes())
		})
	}
}

func TestBaseX_Decode_String(t *testing.T) {
	for index, test := range baseXTests {
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

func TestBaseX_Encode_Bytes(t *testing.T) {
	for index, test := range baseXTests {
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

func TestBaseX_Decode_Bytes(t *testing.T) {
	for index, test := range baseXTests {
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

		t.Run(fmt.Sprintf(test.baseX+"_test_%d", index), func(t *testing.T) {
			assert.Nil(t, d.Error)
			assert.Equal(t, []byte(test.input), d.ToBytes(), "Current test index is "+strconv.Itoa(index))
		})
	}
}

func TestBaseX_Decoding_Error(t *testing.T) {
	tests := []struct {
		baseX string
		input string // 输入值
		error error  // 期望值
	}{
		{"hex", "xxxxxx", invalidDecodingError("hex")},
		{"base16", "xxxxxx", invalidDecodingError("base16")},
		{"base32", "xxxxxx", invalidDecodingError("base32")},
		{"base62", "~_1H=x_t{ |$AjJX(nMFdjL~:?1b3HgM", invalidDecodingError("base62")},
		{"base64", "xxxxxx", invalidDecodingError("base64")},
		{"base64URL", "xxxxxx", invalidDecodingError("base64URL")},
		{"base85", "xxxxxx", invalidDecodingError("base85")},
		{"base91", "'", invalidDecodingError("base91")},
		{"base91", "-", invalidDecodingError("base91")},
		{"base100", "\\", invalidDecodingError("base100")},
		{"base100", "~_1H=x_t{ |$AjJX(nMFdjL~:?1b3HgM", invalidDecodingError("base100")},
	}

	for index, test := range tests {
		d1 := Decode.FromString(test.input)
		d2 := Decode.FromBytes([]byte(test.input))

		switch test.baseX {
		case "hex":
			d1 = d1.ByHex()
			d2 = d2.ByHex()
		case "base16":
			d1 = d1.ByBase16()
			d2 = d2.ByBase16()
		case "base32":
			d1 = d1.ByBase32()
			d2 = d2.ByBase32()
		case "base58":
			d1 = d1.ByBase58()
			d2 = d2.ByBase58()
		case "base62":
			d1 = d1.ByBase62()
			d2 = d2.ByBase62()
		case "base64":
			d1 = d1.ByBase64()
			d2 = d2.ByBase64()
		case "base64URL":
			d1 = d1.ByBase64URL()
			d2 = d2.ByBase64URL()
		case "base85":
			d1 = d1.ByBase85()
			d2 = d2.ByBase85()
		case "base91":
			d1 = d1.ByBase91()
			d2 = d2.ByBase91()
		case "base100":
			d1 = d1.ByBase100()
			d2 = d2.ByBase100()
		}

		t.Run(fmt.Sprintf(test.baseX+"_d1_test_%d", index), func(t *testing.T) {
			assert.Equal(t, invalidDecodingError(test.baseX), d1.Error)
		})

		t.Run(fmt.Sprintf(test.baseX+"_d2_test_%d", index), func(t *testing.T) {
			assert.Equal(t, invalidDecodingError(test.baseX), d2.Error)
		})
	}
}
