package dongle

// fork from https://github.com/martinlindhe/morse

import (
	"errors"
	"fmt"
	"strings"
)

var (
	morseLetter = map[string]string{
		"a":  ".-",
		"b":  "-...",
		"c":  "-.-.",
		"d":  "-..",
		"e":  ".",
		"f":  "..-.",
		"g":  "--.",
		"h":  "....",
		"i":  "..",
		"j":  ".---",
		"k":  "-.-",
		"l":  ".-..",
		"m":  "--",
		"n":  "-.",
		"o":  "---",
		"p":  ".--.",
		"q":  "--.-",
		"r":  ".-.",
		"s":  "...",
		"t":  "-",
		"u":  "..-",
		"v":  "...-",
		"w":  ".--",
		"x":  "-..-",
		"y":  "-.--",
		"z":  "--..",
		"ä":  ".-.-",
		"å":  ".-.-",
		"ç":  "-.-..",
		"ĉ":  "-.-..",
		"ö":  "-.-..",
		"ø":  "---.",
		"ð":  "..--.",
		"ü":  "..--",
		"ŭ":  "..--",
		"ch": "----",
		"0":  "-----",
		"1":  ".----",
		"2":  "..---",
		"3":  "...--",
		"4":  "....-",
		"5":  ".....",
		"6":  "-....",
		"7":  "--...",
		"8":  "---..",
		"9":  "----.",
		".":  ".-.-.-",
		",":  "--..--",
		"`":  ".----.",
		"?":  "..--..",
		"!":  "..--.",
		":":  "---...",
		";":  "-.-.-",
		"\"": ".-..-.",
		"'":  ".----.",
		"=":  "-...-",
		"(":  "-.--.",
		")":  "-.--.-",
		"$":  "...-..-",
		"&":  ".-...",
		"@":  ".--.-.",
		"+":  ".-.-.",
		"-":  "-....-",
		"/":  "-..-.",
	}
)

// morseEncode encodes clear text using `alphabet` mapping
func morseEncode(b []byte, separator string) (dst string, err error) {
	s := strings.ToLower(bytes2string(b))
	if strings.Contains(s, " ") {
		return dst, errors.New("can't contain spaces")
	}
	for _, letter := range s {
		let := string(letter)
		if morseLetter[let] != "" {
			dst += morseLetter[let] + separator
		}
	}
	dst = strings.Trim(dst, separator)
	return
}

// morseDecode decodes morse code using `alphabet` mapping
func morseDecode(b []byte, separator string) (dst string, err error) {
	for _, part := range strings.Split(bytes2string(b), separator) {
		found := false
		for key, letter := range morseLetter {
			if letter == part {
				dst += key
				found = true
				break
			}
		}
		if !found {
			return dst, fmt.Errorf("unknown character " + part)
		}
	}
	return
}

// ByMorse encodes by morse.
// 通过 morse 编码
func (e encoder) ByMorse(separator ...string) encoder {
	if len(e.src) == 0 {
		return e
	}
	if len(separator) == 0 {
		separator = []string{"/"}
	}
	dst, err := morseEncode(e.src, separator[0])
	if err != nil {
		e.Error = invalidMorseSrcError()
		return e
	}
	e.dst = string2bytes(dst)
	return e
}

// ByMorse decodes by morse.
// 通过 morse 解码
func (d decoder) ByMorse(separator ...string) decoder {
	if len(d.src) == 0 || d.Error != nil {
		return d
	}
	if len(separator) == 0 {
		separator = []string{"/"}
	}
	dst, err := morseDecode(d.src, separator[0])
	if err != nil {
		d.Error = morseDecodeError()
		return d
	}
	d.dst = string2bytes(dst)
	return d
}
