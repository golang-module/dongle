package dongle

import (
	"crypto/cipher"
)

// decrypt defines decrypt struct
type decrypt struct {
	dongle
	Cipher *Cipher
}

// newDecrypt returns a new decrypt instance.
func newDecrypt() decrypt {
	return decrypt{Cipher: NewCipher()}
}

// FromString encrypts from string.
func (d decrypt) FromString(s string, decodingMode ...string) decrypt {
	if s == emptyString {
		return d
	}
	mode := HEX
	if len(decodingMode) > 0 {
		mode = decodingMode[len(decodingMode)-1]
	}
	d.input = string2bytes(s)
	d.input, d.Error = d.decode(mode)
	return d
}

// FromBytes encrypts from byte slice.
func (d decrypt) FromBytes(b []byte, decodingMode ...string) decrypt {
	if len(b) == 0 {
		return d
	}
	mode := HEX
	if len(decodingMode) > 0 {
		mode = decodingMode[len(decodingMode)-1]
	}
	d.input = b
	d.input, d.Error = d.decode(mode)
	return d
}

// ToString outputs as string.
func (d decrypt) ToString() string {
	input, output := d.input, d.output
	if len(input) == 0 {
		return emptyString
	}
	return bytes2string(output)
}

// ToBytes outputs as byte slice.
func (d decrypt) ToBytes() []byte {
	input, output := d.input, d.output
	if len(input) == 0 {
		return emptyBytes
	}
	return output
}

func (d decrypt) decode(decodingMode string) ([]byte, error) {
	switch decodingMode {
	case HEX:
		hex := Decode.FromBytes(d.input).ByHex()
		return hex.ToBytes(), hex.Error
	case BASE32:
		base32 := Decode.FromBytes(d.input).ByBase32()
		return base32.ToBytes(), base32.Error
	case BASE58:
		base58 := Decode.FromBytes(d.input).ByBase58()
		return base58.ToBytes(), base58.Error
	case BASE64:
		base64 := Decode.FromBytes(d.input).ByBase64()
		return base64.ToBytes(), base64.Error
	default:
		return emptyBytes, invalidDecodingModeError(decodingMode)
	}
}

func (d decrypt) decrypt(cipher *Cipher, block cipher.Block) ([]byte, error) {
	dst := emptyBytes
	switch cipher.mode {
	case CBC:
		dst = cipher.CBCDecrypt(block, d.input)
	default:
		return dst, invalidGroupModeError(cipher.mode)
	}
	switch cipher.padding {
	case ZeroPadding:
		return cipher.ZeroUnPadding(dst), nil
	case PKCS5Padding:
		return cipher.PKCS5UnPadding(dst), nil
	case PKCS7Padding:
		return cipher.PKCS7UnPadding(dst), nil
	}
	return dst, invalidPaddingModeError(cipher.padding)
}
