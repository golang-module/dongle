package dongle

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	desKey = []byte("12345678")
	desIV  = []byte("12345678")
)

var desTests = []struct {
	mode     cipherMode
	padding  cipherPadding
	input    string
	toHex    string
	toBase64 string
}{
	{CBC, PKCS7, "", "", ""},
	{CBC, No, "hello world, go!", "0b2a92e81fb49ce1b161ab1a2e35d7ce", "CyqS6B+0nOGxYasaLjXXzg=="},
	{CBC, Empty, "hello world", "0b2a92e81fb49ce1dcb45f808e7b4075", "CyqS6B+0nOHctF+AjntAdQ=="},
	{CBC, Zero, "hello world", "0b2a92e81fb49ce1e02ecfa986df2ec8", "CyqS6B+0nOHgLs+pht8uyA=="},
	{CBC, PKCS5, "hello world", "0b2a92e81fb49ce1a43266aacaea7b81", "CyqS6B+0nOGkMmaqyup7gQ=="},
	{CBC, PKCS7, "hello world", "0b2a92e81fb49ce1a43266aacaea7b81", "CyqS6B+0nOGkMmaqyup7gQ=="},
	{CBC, AnsiX923, "hello world", "0b2a92e81fb49ce1c450d98b1e527b4a", "CyqS6B+0nOHEUNmLHlJ7Sg=="},
	{CBC, ISO97971, "hello world", "0b2a92e81fb49ce1cd126e4be2f294cd", "CyqS6B+0nOHNEm5L4vKUzQ=="},

	{CFB, PKCS7, "", "", ""},
	{CFB, No, "hello world, go!", "feb56ee417f5fbe61a4e5700b93c340b", "/rVu5Bf1++YaTlcAuTw0Cw=="},
	{CFB, Empty, "hello world", "feb56ee417f5fbe61a4e570cb97b7b0a", "/rVu5Bf1++YaTlcMuXt7Cg=="},
	{CFB, Zero, "hello world", "feb56ee417f5fbe61a4e572c995b5b2a", "/rVu5Bf1++YaTlcsmVtbKg=="},
	{CFB, PKCS5, "hello world", "feb56ee417f5fbe61a4e57299c5e5e2f", "/rVu5Bf1++YaTlcpnF5eLw=="},
	{CFB, PKCS7, "hello world", "feb56ee417f5fbe61a4e57299c5e5e2f", "/rVu5Bf1++YaTlcpnF5eLw=="},
	{CFB, AnsiX923, "hello world", "feb56ee417f5fbe61a4e572c995b5b2f", "/rVu5Bf1++YaTlcsmVtbLw=="},
	{CFB, ISO97971, "hello world", "feb56ee417f5fbe61a4e57ac995b5b2a", "/rVu5Bf1++YaTlesmVtbKg=="},

	{OFB, PKCS7, "", "", ""},
	{OFB, No, "hello world, go!", "feb56ee417f5fbe69b19fcf309087d8b", "/rVu5Bf1++abGfzzCQh9iw=="},
	{OFB, Empty, "hello world", "feb56ee417f5fbe69b19fcff094f328a", "/rVu5Bf1++abGfz/CU8yig=="},
	{OFB, Zero, "hello world", "feb56ee417f5fbe69b19fcdf296f12aa", "/rVu5Bf1++abGfzfKW8Sqg=="},
	{OFB, PKCS5, "hello world", "feb56ee417f5fbe69b19fcda2c6a17af", "/rVu5Bf1++abGfzaLGoXrw=="},
	{OFB, PKCS7, "hello world", "feb56ee417f5fbe69b19fcda2c6a17af", "/rVu5Bf1++abGfzaLGoXrw=="},
	{OFB, AnsiX923, "hello world", "feb56ee417f5fbe69b19fcdf296f12af", "/rVu5Bf1++abGfzfKW8Srw=="},
	{OFB, ISO97971, "hello world", "feb56ee417f5fbe69b19fc5f296f12aa", "/rVu5Bf1++abGfxfKW8Sqg=="},

	{CTR, PKCS7, "", "", ""},
	{CTR, No, "hello world, go!", "feb56ee417f5fbe6337e8b3af91e9929", "/rVu5Bf1++Yzfos6+R6ZKQ=="},
	{CTR, Empty, "hello world", "feb56ee417f5fbe6337e8b36f959d628", "/rVu5Bf1++Yzfos2+VnWKA=="},
	{CTR, Zero, "hello world", "feb56ee417f5fbe6337e8b16d979f608", "/rVu5Bf1++YzfosW2Xn2CA=="},
	{CTR, PKCS5, "hello world", "feb56ee417f5fbe6337e8b13dc7cf30d", "/rVu5Bf1++YzfosT3HzzDQ=="},
	{CTR, PKCS7, "hello world", "feb56ee417f5fbe6337e8b13dc7cf30d", "/rVu5Bf1++YzfosT3HzzDQ=="},
	{CTR, AnsiX923, "hello world", "feb56ee417f5fbe6337e8b16d979f60d", "/rVu5Bf1++YzfosW2Xn2DQ=="},
	{CTR, ISO97971, "hello world", "feb56ee417f5fbe6337e8b96d979f608", "/rVu5Bf1++YzfouW2Xn2CA=="},

	{ECB, PKCS7, "", "", ""},
	{ECB, No, "hello world, go!", "28dba02eb5f6dd47e5a66298af31d4ca", "KNugLrX23UflpmKYrzHUyg=="},
	{ECB, Empty, "hello world", "28dba02eb5f6dd47a66e845b5ef12e13", "KNugLrX23UemboRbXvEuEw=="},
	{ECB, Zero, "hello world", "28dba02eb5f6dd476042daebfa59687a", "KNugLrX23UdgQtrr+lloeg=="},
	{ECB, PKCS5, "hello world", "28dba02eb5f6dd475d82e3681c83bb77", "KNugLrX23UddguNoHIO7dw=="},
	{ECB, PKCS7, "hello world", "28dba02eb5f6dd475d82e3681c83bb77", "KNugLrX23UddguNoHIO7dw=="},
	{ECB, AnsiX923, "hello world", "28dba02eb5f6dd47d33696d839c770b2", "KNugLrX23UfTNpbYOcdwsg=="},
	{ECB, ISO97971, "hello world", "28dba02eb5f6dd4706b5c56593dcbe2c", "KNugLrX23UcGtcVlk9y+LA=="},
}

func TestDes_Encrypt_String(t *testing.T) {
	for index, test := range desTests {
		raw := Decode.FromString(test.toHex).ByHex().ToString()
		e := Encrypt.FromString(test.input).ByDes(getCipher(test.mode, test.padding, desKey, desIV))

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

func TestDes_Decrypt_String(t *testing.T) {
	for index, test := range desTests {
		raw := Decode.FromString(test.toHex).ByHex().ToString()
		e := Decrypt.FromRawString(raw).ByDes(getCipher(test.mode, test.padding, desKey, desIV))

		t.Run(fmt.Sprintf(string(test.mode)+"_"+string(test.padding)+"_raw_test_%d", index), func(t *testing.T) {
			assert.Nil(t, e.Error)
			assert.Equal(t, test.input, e.ToString())
			assert.Equal(t, test.input, fmt.Sprintf("%s", e))
		})
	}

	for index, test := range desTests {
		e := Decrypt.FromHexString(test.toHex).ByDes(getCipher(test.mode, test.padding, desKey, desIV))

		t.Run(fmt.Sprintf(string(test.mode)+"_"+string(test.padding)+"_hex_test_%d", index), func(t *testing.T) {
			assert.Nil(t, e.Error)
			assert.Equal(t, test.input, e.ToString())
			assert.Equal(t, test.input, fmt.Sprintf("%s", e))
		})
	}

	for index, test := range desTests {
		e := Decrypt.FromBase64String(test.toBase64).ByDes(getCipher(test.mode, test.padding, desKey, desIV))

		t.Run(fmt.Sprintf(string(test.mode)+"_"+string(test.padding)+"_base64_test_%d", index), func(t *testing.T) {
			assert.Nil(t, e.Error)
			assert.Equal(t, test.input, e.ToString())
			assert.Equal(t, test.input, fmt.Sprintf("%s", e))
		})
	}
}

func TestDes_Encrypt_Bytes(t *testing.T) {
	for index, test := range desTests {
		raw := Decode.FromBytes([]byte(test.toHex)).ByHex().ToBytes()
		e := Encrypt.FromBytes([]byte(test.input)).ByDes(getCipher(test.mode, test.padding, desKey, desIV))

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

func TestDes_Decrypt_Bytes(t *testing.T) {
	for index, test := range desTests {
		raw := Decode.FromBytes([]byte(test.toHex)).ByHex().ToBytes()
		e := Decrypt.FromRawBytes(raw).ByDes(getCipher(test.mode, test.padding, desKey, desIV))

		t.Run(fmt.Sprintf(string(test.mode)+"_"+string(test.padding)+"_raw_test_%d", index), func(t *testing.T) {
			assert.Nil(t, e.Error)
			assert.Equal(t, []byte(test.input), e.ToBytes())
		})
	}

	for index, test := range desTests {
		e := Decrypt.FromHexBytes([]byte(test.toHex)).ByDes(getCipher(test.mode, test.padding, desKey, desIV))

		t.Run(fmt.Sprintf(string(test.mode)+"_"+string(test.padding)+"_hex_test_%d", index), func(t *testing.T) {
			assert.Nil(t, e.Error)
			assert.Equal(t, []byte(test.input), e.ToBytes())
		})
	}

	for index, test := range desTests {
		e := Decrypt.FromBase64Bytes([]byte(test.toBase64)).ByDes(getCipher(test.mode, test.padding, desKey, desIV))

		t.Run(fmt.Sprintf(string(test.mode)+"_"+string(test.padding)+"_base64_test_%d", index), func(t *testing.T) {
			assert.Nil(t, e.Error)
			assert.Equal(t, []byte(test.input), e.ToBytes())
		})
	}
}

func TestDes_Key_Error(t *testing.T) {
	key, iv := []byte("xxxx"), desIV

	e := Encrypt.FromString("hello world").ByDes(getCipher(CBC, PKCS7, key, iv))
	assert.Equal(t, desError.KeyError(), e.Error)

	d := Decrypt.FromHexString("0b2a92e81fb49ce1a43266aacaea7b81").ByDes(getCipher(CBC, PKCS7, key, iv))
	assert.Equal(t, desError.KeyError(), d.Error)
}

func TestDes_IV_Error(t *testing.T) {
	key, iv := desKey, []byte("xxxx")

	e := Encrypt.FromString("hello world").ByDes(getCipher(OFB, PKCS7, key, iv))
	assert.Equal(t, desError.IvError(), e.Error)

	d := Decrypt.FromHexString("0b2a92e81fb49ce1a43266aacaea7b81").ByDes(getCipher(CBC, PKCS7, key, iv))
	assert.Equal(t, desError.IvError(), d.Error)
}

func TestDes_Src_Error(t *testing.T) {
	e := Encrypt.FromString("hello world").ByDes(getCipher(CFB, No, desKey, desIV))
	assert.Equal(t, desError.SrcError(), e.Error)

	d := Decrypt.FromHexString("68656c6c6f20776f726c64").ByDes(getCipher(CBC, No, desKey, desIV))
	assert.Equal(t, desError.SrcError(), d.Error)
}

func TestDes_Decoding_Error(t *testing.T) {
	d1 := Decrypt.FromHexBytes([]byte("xxxx")).ByDes(getCipher(CTR, Zero, desKey, desIV))
	assert.Equal(t, decodeError.ModeError("hex"), d1.Error)

	d2 := Decrypt.FromBase64Bytes([]byte("xxxxxx")).ByDes(getCipher(CFB, PKCS7, desKey, desIV))
	assert.Equal(t, decodeError.ModeError("base64"), d2.Error)
}
