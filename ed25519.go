package dongle

import (
	"crypto/ed25519"
	"fmt"
)

type encodingMode string

type Ed25519Error struct {
}

func NewEd25519Error() Ed25519Error {
	return Ed25519Error{}
}

func (e Ed25519Error) PrivateKeyError() error {
	return fmt.Errorf("ed25519: invalid private key, please make sure the private key is valid")
}

func (e Ed25519Error) PublicKeyError() error {
	return fmt.Errorf("ed25519: invalid public key, please make sure the public key is valid")
}

func (e Ed25519Error) SignatureError() error {
	return fmt.Errorf("ed25519: invalid signature, please make sure the signature is valid")
}

const (
	Raw    encodingMode = "raw"
	Hex    encodingMode = "hex"
	Base64 encodingMode = "base64"
)

var ed25519Error = NewEd25519Error()

// ByEd25519 signs by ed25519.
func (s Signer) ByEd25519(privateKey []byte, mode encodingMode) Signer {
	if len(s.src) == 0 || s.Error != nil {
		return s
	}
	pri, err := mode.getDecodedKey(privateKey)
	if err != nil {
		s.Error = err
		return s
	}
	if len(pri) != ed25519.PrivateKeySize {
		s.Error = ed25519Error.PrivateKeyError()
		return s
	}
	s.dst = ed25519.Sign(pri, s.src)
	return s
}

// ByEd25519 verify by ed25519.
func (v Verifier) ByEd25519(publicKey []byte, mode encodingMode) Verifier {
	if len(v.src) == 0 || v.Error != nil {
		return v
	}
	pub, err := mode.getDecodedKey(publicKey)
	if err != nil {
		v.Error = err
		return v
	}
	if len(pub) != ed25519.PublicKeySize {
		v.Error = ed25519Error.PublicKeyError()
		return v
	}
	if ed25519.Verify(pub, v.src, v.sign) == false {
		v.Error = ed25519Error.SignatureError()
		return v
	}
	return v
}

// gets the decoded key
func (mode encodingMode) getDecodedKey(key []byte) (dst []byte, err error) {
	var decode Decoder
	switch mode {
	case Raw:
		dst = key
		return
	case Hex:
		decode = Decode.FromBytes(key).ByHex()
	case Base64:
		decode = Decode.FromBytes(key).ByBase64()
	}
	if decode.Error != nil {
		err = decode.Error
		return
	}
	dst = decode.ToBytes()
	return
}
