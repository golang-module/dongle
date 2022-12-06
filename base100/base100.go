// Package base100 implements base100 encoding, fork from https://github.com/stek29/base100
package base100

// Encode encodes by base100.
// 通过 base100 编码
func Encode(src []byte) []byte {
	buf := make([]byte, len(src)*4)
	for i, v := range src {
		buf[i*4+0] = 0xf0
		buf[i*4+1] = 0x9f
		buf[i*4+2] = byte((uint16(v)+55)/64 + 0x8f)
		buf[i*4+3] = (v+55)%64 + 0x80
	}
	return buf
}

// Decode decodes by base100.
// 通过 base100 解码
func Decode(src []byte) ([]byte, error) {
	if len(src)%4 != 0 {
		return nil, invalidLengthError
	}
	buf := make([]byte, len(src)/4)
	for i := 0; i != len(src); i += 4 {
		if src[i+0] != 0xf0 || src[i+1] != 0x9f {
			return nil, invalidDataError
		}
		buf[i/4] = (src[i+2]-0x8f)*64 + src[i+3] - 0x80 - 55
	}
	return buf, nil
}

// invalidInputError is returned when Decode fails
type invalidInputError struct {
	message string
}

func (e invalidInputError) Error() string {
	return e.message
}

// invalidLengthError is returned when length of string being decoded is
// not divisible by four
var invalidLengthError = invalidInputError{"len(data) should be divisible by 4"}

// invalidDataError is returned if data is not a valid base100 string
var invalidDataError = invalidInputError{"data is invalid"}
