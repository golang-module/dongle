package dongle

type Encrypter struct {
	dongle
}

// newEncrypter returns a new Encrypter instance.
func newEncrypter() Encrypter {
	return Encrypter{}
}

// FromString encrypts from string.
func (e Encrypter) FromString(s string) Encrypter {
	e.src = string2bytes(s)
	return e
}

// FromBytes encrypts from byte slice.
func (e Encrypter) FromBytes(b []byte) Encrypter {
	e.src = b
	return e
}

// String implements Stringer interface for Encrypter struct.
func (e Encrypter) String() string {
	return e.ToRawString()
}

// ToRawString outputs as raw string without encoding.
func (e Encrypter) ToRawString() string {
	return bytes2string(e.dst)
}

// ToHexString outputs as string with hex encoding.
func (e Encrypter) ToHexString() string {
	return Encode.FromBytes(e.dst).ByHex().ToString()
}

// ToBase64String outputs as string with base64 encoding.
func (e Encrypter) ToBase64String() string {
	return Encode.FromBytes(e.dst).ByBase64().ToString()
}

// ToRawBytes outputs as raw byte slice without encoding.
func (e Encrypter) ToRawBytes() []byte {
	if len(e.dst) == 0 {
		return []byte("")
	}
	return e.dst
}

// ToHexBytes outputs as byte with hex encoding.
func (e Encrypter) ToHexBytes() []byte {
	return Encode.FromBytes(e.dst).ByHex().ToBytes()
}

// ToBase64Bytes outputs as byte slice with base64 encoding.
func (e Encrypter) ToBase64Bytes() []byte {
	return Encode.FromBytes(e.dst).ByBase64().ToBytes()
}
