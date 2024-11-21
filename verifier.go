package dongle

type Verifier struct {
	dongle
	sign []byte
}

// newVerifier returns a new Verifier instance.
func newVerifier() Verifier {
	return Verifier{}
}

// FromRawString verify signature from raw string.
func (v Verifier) FromRawString(signature, message string) Verifier {
	v.src = string2bytes(message)
	v.sign = string2bytes(signature)
	return v
}

// FromHexString decrypts from string with hex encoding.
func (v Verifier) FromHexString(signature, message string) Verifier {
	d := Decode.FromString(signature).ByHex()
	if d.Error != nil {
		v.Error = NewDecodeError().ModeError("hex")
		return v
	}
	v.sign = d.ToBytes()
	v.src = string2bytes(message)
	return v
}

// FromBase64String decrypts from string with base64 encoding.
func (v Verifier) FromBase64String(signature, message string) Verifier {
	d := Decode.FromString(signature).ByBase64()
	if d.Error != nil {
		v.Error = NewDecodeError().ModeError("base64")
		return v
	}
	v.sign = d.ToBytes()
	v.src = string2bytes(message)
	return v
}

// FromRawBytes verify signature from raw byte slice.
func (v Verifier) FromRawBytes(signature, message []byte) Verifier {
	v.sign = signature
	v.src = message
	return v
}

// FromHexBytes verify from byte slice with hex encoding.
func (v Verifier) FromHexBytes(signature, message []byte) Verifier {
	d := Decode.FromBytes(signature).ByHex()
	if d.Error != nil {
		v.Error = NewDecodeError().ModeError("hex")
		return v
	}
	v.sign = d.ToBytes()
	v.src = message
	return v
}

// FromBase64Bytes verify from byte slice with base64 encoding.
func (v Verifier) FromBase64Bytes(signature, message []byte) Verifier {
	d := Decode.FromBytes(signature).ByBase64()
	if d.Error != nil {
		v.Error = NewDecodeError().ModeError("base64")
		return v
	}
	v.sign = d.ToBytes()
	v.src = message
	return v
}

// ToBool outputs as bool.
func (v Verifier) ToBool() bool {
	if len(v.src) == 0 || len(v.sign) == 0 {
		return false
	}
	return v.Error == nil
}
