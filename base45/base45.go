// Package base45 implements base45 encoding, fork from https://github.com/xkmsoft/base45
package base45

import (
	"encoding/binary"
	"fmt"
	"strings"
)

const (
	base       = 45
	baseSquare = 45 * 45
	maxUint16  = 0xFFFF
)

type InvalidLengthError struct {
	length int
	mod    int
}

func (e InvalidLengthError) Error() string {
	return fmt.Sprintf("invalid length n=%d. It should be n mod 3 = [0, 2] NOT n mod 3 = %d", e.length, e.mod)
}

type InvalidCharacterError struct {
	char     rune
	position int
}

func (e InvalidCharacterError) Error() string {
	return fmt.Sprintf("invalid character %s at position: %d\n", string(e.char), e.position)
}

type IllegalBase45ByteError struct {
	position int
}

func (e IllegalBase45ByteError) Error() string {
	return fmt.Sprintf("illegal base45 data at byte position %d\n", e.position)
}

var encodingMap = map[byte]rune{
	byte(0):  '0',
	byte(1):  '1',
	byte(2):  '2',
	byte(3):  '3',
	byte(4):  '4',
	byte(5):  '5',
	byte(6):  '6',
	byte(7):  '7',
	byte(8):  '8',
	byte(9):  '9',
	byte(10): 'A',
	byte(11): 'B',
	byte(12): 'C',
	byte(13): 'D',
	byte(14): 'E',
	byte(15): 'F',
	byte(16): 'G',
	byte(17): 'H',
	byte(18): 'I',
	byte(19): 'J',
	byte(20): 'K',
	byte(21): 'L',
	byte(22): 'M',
	byte(23): 'N',
	byte(24): 'O',
	byte(25): 'P',
	byte(26): 'Q',
	byte(27): 'R',
	byte(28): 'S',
	byte(29): 'T',
	byte(30): 'U',
	byte(31): 'V',
	byte(32): 'W',
	byte(33): 'X',
	byte(34): 'Y',
	byte(35): 'Z',
	byte(36): ' ',
	byte(37): '$',
	byte(38): '%',
	byte(39): '*',
	byte(40): '+',
	byte(41): '-',
	byte(42): '.',
	byte(43): '/',
	byte(44): ':',
}

var decodingMap = map[rune]byte{
	'0': byte(0),
	'1': byte(1),
	'2': byte(2),
	'3': byte(3),
	'4': byte(4),
	'5': byte(5),
	'6': byte(6),
	'7': byte(7),
	'8': byte(8),
	'9': byte(9),
	'A': byte(10),
	'B': byte(11),
	'C': byte(12),
	'D': byte(13),
	'E': byte(14),
	'F': byte(15),
	'G': byte(16),
	'H': byte(17),
	'I': byte(18),
	'J': byte(19),
	'K': byte(20),
	'L': byte(21),
	'M': byte(22),
	'N': byte(23),
	'O': byte(24),
	'P': byte(25),
	'Q': byte(26),
	'R': byte(27),
	'S': byte(28),
	'T': byte(29),
	'U': byte(30),
	'V': byte(31),
	'W': byte(32),
	'X': byte(33),
	'Y': byte(34),
	'Z': byte(35),
	' ': byte(36),
	'$': byte(37),
	'%': byte(38),
	'*': byte(39),
	'+': byte(40),
	'-': byte(41),
	'.': byte(42),
	'/': byte(43),
	':': byte(44),
}

// Encode
//
//  4. The Base45 Encoding
//     A 45-character subset of US-ASCII is used; the 45 characters usable
//     in a QR code in Alphanumeric mode.  Base45 encodes 2 bytes in 3
//     characters, compared to Base64, which encodes 3 bytes in 4
//     characters.
//
//     For encoding two bytes [a, b] MUST be interpreted as a number n in
//     base 256, i.e. as an unsigned integer over 16 bits so that the number
//     n = (a*256) + b.
//
//     This number n is converted to base 45 [c, d, e] so that n = c +
//     (d*45) + (e*45*45).  Note the order of c, d and e which are chosen so
//     that the left-most [c] is the least significant.
//
//     The values c, d and e are then looked up in Table 1 to produce a
//     three character string.  The process is reversed when decoding.
//
//     For encoding a single byte [a], it MUST be interpreted as a base 256
//     number, i.e. as an unsigned integer over 8 bits.  That integer MUST
//     be converted to base 45 [c d] so that a = c + (45*d).  The values c
//     and d are then looked up in Table 1 to produce a two character
//     string.
//
//     A byte string [a b c d ... x y z] with arbitrary content and
//     arbitrary length MUST be encoded as follows: From left to right pairs
//     of bytes are encoded as described above.  If the number of bytes is
//     even, then the encoded form is a string with a length which is evenly
//     divisible by 3.  If the number of bytes is odd, then the last
//     (rightmost) byte is encoded on two characters as described above.
//
//     For decoding a Base45 encoded string the inverse operations are
//     performed.
func Encode(in string) string {
	bytes := []byte(in)
	pairs := encodePairs(bytes)
	var builder strings.Builder
	for i, pair := range pairs {
		res := encodeBase45(pair)
		if i+1 == len(pairs) && res[2] == 0 {
			for _, b := range res[:2] {
				if c, ok := encodingMap[b]; ok {
					builder.WriteRune(c)
				}
			}
		} else {
			for _, b := range res {
				if c, ok := encodingMap[b]; ok {
					builder.WriteRune(c)
				}
			}
		}
	}
	return builder.String()
}

// Decode
//
//	Decoding example 1: The string "QED8WEX0" represents, when looked up
//	in Table 1, the values [26 14 13 8 32 14 33 0].  We arrange the
//	numbers in chunks of three, except for the last one which can be two,
//	and get [[26 14 13] [8 32 14] [33 0]].  In base 45 we get [26981
//	29798 33] where the bytes are [[105 101] [116 102] [33]].  If we look
//	at the ASCII values we get the string "ietf!".
func Decode(in string) (string, error) {
	size := len(in)
	mod := size % 3
	if mod != 0 && mod != 2 {
		return "", InvalidLengthError{
			length: size,
			mod:    mod,
		}
	}
	bytes := make([]byte, 0, size)
	for pos, char := range in {
		v, ok := decodingMap[char]
		if !ok {
			return "", InvalidCharacterError{
				char:     char,
				position: pos,
			}
		}
		bytes = append(bytes, v)
	}
	chunks := decodeChunks(bytes)
	triplets, err := decodeTriplets(chunks)
	if err != nil {
		return "", err
	}
	tripletsLength := len(triplets)
	decoded := make([]byte, 0, tripletsLength*2)
	for i := 0; i < tripletsLength-1; i++ {
		bytes := uint16ToBytes(triplets[i])
		decoded = append(decoded, bytes[0])
		decoded = append(decoded, bytes[1])
	}
	if mod == 2 {
		bytes := uint16ToBytes(triplets[tripletsLength-1])
		decoded = append(decoded, bytes[1])
	} else {
		bytes := uint16ToBytes(triplets[tripletsLength-1])
		decoded = append(decoded, bytes[0])
		decoded = append(decoded, bytes[1])
	}
	return string(decoded), nil
}

func uint16ToBytes(in uint16) []byte {
	bytes := make([]byte, 2)
	binary.BigEndian.PutUint16(bytes, in)
	return bytes
}

func decodeChunks(in []byte) [][]byte {
	size := len(in)
	ret := make([][]byte, 0, size/2)
	for i := 0; i < size; i += 3 {
		var f, s, l byte
		if i+2 < size {
			f = in[i]
			s = in[i+1]
			l = in[i+2]
			ret = append(ret, []byte{f, s, l})
		} else {
			f = in[i]
			s = in[i+1]
			ret = append(ret, []byte{f, s})
		}
	}
	return ret
}

func encodePairs(in []byte) [][]byte {
	size := len(in)
	ret := make([][]byte, 0, size/2)
	for i := 0; i < size; i += 2 {
		var high, low byte
		if i+1 < size {
			high = in[i]
			low = in[i+1]
		} else {
			low = in[i]
		}
		ret = append(ret, []byte{high, low})
	}
	return ret
}

func encodeBase45(in []byte) []byte {
	n := binary.BigEndian.Uint16(in)
	c := n % base
	e := (n - c) / (baseSquare)
	d := (n - (c + (e * baseSquare))) / base
	return []byte{byte(c), byte(d), byte(e)}
}

func decodeTriplets(in [][]byte) ([]uint16, error) {
	size := len(in)
	ret := make([]uint16, 0, size)
	for pos, chunk := range in {
		if len(chunk) == 3 {
			// n = c + (d*45) + (e*45*45)
			c := int(chunk[0])
			d := int(chunk[1])
			e := int(chunk[2])
			n := c + (d * base) + (e * baseSquare)
			if n > maxUint16 {
				return nil, IllegalBase45ByteError{position: pos}
			}
			ret = append(ret, uint16(n))
		}
		if len(chunk) == 2 {
			// n = c + (d*45)
			c := uint16(chunk[0])
			d := uint16(chunk[1])
			n := c + (d * base)
			ret = append(ret, n)
		}
	}
	return ret, nil
}
