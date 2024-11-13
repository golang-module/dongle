package dongle

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var safeUrlTestS = []struct {
	input  string
	output string
}{
	{"", ""},
	{"www.gouguoyin.cn?sex=男&age=18", "www.gouguoyin.cn%3Fsex%3D%E7%94%B7%26age%3D18"},
}

func TestSafeURL_Encrypt_String(t *testing.T) {
	for index, test := range safeUrlTestS {
		e := Encode.FromString(test.input).BySafeURL()

		t.Run(fmt.Sprintf("test_%d", index), func(t *testing.T) {
			assert.Nil(t, e.Error)
			assert.Equal(t, test.output, e.ToString())
		})
	}
}

func TestSafeURL_Decrypt_String(t *testing.T) {
	for index, test := range safeUrlTestS {
		e := Decode.FromString(test.output).BySafeURL()

		t.Run(fmt.Sprintf("test_%d", index), func(t *testing.T) {
			assert.Nil(t, e.Error)
			assert.Equal(t, test.input, e.ToString())
		})
	}
}

func TestSafeURL_Encrypt_Bytes(t *testing.T) {
	for index, test := range safeUrlTestS {
		e := Encode.FromBytes([]byte(test.input)).BySafeURL()

		t.Run(fmt.Sprintf("test_%d", index), func(t *testing.T) {
			assert.Nil(t, e.Error)
			assert.Equal(t, []byte(test.output), e.ToBytes())
		})
	}
}

func TestSafeURL_Decrypt_Bytes(t *testing.T) {
	for index, test := range safeUrlTestS {
		e := Decode.FromBytes([]byte(test.output)).BySafeURL()

		t.Run(fmt.Sprintf("test_%d", index), func(t *testing.T) {
			assert.Nil(t, e.Error)
			assert.Equal(t, []byte(test.input), e.ToBytes())
		})
	}
}
