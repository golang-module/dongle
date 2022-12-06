// Package base91 implements base91 encoding, fork from https://github.com/mtraver/base91
package base91

import (
	"fmt"
	"math"
)

// An Encoding is a base 91 encoding/decoding scheme defined by a 91-character alphabet.
type Encoding struct {
	encode    [91]byte
	decodeMap [256]byte
}

// encodeStd is the standard base91 encoding alphabet (that is, the one specified
// at http://base91.sourceforge.net). Of the 95 printable ASCII characters, the
// following four are omitted: space (0x20), apostrophe (0x27), hyphen (0x2d),
// and backslash (0x5c).
const encodeStd = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!#$%&()*+,./:;<=>?@[]^_`{|}~\""

// newEncoding returns a new Encoding defined by the given alphabet, which must
// be a 91-byte string that does not contain CR or LF ('\r', '\n').
func newEncoding(encoder string) *Encoding {
	e := new(Encoding)
	copy(e.encode[:], encoder)

	for i := 0; i < len(e.decodeMap); i++ {
		// 0xff indicates that this entry in the decode map is not in the encoding alphabet.
		e.decodeMap[i] = 0xff
	}
	for i := 0; i < len(encoder); i++ {
		e.decodeMap[encoder[i]] = byte(i)
	}
	return e
}

// StdEncoding is the standard base91 encoding (that is, the one specified
// at http://base91.sourceforge.net). Of the 95 printable ASCII characters,
// the following four are omitted: space (0x20), apostrophe (0x27),
// hyphen (0x2d), and backslash (0x5c).
var StdEncoding = newEncoding(encodeStd)

// Encode encodes src using the encoding enc, writing bytes to dst.
// It returns the number of bytes written, because the exact output size cannot
// be known before encoding takes place. EncodedLen(len(src)) may be used to
// determine an upper bound on the output size when allocating a dst slice.
func (enc *Encoding) Encode(dst, src []byte) int {
	var queue, numBits uint

	n := 0
	for i := 0; i < len(src); i++ {
		queue |= uint(src[i]) << numBits
		numBits += 8
		if numBits > 13 {
			var v = queue & 8191

			if v > 88 {
				queue >>= 13
				numBits -= 13
			} else {
				// We can take 14 bits.
				v = queue & 16383
				queue >>= 14
				numBits -= 14
			}
			dst[n] = enc.encode[v%91]
			n++
			dst[n] = enc.encode[v/91]
			n++
		}
	}

	if numBits > 0 {
		dst[n] = enc.encode[queue%91]
		n++

		if numBits > 7 || queue > 90 {
			dst[n] = enc.encode[queue/91]
			n++
		}
	}

	return n
}

// EncodedLen returns an upper bound on the length in bytes of the base91 encoding
// of an input buffer of length n. The true encoded length may be shorter.
func (enc *Encoding) EncodedLen(n int) int {
	// At worst, base91 encodes 13 bits into 16 bits. Even though 14 bits can
	// sometimes be encoded into 16 bits, assume the worst case to get the upper
	// bound on encoded length.
	return int(math.Ceil(float64(n) * 16.0 / 13.0))
}

// A invalidCharacterError is returned if invalid base91 data is encountered during decoding.

var invalidCharacterError = func() error {
	return fmt.Errorf("base91: invalid character, the he character is not in the encoding alphabet")
}

// Decode decodes src using the encoding enc. It writes at most DecodedLen(len(src))
// bytes to dst and returns the number of bytes written. If src contains invalid base91
// data, it will return the number of bytes successfully written and CorruptInputError.
func (enc *Encoding) Decode(dst, src []byte) (int, error) {
	var queue, numBits uint
	var v = -1

	n := 0
	for i := 0; i < len(src); i++ {
		if enc.decodeMap[src[i]] == 0xff {
			// The character is not in the encoding alphabet.
			return n, invalidCharacterError()
		}

		if v == -1 {
			// Start the next value.
			v = int(enc.decodeMap[src[i]])
		} else {
			v += int(enc.decodeMap[src[i]]) * 91
			queue |= uint(v) << numBits

			if (v & 8191) > 88 {
				numBits += 13
			} else {
				numBits += 14
			}

			for ok := true; ok; ok = numBits > 7 {
				dst[n] = byte(queue)
				n++

				queue >>= 8
				numBits -= 8
			}

			// Mark this value complete.
			v = -1
		}
	}

	if v != -1 {
		dst[n] = byte(queue | uint(v)<<numBits)
		n++
	}

	return n, nil
}

// DecodedLen returns the maximum length in bytes of the decoded data
// corresponding to n bytes of base91-encoded data.
func (enc *Encoding) DecodedLen(n int) int {
	// At best, base91 encodes 14 bits into 16 bits, so assume that the input is
	// optimally encoded to get the upper bound on decoded length.
	return int(math.Ceil(float64(n) * 14.0 / 16.0))
}
