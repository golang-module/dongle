package dongle

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	hmacSha224Input, hmacSha224Key = "hello world", "dongle"
	hmacSha224HexExpected          = "e15b9e5a7eccb1f17dc81dc07c909a891936dc3429dc0d940accbcec"
	hmacSha224Base32Expected       = "4FNZ4WT6ZSY7C7OIDXAHZEE2REMTNXBUFHOA3FAKZS6OY==="
	hmacSha224Base64Expected       = "4VueWn7MsfF9yB3AfJCaiRk23DQp3A2UCsy87A=="
)

func TestEncrypt_ByHmacSha224_FromStringToString(t *testing.T) {
	hexTests := []struct {
		input1   string // 输入值1
		input2   string // 输入值2
		expected string // 期望值
	}{
		{"", "", ""},
		{hmacSha224Input, hmacSha224Key, hmacSha224HexExpected},
	}

	for index, test := range hexTests {
		e := Encrypt.FromString(test.input1).ByHmacSha224(test.input2)
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToHexString(), "Current test index is "+strconv.Itoa(index))
	}

	base32Tests := []struct {
		input1   string // 输入值1
		input2   string // 输入值2
		expected string // 期望值
	}{
		{"", "", ""},
		{hmacSha224Input, hmacSha224Key, hmacSha224Base32Expected},
	}

	for index, test := range base32Tests {
		e := Encrypt.FromString(test.input1).ByHmacSha224(test.input2)
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToBase32String(), "Current test index is "+strconv.Itoa(index))
	}

	base64Tests := []struct {
		input1   string // 输入值1
		input2   string // 输入值2
		expected string // 期望值
	}{
		{"", "", ""},
		{hmacSha224Input, hmacSha224Key, hmacSha224Base64Expected},
	}

	for index, test := range base64Tests {
		e := Encrypt.FromString(test.input1).ByHmacSha224(test.input2)
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToBase64String(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestEncrypt_ByHmacSha224_FromBytesToBytes(t *testing.T) {
	hexTests := []struct {
		input1   []byte // 输入值1
		input2   []byte // 输入值2
		expected []byte // 期望值
	}{
		{[]byte(""), []byte(""), []byte("")},
		{[]byte(hmacSha224Input), []byte(hmacSha224Key), []byte(hmacSha224HexExpected)},
	}

	for index, test := range hexTests {
		e := Encrypt.FromBytes(test.input1).ByHmacSha224(test.input2)
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToHexBytes(), "Current test index is "+strconv.Itoa(index))
	}

	base32Tests := []struct {
		input1   []byte // 输入值1
		input2   []byte // 输入值2
		expected []byte // 期望值
	}{
		{[]byte(""), []byte(""), []byte("")},
		{[]byte(hmacSha224Input), []byte(hmacSha224Key), []byte(hmacSha224Base32Expected)},
	}

	for index, test := range base32Tests {
		e := Encrypt.FromBytes(test.input1).ByHmacSha224(test.input2)
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToBase32Bytes(), "Current test index is "+strconv.Itoa(index))
	}

	base64Tests := []struct {
		input1   []byte // 输入值1
		input2   []byte // 输入值2
		expected []byte // 期望值
	}{
		{[]byte(""), []byte(""), []byte("")},
		{[]byte(hmacSha224Input), []byte(hmacSha224Key), []byte(hmacSha224Base64Expected)},
	}

	for index, test := range base64Tests {
		e := Encrypt.FromBytes(test.input1).ByHmacSha224(test.input2)
		assert.Nil(t, e.Error)
		assert.Equal(t, test.expected, e.ToBase64Bytes(), "Current test index is "+strconv.Itoa(index))
	}
}
