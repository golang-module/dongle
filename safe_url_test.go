package dongle

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

var safeUrlTest = []struct {
	input  string
	output string
}{
	{"", ""},
	{"www.gouguoyin.cn?sex=男&age=18", "www.gouguoyin.cn%3Fsex%3D%E7%94%B7%26age%3D18"},
}

func TestSafeURL_Encrypt_ToString(t *testing.T) {
	for index, test := range safeUrlTest {
		e := Encode.FromString(test.input).BySafeURL()
		assert.Nil(t, e.Error)

		assert.Equal(t, test.output, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestSafeURL_Decrypt_ToString(t *testing.T) {
	for index, test := range safeUrlTest {
		e := Decode.FromString(test.output).BySafeURL()
		assert.Nil(t, e.Error)

		assert.Equal(t, test.input, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestSafeURL_Encrypt_ToBytes(t *testing.T) {
	for index, test := range safeUrlTest {
		e := Encode.FromBytes([]byte(test.input)).BySafeURL()
		assert.Nil(t, e.Error)

		assert.Equal(t, []byte(test.output), e.ToBytes(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestSafeURL_Decrypt_ToBytes(t *testing.T) {
	for index, test := range safeUrlTest {
		e := Decode.FromBytes([]byte(test.output)).BySafeURL()
		assert.Nil(t, e.Error)

		assert.Equal(t, []byte(test.input), e.ToBytes(), "Current test index is "+strconv.Itoa(index))
	}
}
