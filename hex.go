package dongle

import (
	"encoding/hex"
)

// ByHex encodes by hex.
func (e encode) ByHex() encode {
	input, output := e.input, e.output
	if input == nil {
		return e
	}
	if output != nil {
		input = output
	}
	e.output = string2bytes(hex.EncodeToString(input))
	return e
}

// ByHex decodes by base64.
func (d decode) ByHex() decode {
	input, output := d.input, d.output
	if input == nil {
		return d
	}
	if output != nil {
		input = output
	}
	d.output, d.Error = hex.DecodeString(bytes2string(input))
	return d
}
