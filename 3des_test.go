package dongle

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

var (
	tripleDesKey = "0123456789abcdefghijklmn"
	tripleDesIV  = "12345678"
)

var tripleDesTest = []struct {
	mode     cipherMode
	padding  cipherPadding
	input    string
	toHex    string
	toBase32 string
	toBase64 string
}{
	{CBC, "", "", "", "", ""},
	{CBC, No, "hello world, go!", "7e9194cc827a325d49111eaa503110fe", "P2IZJTECPIZF2SIRD2VFAMIQ7Y======", "fpGUzIJ6Ml1JER6qUDEQ/g=="},
	{CBC, Zero, "hello world", "7e9194cc827a325ddaee992a89c5cd8d", "P2IZJTECPIZF3WXOTEVITRONRU======", "fpGUzIJ6Ml3a7pkqicXNjQ=="},
	{CBC, PKCS5, "hello world", "7e9194cc827a325db9c765859716cc97", "P2IZJTECPIZF3OOHMWCZOFWMS4======", "fpGUzIJ6Ml25x2WFlxbMlw=="},
	{CBC, PKCS7, "hello world", "7e9194cc827a325db9c765859716cc97", "P2IZJTECPIZF3OOHMWCZOFWMS4======", "fpGUzIJ6Ml25x2WFlxbMlw=="},

	{CFB, PKCS5, "", "", "", ""},
	{CFB, No, "hello world, go!", "f52dc896ceff0ecc9393c19d4e1e7591", "6UW4RFWO74HMZE4TYGOU4HTVSE======", "9S3Ils7/DsyTk8GdTh51kQ=="},
	{CFB, Zero, "hello world", "f52dc896ceff0ecc9393c1b16e791ab0", "6UW4RFWO74HMZE4TYGYW46I2WA======", "9S3Ils7/DsyTk8GxbnkasA=="},
	{CFB, PKCS5, "hello world", "f52dc896ceff0ecc9393c1b46b7c1fb5", "6UW4RFWO74HMZE4TYG2GW7A7WU======", "9S3Ils7/DsyTk8G0a3wftQ=="},
	{CFB, PKCS7, "hello world", "f52dc896ceff0ecc9393c1b46b7c1fb5", "6UW4RFWO74HMZE4TYG2GW7A7WU======", "9S3Ils7/DsyTk8G0a3wftQ=="},

	{OFB, PKCS7, "", "", "", ""},
	{OFB, No, "hello world, go!", "f52dc896ceff0eccf869cb59735c3766", "6UW4RFWO74HMZ6DJZNMXGXBXMY======", "9S3Ils7/Dsz4actZc1w3Zg=="},
	{OFB, Zero, "hello world", "f52dc896ceff0eccf869cb75533b5847", "6UW4RFWO74HMZ6DJZN2VGO2YI4======", "9S3Ils7/Dsz4act1UztYRw=="},
	{OFB, PKCS5, "hello world", "f52dc896ceff0eccf869cb70563e5d42", "6UW4RFWO74HMZ6DJZNYFMPS5II======", "9S3Ils7/Dsz4actwVj5dQg=="},
	{OFB, PKCS7, "hello world", "f52dc896ceff0eccf869cb70563e5d42", "6UW4RFWO74HMZ6DJZNYFMPS5II======", "9S3Ils7/Dsz4actwVj5dQg=="},

	{CTR, PKCS5, "", "", "", ""},
	{CTR, No, "hello world, go!", "f52dc896ceff0ecc366b2281038f6f7f", "6UW4RFWO74HMYNTLEKAQHD3PP4======", "9S3Ils7/Dsw2ayKBA49vfw=="},
	{CTR, Zero, "hello world", "f52dc896ceff0ecc366b22ad23e8005e", "6UW4RFWO74HMYNTLEKWSH2AALY======", "9S3Ils7/Dsw2ayKtI+gAXg=="},
	{CTR, PKCS5, "hello world", "f52dc896ceff0ecc366b22a826ed055b", "6UW4RFWO74HMYNTLEKUCN3IFLM======", "9S3Ils7/Dsw2ayKoJu0FWw=="},
	{CTR, PKCS7, "hello world", "f52dc896ceff0ecc366b22a826ed055b", "6UW4RFWO74HMYNTLEKUCN3IFLM======", "9S3Ils7/Dsw2ayKoJu0FWw=="},

	{ECB, PKCS7, "", "", "", ""},
	{ECB, No, "hello world, go!", "b8097975c76319c623be7c7aa6e0f3fc", "XAEXS5OHMMM4MI56PR5KNYHT7Q======", "uAl5dcdjGcYjvnx6puDz/A=="},
	{ECB, Zero, "hello world", "b8097975c76319c61971a986e579cdf9", "XAEXS5OHMMM4MGLRVGDOK6ON7E======", "uAl5dcdjGcYZcamG5XnN+Q=="},
	{ECB, PKCS5, "hello world", "b8097975c76319c6172687e0d90fd4d1", "XAEXS5OHMMM4MFZGQ7QNSD6U2E======", "uAl5dcdjGcYXJofg2Q/U0Q=="},
	{ECB, PKCS7, "hello world", "b8097975c76319c6172687e0d90fd4d1", "XAEXS5OHMMM4MFZGQ7QNSD6U2E======", "uAl5dcdjGcYXJofg2Q/U0Q=="},
}

func Test3Des_Encrypt_ToString(t *testing.T) {
	for index, test := range tripleDesTest {
		e := Encrypt.FromString(test.input).By3Des(getCipher(test.mode, test.padding, tripleDesKey, tripleDesIV))
		assert.Nil(t, e.Error)

		assert.Equal(t, test.toHex, e.ToHexString(), "Hex test index is "+strconv.Itoa(index))
		assert.Equal(t, test.toBase32, e.ToBase32String(), "Base32 test index is "+strconv.Itoa(index))
		assert.Equal(t, test.toBase64, e.ToBase64String(), "Base64 test index is "+strconv.Itoa(index))

		assert.Equal(t, Decode.FromString(test.toHex).ByHex().ToString(), fmt.Sprintf("%s", e), "Hex test index is "+strconv.Itoa(index))
		assert.Equal(t, Decode.FromString(test.toHex).ByHex().ToString(), e.ToString(), "Raw test index is "+strconv.Itoa(index))
	}
}

func Test3Des_Decrypt_ToString(t *testing.T) {
	for index, test := range tripleDesTest {
		e := Decrypt.FromHexString(test.toHex).By3Des(getCipher(test.mode, test.padding, tripleDesKey, tripleDesIV))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.input, e.ToString(), "Hex test index is "+strconv.Itoa(index))
		assert.Equal(t, test.input, fmt.Sprintf("%s", e), "Hex test index is "+strconv.Itoa(index))
	}

	for index, test := range tripleDesTest {
		e := Decrypt.FromBase32String(test.toBase32).By3Des(getCipher(test.mode, test.padding, tripleDesKey, tripleDesIV))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.input, e.ToString(), "Base32 test index is "+strconv.Itoa(index))
	}

	for index, test := range tripleDesTest {
		e := Decrypt.FromBase64String(test.toBase64).By3Des(getCipher(test.mode, test.padding, tripleDesKey, tripleDesIV))
		assert.Nil(t, e.Error)
		assert.Equal(t, test.input, e.ToString(), "Base64 test index is "+strconv.Itoa(index))
		assert.Equal(t, test.input, fmt.Sprintf("%s", e), "Hex test index is "+strconv.Itoa(index))
	}
}

func Test3Des_Encrypt_ToBytes(t *testing.T) {
	for index, test := range tripleDesTest {
		e := Encrypt.FromString(test.input).By3Des(getCipher(test.mode, test.padding, []byte(tripleDesKey), []byte(tripleDesIV)))
		assert.Nil(t, e.Error)

		assert.Equal(t, []byte(test.toHex), e.ToHexBytes(), "Hex test index is "+strconv.Itoa(index))
		assert.Equal(t, []byte(test.toBase32), e.ToBase32Bytes(), "Base32 test index is "+strconv.Itoa(index))
		assert.Equal(t, []byte(test.toBase64), e.ToBase64Bytes(), "Base64 test index is "+strconv.Itoa(index))

		assert.Equal(t, Decode.FromString(test.toHex).ByHex().ToBytes(), e.ToBytes(), "Raw test index is "+strconv.Itoa(index))
	}
}

func Test3Des_Decrypt_ToBytes(t *testing.T) {
	for index, test := range tripleDesTest {
		e := Decrypt.FromBytes(Decode.FromString(test.toHex).ByHex().ToBytes()).By3Des(getCipher(test.mode, test.padding, []byte(tripleDesKey), []byte(tripleDesIV)))
		assert.Nil(t, e.Error)
		assert.Equal(t, []byte(test.input), e.ToBytes(), "Raw test index is "+strconv.Itoa(index))
	}

	for index, test := range tripleDesTest {
		e := Decrypt.FromHexBytes([]byte(test.toHex)).By3Des(getCipher(test.mode, test.padding, []byte(tripleDesKey), []byte(tripleDesIV)))
		assert.Nil(t, e.Error)
		assert.Equal(t, []byte(test.input), e.ToBytes(), "Hex test index is "+strconv.Itoa(index))
	}

	for index, test := range tripleDesTest {
		e := Decrypt.FromBase32Bytes([]byte(test.toBase32)).By3Des(getCipher(test.mode, test.padding, []byte(tripleDesKey), []byte(tripleDesIV)))
		assert.Nil(t, e.Error)
		assert.Equal(t, []byte(test.input), e.ToBytes(), "Base32 test index is "+strconv.Itoa(index))
	}

	for index, test := range tripleDesTest {
		e := Decrypt.FromBase64Bytes([]byte(test.toBase64)).By3Des(getCipher(test.mode, test.padding, []byte(tripleDesKey), []byte(tripleDesIV)))
		assert.Nil(t, e.Error)
		assert.Equal(t, []byte(test.input), e.ToBytes(), "Base64 test index is "+strconv.Itoa(index))
	}
}

func Test3Des_Key_Error(t *testing.T) {
	e := Encrypt.FromString("hello world").By3Des(getCipher(CBC, PKCS7, "xxxx", tripleDesIV))
	assert.Equal(t, invalid3DesKeyError(), e.Error)

	d := Decrypt.FromHexString("68656c6c6f20776f726c64").By3Des(getCipher(CBC, PKCS7, "xxxx", tripleDesIV))
	assert.Equal(t, invalid3DesKeyError(), d.Error)
}

func Test3Des_IV_Error(t *testing.T) {
	e := Encrypt.FromString("hello world").By3Des(getCipher(OFB, PKCS7, tripleDesKey, "xxxx"))
	assert.Equal(t, invalid3DesIVError(), e.Error)

	d := Decrypt.FromHexString("68656c6c6f20776f726c64").By3Des(getCipher(CBC, PKCS7, tripleDesKey, "xxxx"))
	assert.Equal(t, invalid3DesIVError(), d.Error)
}

func Test3Des_Mode_Error(t *testing.T) {
	e := Encrypt.FromString("hello world").By3Des(getCipher("xxxx", PKCS7, tripleDesKey, tripleDesIV))
	assert.Equal(t, invalidModeError("xxxx"), e.Error)

	d := Decrypt.FromHexString("68656c6c6f20776f726c64").By3Des(getCipher("xxxx", PKCS7, tripleDesKey, tripleDesIV))
	assert.Equal(t, invalidModeError("xxxx"), d.Error)
}

func Test3Des_Padding_Error(t *testing.T) {
	e := Encrypt.FromString("hello world").By3Des(getCipher(CFB, "xxxx", tripleDesKey, tripleDesIV))
	assert.Equal(t, invalidPaddingError("xxxx"), e.Error)

	d := Decrypt.FromHexString("68656c6c6f20776f726c64").By3Des(getCipher(CBC, "xxxx", tripleDesKey, tripleDesIV))
	assert.Equal(t, invalidPaddingError("xxxx"), d.Error)
}

func Test3Des_Plaintext_Error(t *testing.T) {
	e := Encrypt.FromString("hello world").By3Des(getCipher(CFB, No, tripleDesKey, tripleDesIV))
	assert.Equal(t, invalid3DesSrcError(), e.Error)

	d := Decrypt.FromHexString("68656c6c6f20776f726c64").By3Des(getCipher(CBC, No, tripleDesKey, tripleDesIV))
	assert.Equal(t, invalid3DesSrcError(), d.Error)
}

func Test3Des_Ciphertext_Error(t *testing.T) {
	d1 := Decrypt.FromHexBytes([]byte("xxxx")).By3Des(getCipher(CTR, Zero, []byte(tripleDesKey), []byte(tripleDesIV)))
	assert.Equal(t, invalidDecodingError("hex"), d1.Error)

	d2 := Decrypt.FromBase32Bytes([]byte("xxxx")).By3Des(getCipher(CBC, PKCS5, []byte(tripleDesKey), []byte(tripleDesIV)))
	assert.Equal(t, invalidDecodingError("base32"), d2.Error)

	d3 := Decrypt.FromBase64Bytes([]byte("xxxxxx")).By3Des(getCipher(CFB, PKCS7, []byte(tripleDesKey), []byte(tripleDesIV)))
	assert.Equal(t, invalidDecodingError("base64"), d3.Error)
}
