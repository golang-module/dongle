package dongle

import (
	"encoding/base32"
)

// ByBase32 encodes by base32.
func (e encode) ByBase32() encode {
	input, output := e.input, e.output
	if input == nil {
		return e
	}
	if output != nil {
		input = output
	}
	e.output = string2bytes(base32.StdEncoding.EncodeToString(input))
	return e
}

// ByBase32 encodes by base32.
func (d decode) ByBase32() decode {
	input, output := d.input, d.output
	if input == nil {
		return d
	}
	if output != nil {
		input = output
	}
	d.output, d.Error = base32.StdEncoding.DecodeString(bytes2string(input))
	return d
}
