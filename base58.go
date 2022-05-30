package dongle

import (
	"bytes"
	"math/big"
)

const base58table = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"

// ByBase58 encodes by base58.
// 通过 base58 编码
func (e encode) ByBase58() encode {
	if len(e.src) == 0 {
		return e
	}
	intBytes := big.NewInt(0).SetBytes(e.src)
	int0, int58 := big.NewInt(0), big.NewInt(58)
	for intBytes.Cmp(big.NewInt(0)) > 0 {
		intBytes.DivMod(intBytes, int58, int0)
		e.dst = append(e.dst, string2bytes(base58table)[int0.Int64()])
	}
	e.dst = reverseBytes(e.dst)
	return e
}

// ByBase58 decodes by base58.
// 通过 base58 解码
func (d decode) ByBase58() decode {
	if len(d.src) == 0 {
		return d
	}
	Int0 := big.NewInt(0)
	for _, v := range d.src {
		index := bytes.IndexByte([]byte(base58table), v)
		Int0.Mul(Int0, big.NewInt(58))
		Int0.Add(Int0, big.NewInt(int64(index)))
	}
	d.dst = Int0.Bytes()
	return d
}

// reverses byte slice.
// 反转字节切片
func reverseBytes(b []byte) []byte {
	for i := 0; i < len(b)/2; i++ {
		b[i], b[len(b)-1-i] = b[len(b)-1-i], b[i]
	}
	return b
}
