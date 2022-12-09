// Package morse implements morse encoding, fork from https://github.com/martinlindhe/morse
package morse

import (
	"errors"
	"fmt"
	"strings"
	"unsafe"
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

// Encode encodes clear text using `alphabet` mapping
func Encode(b []byte, separator string) (dst string, err error) {
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

// Decode decodes morse code using `alphabet` mapping
func Decode(b []byte, separator string) (dst string, err error) {
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

// converts string to byte slice without a memory allocation.
// 将字符串转换为字节切片
func string2bytes(s string) []byte {
	if len(s) == 0 {
		return nil
	}
	return *(*[]byte)(unsafe.Pointer(
		&struct {
			string
			Cap int
		}{s, len(s)},
	))
}

// converts byte slice to string without a memory allocation.
// 将字节切片转换为字符串
func bytes2string(b []byte) string {
	if len(b) == 0 {
		return ""
	}
	return *(*string)(unsafe.Pointer(&b))
}
