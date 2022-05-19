package dongle

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	hmacSha384Input, hmacSha384Key = "hello world", "dongle"
	hmacSha384HexExpected          = "421fcaa740216a31bbcd1f86f2212e0c68aa4b156a8ebc2ae55b3e75c4ee0509ea0325a0570ae739006b61d91d817fe8"
	hmacSha384Base32Expected       = "IIP4VJ2AEFVDDO6ND6DPEIJOBRUKUSYVNKHLYKXFLM7HLRHOAUE6UAZFUBLQVZZZABVWDWI5QF76Q==="
	hmacSha384Base64Expected       = "Qh/Kp0AhajG7zR+G8iEuDGiqSxVqjrwq5Vs+dcTuBQnqAyWgVwrnOQBrYdkdgX/o"
)

func TestEncrypt_ByHmacSha384_FromStringToString(t *testing.T) {
	assert := assert.New(t)

	hexTests := []struct {
		input1   string // 输入值1
		input2   string // 输入值2
		expected string // 期望值
	}{
		{"", "", ""},
		{hmacSha384Input, hmacSha384Key, hmacSha384HexExpected},
	}

	for index, test := range hexTests {
		e := Encrypt.FromString(test.input1).ByHmacSha384(test.input2)
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToHexString(), "Current test index is "+strconv.Itoa(index))
	}

	base32Tests := []struct {
		input1   string // 输入值1
		input2   string // 输入值2
		expected string // 期望值
	}{
		{"", "", ""},
		{hmacSha384Input, hmacSha384Key, hmacSha384Base32Expected},
	}

	for index, test := range base32Tests {
		e := Encrypt.FromString(test.input1).ByHmacSha384(test.input2)
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToBase32String(), "Current test index is "+strconv.Itoa(index))
	}

	base64Tests := []struct {
		input1   string // 输入值1
		input2   string // 输入值2
		expected string // 期望值
	}{
		{"", "", ""},
		{hmacSha384Input, hmacSha384Key, hmacSha384Base64Expected},
	}

	for index, test := range base64Tests {
		e := Encrypt.FromString(test.input1).ByHmacSha384(test.input2)
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToBase64String(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestEncrypt_ByHmacSha384_FromBytesToBytes(t *testing.T) {
	assert := assert.New(t)

	hexTests := []struct {
		input1   []byte // 输入值1
		input2   []byte // 输入值2
		expected []byte // 期望值
	}{
		{[]byte(""), []byte(""), []byte("")},
		{[]byte(hmacSha384Input), []byte(hmacSha384Key), []byte(hmacSha384HexExpected)},
	}

	for index, test := range hexTests {
		e := Encrypt.FromBytes(test.input1).ByHmacSha384(test.input2)
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToHexBytes(), "Current test index is "+strconv.Itoa(index))
	}

	base32Tests := []struct {
		input1   []byte // 输入值1
		input2   []byte // 输入值2
		expected []byte // 期望值
	}{
		{[]byte(""), []byte(""), []byte("")},
		{[]byte(hmacSha384Input), []byte(hmacSha384Key), []byte(hmacSha384Base32Expected)},
	}

	for index, test := range base32Tests {
		e := Encrypt.FromBytes(test.input1).ByHmacSha384(test.input2)
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToBase32Bytes(), "Current test index is "+strconv.Itoa(index))
	}

	base64Tests := []struct {
		input1   []byte // 输入值1
		input2   []byte // 输入值2
		expected []byte // 期望值
	}{
		{[]byte(""), []byte(""), []byte("")},
		{[]byte(hmacSha384Input), []byte(hmacSha384Key), []byte(hmacSha384Base64Expected)},
	}

	for index, test := range base64Tests {
		e := Encrypt.FromBytes(test.input1).ByHmacSha384(test.input2)
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToBase64Bytes(), "Current test index is "+strconv.Itoa(index))
	}
}
