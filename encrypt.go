package dongle

import (
	"crypto/cipher"
	"io/ioutil"
)

// encrypt defines a encrypt struct.
type encrypt struct {
	dongle
}

// newEncrypt returns a new encrypt instance.
func newEncrypt() encrypt {
	return encrypt{}
}

// FromString encrypts from string.
func (e encrypt) FromString(s string) encrypt {
	if s == "" {
		return e
	}
	e.input = string2bytes(s)
	return e
}

// FromFile encrypts from file.
func (e encrypt) FromFile(f string) encrypt {
	if f == "" {
		return e
	}
	bytes, err := ioutil.ReadFile(f)
	if err != nil {
		e.Error = invalidFileError(f)
		return e
	}
	e.input = bytes
	return e
}

// FromBytes encrypts from byte slice.
func (e encrypt) FromBytes(b []byte) encrypt {
	if len(b) > 0 {
		e.input = b
	}
	return e
}

// ToString outputs as string with given encoding mode.
func (e encrypt) ToString(encodingMode ...string) string {
	if len(e.output) == 0 {
		return emptyString
	}
	mode := HEX
	if len(encodingMode) > 0 {
		mode = encodingMode[len(encodingMode)-1]
	}
	return e.encode(mode)
}

// ToBytes outputs as byte slice with given encoding mode.
func (e encrypt) ToBytes(encodingMode ...string) []byte {
	bytes := string2bytes(e.ToString(encodingMode...))
	if len(bytes) == 0 {
		return emptyBytes
	}
	return bytes
}

func (e encrypt) encode(encodingMode string) string {
	switch encodingMode {
	case HEX:
		return Encode.FromBytes(e.output).ByHex().ToString()
	case BASE32:
		return Encode.FromBytes(e.output).ByBase32().ToString()
	case BASE58:
		return Encode.FromBytes(e.output).ByBase58().ToString()
	case BASE64:
		return Encode.FromBytes(e.output).ByBase64().ToString()
	}
	return emptyString
}

func (e encrypt) encrypt(cipher *Cipher, block cipher.Block) ([]byte, error) {
	src := emptyBytes
	switch cipher.padding {
	case ZeroPadding:
		src = cipher.ZeroPadding(e.input, block.BlockSize())
	case PKCS5Padding:
		src = cipher.PKCS5Padding(e.input)
	case PKCS7Padding:
		src = cipher.PKCS7Padding(e.input, block.BlockSize())
	default:
		return src, invalidPaddingModeError(cipher.padding)
	}
	switch cipher.mode {
	case CBC:
		return cipher.CBCEncrypt(block, src), nil
	}
	return src, invalidGroupModeError(cipher.mode)
}
