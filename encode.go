package dongle

const (
	// BASE32 base32 encoding mode
	BASE32 = "base32"
	// BASE64 base64 encoding mode
	BASE64 = "base64"
	// HEX hex encoding mode
	HEX = "hex"
)

// encode defines encode struct.
type encode struct {
	dongle
}

// newEncode returns a new encode instance.
func newEncode() encode {
	return encode{}
}

// FromString encrypts from string.
// 对字符串进行加密
func (e encode) FromString(s string) encode {
	if s == "" {
		return e
	}
	e.input = string2bytes(s)
	return e
}

// FromBytes encrypts from byte slice.
func (e encode) FromBytes(b []byte) encode {
	if len(b) > 0 {
		e.input = b
	}
	return e
}

// ToString outputs as string.
func (e encode) ToString() string {
	output := e.output
	if output == nil {
		output = e.input
	}
	return bytes2string(output)
}

// ToBytes outputs as byte slice.
func (e encode) ToBytes() []byte {
	input, output := e.input, e.output
	if len(input) == 0 {
		return []byte("")
	}
	return output
}
