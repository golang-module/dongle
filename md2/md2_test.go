package md2

import (
	"fmt"
	"io"
	"testing"
)

var tests = []struct {
	input  string
	output string
}{
	{"", "8350e5a3e24c153df2275c9f80692773"},
	{"a", "32ec01ec4a6dac72c0ab96fb34c0b5d1"},
	{"abc", "da853b0d3f88d99b30283a69e6ded6bb"},
	{"hello world", "d9cce882ee690a5c1ce70beff3a78c77"},
	{"message digest", "ab4f496bfb2a530b219ff33031fe06b0"},
	{"abcdefghijklmnopqrstuvwxyz", "4e8ddff3650292ab5a4108c3aa47940b"},
	{"ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789", "da33def2a42df13975352846c30338cd"},
}

func TestMd2(t *testing.T) {
	for i := 0; i < len(tests); i++ {
		g := tests[i]
		c := New()
		for j := 0; j < 3; j++ {
			if j < 2 {
				io.WriteString(c, g.input)
			} else {
				io.WriteString(c, g.input[0:len(g.input)/2])
				c.Sum(nil)
				io.WriteString(c, g.input[len(g.input)/2:])
			}
			s := fmt.Sprintf("%x", c.Sum(nil))
			if s != g.output {
				t.Fatalf("md2[%d](%s) = %s want %s", j, g.input, s, g.output)
			} else {
				fmt.Printf("md2[%d](%s) = %s want %s - Passed\n", j, g.input, s, g.output)
			}
			c.Reset()
		}
	}
}

func TestDigest_Size(t *testing.T) {
	d := New()
	got, want := d.Size(), Size
	if got != want {
		t.Errorf("The values of %v is not %v\n", got, want)
	}
}

func TestDigest_BlockSize(t *testing.T) {
	d := New()
	got, want := d.BlockSize(), BlockSize
	if got != want {
		t.Errorf("The values of %v is not %v\n", got, want)
	}
}
