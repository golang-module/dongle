package dongle

// defines a decode struct.
type decode struct {
	dongle
}

// returns a new decode instance.
func newDecode() decode {
	return decode{}
}

// FromString decodes from string.
func (d decode) FromString(s string) decode {
	d.src = string2bytes(s)
	return d
}

// FromBytes decodes from byte slice.
func (d decode) FromBytes(b []byte) decode {
	d.src = b
	return d
}

// String implements the interface Stringer for decode struct.
// 实现 Stringer 接口
func (d decode) String() string {
	return d.ToString()
}

// ToString decodes as string.
func (d decode) ToString() string {
	return bytes2string(d.dst)
}

// ToBytes decodes as byte slice.
func (d decode) ToBytes() []byte {
	if len(d.dst) == 0 {
		return []byte("")
	}
	return d.dst
}
