package base62

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEncode(t *testing.T) {
	for _, s := range SamplesStd {
		encoded := StdEncoding.Encode([]byte(s.source))
		if bytes.Equal(encoded, s.targetBytes) {
			t.Logf("source: %-15s\ttarget: %s", s.source, s.target)
		} else {
			str := string(encoded)
			t.Errorf("source: %-15s\texpected target: %s(%d)\tactual target: %s(%d)", s.source, s.target, len(s.target), str, len(str))
		}
	}
}

func TestDecode(t *testing.T) {
	for _, s := range SamplesStd {
		decoded, err := StdEncoding.Decode(s.targetBytes)
		if err != nil {
			t.Error(err)
			continue
		}
		if bytes.Equal(decoded, s.sourceBytes) {
			t.Logf("target: %-15s\tsource: %s", s.target, s.source)
		} else {
			str := string(decoded)
			t.Errorf("target: %-15s\texpected source: %s(%d)\tactual source: %s(%d)", s.target, s.source, len(s.source), str, len(str))
		}
	}
}

func TestDecodeWithNewLine(t *testing.T) {
	for _, s := range SamplesWithNewLine {
		decoded, err := StdEncoding.Decode(s.targetBytes)
		if err != nil {
			t.Error(err)
			continue
		}
		if bytes.Equal(decoded, s.sourceBytes) {
			t.Logf("target: %-15s\tsource: %s", s.target, s.source)
		} else {
			str := string(decoded)
			t.Errorf("target: %-15s\texpected source: %s(%d)\tactual source: %s(%d)", s.target, s.source, len(s.source), str, len(str))
		}
	}
}

func TestDecodeError(t *testing.T) {
	for _, s := range SamplesErr {
		decoded, err := StdEncoding.Decode(s.targetBytes)
		if err != nil {
			t.Logf("%s: \"%c\"", err.Error(), err)
			continue
		}
		str := string(decoded)
		t.Errorf("An error should have occurred, instead of returning \"%s\"", str)
	}
}

func TestEncodeWithCustomAlphabet(t *testing.T) {
	for _, s := range SamplesWithAlphabet {
		encoded := NewEncoding(s.alphabet).Encode([]byte(s.source))
		if bytes.Equal(encoded, s.targetBytes) {
			t.Logf("source: %-15s\ttarget: %s", s.source, s.target)
		} else {
			str := string(encoded)
			t.Errorf("source: %-15s\texpected target: %s(%d)\tactual target: %s(%d)", s.source, s.target, len(s.target), str, len(str))
		}
	}
}

func TestDecodeWithCustomAlphabet(t *testing.T) {
	for _, s := range SamplesWithAlphabet {
		decoded, err := NewEncoding(s.alphabet).Decode(s.targetBytes)
		if err != nil {
			t.Error(err)
			continue
		}
		if bytes.Equal(decoded, s.sourceBytes) {
			t.Logf("target: %-15s\tsource: %s", s.target, s.source)
		} else {
			str := string(decoded)
			t.Errorf("target: %-15s\texpected source: %s(%d)\tactual source: %s(%d)", s.target, s.source, len(s.source), str, len(str))
		}
	}
}

func NewSample(source, target string) *Sample {
	return &Sample{source: source, target: target, sourceBytes: []byte(source), targetBytes: []byte(target)}
}

func NewSampleWithAlphabet(source, target, alphabet string) *Sample {
	return &Sample{source: source, target: target, sourceBytes: []byte(source), targetBytes: []byte(target), alphabet: alphabet}
}

type Sample struct {
	source      string
	target      string
	sourceBytes []byte
	targetBytes []byte
	alphabet    string
}

var SamplesStd = []*Sample{
	NewSample("", ""),
	NewSample("f", "1e"),
	NewSample("fo", "6ox"),
	NewSample("foo", "SAPP"),
	NewSample("foob", "1sIyuo"),
	NewSample("fooba", "7kENWa1"),
	NewSample("foobar", "VytN8Wjy"),

	NewSample("su", "7gj"),
	NewSample("sur", "VkRe"),
	NewSample("sure", "275mAn"),
	NewSample("sure.", "8jHquZ4"),
	NewSample("asure.", "UQPPAab8"),
	NewSample("easure.", "26h8PlupSA"),
	NewSample("leasure.", "9IzLUOIY2fe"),

	NewSample("=", "z"),
	NewSample(">", "10"),
	NewSample("?", "11"),
	NewSample("11", "3H7"),
	NewSample("111", "DWfh"),
	NewSample("1111", "tquAL"),
	NewSample("11111", "3icRuhV"),
	NewSample("111111", "FMElG7cn"),

	NewSample("Hello, World!", "1wJfrzvdbtXUOlUjUf"),
	NewSample("你好，世界！", "1ugmIChyMAcCbDRpROpAtpXdp"),
	NewSample("こんにちは", "1fyB0pNlcVqP3tfXZ1FmB"),
	NewSample("안녕하십니까", "1yl6dfHPaO9hroEXU9qFioFhM"),
}

var SamplesWithAlphabet = []*Sample{
	NewSampleWithAlphabet("", "", "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"),
	NewSampleWithAlphabet("f", "Bo", "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"),
	NewSampleWithAlphabet("fo", "Gy7", "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"),
	NewSampleWithAlphabet("foo", "cKZZ", "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"),
	NewSampleWithAlphabet("foob", "B2S84y", "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"),
	NewSampleWithAlphabet("fooba", "HuOXgkB", "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"),
	NewSampleWithAlphabet("foobar", "f83XIgt8", "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"),

	NewSampleWithAlphabet("su", "Hqt", "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"),
	NewSampleWithAlphabet("sur", "fubo", "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"),
	NewSampleWithAlphabet("sure", "CHFwKx", "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"),
	NewSampleWithAlphabet("sure.", "ItR04jE", "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"),
	NewSampleWithAlphabet("asure.", "eaZZKklI", "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"),
	NewSampleWithAlphabet("easure.", "CGrIZv4zcK", "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"),
	NewSampleWithAlphabet("leasure.", "JS9VeYSiCpo", "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"),

	NewSampleWithAlphabet("=", "9", "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"),
	NewSampleWithAlphabet(">", "BA", "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"),
	NewSampleWithAlphabet("?", "BB", "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"),
	NewSampleWithAlphabet("11", "DRH", "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"),
	NewSampleWithAlphabet("111", "Ngpr", "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"),
	NewSampleWithAlphabet("1111", "304KV", "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"),
	NewSampleWithAlphabet("11111", "Dsmb4rf", "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"),
	NewSampleWithAlphabet("111111", "PWOvQHmx", "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"),

	NewSampleWithAlphabet("Hello, World!", "B6Tp195nl3heYvetep", "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"),
	NewSampleWithAlphabet("你好，世界！", "B4qwSMr8WKmMlNbzbYzK3zhnz", "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"),
	NewSampleWithAlphabet("こんにちは", "Bp8LAzXvmf0ZD3phjBPwL", "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"),
	NewSampleWithAlphabet("안녕하십니까", "B8vGnpRZkYJr1yOheJ0PsyPrW", "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"),
}

var SamplesWithNewLine = []*Sample{
	NewSample("111111", "FMEl\rG7cn"),
	NewSample("Hello, World!", "1wJfrzvdb\ntXUOlUjUf"),
	NewSample("你好，世界！", "1ugmIChyMAcC\rbDRpROpAtpXdp"),
}

var SamplesErr = []*Sample{
	NewSample("", "Hello, World!"),
	NewSample("", "哈哈"),
	NewSample("", "はは"),
}

func TestEncodeAlphabetErr(t *testing.T) {
	alphabet := "xxxx"
	enc := NewEncoding(alphabet)
	fmt.Println(enc.Error)
	assert.Equal(t, InvalidAlphabetError(), enc.Error)
}
