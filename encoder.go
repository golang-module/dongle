package dongle

type Encoder struct {
	dongle
}

// newEncoder returns a new Encoder instance.
func newEncoder() Encoder {
	return Encoder{}
}

// FromString encodes from string.
func (e Encoder) FromString(s string) Encoder {
	e.src = string2bytes(s)
	return e
}

// FromBytes encodes from byte slice.
func (e Encoder) FromBytes(b []byte) Encoder {
	e.src = b
	return e
}

// String implements Stringer interface for Encoder struct.
func (e Encoder) String() string {
	return e.ToString()
}

// ToString outputs as string.
func (e Encoder) ToString() string {
	return bytes2string(e.dst)
}

// ToBytes outputs as byte slice.
func (e Encoder) ToBytes() []byte {
	if len(e.dst) == 0 {
		return []byte("")
	}
	return e.dst
}
