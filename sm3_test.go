package dongle

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var sm3Tests = []struct {
	input    string
	toHex    string
	toBase64 string
}{
	{"", "", ""},
	{"hello world", "44f0061e69fa6fdfc290c494654a05dc0c053da7e5c52b84ef93a9d67d3fff88", "RPAGHmn6b9/CkMSUZUoF3AwFPaflxSuE75Op1n0//4g="},
}

func TestSm3_Encrypt_String(t *testing.T) {
	for index, test := range sm3Tests {
		e := Encrypt.FromString(test.input).BySm3()

		t.Run(fmt.Sprintf("test_%d", index), func(t *testing.T) {
			assert.Nil(t, e.Error)
			assert.Equal(t, test.toHex, e.ToHexString())
			assert.Equal(t, test.toBase64, e.ToBase64String())
		})
	}
}

func TestSm3_Encrypt_Bytes(t *testing.T) {
	for index, test := range sm3Tests {
		e := Encrypt.FromBytes([]byte(test.input)).BySm3()

		t.Run(fmt.Sprintf("test_%d", index), func(t *testing.T) {
			assert.Nil(t, e.Error)
			assert.Equal(t, []byte(test.toHex), e.ToHexBytes())
			assert.Equal(t, []byte(test.toBase64), e.ToBase64Bytes())
		})
	}
}
