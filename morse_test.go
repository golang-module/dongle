package dongle

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var morseTests = []struct {
	input     string // 输入值
	separator string // 分隔符
	output    string // 期望值
}{
	{"", "", ""},
	{"1", "/", ".----"},
	{"F", "/", "..-."},
	{"dongle", "|", "-..|---|-.|--.|.-..|."},
	{"SOS", "/", ".../---/..."},
}

func TestMorse_Encode_String(t *testing.T) {
	for index, test := range morseTests {
		e := Encode.FromString(test.input).ByMorse(test.separator)
		t.Run(fmt.Sprintf("test_%d", index), func(t *testing.T) {
			assert.Nil(t, e.Error)
			assert.Equal(t, test.output, e.ToString())
		})
	}
}

func TestMorse_Decode_String(t *testing.T) {
	for index, test := range morseTests {
		d := Decode.FromString(test.output).ByMorse(test.separator)
		t.Run(fmt.Sprintf("test_%d", index), func(t *testing.T) {
			assert.Nil(t, d.Error)
			assert.Equal(t, strings.ToLower(test.input), d.ToString())
		})
	}
}

func TestMorse_Encode_Bytes(t *testing.T) {
	for index, test := range morseTests {
		e := Encode.FromBytes([]byte(test.input)).ByMorse(test.separator)
		t.Run(fmt.Sprintf("test_%d", index), func(t *testing.T) {
			assert.Nil(t, e.Error)
			assert.Equal(t, []byte(test.output), e.ToBytes())
		})
	}
}

func TestMorse_Decode_Bytes(t *testing.T) {
	for index, test := range morseTests {
		d := Decode.FromBytes([]byte(test.output)).ByMorse(test.separator)
		t.Run(fmt.Sprintf("test_%d", index), func(t *testing.T) {
			assert.Nil(t, d.Error)
			assert.Equal(t, []byte(strings.ToLower(test.input)), d.ToBytes())
		})
	}
}

func TestMorse_Src_Error(t *testing.T) {
	e := Encode.FromString("hello world").ByMorse()
	assert.Equal(t, morseError.SrcError(), e.Error)
}

func TestMorse_Decoding_Error(t *testing.T) {
	e := Decode.FromString("hello world").ByMorse()
	assert.Equal(t, morseError.DecodeError(), e.Error)
}
