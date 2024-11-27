package dongle

import "fmt"

type Decoder struct {
	dongle
}

type DecodeError struct {
}

func NewDecodeError() DecodeError {
	return DecodeError{}
}

func (e DecodeError) ModeError(mode string) error {
	return fmt.Errorf("decode: invalid decoding mode, the src can't be decoded by %s", mode)
}

var decodeError = NewDecodeError()

// newDecoder returns a new Decoder instance.
func newDecoder() Decoder {
	return Decoder{}
}

// FromString decodes from string.
func (d Decoder) FromString(s string) Decoder {
	d.src = string2bytes(s)
	return d
}

// FromBytes decodes from byte slice.
func (d Decoder) FromBytes(b []byte) Decoder {
	d.src = b
	return d
}

// String implements Stringer interface for Decoder struct.
func (d Decoder) String() string {
	return d.ToString()
}

// ToString outputs as string.
func (d Decoder) ToString() string {
	return bytes2string(d.dst)
}

// ToBytes outputs as byte slice.
func (d Decoder) ToBytes() []byte {
	if len(d.dst) == 0 {
		return []byte("")
	}
	return d.dst
}
