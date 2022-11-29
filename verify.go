package dongle

import (
	"crypto"
)

// defines a verify struct.
// 定义 verify 结构体
type verify struct {
	dongle
	hash crypto.Hash
}

// returns a new verify instance.
// 初始化 verify 结构体
func newVerify() verify {
	return verify{}
}

// FromString verify from string.
// 对字符串进行验签
func (v verify) FromString(s string) verify {
	v.src = string2bytes(s)
	return v
}

// FromBytes decodes from byte slice.
// 对字节切片进行验签
func (v verify) FromBytes(b []byte) verify {
	v.src = b
	return v
}
