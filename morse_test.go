package dongle

import (
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var morseTest = []struct {
	input  string // 输入值
	output string // 期望值
}{
	{"", ""},
	{"1", ".----"},
	{"F", "..-."},
	{"dongle", "-../---/-./--./.-../."},
	{"SOS", ".../---/..."},
}

func TestMorse_Encode_ToString(t *testing.T) {
	for index, test := range morseTest {
		e := Encode.FromString(test.input).ByMorse()
		assert.Nil(t, e.Error)
		assert.Equal(t, test.output, e.ToString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestMorse_Decode_ToString(t *testing.T) {
	for index, test := range morseTest {
		d := Decode.FromString(test.output).ByMorse()
		assert.Nil(t, d.Error)
		assert.Equal(t, strings.ToLower(test.input), d.ToString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestMorse_Encode_ToBytes(t *testing.T) {
	for index, test := range morseTest {
		e := Encode.FromBytes([]byte(test.input)).ByMorse()
		assert.Nil(t, e.Error)
		assert.Equal(t, []byte(test.output), e.ToBytes(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestMorse_Decode_ToBytes(t *testing.T) {
	for index, test := range morseTest {
		d := Decode.FromBytes([]byte(test.output)).ByMorse()
		assert.Nil(t, d.Error)
		assert.Equal(t, []byte(strings.ToLower(test.input)), d.ToBytes(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestMorse_Plaintext_Error(t *testing.T) {
	e := Encode.FromString("hello world").ByMorse()
	assert.Equal(t, invalidMorsePlaintextError(), e.Error)
}

func TestMorse_Ciphertext_Error(t *testing.T) {
	e := Decode.FromString("hello world").ByMorse()
	assert.Equal(t, invalidMorseCiphertextError(), e.Error)
}
