package dongle

type Signer struct {
	dongle
}

// newSigner returns a new Signer instance.
func newSigner() Signer {
	return Signer{}
}

// FromString signs from string.
func (s Signer) FromString(message string) Signer {
	s.src = string2bytes(message)
	return s
}

// FromBytes signs from byte slice.
func (s Signer) FromBytes(message []byte) Signer {
	s.src = message
	return s
}

// String implements Stringer interface for Signer struct.
func (s Signer) String() string {
	return s.ToRawString()
}

// ToRawString outputs as raw string without encoding.
func (s Signer) ToRawString() string {
	return bytes2string(s.dst)
}

// ToHexString outputs as string with hex encoding.
func (s Signer) ToHexString() string {
	return Encode.FromBytes(s.dst).ByHex().ToString()
}

// ToBase64String outputs as string with base64 encoding.
func (s Signer) ToBase64String() string {
	return Encode.FromBytes(s.dst).ByBase64().ToString()
}

// ToRawBytes outputs as raw byte slice without encoding.
func (s Signer) ToRawBytes() []byte {
	if len(s.dst) == 0 {
		return []byte("")
	}
	return s.dst
}

// ToHexBytes outputs as byte with hex encoding.
func (s Signer) ToHexBytes() []byte {
	return Encode.FromBytes(s.dst).ByHex().ToBytes()
}

// ToBase64Bytes outputs as byte slice with base64 encoding.
func (s Signer) ToBase64Bytes() []byte {
	return Encode.FromBytes(s.dst).ByBase64().ToBytes()
}
