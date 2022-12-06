// Package base62 implements base62 encoding, fork from https://github.com/yihleego/base62
package base62

import (
	"math"
	"strconv"
)

// An Encoding is a radix 62 encoding/decoding scheme, defined by a 62-character alphabet.
type Encoding struct {
	encode    [62]byte
	decodeMap [256]byte
}

const encodeStd = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

// newEncoding returns a new padded Encoding defined by the given alphabet,
// which must be a 62-byte string that does not contain the padding character
// or CR / LF ('\r', '\n').
func newEncoding(encoder string) *Encoding {
	e := new(Encoding)
	copy(e.encode[:], encoder)

	for i := 0; i < len(e.decodeMap); i++ {
		e.decodeMap[i] = 0xFF
	}
	for i := 0; i < len(encoder); i++ {
		e.decodeMap[encoder[i]] = byte(i)
	}
	return e
}

// StdEncoding is the standard base62 encoding.
var StdEncoding = newEncoding(encodeStd)

// Encode encodes src using the encoding enc.
func (enc *Encoding) Encode(src []byte) []byte {
	if len(src) == 0 {
		return nil
	}

	// enc is a pointer receiver, so the use of enc.encode within the hot
	// loop below means a nil check at every operation. Lift that nil check
	// outside the loop to speed up the encoder.
	_ = enc.encode

	rs := 0
	cs := int(math.Ceil(math.Log(256) / math.Log(62) * float64(len(src))))
	dst := make([]byte, cs)
	for i := range src {
		c := 0
		v := int(src[i])
		for j := cs - 1; j >= 0 && (v != 0 || c < rs); j-- {
			v += 256 * int(dst[j])
			dst[j] = byte(v % 62)
			v /= 62
			c++
		}
		rs = c
	}
	for i := range dst {
		dst[i] = enc.encode[dst[i]]
	}
	if cs > rs {
		return dst[cs-rs:]
	}
	return dst
}

// Decode decodes src using the encoding enc.
// If src contains invalid base62 data, it will return the
// number of bytes successfully written and CorruptInputError.
// New line characters (\r and \n) are ignored.
func (enc *Encoding) Decode(src []byte) ([]byte, error) {
	if len(src) == 0 {
		return nil, nil
	}

	// Lift the nil check outside the loop. enc.decodeMap is directly
	// used later in this function, to let the compiler know that the
	// receiver can't be nil.
	_ = enc.decodeMap

	rs := 0
	cs := int(math.Ceil(math.Log(62) / math.Log(256) * float64(len(src))))
	dst := make([]byte, cs)
	for i := range src {
		if src[i] == '\n' || src[i] == '\r' {
			continue
		}
		c := 0
		v := int(enc.decodeMap[src[i]])
		if v == 255 {
			return nil, CorruptInputError(src[i])
		}
		for j := cs - 1; j >= 0 && (v != 0 || c < rs); j-- {
			v += 62 * int(dst[j])
			dst[j] = byte(v % 256)
			v /= 256
			c++
		}
		rs = c
	}
	if cs > rs {
		return dst[cs-rs:], nil
	}
	return dst, nil
}

type CorruptInputError int64

func (e CorruptInputError) Error() string {
	return "illegal base62 data at input byte " + strconv.FormatInt(int64(e), 10)
}
