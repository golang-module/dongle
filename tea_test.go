package dongle

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

var teaTests = []struct {
	input    string
	key      []byte
	rounds   int
	toHex    string
	toBase64 string
}{
	{"", []byte(""), 64, "", ""},
	{"", []byte("0123456789abcdef"), 64, "", ""},
	{"hello go", []byte("0123456789abcdef"), 0, "3f9a9d3f2f58277f", "P5qdPy9YJ38="},
	{"hello world", []byte("0123456789abcdef"), 32, "e63b448416433bda1610497542fa66a7", "5jtEhBZDO9oWEEl1Qvpmpw=="},
	{"hello world, hello world, hello world, hello world, hello world, hello world, hello world, hello world, hello world, hello world, hello world, hello world, hello world, hello world, hello world, hello world, hello world, hello world, hello world, hello world, hello world, hello world, hello world, hello world, hello world, hello world, hello world, hello world, hello world, hello world, hello world, hello world, hello world, hello world, hello world, hello world, hello world, hello world, hello world, hello world, hello world, hello world, hello world, hello world, hello world, hello world, hello world, hello world, hello world, hello world, hello world, hello world, hello world, hello world, hello world, hello world, hello world, hello world, hello world, hello world, hello world, hello world, hello world, hello world, hello world, hello world, hello world, hello world", []byte("0123456789abcdef"), 64, "bfa8d956247c0dce3d10237bc95b856f16a79755d41c8503c93450b273110d9d93722af35615315b71cd2258c94efdb3527b3b10a4d8297e3f580fe6b42b368aed5195ce34ea5edcd0f74478ba9369708023b1a013f199ac10e3070b4134003e563adc981e1f34dcbfa8d956247c0dce3d10237bc95b856f16a79755d41c8503c93450b273110d9d93722af35615315b71cd2258c94efdb3527b3b10a4d8297e3f580fe6b42b368aed5195ce34ea5edcd0f74478ba9369708023b1a013f199ac10e3070b4134003e563adc981e1f34dcbfa8d956247c0dce3d10237bc95b856f16a79755d41c8503c93450b273110d9d93722af35615315b71cd2258c94efdb3527b3b10a4d8297e3f580fe6b42b368aed5195ce34ea5edcd0f74478ba9369708023b1a013f199ac10e3070b4134003e563adc981e1f34dcbfa8d956247c0dce3d10237bc95b856f16a79755d41c8503c93450b273110d9d93722af35615315b71cd2258c94efdb3527b3b10a4d8297e3f580fe6b42b368aed5195ce34ea5edcd0f74478ba9369708023b1a013f199ac10e3070b4134003e563adc981e1f34dcbfa8d956247c0dce3d10237bc95b856f16a79755d41c8503c93450b273110d9d93722af35615315b71cd2258c94efdb3527b3b10a4d8297e3f580fe6b42b368aed5195ce34ea5edcd0f74478ba9369708023b1a013f199ac10e3070b4134003e563adc981e1f34dcbfa8d956247c0dce3d10237bc95b856f16a79755d41c8503c93450b273110d9d93722af35615315b71cd2258c94efdb3527b3b10a4d8297e3f580fe6b42b368aed5195ce34ea5edcd0f74478ba9369708023b1a013f199ac10e3070b4134003e563adc981e1f34dcbfa8d956247c0dce3d10237bc95b856f16a79755d41c8503c93450b273110d9d93722af35615315b71cd2258c94efdb3527b3b10a4d8297e3f580fe6b42b368aed5195ce34ea5edcd0f74478ba9369708023b1a013f199ac10e3070b4134003e563adc981e1f34dcbfa8d956247c0dce3d10237bc95b856f16a79755d41c8503c93450b273110d9d93722af35615315b71cd2258c94efdb3527b3b10a4d8297e3f580fe6b42b368aed5195ce34ea5edcd0f74478ba9369708023b1a013f199ac10e3070b4134003e563adc981e1f34dcbfa8d956247c0dce3d10237bc95b856f16a79755d41c8503c93450b273110d9d93722af35615315b71cd2258c94efdb3a831e8fcd2166d8e", "v6jZViR8Dc49ECN7yVuFbxanl1XUHIUDyTRQsnMRDZ2TcirzVhUxW3HNIljJTv2zUns7EKTYKX4/WA/mtCs2iu1Rlc406l7c0PdEeLqTaXCAI7GgE/GZrBDjBwtBNAA+VjrcmB4fNNy/qNlWJHwNzj0QI3vJW4VvFqeXVdQchQPJNFCycxENnZNyKvNWFTFbcc0iWMlO/bNSezsQpNgpfj9YD+a0KzaK7VGVzjTqXtzQ90R4upNpcIAjsaAT8ZmsEOMHC0E0AD5WOtyYHh803L+o2VYkfA3OPRAje8lbhW8Wp5dV1ByFA8k0ULJzEQ2dk3Iq81YVMVtxzSJYyU79s1J7OxCk2Cl+P1gP5rQrNortUZXONOpe3ND3RHi6k2lwgCOxoBPxmawQ4wcLQTQAPlY63JgeHzTcv6jZViR8Dc49ECN7yVuFbxanl1XUHIUDyTRQsnMRDZ2TcirzVhUxW3HNIljJTv2zUns7EKTYKX4/WA/mtCs2iu1Rlc406l7c0PdEeLqTaXCAI7GgE/GZrBDjBwtBNAA+VjrcmB4fNNy/qNlWJHwNzj0QI3vJW4VvFqeXVdQchQPJNFCycxENnZNyKvNWFTFbcc0iWMlO/bNSezsQpNgpfj9YD+a0KzaK7VGVzjTqXtzQ90R4upNpcIAjsaAT8ZmsEOMHC0E0AD5WOtyYHh803L+o2VYkfA3OPRAje8lbhW8Wp5dV1ByFA8k0ULJzEQ2dk3Iq81YVMVtxzSJYyU79s1J7OxCk2Cl+P1gP5rQrNortUZXONOpe3ND3RHi6k2lwgCOxoBPxmawQ4wcLQTQAPlY63JgeHzTcv6jZViR8Dc49ECN7yVuFbxanl1XUHIUDyTRQsnMRDZ2TcirzVhUxW3HNIljJTv2zUns7EKTYKX4/WA/mtCs2iu1Rlc406l7c0PdEeLqTaXCAI7GgE/GZrBDjBwtBNAA+VjrcmB4fNNy/qNlWJHwNzj0QI3vJW4VvFqeXVdQchQPJNFCycxENnZNyKvNWFTFbcc0iWMlO/bNSezsQpNgpfj9YD+a0KzaK7VGVzjTqXtzQ90R4upNpcIAjsaAT8ZmsEOMHC0E0AD5WOtyYHh803L+o2VYkfA3OPRAje8lbhW8Wp5dV1ByFA8k0ULJzEQ2dk3Iq81YVMVtxzSJYyU79s6gx6PzSFm2O"},
}

func TestTea_Encrypt_String(t *testing.T) {
	for index, test := range teaTests {
		e1 := Encrypt.FromString(test.input).ByTea(test.key, test.rounds)
		if test.rounds == 0 {
			e1 = Encrypt.FromString(test.input).ByTea(test.key)
		}
		assert.Nil(t, e1.Error)
		assert.Equal(t, test.toHex, e1.ToHexString(), "Hex test index is "+strconv.Itoa(index))
		assert.Equal(t, test.toBase64, e1.ToBase64String(), "Base64 test index is "+strconv.Itoa(index))

		e2 := Encrypt.FromString(test.input).ByTea(test.key, test.rounds)
		if test.rounds == 0 {
			e2 = Encrypt.FromString(test.input).ByTea(test.key)
		}
		assert.Nil(t, e2.Error)
		assert.Equal(t, test.toHex, e2.ToHexString(), "Hex test index is "+strconv.Itoa(index))
		assert.Equal(t, test.toBase64, e2.ToBase64String(), "Base64 test index is "+strconv.Itoa(index))
	}
}

func TestTea_Decrypt_String(t *testing.T) {
	for index, test := range teaTests {
		d1 := Decrypt.FromHexString(test.toHex).ByTea(test.key, test.rounds)
		if test.rounds == 0 {
			d1 = Decrypt.FromHexString(test.toHex).ByTea(test.key)
		}
		assert.Nil(t, d1.Error)
		assert.Equal(t, test.input, d1.ToString(), "Hex test index is "+strconv.Itoa(index))

		d2 := Decrypt.FromBase64String(test.toBase64).ByTea(test.key, test.rounds)
		if test.rounds == 0 {
			d2 = Decrypt.FromBase64String(test.toBase64).ByTea(test.key)
		}
		assert.Nil(t, d2.Error)
		assert.Equal(t, test.input, d2.ToString(), "base64 test index is "+strconv.Itoa(index))
	}
}

func TestTea_Encrypt_Bytes(t *testing.T) {
	for index, test := range teaTests {
		e := Encrypt.FromBytes([]byte(test.input)).ByTea(test.key, test.rounds)
		if test.rounds == 0 {
			e = Encrypt.FromBytes([]byte(test.input)).ByTea(test.key)
		}
		assert.Nil(t, e.Error)
		assert.Equal(t, []byte(test.toHex), e.ToHexBytes(), "Hex test test index is "+strconv.Itoa(index))
		assert.Equal(t, []byte(test.toBase64), e.ToBase64Bytes(), "Base64 test index is "+strconv.Itoa(index))
	}
}

func TestTea_Decrypt_Bytes(t *testing.T) {
	for index, test := range teaTests {
		e1 := Decrypt.FromHexBytes([]byte(test.toHex)).ByTea(test.key, test.rounds)
		if test.rounds == 0 {
			e1 = Decrypt.FromHexBytes([]byte(test.toHex)).ByTea(test.key)
		}
		assert.Nil(t, e1.Error)
		assert.Equal(t, []byte(test.input), e1.ToBytes(), "Hex test index is "+strconv.Itoa(index))

		e2 := Decrypt.FromBase64Bytes([]byte(test.toBase64)).ByTea(test.key, test.rounds)
		if test.rounds == 0 {
			e2 = Decrypt.FromBase64Bytes([]byte(test.toBase64)).ByTea(test.key)
		}
		assert.Nil(t, e2.Error)
		assert.Equal(t, []byte(test.input), e2.ToBytes(), "base64 test index is "+strconv.Itoa(index))
	}
}

func TestTea_Key_Error(t *testing.T) {
	key := []byte("xxx")

	e := Encrypt.FromString("hello go").ByTea(key, 8)
	assert.Equal(t, teaError.KeyError(key), e.Error)

	d := Decrypt.FromRawString("hello go").ByTea(key, 8)
	assert.Equal(t, teaError.KeyError(key), d.Error)
}

func TestTea_Rounds_Error(t *testing.T) {
	key := []byte("0123456789abcdefghijklmn")

	e := Encrypt.FromString("hello go").ByTea(key, 1)
	assert.Equal(t, teaError.RoundsError(), e.Error)

	d := Decrypt.FromRawString("hello go").ByTea(key, 1)
	assert.Equal(t, teaError.RoundsError(), d.Error)
}
