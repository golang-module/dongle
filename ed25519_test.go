package dongle

import (
	"crypto/ed25519"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	ed25519Input = "hello world"
)

func TestEd25519_String(t *testing.T) {
	var publicKey, privateKey []byte
	publicKey, privateKey, _ = ed25519.GenerateKey(nil)

	rawPrivateKey := string(privateKey)
	rawPublicKey := string(publicKey)
	s1 := Sign.FromString(ed25519Input).ByEd25519(rawPrivateKey, Raw)
	assert.Nil(t, s1.Error)
	v1 := Verify.FromRawString(s1.ToRawString(), "hello world").ByEd25519(rawPublicKey, Raw)
	assert.Nil(t, v1.Error)
	assert.Equal(t, true, v1.ToBool())

	hexPrivateKey := Encode.FromString(rawPrivateKey).ByHex().ToString()
	hexPublicKey := Encode.FromString(rawPublicKey).ByHex().ToString()
	s2 := Sign.FromString(ed25519Input).ByEd25519(hexPrivateKey, Hex)
	assert.Nil(t, s2.Error)
	v2 := Verify.FromHexString(s1.ToHexString(), "hello world").ByEd25519(hexPublicKey, Hex)
	assert.Nil(t, v2.Error)
	assert.Equal(t, true, v2.ToBool())

	base64PrivateKey := Encode.FromBytes(privateKey).ByBase64().ToString()
	base64PublicKey := Encode.FromBytes(publicKey).ByBase64().ToString()
	s3 := Sign.FromString(ed25519Input).ByEd25519(base64PrivateKey, Base64)
	assert.Nil(t, s3.Error)
	v3 := Verify.FromHexString(s1.ToHexString(), "hello world").ByEd25519(base64PublicKey, Base64)
	assert.Nil(t, v3.Error)
	assert.Equal(t, true, v3.ToBool())
}

func TestEd25519_Sign_Bytes(t *testing.T) {
	var publicKey, privateKey []byte
	publicKey, privateKey, _ = ed25519.GenerateKey(nil)

	rawPrivateKey := privateKey
	rawPublicKey := publicKey
	s1 := Sign.FromBytes([]byte(ed25519Input)).ByEd25519(rawPrivateKey, Raw)
	assert.Nil(t, s1.Error)
	v1 := Verify.FromRawBytes(s1.ToRawBytes(), []byte("hello world")).ByEd25519(rawPublicKey, Raw)
	assert.Nil(t, v1.Error)
	assert.Equal(t, true, v1.ToBool())

	hexPrivateKey := Encode.FromBytes(rawPrivateKey).ByHex().ToBytes()
	hexPublicKey := Encode.FromBytes(rawPublicKey).ByHex().ToBytes()
	s2 := Sign.FromBytes([]byte(ed25519Input)).ByEd25519(hexPrivateKey, Hex)
	assert.Nil(t, s2.Error)
	v2 := Verify.FromHexBytes(s1.ToHexBytes(), []byte("hello world")).ByEd25519(hexPublicKey, Hex)
	assert.Nil(t, v2.Error)
	assert.Equal(t, true, v2.ToBool())

	base64PrivateKey := Encode.FromBytes(privateKey).ByBase64().ToBytes()
	base64PublicKey := Encode.FromBytes(publicKey).ByBase64().ToBytes()
	s3 := Sign.FromBytes([]byte(ed25519Input)).ByEd25519(base64PrivateKey, Base64)
	assert.Nil(t, s3.Error)
	v3 := Verify.FromBase64Bytes(s1.ToBase64Bytes(), []byte("hello world")).ByEd25519(base64PublicKey, Base64)
	assert.Nil(t, v3.Error)
	assert.Equal(t, true, v3.ToBool())
}

func TestEd25519_Empty(t *testing.T) {
	var publicKey, privateKey []byte
	publicKey, privateKey, _ = ed25519.GenerateKey(nil)

	s1 := Sign.FromString("").ByEd25519(privateKey, Raw)
	assert.Nil(t, s1.Error)
	v1 := Verify.FromRawString(s1.ToRawString(), "").ByEd25519(publicKey, Raw)
	assert.Nil(t, v1.Error)
	assert.Equal(t, false, v1.ToBool())

	s2 := Sign.FromString("").ByEd25519(privateKey, Hex)
	assert.Nil(t, s2.Error)
	v2 := Verify.FromHexString(s2.ToHexString(), "").ByEd25519(publicKey, Hex)
	assert.Nil(t, v2.Error)
	assert.Equal(t, false, v2.ToBool())

	s3 := Sign.FromString("").ByEd25519(privateKey, Hex)
	assert.Nil(t, s3.Error)
	v3 := Verify.FromBase64String(s3.ToBase64String(), "").ByEd25519(publicKey, Hex)
	assert.Nil(t, v3.Error)
	assert.Equal(t, false, v3.ToBool())
}

func TestEd25519_PrivateKey_Error(t *testing.T) {
	s := Sign.FromString(ed25519Input).ByEd25519("xxxx", Raw)
	assert.Equal(t, invalidEd25519PrivateKeyError(), s.Error)
}

func TestEd25519_PublicKey_Error(t *testing.T) {
	s := Verify.FromBase64String("", ed25519Input).ByEd25519("xxxx", Raw)
	assert.Equal(t, invalidEd25519PublicKeyError(), s.Error)
}

func TestEd25519_Signature_Error(t *testing.T) {
	var publicKey []byte
	publicKey, _, _ = ed25519.GenerateKey(nil)

	s := Verify.FromRawString("xxxx", ed25519Input).ByEd25519(publicKey, Raw)
	assert.Equal(t, invalidEd25519SignatureError(), s.Error)
}

func TestEd25519_Decoding_Error(t *testing.T) {
	s1 := Sign.FromString(ed25519Input).ByEd25519("xxxx", Hex)
	assert.Equal(t, invalidDecodingError("hex"), s1.Error)
	v1 := Verify.FromHexString("68656c6c6f20776f726c64", ed25519Input).ByEd25519("xxxx", Hex)
	assert.Equal(t, invalidDecodingError("hex"), v1.Error)

	s2 := Sign.FromString(ed25519Input).ByEd25519("xxxxxx", Base64)
	assert.Equal(t, invalidDecodingError("base64"), s2.Error)
	v2 := Verify.FromBase64String("aGVsbG8gd29ybGQ=", ed25519Input).ByEd25519("xxxxxx", Base64)
	assert.Equal(t, invalidDecodingError("base64"), v2.Error)
}
