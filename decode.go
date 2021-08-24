package dongle

// decode defines a decode struct.
type decode struct {
	dongle
}

// newDecode returns a new decode instance.
func newDecode() decode {
	return decode{}
}

// FromString decodes from string.
func (d decode) FromString(s string) decode {
	if s == "" {
		return d
	}
	d.input = string2bytes(s)
	return d
}

// FromBytes decodes from byte slice.
func (d decode) FromBytes(b []byte) decode {
	if len(b) > 0 {
		d.input = b
	}
	return d
}

// ToString decodes as string.
func (d decode) ToString() string {
	output := d.output
	if output == nil {
		output = d.input
	}
	return bytes2string(output)
}

// ToBytes decodes as byte slice.
func (d decode) ToBytes() []byte {
	input, output := d.input, d.output
	if len(input) == 0 {
		return emptyBytes
	}
	return output
}
