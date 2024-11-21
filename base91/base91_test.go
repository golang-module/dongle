package base91

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type pair struct {
	decoded, encoded string
}

var pairs = []pair{
	{
		"May your trails be crooked, winding, lonesome, dangerous, leading to the most amazing view. May your mountains rise into and above the clouds.",
		"8D9KR`0eLUd/ZQFl62>vb,1RL%%&~8bju\"sQ;mmaU=UfU)1T70<^rm?i;Ct)/p;R(&^m5PKimf2+H[QSd/[E<oTPgZh>DZ%y;#,aIl]U>vP:3pPIqSwPmLwre3:W.{6U)/wP;mYBxgP[UCsS)/[EOiqMgZR*Sk<Rd/=8jL=ibg7+b[C",
	},

	// Random bytes
	{
		"\x35\x5e\x56\xe0\xc6\x29\x38\xf4\x81\x00\xab\x81\x7e\xd7\x08\x95\x62\x20\xa7\xda\x64\xa2\xce\xb3\xc5",
		"~_1H=x_t{|$AjJX(nMFdjL~:?1b3HgM",
	},

	// RFC 3548 examples (adapted from base64 to base91)
	{"\x14\xfb\x9c\x03\xd9\x7e", "Q<c[2!,C"},
	{"\x14\xfb\x9c\x03\xd9", "Q<c[2!B"},
	{"\x14\xfb\x9c\x03", "Q<c[A"},

	// RFC 4648 examples (adapted from base64 to base91)
	{"", ""},
	{"f", "LB"},
	{"fo", "drD"},
	{"foo", "dr.J"},
	{"foob", "dr/2Y"},
	{"fooba", "dr/2s)A"},
	{"foobar", "dr/2s)uC"},
}

func TestEncode(t *testing.T) {
	for i, p := range pairs {
		t.Run(fmt.Sprintf("case_%d", i), func(t *testing.T) {
			dst := make([]byte, StdEncoding.EncodedLen(len(p.decoded)))

			n := StdEncoding.Encode(dst, []byte(p.decoded))
			got := dst[:n]
			if !bytes.Equal(got, []byte(p.encoded)) {
				assert.Equal(t, []byte(p.encoded), got)
			}
		})
	}
}

func TestDecode(t *testing.T) {
	for i, p := range pairs {
		t.Run(fmt.Sprintf("case_%d", i), func(t *testing.T) {
			dst := make([]byte, StdEncoding.DecodedLen(len(p.encoded)))

			n, err := StdEncoding.Decode(dst, []byte(p.encoded))
			if err != nil {
				t.Errorf("Got decoding error: %v", err)
			} else {
				got := dst[:n]
				if !bytes.Equal(got, []byte(p.decoded)) {
					assert.Equal(t, []byte(p.decoded), got)
				}
			}
		})
	}
}

func TestError(t *testing.T) {
	tests := []string{
		"~_1H=x_t{ |$AjJX(nMFdjL~:?1b3HgM", // Spaces are not in the standard encoding alphabet.
		"-", "\\", "'",                     // These characters are not in the standard encoding alphabet.
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("case_%d", i), func(t *testing.T) {
			dst := make([]byte, StdEncoding.DecodedLen(len(test)))

			_, err := StdEncoding.Decode(dst, []byte(test))
			assert.Equal(t, invalidCharacterError(), err)
		})
	}
}
