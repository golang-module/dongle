package dongle

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	blowfishKey = []byte("0123456789abcdef")
	blowfishIV  = []byte("12345678")
)

var blowfishTests = []struct {
	mode     cipherMode
	padding  cipherPadding
	input    string
	toHex    string
	toBase64 string
}{
	{CBC, PKCS7, "", "", ""},
	{CBC, No, "hello world, go!", "92c3724d720b5998615b123dd9774323", "ksNyTXILWZhhWxI92XdDIw=="},
	{CBC, Empty, "hello world", "92c3724d720b59989fcf8696627e6c0e", "ksNyTXILWZifz4aWYn5sDg=="},
	{CBC, Zero, "hello world", "92c3724d720b599812a2da1e70f954ff", "ksNyTXILWZgSotoecPlU/w=="},
	{CBC, PKCS5, "hello world", "92c3724d720b59982d21aab9a7bb9f40", "ksNyTXILWZgtIaq5p7ufQA=="},
	{CBC, PKCS7, "hello world", "92c3724d720b59982d21aab9a7bb9f40", "ksNyTXILWZgtIaq5p7ufQA=="},
	{CBC, AnsiX923, "hello world", "92c3724d720b59986113306d58791c7f", "ksNyTXILWZhhEzBtWHkcfw=="},
	{CBC, ISO97971, "hello world", "92c3724d720b599836aa40b4661e1e29", "ksNyTXILWZg2qkC0Zh4eKQ=="},

	{CFB, PKCS7, "", "", ""},
	{CFB, No, "hello world, go!", "f0ea2f993584f026d070683357c9205c", "8OovmTWE8CbQcGgzV8kgXA=="},
	{CFB, Empty, "hello world", "f0ea2f993584f026d070683f578e6f5d", "8OovmTWE8CbQcGg/V45vXQ=="},
	{CFB, Zero, "hello world", "f0ea2f993584f026d070681f77ae4f7d", "8OovmTWE8CbQcGgfd65PfQ=="},
	{CFB, PKCS5, "hello world", "f0ea2f993584f026d070681a72ab4a78", "8OovmTWE8CbQcGgacqtKeA=="},
	{CFB, PKCS7, "hello world", "f0ea2f993584f026d070681a72ab4a78", "8OovmTWE8CbQcGgacqtKeA=="},
	{CFB, AnsiX923, "hello world", "f0ea2f993584f026d070681f77ae4f78", "8OovmTWE8CbQcGgfd65PeA=="},
	{CFB, ISO97971, "hello world", "f0ea2f993584f026d070689f77ae4f7d", "8OovmTWE8CbQcGifd65PfQ=="},

	{OFB, PKCS7, "", "", ""},
	{OFB, No, "hello world, go!", "f0ea2f993584f0267e628f1ec7f5d21d", "8OovmTWE8CZ+Yo8ex/XSHQ=="},
	{OFB, Empty, "hello world", "f0ea2f993584f0267e628f12c7b29d1c", "8OovmTWE8CZ+Yo8Sx7KdHA=="},
	{OFB, Zero, "hello world", "f0ea2f993584f0267e628f32e792bd3c", "8OovmTWE8CZ+Yo8y55K9PA=="},
	{OFB, PKCS5, "hello world", "f0ea2f993584f0267e628f37e297b839", "8OovmTWE8CZ+Yo834pe4OQ=="},
	{OFB, PKCS7, "hello world", "f0ea2f993584f0267e628f37e297b839", "8OovmTWE8CZ+Yo834pe4OQ=="},
	{OFB, AnsiX923, "hello world", "f0ea2f993584f0267e628f32e792bd39", "8OovmTWE8CZ+Yo8y55K9OQ=="},
	{OFB, ISO97971, "hello world", "f0ea2f993584f0267e628fb2e792bd3c", "8OovmTWE8CZ+Yo+y55K9PA=="},

	{CTR, PKCS7, "", "", ""},
	{CTR, No, "hello world, go!", "f0ea2f993584f026d6f92ffda9d639be", "8OovmTWE8CbW+S/9qdY5vg=="},
	{CTR, Empty, "hello world", "f0ea2f993584f026d6f92ff1a99176bf", "8OovmTWE8CbW+S/xqZF2vw=="},
	{CTR, Zero, "hello world", "f0ea2f993584f026d6f92fd189b1569f", "8OovmTWE8CbW+S/RibFWnw=="},
	{CTR, PKCS5, "hello world", "f0ea2f993584f026d6f92fd48cb4539a", "8OovmTWE8CbW+S/UjLRTmg=="},
	{CTR, PKCS7, "hello world", "f0ea2f993584f026d6f92fd48cb4539a", "8OovmTWE8CbW+S/UjLRTmg=="},
	{CTR, AnsiX923, "hello world", "f0ea2f993584f026d6f92fd189b1569a", "8OovmTWE8CbW+S/RibFWmg=="},
	{CTR, ISO97971, "hello world", "f0ea2f993584f026d6f92f5189b1569f", "8OovmTWE8CbW+S9RibFWnw=="},

	{ECB, PKCS7, "", "", ""},
	{ECB, No, "hello world, go!", "954ddb4b65c05c90112006edaeda25f5", "lU3bS2XAXJARIAbtrtol9Q=="},
	{ECB, Empty, "hello world", "954ddb4b65c05c9002d465bae4ae9dc1", "lU3bS2XAXJAC1GW65K6dwQ=="},
	{ECB, Zero, "hello world", "954ddb4b65c05c90c8cb0c6bbb3ec912", "lU3bS2XAXJDIywxruz7JEg=="},
	{ECB, PKCS5, "hello world", "954ddb4b65c05c9022457a6f07c5859c", "lU3bS2XAXJAiRXpvB8WFnA=="},
	{ECB, PKCS7, "hello world", "954ddb4b65c05c9022457a6f07c5859c", "lU3bS2XAXJAiRXpvB8WFnA=="},
	{ECB, AnsiX923, "hello world", "954ddb4b65c05c90b4c296978984f345", "lU3bS2XAXJC0wpaXiYTzRQ=="},
	{ECB, ISO97971, "hello world", "954ddb4b65c05c907227aaa09f2e4ddb", "lU3bS2XAXJByJ6qgny5N2w=="},
}

func TestBlowfish_Encrypt_String(t *testing.T) {
	for index, test := range blowfishTests {
		raw := Decode.FromString(test.toHex).ByHex().ToString()
		e := Encrypt.FromString(test.input).ByBlowfish(getCipher(test.mode, test.padding, blowfishKey, blowfishIV))

		t.Run(fmt.Sprintf(string(test.mode)+"_"+string(test.padding)+"_raw_test_%d", index), func(t *testing.T) {
			assert.Nil(t, e.Error)
			assert.Equal(t, raw, e.ToRawString())
		})
		t.Run(fmt.Sprintf(string(test.mode)+"_"+string(test.padding)+"_hex_test_%d", index), func(t *testing.T) {
			assert.Nil(t, e.Error)
			assert.Equal(t, test.toHex, e.ToHexString())
		})
		t.Run(fmt.Sprintf(string(test.mode)+"_"+string(test.padding)+"_base64_test_%d", index), func(t *testing.T) {
			assert.Nil(t, e.Error)
			assert.Equal(t, test.toBase64, e.ToBase64String())
		})
	}
}

func TestBlowfish_Decrypt_String(t *testing.T) {
	for index, test := range blowfishTests {
		raw := Decode.FromString(test.toHex).ByHex().ToString()
		e := Decrypt.FromRawString(raw).ByBlowfish(getCipher(test.mode, test.padding, blowfishKey, blowfishIV))

		t.Run(fmt.Sprintf(string(test.mode)+"_"+string(test.padding)+"_raw_test_%d", index), func(t *testing.T) {
			assert.Nil(t, e.Error)
			assert.Equal(t, test.input, e.ToString())
			assert.Equal(t, test.input, fmt.Sprintf("%s", e))
		})
	}

	for index, test := range blowfishTests {
		e := Decrypt.FromHexString(test.toHex).ByBlowfish(getCipher(test.mode, test.padding, blowfishKey, blowfishIV))

		t.Run(fmt.Sprintf(string(test.mode)+"_"+string(test.padding)+"_hex_test_%d", index), func(t *testing.T) {
			assert.Nil(t, e.Error)
			assert.Equal(t, test.input, e.ToString())
			assert.Equal(t, test.input, fmt.Sprintf("%s", e))
		})
	}

	for index, test := range blowfishTests {
		e := Decrypt.FromBase64String(test.toBase64).ByBlowfish(getCipher(test.mode, test.padding, blowfishKey, blowfishIV))

		t.Run(fmt.Sprintf(string(test.mode)+"_"+string(test.padding)+"_base64_test_%d", index), func(t *testing.T) {
			assert.Nil(t, e.Error)
			assert.Equal(t, test.input, e.ToString())
			assert.Equal(t, test.input, fmt.Sprintf("%s", e))
		})
	}
}

func TestBlowfish_Encrypt_Bytes(t *testing.T) {
	for index, test := range blowfishTests {
		raw := Decode.FromBytes([]byte(test.toHex)).ByHex().ToBytes()
		e := Encrypt.FromBytes([]byte(test.input)).ByBlowfish(getCipher(test.mode, test.padding, blowfishKey, blowfishIV))

		t.Run(fmt.Sprintf(string(test.mode)+"_"+string(test.padding)+"_raw_test_%d", index), func(t *testing.T) {
			assert.Nil(t, e.Error)
			assert.Equal(t, raw, e.ToRawBytes())
		})
		t.Run(fmt.Sprintf(string(test.mode)+"_"+string(test.padding)+"_hex_test_%d", index), func(t *testing.T) {
			assert.Nil(t, e.Error)
			assert.Equal(t, []byte(test.toHex), e.ToHexBytes())
		})
		t.Run(fmt.Sprintf(string(test.mode)+"_"+string(test.padding)+"_base64_test_%d", index), func(t *testing.T) {
			assert.Nil(t, e.Error)
			assert.Equal(t, []byte(test.toBase64), e.ToBase64Bytes())
		})
	}
}

func TestBlowfish_Decrypt_Bytes(t *testing.T) {
	for index, test := range blowfishTests {
		raw := Decode.FromBytes([]byte(test.toHex)).ByHex().ToBytes()
		e := Decrypt.FromRawBytes(raw).ByBlowfish(getCipher(test.mode, test.padding, blowfishKey, blowfishIV))

		t.Run(fmt.Sprintf(string(test.mode)+"_"+string(test.padding)+"_raw_test_%d", index), func(t *testing.T) {
			assert.Nil(t, e.Error)
			assert.Equal(t, []byte(test.input), e.ToBytes())
		})
	}

	for index, test := range blowfishTests {
		e := Decrypt.FromHexBytes([]byte(test.toHex)).ByBlowfish(getCipher(test.mode, test.padding, blowfishKey, blowfishIV))

		t.Run(fmt.Sprintf(string(test.mode)+"_"+string(test.padding)+"_hex_test_%d", index), func(t *testing.T) {
			assert.Nil(t, e.Error)
			assert.Equal(t, []byte(test.input), e.ToBytes())
		})
	}

	for index, test := range blowfishTests {
		e := Decrypt.FromBase64Bytes([]byte(test.toBase64)).ByBlowfish(getCipher(test.mode, test.padding, blowfishKey, blowfishIV))

		t.Run(fmt.Sprintf(string(test.mode)+"_"+string(test.padding)+"_base64_test_%d", index), func(t *testing.T) {
			assert.Nil(t, e.Error)
			assert.Equal(t, []byte(test.input), e.ToBytes())
		})
	}
}

func TestBlowfish_Key_Error(t *testing.T) {
	key, iv := []byte(""), blowfishIV

	e := Encrypt.FromString("hello world").ByBlowfish(getCipher(CBC, PKCS7, key, iv))
	assert.Equal(t, blowfishError.KeyError([]byte("")), e.Error)

	d := Decrypt.FromHexString("c1e9b4529aac9793010f4677f6358efe").ByBlowfish(getCipher(CBC, PKCS7, key, iv))
	assert.Equal(t, blowfishError.KeyError([]byte("")), d.Error)
}

func TestBlowfish_IV_Error(t *testing.T) {
	key, iv := blowfishKey, []byte("xxxx")
	err := NewBlowfishError()
	e := Encrypt.FromString("hello world").ByBlowfish(getCipher(OFB, PKCS7, key, iv))
	assert.Equal(t, err.IvError(iv), e.Error)

	d := Decrypt.FromHexString("c1e9b4529aac9793010f4677f6358efec1e9b4529aac9793010f4677f6358efe").ByBlowfish(getCipher(CBC, PKCS7, key, iv))
	assert.Equal(t, err.IvError(iv), d.Error)
}

func TestBlowfish_Src_Error(t *testing.T) {
	e := Encrypt.FromString("hello world").ByBlowfish(getCipher(CFB, No, blowfishKey, blowfishIV))
	assert.Equal(t, blowfishError.SrcError(), e.Error)

	d := Decrypt.FromHexString("68656c6c6f20776f726c64").ByBlowfish(getCipher(CBC, No, blowfishKey, blowfishIV))
	assert.Equal(t, blowfishError.SrcError(), d.Error)
}

func TestBlowfish_Decoding_Error(t *testing.T) {
	d1 := Decrypt.FromHexString("xxxx").ByBlowfish(getCipher(CTR, Zero, blowfishKey, blowfishIV))
	assert.Equal(t, decodeError.ModeError("hex"), d1.Error)
	d2 := Decrypt.FromHexBytes([]byte("xxxx")).ByBlowfish(getCipher(CTR, Zero, blowfishKey, blowfishIV))
	assert.Equal(t, decodeError.ModeError("hex"), d2.Error)

	d3 := Decrypt.FromBase64String("xxxxxx").ByBlowfish(getCipher(CFB, PKCS7, blowfishKey, blowfishIV))
	assert.Equal(t, decodeError.ModeError("base64"), d3.Error)
	d4 := Decrypt.FromBase64Bytes([]byte("xxxxxx")).ByBlowfish(getCipher(CFB, PKCS7, blowfishKey, blowfishIV))
	assert.Equal(t, decodeError.ModeError("base64"), d4.Error)
}
