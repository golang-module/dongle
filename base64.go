package dongle

import (
	"encoding/base64"
)

// ByBase64 encodes by base64.
func (e encode) ByBase64() encode {
	input, output := e.input, e.output
	if input == nil {
		return e
	}
	if output != nil {
		input = output
	}
	e.output = string2bytes(base64.StdEncoding.EncodeToString(input))
	return e
}

// ByBase64 encodes by base64.
func (d decode) ByBase64() decode {
	input, output := d.input, d.output
	if input == nil {
		return d
	}
	if output != nil {
		input = output
	}
	d.output, d.Error = base64.StdEncoding.DecodeString(bytes2string(input))
	return d
}
