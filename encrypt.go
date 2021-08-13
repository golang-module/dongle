package dongle

import (
	"encoding/base32"
	"encoding/base64"
	"encoding/hex"
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
func (e encrypt) ToString(encodeMode ...string) string {
	if len(e.output) == 0 {
		return ""
	}
	mode := HEX
	if len(encodeMode) > 0 {
		mode = encodeMode[len(encodeMode)-1]
	}
	if mode == HEX {
		return hex.EncodeToString(e.output)
	}
	if mode == BASE32 {
		return base32.StdEncoding.EncodeToString(e.output)
	}
	if mode == BASE64 {
		return base64.StdEncoding.EncodeToString(e.output)
	}
	return ""
}

// ToBytes outputs as byte slice with given encoding mode.
func (e encrypt) ToBytes(encodeMode ...string) []byte {
	bytes := string2bytes(e.ToString(encodeMode...))
	if len(bytes) == 0 {
		return []byte("")
	}
	return bytes
}
