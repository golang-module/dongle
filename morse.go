package dongle

// reference: https://github.com/boratanrikulu/morse

import (
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
func morseEncode(b []byte, letterSeparator string) (dst string, err error) {
	s := strings.ToLower(bytes2string(b))
	if strings.Contains(s, " ") {
		return dst, invalidMorsePlaintextError()
	}
	for _, letter := range s {
		let := string(letter)
		if morseLetter[let] != "" {
			dst += morseLetter[let] + letterSeparator
		}
	}
	dst = strings.Trim(dst, letterSeparator)
	return
}

// morseDecode decodes morse code using `alphabet` mapping
func morseDecode(b []byte, letterSeparator string) (dst string, err error) {
	for _, part := range strings.Split(bytes2string(b), letterSeparator) {
		found := false
		for key, letter := range morseLetter {
			if letter == part {
				dst += key
				found = true
				break
			}
		}
		if !found {
			return dst, invalidMorseCiphertextError()
		}
	}
	return
}

// ByMorse encodes by morse.
// 通过 morse 编码
func (e encode) ByMorse() encode {
	if len(e.src) == 0 {
		return e
	}
	dst, err := morseEncode(e.src, "/")
	if err != nil {
		e.Error = err
		return e
	}
	e.dst = string2bytes(dst)
	return e
}

// ByMorse decodes by morse.
// 通过 morse 解码
func (d decode) ByMorse() decode {
	if len(d.src) == 0 || d.Error != nil {
		return d
	}
	dst, err := morseDecode(d.src, "/")
	if err != nil {
		d.Error = err
		return d
	}
	d.dst = string2bytes(dst)
	return d
}
