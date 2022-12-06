package morse

import (
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var Tests = []struct {
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

func TestEncode(t *testing.T) {
	for index, test := range Tests {
		dst, err := Encode(string2bytes(test.input), test.separator)

		t.Run(fmt.Sprintf("test_%d", index), func(t *testing.T) {
			assert.Nil(t, err)
			assert.Equal(t, test.output, dst)
		})
	}
}

func TestDecode(t *testing.T) {
	for index, test := range Tests {
		dst, err := Decode(string2bytes(test.output), test.separator)

		t.Run(fmt.Sprintf("test_%d", index), func(t *testing.T) {
			assert.Nil(t, err)
			assert.Equal(t, strings.ToLower(test.input), dst)
		})
	}
}

func TestError(t *testing.T) {
	_, err1 := Encode([]byte("hello world"), "/")
	assert.Equal(t, errors.New("can't contain spaces"), err1)

	_, err2 := Decode([]byte("hello world"), "/")
	assert.Equal(t, fmt.Errorf("unknown character hello world"), err2)
}
