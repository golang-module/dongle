package dongle

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	aesKey = []byte("0123456789abcdef")
	aesIV  = []byte("0123456789abcdef")
)

var aesTests = []struct {
	mode     cipherMode
	padding  cipherPadding
	input    string
	toHex    string
	toBase64 string
}{
	{CBC, PKCS7, "", "", ""},
	{CBC, No, "hello world, go!", "77aa39926f9b2f3f22254bfd422fa75d", "d6o5km+bLz8iJUv9Qi+nXQ=="},
	{CBC, Empty, "hello world", "956f2b73603b1a97024a7e5aa73ad7e2", "lW8rc2A7GpcCSn5apzrX4g=="},
	{CBC, Zero, "hello world", "889935b7a0c64b0333d713cafaee08fe", "iJk1t6DGSwMz1xPK+u4I/g=="},
	{CBC, PKCS5, "hello world", "c1e9b4529aac9793010f4677f6358efe", "wem0Upqsl5MBD0Z39jWO/g=="},
	{CBC, PKCS7, "hello world", "c1e9b4529aac9793010f4677f6358efe", "wem0Upqsl5MBD0Z39jWO/g=="},
	{CBC, AnsiX923, "hello world", "19267ee4e262addb6657840738d3b93d", "GSZ+5OJirdtmV4QHONO5PQ=="},
	{CBC, ISO97971, "hello world", "fd4fd654217962e0cfa285aa4da354e7", "/U/WVCF5YuDPooWqTaNU5w=="},

	{CFB, PKCS7, "", "", ""},
	{CFB, No, "hello world, go!", "1a1712e471fc8a6e72cb7c44596eda44", "GhcS5HH8im5yy3xEWW7aRA=="},
	{CFB, Empty, "hello world", "1a1712e471fc8a6e72cb7c4859299545", "GhcS5HH8im5yy3xIWSmVRQ=="},
	{CFB, Zero, "hello world", "1a1712e471fc8a6e72cb7c687909b565", "GhcS5HH8im5yy3xoeQm1ZQ=="},
	{CFB, PKCS5, "hello world", "1a1712e471fc8a6e72cb7c6d7c0cb060", "GhcS5HH8im5yy3xtfAywYA=="},
	{CFB, PKCS7, "hello world", "1a1712e471fc8a6e72cb7c6d7c0cb060", "GhcS5HH8im5yy3xtfAywYA=="},
	{CFB, AnsiX923, "hello world", "1a1712e471fc8a6e72cb7c687909b560", "GhcS5HH8im5yy3xoeQm1YA=="},
	{CFB, ISO97971, "hello world", "1a1712e471fc8a6e72cb7ce87909b565", "GhcS5HH8im5yy3zoeQm1ZQ=="},

	{OFB, PKCS7, "", "", ""},
	{OFB, No, "hello world, go!", "1a1712e471fc8a6e72cb7c44596eda44", "GhcS5HH8im5yy3xEWW7aRA=="},
	{OFB, Empty, "hello world", "1a1712e471fc8a6e72cb7c4859299545", "GhcS5HH8im5yy3xIWSmVRQ=="},
	{OFB, Zero, "hello world", "1a1712e471fc8a6e72cb7c687909b565", "GhcS5HH8im5yy3xoeQm1ZQ=="},
	{OFB, PKCS5, "hello world", "1a1712e471fc8a6e72cb7c6d7c0cb060", "GhcS5HH8im5yy3xtfAywYA=="},
	{OFB, PKCS7, "hello world", "1a1712e471fc8a6e72cb7c6d7c0cb060", "GhcS5HH8im5yy3xtfAywYA=="},
	{OFB, AnsiX923, "hello world", "1a1712e471fc8a6e72cb7c687909b560", "GhcS5HH8im5yy3xoeQm1YA=="},
	{OFB, ISO97971, "hello world", "1a1712e471fc8a6e72cb7ce87909b565", "GhcS5HH8im5yy3zoeQm1ZQ=="},

	{CTR, PKCS7, "", "", ""},
	{CTR, No, "hello world, go!", "1a1712e471fc8a6e72cb7c44596eda44", "GhcS5HH8im5yy3xEWW7aRA=="},
	{CTR, Empty, "hello world", "1a1712e471fc8a6e72cb7c4859299545", "GhcS5HH8im5yy3xIWSmVRQ=="},
	{CTR, Zero, "hello world", "1a1712e471fc8a6e72cb7c687909b565", "GhcS5HH8im5yy3xoeQm1ZQ=="},
	{CTR, PKCS5, "hello world", "1a1712e471fc8a6e72cb7c6d7c0cb060", "GhcS5HH8im5yy3xtfAywYA=="},
	{CTR, PKCS7, "hello world", "1a1712e471fc8a6e72cb7c6d7c0cb060", "GhcS5HH8im5yy3xtfAywYA=="},
	{CTR, AnsiX923, "hello world", "1a1712e471fc8a6e72cb7c687909b560", "GhcS5HH8im5yy3xoeQm1YA=="},
	{CTR, ISO97971, "hello world", "1a1712e471fc8a6e72cb7ce87909b565", "GhcS5HH8im5yy3zoeQm1ZQ=="},

	{ECB, PKCS7, "", "", ""},
	{ECB, No, "hello world, go!", "f82a4c0db7a82f70c7fa84c39fa7627b", "+CpMDbeoL3DH+oTDn6diew=="},
	{ECB, Empty, "hello world", "5868584176d2aca3c257592a9d8a020d", "WGhYQXbSrKPCV1kqnYoCDQ=="},
	{ECB, Zero, "hello world", "769c326290511c93bd59bba9c24d8904", "dpwyYpBRHJO9Wbupwk2JBA=="},
	{ECB, PKCS5, "hello world", "8169bed4ef49a8874559c5b200daade7", "gWm+1O9JqIdFWcWyANqt5w=="},
	{ECB, PKCS7, "hello world", "8169bed4ef49a8874559c5b200daade7", "gWm+1O9JqIdFWcWyANqt5w=="},
	{ECB, AnsiX923, "hello world", "1e33dfd0c42e440761065e2705e7f0f7", "HjPf0MQuRAdhBl4nBefw9w=="},
	{ECB, ISO97971, "hello world", "49153a0969bae6246d1f6f7e75628eea", "SRU6CWm65iRtH29+dWKO6g=="},
}

func TestAes_Encrypt_String(t *testing.T) {
	for index, test := range aesTests {
		raw := Decode.FromString(test.toHex).ByHex().ToString()
		e := Encrypt.FromString(test.input).ByAes(getCipher(test.mode, test.padding, aesKey, aesIV))

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

func TestAes_Decrypt_String(t *testing.T) {
	for index, test := range aesTests {
		raw := Decode.FromString(test.toHex).ByHex().ToString()
		e := Decrypt.FromRawString(raw).ByAes(getCipher(test.mode, test.padding, aesKey, aesIV))

		t.Run(fmt.Sprintf(string(test.mode)+"_"+string(test.padding)+"_raw_test_%d", index), func(t *testing.T) {
			assert.Nil(t, e.Error)
			assert.Equal(t, test.input, e.ToString())
			assert.Equal(t, test.input, fmt.Sprintf("%s", e))
		})
	}

	for index, test := range aesTests {
		e := Decrypt.FromHexString(test.toHex).ByAes(getCipher(test.mode, test.padding, aesKey, aesIV))

		t.Run(fmt.Sprintf(string(test.mode)+"_"+string(test.padding)+"_hex_test_%d", index), func(t *testing.T) {
			assert.Nil(t, e.Error)
			assert.Equal(t, test.input, e.ToString())
			assert.Equal(t, test.input, fmt.Sprintf("%s", e))
		})
	}

	for index, test := range aesTests {
		e := Decrypt.FromBase64String(test.toBase64).ByAes(getCipher(test.mode, test.padding, aesKey, aesIV))

		t.Run(fmt.Sprintf(string(test.mode)+"_"+string(test.padding)+"_base64_test_%d", index), func(t *testing.T) {
			assert.Nil(t, e.Error)
			assert.Equal(t, test.input, e.ToString())
			assert.Equal(t, test.input, fmt.Sprintf("%s", e))
		})
	}
}

func TestAes_Encrypt_Bytes(t *testing.T) {
	for index, test := range aesTests {
		raw := Decode.FromBytes([]byte(test.toHex)).ByHex().ToBytes()
		e := Encrypt.FromBytes([]byte(test.input)).ByAes(getCipher(test.mode, test.padding, aesKey, aesIV))

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

func TestAes_Decrypt_Bytes(t *testing.T) {
	for index, test := range aesTests {
		raw := Decode.FromBytes([]byte(test.toHex)).ByHex().ToBytes()
		e := Decrypt.FromRawBytes(raw).ByAes(getCipher(test.mode, test.padding, aesKey, aesIV))

		t.Run(fmt.Sprintf(string(test.mode)+"_"+string(test.padding)+"_raw_test_%d", index), func(t *testing.T) {
			assert.Nil(t, e.Error)
			assert.Equal(t, []byte(test.input), e.ToBytes())
		})
	}

	for index, test := range aesTests {
		e := Decrypt.FromHexBytes([]byte(test.toHex)).ByAes(getCipher(test.mode, test.padding, aesKey, aesIV))

		t.Run(fmt.Sprintf(string(test.mode)+"_"+string(test.padding)+"_hex_test_%d", index), func(t *testing.T) {
			assert.Nil(t, e.Error)
			assert.Equal(t, []byte(test.input), e.ToBytes())
		})
	}

	for index, test := range aesTests {
		e := Decrypt.FromBase64Bytes([]byte(test.toBase64)).ByAes(getCipher(test.mode, test.padding, aesKey, aesIV))

		t.Run(fmt.Sprintf(string(test.mode)+"_"+string(test.padding)+"_base64_test_%d", index), func(t *testing.T) {
			assert.Nil(t, e.Error)
			assert.Equal(t, []byte(test.input), e.ToBytes())
		})
	}
}

func TestAes_Key_Error(t *testing.T) {
	key, iv := []byte("xxxx"), aesIV

	e := Encrypt.FromString("hello world").ByAes(getCipher(CBC, PKCS7, key, iv))
	assert.Equal(t, aesError.KeyError(), e.Error)

	d := Decrypt.FromHexString("c1e9b4529aac9793010f4677f6358efe").ByAes(getCipher(CBC, PKCS7, key, iv))
	assert.Equal(t, aesError.KeyError(), d.Error)
}

func TestAes_IV_Error(t *testing.T) {
	key, iv := aesKey, []byte("xxxx")

	e := Encrypt.FromString("hello world").ByAes(getCipher(OFB, PKCS7, key, iv))
	assert.Equal(t, aesError.IvError(), e.Error)

	d := Decrypt.FromHexString("c1e9b4529aac9793010f4677f6358efec1e9b4529aac9793010f4677f6358efe").ByAes(getCipher(CBC, PKCS7, aesKey, iv))
	assert.Equal(t, aesError.IvError(), d.Error)
}

func TestAes_Src_Error(t *testing.T) {
	e := Encrypt.FromString("hello world").ByAes(getCipher(CFB, No, aesKey, aesIV))
	assert.Equal(t, aesError.SrcError(), e.Error)

	d := Decrypt.FromHexString("68656c6c6f20776f726c64").ByAes(getCipher(CBC, No, aesKey, aesIV))
	assert.Equal(t, aesError.SrcError(), d.Error)
}

func TestAes_Decoding_Error(t *testing.T) {
	d1 := Decrypt.FromHexString("xxxx").ByAes(getCipher(CTR, Zero, aesKey, aesIV))
	assert.Equal(t, decodeError.ModeError("hex"), d1.Error)
	d2 := Decrypt.FromHexBytes([]byte("xxxx")).ByAes(getCipher(CTR, Zero, aesKey, aesIV))
	assert.Equal(t, decodeError.ModeError("hex"), d2.Error)

	d3 := Decrypt.FromBase64String("xxxxxx").ByAes(getCipher(CFB, PKCS7, aesKey, aesIV))
	assert.Equal(t, decodeError.ModeError("base64"), d3.Error)
	d4 := Decrypt.FromBase64Bytes([]byte("xxxxxx")).ByAes(getCipher(CFB, PKCS7, aesKey, aesIV))
	assert.Equal(t, decodeError.ModeError("base64"), d4.Error)
}

// gets Cipher instance.
// 获取 Cipher 对象
func getCipher(mode cipherMode, padding cipherPadding, key, iv []byte) (cipher *Cipher) {
	cipher = NewCipher()
	cipher.SetMode(mode)
	cipher.SetPadding(padding)
	cipher.SetKey(bytes2string(key))
	cipher.SetIV(bytes2string(iv))
	cipher.WithKey(key)
	cipher.WithIV(iv)
	return
}
