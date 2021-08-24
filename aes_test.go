package dongle

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncrypt_ByAesWithCBC_FromString(t *testing.T) {
	assert := assert.New(t)
	input, key, iv := "hello world", "0123456789abcdef", "0123456789abcdef"

	tests := []struct {
		paddingMode  string // 填充模式
		encodingMode string // 编码模式
		expected     string // 期望值
	}{
		{ZeroPadding, HEX, "889935b7a0c64b0333d713cafaee08fe"},
		{PKCS5Padding, HEX, "c1e9b4529aac9793010f4677f6358efe"},
		{PKCS7Padding, HEX, "c1e9b4529aac9793010f4677f6358efe"},

		{ZeroPadding, BASE32, "RCMTLN5AYZFQGM6XCPFPV3QI7Y======"},
		{PKCS5Padding, BASE32, "YHU3IUU2VSLZGAIPIZ37MNMO7Y======"},
		{PKCS7Padding, BASE32, "YHU3IUU2VSLZGAIPIZ37MNMO7Y======"},

		{ZeroPadding, BASE58, "HsL9oy8dws8BDiaCnhxcW9"},
		{PKCS5Padding, BASE58, "QwpdQepaZ8PxuVpFdiVSER"},
		{PKCS7Padding, BASE58, "QwpdQepaZ8PxuVpFdiVSER"},

		{ZeroPadding, BASE64, "iJk1t6DGSwMz1xPK+u4I/g=="},
		{PKCS5Padding, BASE64, "wem0Upqsl5MBD0Z39jWO/g=="},
		{PKCS7Padding, BASE64, "wem0Upqsl5MBD0Z39jWO/g=="},
	}

	cipher := NewCipher()
	for index, test := range tests {
		cipher.SetMode(CBC)
		cipher.SetPadding(test.paddingMode)
		cipher.SetKey(key)
		cipher.SetIV(iv)
		e := Encrypt.FromString(input).ByAes(cipher)
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToString(test.encodingMode), "Current test id is "+strconv.Itoa(index))
	}
}

func TestEncrypt_ByAesWithCBC_FromBytes(t *testing.T) {
	assert := assert.New(t)
	input, key, iv := "hello world", "0123456789abcdef", "0123456789abcdef"

	tests := []struct {
		paddingMode  string // 填充模式
		encodingMode string // 编码模式
		expected     string // 期望值
	}{
		{ZeroPadding, HEX, "889935b7a0c64b0333d713cafaee08fe"},
		{PKCS5Padding, HEX, "c1e9b4529aac9793010f4677f6358efe"},
		{PKCS7Padding, HEX, "c1e9b4529aac9793010f4677f6358efe"},

		{ZeroPadding, BASE32, "RCMTLN5AYZFQGM6XCPFPV3QI7Y======"},
		{PKCS5Padding, BASE32, "YHU3IUU2VSLZGAIPIZ37MNMO7Y======"},
		{PKCS7Padding, BASE32, "YHU3IUU2VSLZGAIPIZ37MNMO7Y======"},

		{ZeroPadding, BASE58, "HsL9oy8dws8BDiaCnhxcW9"},
		{PKCS5Padding, BASE58, "QwpdQepaZ8PxuVpFdiVSER"},
		{PKCS7Padding, BASE58, "QwpdQepaZ8PxuVpFdiVSER"},

		{ZeroPadding, BASE64, "iJk1t6DGSwMz1xPK+u4I/g=="},
		{PKCS5Padding, BASE64, "wem0Upqsl5MBD0Z39jWO/g=="},
		{PKCS7Padding, BASE64, "wem0Upqsl5MBD0Z39jWO/g=="},
	}

	cipher := NewCipher()
	for index, test := range tests {
		cipher.SetMode(CBC)
		cipher.SetPadding(test.paddingMode)
		cipher.SetKey(key)
		cipher.SetIV(iv)
		e := Encrypt.FromBytes([]byte(input)).ByAes(cipher)
		assert.Nil(e.Error)
		assert.Equal([]byte(test.expected), e.ToBytes(test.encodingMode), "Current test id is "+strconv.Itoa(index))
	}
}

func TestDecrypt_ByAesWithCBC_FromString(t *testing.T) {
	assert := assert.New(t)
	expected, key, iv := "hello world", "0123456789abcdef", "0123456789abcdef"

	tests := []struct {
		input        string // 期望值
		decodingMode string // 解码模式
		paddingMode  string // 填充模式
	}{
		{"889935b7a0c64b0333d713cafaee08fe", HEX, ZeroPadding},
		{"c1e9b4529aac9793010f4677f6358efe", HEX, PKCS5Padding},
		{"c1e9b4529aac9793010f4677f6358efe", HEX, PKCS7Padding},

		{"RCMTLN5AYZFQGM6XCPFPV3QI7Y======", BASE32, ZeroPadding},
		{"YHU3IUU2VSLZGAIPIZ37MNMO7Y======", BASE32, PKCS5Padding},
		{"YHU3IUU2VSLZGAIPIZ37MNMO7Y======", BASE32, PKCS7Padding},

		{"HsL9oy8dws8BDiaCnhxcW9", BASE58, ZeroPadding},
		{"QwpdQepaZ8PxuVpFdiVSER", BASE58, PKCS5Padding},
		{"QwpdQepaZ8PxuVpFdiVSER", BASE58, PKCS7Padding},

		{"iJk1t6DGSwMz1xPK+u4I/g==", BASE64, ZeroPadding},
		{"wem0Upqsl5MBD0Z39jWO/g==", BASE64, PKCS5Padding},
		{"wem0Upqsl5MBD0Z39jWO/g==", BASE64, PKCS7Padding},
	}

	cipher := NewCipher()
	for index, test := range tests {
		cipher.SetMode(CBC)
		cipher.SetPadding(test.paddingMode)
		cipher.SetKey(key)
		cipher.SetIV(iv)
		e := Decrypt.FromString(test.input, test.decodingMode).ByAes(cipher)
		assert.Nil(e.Error)
		assert.Equal(expected, e.ToString(), "Current test id is "+strconv.Itoa(index))
	}
}

func TestDecrypt_ByAesWithCBC_FromBytes(t *testing.T) {
	assert := assert.New(t)
	expected, key, iv := "hello world", "0123456789abcdef", "0123456789abcdef"

	tests := []struct {
		input        string // 期望值
		decodingMode string // 解码模式
		paddingMode  string // 填充模式
	}{
		{"889935b7a0c64b0333d713cafaee08fe", HEX, ZeroPadding},
		{"c1e9b4529aac9793010f4677f6358efe", HEX, PKCS5Padding},
		{"c1e9b4529aac9793010f4677f6358efe", HEX, PKCS7Padding},

		{"RCMTLN5AYZFQGM6XCPFPV3QI7Y======", BASE32, ZeroPadding},
		{"YHU3IUU2VSLZGAIPIZ37MNMO7Y======", BASE32, PKCS5Padding},
		{"YHU3IUU2VSLZGAIPIZ37MNMO7Y======", BASE32, PKCS7Padding},

		{"HsL9oy8dws8BDiaCnhxcW9", BASE58, ZeroPadding},
		{"QwpdQepaZ8PxuVpFdiVSER", BASE58, PKCS5Padding},
		{"QwpdQepaZ8PxuVpFdiVSER", BASE58, PKCS7Padding},

		{"iJk1t6DGSwMz1xPK+u4I/g==", BASE64, ZeroPadding},
		{"wem0Upqsl5MBD0Z39jWO/g==", BASE64, PKCS5Padding},
		{"wem0Upqsl5MBD0Z39jWO/g==", BASE64, PKCS7Padding},
	}

	cipher := NewCipher()
	for index, test := range tests {
		cipher.SetMode(CBC)
		cipher.SetPadding(test.paddingMode)
		cipher.SetKey(key)
		cipher.SetIV(iv)
		e := Decrypt.FromBytes([]byte(test.input), test.decodingMode).ByAes(cipher)
		assert.Nil(e.Error)
		assert.Equal([]byte(expected), e.ToBytes(), "Current test id is "+strconv.Itoa(index))
	}
}

func TestError_Encrypt_Padding(t *testing.T) {
	encrypt, key, iv := "hello world", "0123456789abcdef", "0123456789abcdef"

	cipher := NewCipher()
	cipher.SetPadding("xxx")
	cipher.SetKey(key)
	cipher.SetIV(iv)

	e := Encrypt.FromString(encrypt).ByAes(cipher)
	assert.NotNil(t, e.Error, "It should catch a padding mode exception in ByAes()")
}

func TestError_Encrypt_Mode(t *testing.T) {
	encrypt, key, iv := "hello world", "0123456789abcdef", "0123456789abcdef"

	cipher := NewCipher()
	cipher.SetMode("xxx")
	cipher.SetKey(key)
	cipher.SetIV(iv)

	e := Encrypt.FromString(encrypt).ByAes(cipher)
	assert.NotNil(t, e.Error, "It should catch a group mode exception in ByAes()")
}

func TestError_Decrypt_Mode(t *testing.T) {
	decrypt, key, iv := "889935b7a0c64b0333d713cafaee08fe", "0123456789abcdef", "0123456789abcdef"

	cipher := NewCipher()
	cipher.SetMode("xxx")
	cipher.SetKey(key)
	cipher.SetIV(iv)

	e := Decrypt.FromString(decrypt).ByAes(cipher)
	assert.NotNil(t, e.Error, "It should catch a group mode exception in ByAes()")
}

func TestError_Decrypt_Padding(t *testing.T) {
	decrypt, key, iv := "889935b7a0c64b0333d713cafaee08fe", "0123456789abcdef", "0123456789abcdef"

	cipher := NewCipher()
	cipher.SetPadding("xxx")
	cipher.SetKey(key)
	cipher.SetIV(iv)

	e := Decrypt.FromString(decrypt).ByAes(cipher)
	assert.NotNil(t, e.Error, "It should catch a padding mode exception in ByAes()")
}

func TestError_Decrypt_Decoding(t *testing.T) {
	decrypt, key, iv := "889935b7a0c64b0333d713cafaee08fe", "0123456789abcdef", "0123456789abcdef"

	cipher := NewCipher()
	cipher.SetKey(key)
	cipher.SetIV(iv)

	e := Decrypt.FromString(decrypt, "xxx").ByAes(cipher)
	assert.NotNil(t, e.Error, "It should catch a decoding mode exception in ByAes()")
}
