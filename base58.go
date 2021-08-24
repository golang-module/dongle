package dongle

import (
	"bytes"
	"math/big"
)

const base58table = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"

// ByBase58 encodes by base58.
func (e encode) ByBase58() encode {
	input, output := e.input, e.output
	if input == nil {
		return e
	}
	if output != nil {
		input = output
	}
	var src []byte
	intBytes := big.NewInt(0).SetBytes(input)
	int0, int58 := big.NewInt(0), big.NewInt(58)
	for intBytes.Cmp(big.NewInt(0)) > 0 {
		intBytes.DivMod(intBytes, int58, int0)
		src = append(src, string2bytes(base58table)[int0.Int64()])
	}
	e.output = reverseBytes(src)
	return e
}

// ByBase58 encodes by base58.
func (d decode) ByBase58() decode {
	input, output := d.input, d.output
	if input == nil {
		return d
	}
	if output != nil {
		input = output
	}
	int0 := big.NewInt(0)
	for _, val := range input {
		index := bytes.IndexByte([]byte(base58table), val)
		int0.Mul(int0, big.NewInt(58))
		int0.Add(int0, big.NewInt(int64(index)))
	}
	d.output = int0.Bytes()
	return d
}

// reverseBytes reverses byte slice.
func reverseBytes(b []byte) []byte {
	for i := 0; i < len(b)/2; i++ {
		b[i], b[len(b)-1-i] = b[len(b)-1-i], b[i]
	}
	return b
}
