package dongle

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

var (
	aesKey = "0123456789abcdef"
	aesIV  = "0123456789abcdef"
)

var aesTest = []struct {
	mode     cipherMode
	padding  cipherPadding
	input    string
	toHex    string
	toBase32 string
	toBase64 string
}{
	{CBC, "", "", "", "", ""},
	{CBC, No, "hello world, go!", "77aa39926f9b2f3f22254bfd422fa75d", "O6VDTETPTMXT6IRFJP6UEL5HLU======", "d6o5km+bLz8iJUv9Qi+nXQ=="},
	{CBC, Zero, "hello world", "889935b7a0c64b0333d713cafaee08fe", "RCMTLN5AYZFQGM6XCPFPV3QI7Y======", "iJk1t6DGSwMz1xPK+u4I/g=="},
	{CBC, PKCS5, "hello world", "c1e9b4529aac9793010f4677f6358efe", "YHU3IUU2VSLZGAIPIZ37MNMO7Y======", "wem0Upqsl5MBD0Z39jWO/g=="},
	{CBC, PKCS7, "hello world", "c1e9b4529aac9793010f4677f6358efe", "YHU3IUU2VSLZGAIPIZ37MNMO7Y======", "wem0Upqsl5MBD0Z39jWO/g=="},

	{CFB, "", "", "", "", ""},
	{CFB, No, "hello world, go!", "1a1712e471fc8a6e72cb7c44596eda44", "DILRFZDR7SFG44WLPRCFS3W2IQ======", "GhcS5HH8im5yy3xEWW7aRA=="},
	{CFB, Zero, "hello world", "1a1712e471fc8a6e72cb7c687909b565", "DILRFZDR7SFG44WLPRUHSCNVMU======", "GhcS5HH8im5yy3xoeQm1ZQ=="},
	{CFB, PKCS5, "hello world", "1a1712e471fc8a6e72cb7c6d7c0cb060", "DILRFZDR7SFG44WLPRWXYDFQMA======", "GhcS5HH8im5yy3xtfAywYA=="},
	{CFB, PKCS7, "hello world", "1a1712e471fc8a6e72cb7c6d7c0cb060", "DILRFZDR7SFG44WLPRWXYDFQMA======", "GhcS5HH8im5yy3xtfAywYA=="},

	{OFB, "", "", "", "", ""},
	{OFB, No, "hello world, go!", "1a1712e471fc8a6e72cb7c44596eda44", "DILRFZDR7SFG44WLPRCFS3W2IQ======", "GhcS5HH8im5yy3xEWW7aRA=="},
	{OFB, Zero, "hello world", "1a1712e471fc8a6e72cb7c687909b565", "DILRFZDR7SFG44WLPRUHSCNVMU======", "GhcS5HH8im5yy3xoeQm1ZQ=="},
	{OFB, PKCS5, "hello world", "1a1712e471fc8a6e72cb7c6d7c0cb060", "DILRFZDR7SFG44WLPRWXYDFQMA======", "GhcS5HH8im5yy3xtfAywYA=="},
	{OFB, PKCS7, "hello world", "1a1712e471fc8a6e72cb7c6d7c0cb060", "DILRFZDR7SFG44WLPRWXYDFQMA======", "GhcS5HH8im5yy3xtfAywYA=="},

	{CTR, "", "", "", "", ""},
	{CTR, No, "hello world, go!", "1a1712e471fc8a6e72cb7c44596eda44", "DILRFZDR7SFG44WLPRCFS3W2IQ======", "GhcS5HH8im5yy3xEWW7aRA=="},
	{CTR, Zero, "hello world", "1a1712e471fc8a6e72cb7c687909b565", "DILRFZDR7SFG44WLPRUHSCNVMU======", "GhcS5HH8im5yy3xoeQm1ZQ=="},
	{CTR, PKCS5, "hello world", "1a1712e471fc8a6e72cb7c6d7c0cb060", "DILRFZDR7SFG44WLPRWXYDFQMA======", "GhcS5HH8im5yy3xtfAywYA=="},
	{CTR, PKCS7, "hello world", "1a1712e471fc8a6e72cb7c6d7c0cb060", "DILRFZDR7SFG44WLPRWXYDFQMA======", "GhcS5HH8im5yy3xtfAywYA=="},

	{ECB, "", "", "", "", ""},
	{ECB, No, "hello world, go!", "f82a4c0db7a82f70c7fa84c39fa7627b", "7AVEYDNXVAXXBR72QTBZ7J3CPM======", "+CpMDbeoL3DH+oTDn6diew=="},
	{ECB, Zero, "hello world", "769c326290511c93bd59bba9c24d8904", "O2ODEYUQKEOJHPKZXOU4ETMJAQ======", "dpwyYpBRHJO9Wbupwk2JBA=="},
	{ECB, PKCS5, "hello world", "8169bed4ef49a8874559c5b200daade7", "QFU35VHPJGUIORKZYWZABWVN44======", "gWm+1O9JqIdFWcWyANqt5w=="},
	{ECB, PKCS7, "hello world", "8169bed4ef49a8874559c5b200daade7", "QFU35VHPJGUIORKZYWZABWVN44======", "gWm+1O9JqIdFWcWyANqt5w=="},
}

func TestAes_Encrypt_String(t *testing.T) {
	for index, test := range aesTest {
		e := Encrypt.FromString(test.input).ByAes(getCipher(test.mode, test.padding, aesKey, aesIV))
		assert.Nil(t, e.Error)

		assert.Equal(t, test.toHex, e.ToHexString(), "Hex test index is "+strconv.Itoa(index))
		assert.Equal(t, test.toBase32, e.ToBase32String(), "Base32 test index is "+strconv.Itoa(index))
		assert.Equal(t, test.toBase64, e.ToBase64String(), "Base64 test index is "+strconv.Itoa(index))

		assert.Equal(t, Decode.FromString(test.toHex).ByHex().ToString(), fmt.Sprintf("%s", e), "Raw test index is "+strconv.Itoa(index))
		assert.Equal(t, Decode.FromString(test.toHex).ByHex().ToString(), e.ToString(), "Raw test index is "+strconv.Itoa(index))
	}
}

func TestAes_Decrypt_ToString(t *testing.T) {
	for index, test := range aesTest {
		e := Decrypt.FromHexString(test.toHex).ByAes(getCipher(test.mode, test.padding, aesKey, aesIV))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.input, e.ToString(), "Hex test index is "+strconv.Itoa(index))
		assert.Equal(t, test.input, fmt.Sprintf("%s", e), "Hex test index is "+strconv.Itoa(index))
	}

	for index, test := range aesTest {
		e := Decrypt.FromBase32String(test.toBase32).ByAes(getCipher(test.mode, test.padding, aesKey, aesIV))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.input, e.ToString(), "Base32 test index is "+strconv.Itoa(index))
		assert.Equal(t, test.input, fmt.Sprintf("%s", e), "Hex test index is "+strconv.Itoa(index))
	}

	for index, test := range aesTest {
		e := Decrypt.FromBase64String(test.toBase64).ByAes(getCipher(test.mode, test.padding, aesKey, aesIV))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.input, e.ToString(), "Base64 test index is "+strconv.Itoa(index))
		assert.Equal(t, test.input, fmt.Sprintf("%s", e), "Hex test index is "+strconv.Itoa(index))
	}
}

func TestAes_Encrypt_ToBytes(t *testing.T) {
	for index, test := range aesTest {
		e := Encrypt.FromBytes([]byte(test.input)).ByAes(getCipher(test.mode, test.padding, []byte(aesKey), []byte(aesIV)))
		assert.Nil(t, e.Error)

		assert.Equal(t, []byte(test.toHex), e.ToHexBytes(), "Hex test index is "+strconv.Itoa(index))
		assert.Equal(t, []byte(test.toBase32), e.ToBase32Bytes(), "Base32 test index is "+strconv.Itoa(index))
		assert.Equal(t, []byte(test.toBase64), e.ToBase64Bytes(), "Base64 test index is "+strconv.Itoa(index))

		assert.Equal(t, Decode.FromString(test.toHex).ByHex().ToBytes(), e.ToBytes(), "Raw test index is "+strconv.Itoa(index))
	}
}

func TestAes_Decrypt_ToBytes(t *testing.T) {
	for index, test := range aesTest {
		e := Decrypt.FromBytes(Decode.FromString(test.toHex).ByHex().ToBytes()).ByAes(getCipher(test.mode, test.padding, []byte(aesKey), []byte(aesIV)))
		assert.Nil(t, e.Error)
		assert.Equal(t, []byte(test.input), e.ToBytes(), "Raw test index is "+strconv.Itoa(index))
	}

	for index, test := range aesTest {
		e := Decrypt.FromHexBytes([]byte(test.toHex)).ByAes(getCipher(test.mode, test.padding, []byte(aesKey), []byte(aesIV)))
		assert.Nil(t, e.Error)
		assert.Equal(t, []byte(test.input), e.ToBytes(), "Hex test index is "+strconv.Itoa(index))
	}

	for index, test := range aesTest {
		e := Decrypt.FromBase32Bytes([]byte(test.toBase32)).ByAes(getCipher(test.mode, test.padding, []byte(aesKey), []byte(aesIV)))
		assert.Nil(t, e.Error)
		assert.Equal(t, []byte(test.input), e.ToBytes(), "Base32 test index is "+strconv.Itoa(index))
	}

	for index, test := range aesTest {
		e := Decrypt.FromBase64Bytes([]byte(test.toBase64)).ByAes(getCipher(test.mode, test.padding, []byte(aesKey), []byte(aesIV)))
		assert.Nil(t, e.Error)
		assert.Equal(t, []byte(test.input), e.ToBytes(), "Base64 test index is "+strconv.Itoa(index))
	}
}

func TestAes_Key_Error(t *testing.T) {
	e := Encrypt.FromString("hello world").ByAes(getCipher(CBC, PKCS7, "xxxx", aesIV))
	assert.Equal(t, invalidAesKeyError(), e.Error)

	d := Decrypt.FromHexString("68656c6c6f20776f726c64").ByAes(getCipher(CBC, PKCS7, "xxxx", aesIV))
	assert.Equal(t, invalidAesKeyError(), d.Error)
}

func TestAes_IV_Error(t *testing.T) {
	e := Encrypt.FromString("hello world").ByAes(getCipher(OFB, PKCS7, aesKey, "xxxx"))
	assert.Equal(t, invalidAesIVError(), e.Error)

	d := Decrypt.FromHexString("68656c6c6f20776f726c64").ByAes(getCipher(CBC, PKCS7, aesKey, "xxxx"))
	assert.Equal(t, invalidAesIVError(), d.Error)
}

func TestAes_Mode_Error(t *testing.T) {
	e := Encrypt.FromString("hello world").ByAes(getCipher("xxxx", PKCS7, aesKey, aesIV))
	assert.Equal(t, invalidModeError("xxxx"), e.Error)

	d := Decrypt.FromHexString("68656c6c6f20776f726c64").ByAes(getCipher("xxxx", PKCS7, aesKey, aesIV))
	assert.Equal(t, invalidModeError("xxxx"), d.Error)
}

func TestAes_Padding_Error(t *testing.T) {
	e := Encrypt.FromString("hello world").ByAes(getCipher(CFB, "xxxx", aesKey, aesIV))
	assert.Equal(t, invalidPaddingError("xxxx"), e.Error)

	d := Decrypt.FromHexString("68656c6c6f20776f726c64").ByAes(getCipher(CBC, "xxxx", aesKey, aesIV))
	assert.Equal(t, invalidPaddingError("xxxx"), d.Error)
}

func TestAes_Plaintext_Error(t *testing.T) {
	e := Encrypt.FromString("hello world").ByAes(getCipher(CFB, No, aesKey, aesIV))
	assert.Equal(t, invalidAesSrcError(), e.Error)

	d := Decrypt.FromHexString("68656c6c6f20776f726c64").ByAes(getCipher(CBC, No, aesKey, aesIV))
	assert.Equal(t, invalidAesSrcError(), d.Error)
}

func TestAes_Ciphertext_Error(t *testing.T) {
	d1 := Decrypt.FromHexString("xxxx").ByAes(getCipher(CTR, Zero, aesKey, aesIV))
	assert.Equal(t, invalidDecodingError("hex"), d1.Error)
	d2 := Decrypt.FromHexBytes([]byte("xxxx")).ByAes(getCipher(CTR, Zero, []byte(aesKey), []byte(aesIV)))
	assert.Equal(t, invalidDecodingError("hex"), d2.Error)

	d3 := Decrypt.FromBase32String("xxxx").ByAes(getCipher(CBC, PKCS5, aesKey, aesIV))
	assert.Equal(t, invalidDecodingError("base32"), d3.Error)
	d4 := Decrypt.FromBase32Bytes([]byte("xxxx")).ByAes(getCipher(CBC, PKCS5, []byte(aesKey), []byte(aesIV)))
	assert.Equal(t, invalidDecodingError("base32"), d4.Error)

	d5 := Decrypt.FromBase64String("xxxxxx").ByAes(getCipher(CFB, PKCS7, aesKey, aesIV))
	assert.Equal(t, invalidDecodingError("base64"), d5.Error)
	d6 := Decrypt.FromBase64Bytes([]byte("xxxxxx")).ByAes(getCipher(CFB, PKCS7, []byte(aesKey), []byte(aesIV)))
	assert.Equal(t, invalidDecodingError("base64"), d6.Error)
}
