package dongle

// Decrypter defines Decrypter struct.
type Decrypter struct {
	dongle
}

// newDecrypter returns a new Decrypter instance.
func newDecrypter() Decrypter {
	return Decrypter{}
}

// FromRawString decrypts from raw string without encoding.
func (d Decrypter) FromRawString(s string) Decrypter {
	d.src = string2bytes(s)
	return d
}

// FromHexString decrypts from string with hex encoding.
func (d Decrypter) FromHexString(s string) Decrypter {
	decode := Decode.FromString(s).ByHex()
	if decode.Error != nil {
		d.Error = NewDecodeError().ModeError("hex")
		return d
	}
	d.src = decode.ToBytes()
	return d
}

// FromBase64String decrypts from string with base64 encoding.
func (d Decrypter) FromBase64String(s string) Decrypter {
	decode := Decode.FromString(s).ByBase64()
	if decode.Error != nil {
		d.Error = NewDecodeError().ModeError("base64")
		return d
	}
	d.src = decode.ToBytes()
	return d
}

// FromRawBytes decrypts from raw byte slice without encoding.
func (d Decrypter) FromRawBytes(b []byte) Decrypter {
	d.src = b
	return d
}

// FromHexBytes decrypts from byte slice with hex encoding.
func (d Decrypter) FromHexBytes(b []byte) Decrypter {
	decode := Decode.FromBytes(b).ByHex()
	if decode.Error != nil {
		d.Error = NewDecodeError().ModeError("hex")
		return d
	}
	d.src = decode.ToBytes()
	return d
}

// FromBase64Bytes decrypts from byte slice with base64 encoding.
func (d Decrypter) FromBase64Bytes(b []byte) Decrypter {
	decode := Decode.FromBytes(b).ByBase64()
	if decode.Error != nil {
		d.Error = NewDecodeError().ModeError("base64")
		return d
	}
	d.src = decode.ToBytes()
	return d
}

// String implements Stringer interface for Decrypter struct.
func (d Decrypter) String() string {
	return d.ToString()
}

// ToString outputs as string.
func (d Decrypter) ToString() string {
	return bytes2string(d.dst)
}

// ToBytes outputs as byte slice.
func (d Decrypter) ToBytes() []byte {
	if len(d.dst) == 0 {
		return []byte("")
	}
	return d.dst
}
